//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package api

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/SSHcom/privx-sdk-go/oauth"
)

type Client struct {
	Auth     *oauth.Client
	endpoint string
	http     *http.Client
}

func NewClient(auth *oauth.Client, endpoint string, cert *x509.Certificate,
	verbose bool) (*Client, error) {

	tlsConfig := &tls.Config{}
	if cert != nil {
		pool := x509.NewCertPool()
		pool.AddCert(cert)
		tlsConfig.RootCAs = pool
	}
	return &Client{
		Auth:     auth,
		endpoint: endpoint,
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
	}, nil
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
