//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//
// tClient is an HTTP client instance.
type tClient struct {
	auth     Provider
	endpoint string
	verbose  bool
	retry    int
	http     *http.Client
}

//
// New creates an instance of HTTP client
func New(opts ...Option) Connector {
	client := &tClient{
		http: &http.Client{
			Transport: &http.Transport{
				ReadBufferSize: 128 * 1024,
				Dial: (&net.Dialer{
					Timeout: 10 * time.Second,
				}).Dial,
			},
		},
		retry: 2,
	}

	for _, opt := range opts {
		client = opt(client)
	}

	return client
}

//
func (client *tClient) doWithRetry(req *http.Request) (*http.Response, error) {
	for i := 0; i < client.retry; i++ {
		in, err := client.do(req)
		if err != nil {
			return nil, err
		}

		if in.StatusCode == http.StatusUnauthorized {
			continue
		}

		return in, nil
	}

	return nil, fmt.Errorf("request failed after %d tries", client.retry)
}

func (client *tClient) do(req *http.Request) (*http.Response, error) {
	if client.auth != nil {
		heads, err := client.auth.Headers()
		if err != nil {
			return nil, err
		}
		for h, v := range heads {
			req.Header.Set(h, v)
		}
	}

	return client.http.Do(req)
}

// URL creates URL connector
func (client *tClient) URL(method, url string) CURL {
	return &tCURL{
		client:  client,
		method:  method,
		url:     client.endpoint + url,
		payload: bytes.NewBuffer(nil),
	}
}

// Get creates URL connector
func (client *tClient) Get(templateURL string, args ...interface{}) CURL {
	return client.URL(http.MethodGet, fmt.Sprintf(templateURL, args...))
}

// Put creates URL connector
func (client *tClient) Put(templateURL string, args ...interface{}) CURL {
	return client.URL(http.MethodPut, fmt.Sprintf(templateURL, args...))
}

// Post creates URL connector
func (client *tClient) Post(templateURL string, args ...interface{}) CURL {
	return client.URL(http.MethodPost, fmt.Sprintf(templateURL, args...))
}

// CURL is a builder type, constructs HTTP request
type tCURL struct {
	client  *tClient
	method  string
	url     string
	payload *bytes.Buffer
	output  *http.Response
	fail    error
}

// Send payload to destination URL.
func (curl *tCURL) Send(data interface{}) CURL {
	return curl.encodeJSON(data)
}

func (curl *tCURL) encodeJSON(data interface{}) CURL {
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
func (curl *tCURL) Recv(data interface{}) error {
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
		return ErrorFromResponse(curl.output, body)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}

// RecvStatus payload from target URL and discards it.
func (curl *tCURL) RecvStatus(status ...int) (http.Header, error) {
	curl = curl.unsafeIO()

	if curl.fail != nil {
		return nil, curl.fail
	}

	defer curl.output.Body.Close()
	body, err := ioutil.ReadAll(curl.output.Body)
	if err != nil {
		return nil, err
	}

	expect := http.StatusOK
	if len(status) == 1 {
		expect = status[0]
	}
	if curl.output.StatusCode != expect {
		return nil, ErrorFromResponse(curl.output, body)
	}

	return curl.output.Header, nil
}

//
func (curl *tCURL) unsafeIO() *tCURL {
	if curl.fail != nil {
		return curl
	}

	req, err := http.NewRequest(curl.method, curl.url, curl.payload)
	if curl.fail = err; err != nil {
		return curl
	}

	curl.output, curl.fail = curl.client.doWithRetry(req)
	return curl
}
