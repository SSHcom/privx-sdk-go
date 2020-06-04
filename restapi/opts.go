//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package restapi

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

// Option is configuration applied to the client
type Option func(*tClient) *tClient

// Endpoint defines a target PrivX API endpoint
func Endpoint(endpoint string) Option {
	return func(client *tClient) *tClient {
		client.endpoint = endpoint
		return client
	}
}

// AccessToken setup access token provider for api
func AccessToken(auth Provider) Option {
	return func(client *tClient) *tClient {
		client.auth = auth
		return client
	}
}

// X509 setup trust certificate
func X509(cert *x509.Certificate) Option {
	return func(client *tClient) *tClient {
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

// NoRedirect disables redirect
func NoRedirect() Option {
	return func(client *tClient) *tClient {
		client.http.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
		return client
	}
}

// Verbose enables debug-level logging
func Verbose() Option {
	return func(client *tClient) *tClient {
		client.verbose = true
		return client
	}
}

// Retry HTTP I/O multiple times before failure
func Retry(n int) Option {
	return func(client *tClient) *tClient {
		client.retry = n
		return client
	}
}
