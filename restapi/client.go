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
	"net/url"
	"time"
)

//
// tClient is an HTTP client instance.
type tClient struct {
	auth     Authorizer
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
		token, err := client.auth.AccessToken()
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", token)
	}
	req.Header.Set("User-Agent", UserAgent)

	return client.http.Do(req)
}

// URL creates URL connector
func (client *tClient) URL(method, url string) CURL {
	return &tCURL{
		client:  client,
		method:  method,
		url:     client.endpoint + url,
		header:  http.Header{},
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
	header  http.Header
	payload *bytes.Buffer
	output  *http.Response
	fail    error
}

// Params defines query parameters
func (curl *tCURL) Params(data interface{}) CURL {
	params, err := curl.encodeURL(data)
	if curl.fail = err; err != nil {
		return curl
	}
	curl.url = curl.url + "?" + params.Encode()
	return curl
}

func (curl *tCURL) encodeURL(query interface{}) (url.Values, error) {
	bin, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	var params map[string]string
	if err = json.Unmarshal(bin, &params); err != nil {
		return nil, err
	}

	var values url.Values = make(map[string][]string)
	for key, val := range params {
		values[key] = []string{val}
	}

	return values, nil
}

// With defines HTTP header
func (curl *tCURL) With(head, value string) CURL {
	curl.header.Add(head, value)
	return curl
}

// Send payload to destination URL.
func (curl *tCURL) Send(data interface{}) CURL {
	if curl.fail != nil {
		return curl
	}

	switch curl.header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		return curl.encodeForm(data)
	}

	return curl.encodeJSON(data)
}

func (curl *tCURL) encodeJSON(data interface{}) CURL {
	encoded, err := json.Marshal(data)
	if curl.fail = err; err == nil {
		curl.payload = bytes.NewBuffer(encoded)
	}
	return curl
}

func (curl *tCURL) encodeForm(data interface{}) CURL {
	params, err := curl.encodeURL(data)
	if curl.fail = err; err != nil {
		return curl
	}

	curl.payload = bytes.NewBuffer([]byte(params.Encode()))
	return curl
}

// Recv payload from target URL.
func (curl *tCURL) Recv(data interface{}) (http.Header, error) {
	curl = curl.unsafeIO()

	if curl.fail != nil {
		return nil, curl.fail
	}

	defer curl.output.Body.Close()
	body, err := ioutil.ReadAll(curl.output.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("---")
	fmt.Println(string(body))
	fmt.Println("---")

	if curl.output.StatusCode != http.StatusOK {
		return nil, ErrorFromResponse(curl.output, body)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return curl.output.Header, nil
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

	fmt.Println(curl.url)
	req, err := http.NewRequest(curl.method, curl.url, curl.payload)
	if curl.fail = err; err != nil {
		return curl
	}

	for head := range curl.header {
		req.Header.Set(head, curl.header.Get(head))
	}

	curl.output, curl.fail = curl.client.doWithRetry(req)
	return curl
}
