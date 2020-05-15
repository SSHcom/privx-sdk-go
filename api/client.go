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

type Client struct {
	Auth     *oauth.Client
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

// Authenticator setup credential provider for api
func Authenticator(auth *oauth.Client) Option {
	return func(client *Client) *Client {
		client.Auth = auth
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
func NewClient(opts ...Option) (*Client, error) {
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

	return client, nil
}

func (client *Client) Endpoint() string {
	return client.endpoint
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	retryLimit := 2
	for i := 0; i < retryLimit; i++ {
		token, err := client.Auth.Token()
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
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

//
type Request struct {
	client  *Client
	http    *http.Request
	payload *bytes.Buffer
	fail    error
}

func (req Request) encodeJSON(data ...interface{}) Request {
	if req.fail != nil {
		return req
	}

	if len(data) > 0 {
		encoded, err := json.Marshal(data[1])
		req.payload = bytes.NewBuffer(encoded)
		req.fail = err
	}

	return req
}

//
func (client *Client) Get(url string, data ...interface{}) Request {
	request := Request{client: client}
	request.encodeJSON(data...)

	if request.fail != nil {
		return request
	}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/%s", client.endpoint, url),
		request.payload,
	)

	request.http = req
	request.fail = err

	return request
}

//
func (req Request) Recv(data interface{}) error {
	out, err := req.client.Do(req.http)
	if err != nil {
		return err
	}

	defer out.Body.Close()
	body, err := ioutil.ReadAll(out.Body)
	if err != nil {
		return err
	}

	if out.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", out.Status)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}
