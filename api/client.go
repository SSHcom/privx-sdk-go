//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package api

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/SSHcom/privx-sdk-go/oauth"
)

// Connector is HTTP connector for api
type Connector interface {
	Get(string, ...interface{}) *CURL
	Put(string, ...interface{}) *CURL
}

// Client is an SDK client instance.
type Client struct {
	auth     oauth.Provider
	endpoint string
	http     *http.Client
}

// Option is configuration applied to the client
type Option func(*Client) *Client

// Endpoint defines a target PrivX API endpoint
func Endpoint(endpoint string) Option {
	return func(client *Client) *Client {
		client.endpoint = endpoint
		return client
	}
}

// AccessToken setup access token provider for api
func AccessToken(auth oauth.Provider) Option {
	return func(client *Client) *Client {
		client.auth = auth
		return client
	}
}

// X509 setup trust certificate
func X509(cert *x509.Certificate) Option {
	return func(client *Client) *Client {
		tlsConfig := &tls.Config{}
		if cert != nil {
			pool := x509.NewCertPool()
			pool.AddCert(cert)
			tlsConfig.RootCAs = pool
		}
		client.http.Transport.(*http.Transport).TLSClientConfig = tlsConfig
		return client
	}
}

// Verbose enables debug-level logging
func Verbose() Option {
	return func(client *Client) *Client {
		return client
	}
}

// NewClient creates an instance of PrivX API client
func NewClient(opts ...Option) *Client {
	client := &Client{
		http: &http.Client{
			Transport: &http.Transport{
				ReadBufferSize: 128 * 1024,
				Dial: (&net.Dialer{
					Timeout: 10 * time.Second,
				}).Dial,
			},
		},
	}

	for _, opt := range opts {
		client = opt(client)
	}

	return client
}

//
// Do executes the argument HTTP request and returns it response. The
// Do function handles the API OAuth2 authentication.
func (client *Client) Do(req *http.Request) (*http.Response, error) {
	retryLimit := 2
	for i := 0; i < retryLimit; i++ {

		if client.auth != nil {
			token, err := client.auth.Token()
			if err != nil {
				return nil, err
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		}

		resp, err := client.http.Do(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == http.StatusUnauthorized {
			continue
		}
		return resp, err
	}
	return nil, fmt.Errorf("request failed after %d tries", retryLimit)
}

// CURL is a builder type, constructs HTTP request
type CURL struct {
	client  *Client
	method  string
	url     string
	payload *bytes.Buffer
	output  *http.Response
	fail    error
}

// URL creates URL connector
func (client *Client) URL(method, url string) *CURL {
	return &CURL{
		client:  client,
		method:  method,
		url:     client.endpoint + url,
		payload: bytes.NewBuffer(nil),
	}
}

// Get creates URL connector
func (client *Client) Get(templateURL string, args ...interface{}) *CURL {
	return client.URL(http.MethodGet, fmt.Sprintf(templateURL, args...))
}

// Post creates URL connector
func (client *Client) Post(templateURL string, args ...interface{}) *CURL {
	return client.URL(http.MethodPost, fmt.Sprintf(templateURL, args...))
}

// Put creates URL connector
func (client *Client) Put(templateURL string, args ...interface{}) *CURL {
	return client.URL(http.MethodPut, fmt.Sprintf(templateURL, args...))
}

// Send payload to destination URL.
func (curl *CURL) Send(data interface{}) *CURL {
	return curl.encodeJSON(data)
}

func (curl *CURL) encodeJSON(data interface{}) *CURL {
	if curl.fail != nil {
		return curl
	}

	encoded, err := json.Marshal(data)
	if curl.fail = err; err == nil {
		curl.payload = bytes.NewBuffer(encoded)
	}
	return curl
}

// Recv payload from target URL.
func (curl *CURL) Recv(data interface{}) error {
	curl = curl.unsafeIO()

	if curl.fail != nil {
		return curl.fail
	}

	defer curl.output.Body.Close()
	body, err := ioutil.ReadAll(curl.output.Body)
	if err != nil {
		return err
	}

	if curl.output.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", curl.output.Status)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}

// RecvStatus payload from target URL and discards it.
func (curl *CURL) RecvStatus() error {
	curl = curl.unsafeIO()

	if curl.fail != nil {
		return curl.fail
	}

	defer curl.output.Body.Close()
	_, err := ioutil.ReadAll(curl.output.Body)
	if err != nil {
		return err
	}

	if curl.output.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", curl.output.Status)
	}

	return nil
}

func (curl *CURL) unsafeIO() *CURL {
	if curl.fail != nil {
		return curl
	}

	req, err := http.NewRequest(curl.method, curl.url, curl.payload)
	if curl.fail = err; err != nil {
		return curl
	}

	curl.output, curl.fail = curl.client.Do(req)
	return curl
}
