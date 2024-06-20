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
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// tClient is an HTTP client instance.
type tClient struct {
	auth    Authorizer
	baseURL string
	verbose bool
	retry   int
	http    *http.Client
}

// WriteCounter count bytes for a file download
type WriteCounter struct {
	Total uint64
}

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
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		retry: 2,
	}

	for _, opt := range opts {
		client = opt(client)
	}

	return client
}

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

		if cookie := client.auth.Cookie(); cookie != "" {
			req.Header.Set("Cookie", cookie)
		}
	}
	req.Header.Set("User-Agent", UserAgent)

	return client.http.Do(req)
}

// URL creates a connector to specified endpoint. It is either absolute
// URL or relative path to base url
func (client *tClient) URL(templatePath string, args ...interface{}) CURL {
	target := fmt.Sprintf(templatePath, args...)
	if target[0] == '/' {
		target = client.baseURL + target
	}

	return &tCURL{
		client:  client,
		url:     target,
		header:  http.Header{},
		payload: bytes.NewBuffer(nil),
	}
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

// Query defines URI parameters of the request
func (curl *tCURL) Query(data interface{}) CURL {
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

	var params map[string]interface{}
	if err = json.Unmarshal(bin, &params); err != nil {
		return nil, err
	}

	var values url.Values = make(map[string][]string)
	for key, param := range params {
		var val string
		switch v := param.(type) {
		case int:
			val = strconv.Itoa(v)
		case float64:
			val = fmt.Sprintf("%g", v)
		case string:
			val = v
		case bool:
			val = strconv.FormatBool(v)
		default:
			return nil, fmt.Errorf("wrong format: %T", v)
		}
		values[key] = []string{val}
	}

	return values, nil
}

// Header defines request header
func (curl *tCURL) Header(head, value string) CURL {
	curl.header.Add(head, value)
	return curl
}

// Status payload from target URL and discards it.
func (curl *tCURL) Status(status ...int) (http.Header, error) {
	curl.method = http.MethodGet
	return curl.status()
}

func (curl *tCURL) status(status ...int) (http.Header, error) {
	curl = curl.unsafeIO()
	if curl.fail != nil {
		return nil, curl.fail
	}

	defer curl.output.Body.Close()
	body, err := io.ReadAll(curl.output.Body)
	if err != nil {
		return nil, err
	}

	curl.fail = curl.isSuccess(body, status...)
	return curl.unWrap()
}

func (curl *tCURL) isSuccess(body []byte, status ...int) error {
	if len(status) == 1 {
		if curl.output.StatusCode != status[0] {
			return ErrorFromResponse(curl.output, body)
		}
	} else {
		if curl.output.StatusCode >= http.StatusBadRequest {
			return ErrorFromResponse(curl.output, body)
		}
	}

	return nil
}

// Write increments the counter by the size of the bytes written into it
func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)

	wc.Total += uint64(n)
	wc.printProgress()

	return n, nil
}

func (wc *WriteCounter) printProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 50))

	fmt.Printf("\rDownloading... %s complete", fBytes(wc.Total))
}

// Func insp: https://github.com/dustin/go-humanize
func fBytes(s uint64) string {
	sizes := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1000, sizes)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(math.Log(float64(s)) / math.Log(base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f %s"
	if val < 10 {
		f = "%.1f %s"
	}

	return fmt.Sprintf(f, val, suffix)
}

func writeToFile(filename string, resp *http.Response) error {
	out, err := os.Create(filename + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	err = os.Rename(filename+".tmp", filename)
	if err != nil {
		return err
	}

	return nil
}

// Download dowmload file via http from endpoint
func (curl *tCURL) Download(filename string) error {
	curl.method = http.MethodGet

	req, err := http.NewRequest(curl.method, curl.url, curl.payload)
	if err != nil {
		return err
	}

	resp, err := curl.client.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal(resp.Status)
	}

	err = writeToFile(filename, resp)
	if err != nil {
		return err
	}

	return nil
}

// Get fetches content from endpoint
func (curl *tCURL) Get(in interface{}) (http.Header, error) {
	curl.method = http.MethodGet
	return curl.recv(in)
}

// Put sends content to endpoint
func (curl *tCURL) Put(eg interface{}, in ...interface{}) (http.Header, error) {
	curl.method = http.MethodPut
	curl.send(eg)

	if len(in) > 0 {
		return curl.recv(in[0])
	}

	return curl.status()
}

// Post sends content to endpoint
func (curl *tCURL) Post(eg interface{}, in ...interface{}) (http.Header, error) {
	curl.method = http.MethodPost

	if eg != nil {
		curl.send(eg)
	}

	if len(in) > 0 {
		return curl.recv(in[0])
	}

	return curl.status()
}

// Delete removes content behind url
func (curl *tCURL) Delete(in ...interface{}) (http.Header, error) {
	curl.method = http.MethodDelete

	if len(in) > 0 {
		return curl.recv(in[0])
	}

	return curl.status()
}

// send payload to destination URL.
func (curl *tCURL) send(data interface{}) CURL {
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
	curl.header.Set("Content-Type", "application/json")

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

// recv payload from target URL.
func (curl *tCURL) recv(data interface{}) (http.Header, error) {
	curl = curl.unsafeIO()

	if curl.fail != nil {
		return nil, curl.fail
	}

	defer curl.output.Body.Close()
	body, err := io.ReadAll(curl.output.Body)
	if err != nil {
		return nil, err
	}

	curl.fail = curl.isSuccess(body)

	return curl.unWrapWithData(body, data)
}

// Fetch receive raw content
func (curl *tCURL) Fetch() ([]byte, error) {
	curl.method = http.MethodGet
	curl = curl.unsafeIO()

	if curl.fail != nil {
		return nil, curl.fail
	}

	defer curl.output.Body.Close()
	body, err := io.ReadAll(curl.output.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (curl *tCURL) unsafeIO() *tCURL {
	if curl.fail != nil {
		return curl
	}

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

// unWrap tCURL object to results
func (curl *tCURL) unWrap() (http.Header, error) {
	if curl.fail != nil {
		return nil, curl.fail
	}

	return curl.output.Header, nil
}

// unWrap tCURL object to results and decodes JSON
func (curl *tCURL) unWrapWithData(body []byte, data interface{}) (http.Header, error) {
	if curl.fail != nil {
		return nil, curl.fail
	}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return curl.output.Header, nil
}
