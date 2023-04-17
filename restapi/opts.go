//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package restapi

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

// Option is configuration applied to the client
type Option func(*tClient) *tClient

// BaseURL defines a target PrivX server and possible path prefix
func BaseURL(endpoint string) Option {
	return func(client *tClient) *tClient {
		if endpoint != "" {
			client.baseURL = endpoint
		}
		return client
	}
}

// Auth setup access token provider for api
func Auth(auth Authorizer) Option {
	return func(client *tClient) *tClient {
		client.auth = auth
		return client
	}
}

// TrustAnchor setups X509 certificates to trust TLS connections
func TrustAnchor(cert *x509.Certificate) Option {
	return func(client *tClient) *tClient {
		if cert != nil {
			tlsConfig := &tls.Config{}
			pool := x509.NewCertPool()
			pool.AddCert(cert)
			tlsConfig.RootCAs = pool
			client.http.Transport.(*http.Transport).TLSClientConfig = tlsConfig
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

// UseConfigFile setup rest client from toml file
func UseConfigFile(path string) Option {
	return func(client *tClient) *tClient {
		type config struct {
			BaseURL     string       `toml:"base_url"`
			Certificate *Certificate `toml:"api_ca_crt"`
		}
		var file struct {
			API config
		}

		if path == "" {
			return client
		}

		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}

		if err = toml.Unmarshal(data, &file); err != nil {
			panic(err)
		}

		client = BaseURL(file.API.BaseURL)(client)
		if file.API.Certificate != nil {
			client = TrustAnchor(file.API.Certificate.X509)(client)
		}
		return client
	}
}

// UseEnvironment setups rest client using environment variables
func UseEnvironment() Option {
	return func(client *tClient) *tClient {
		if url, ok := os.LookupEnv("PRIVX_API_BASE_URL"); ok {
			client = BaseURL(url)(client)
		}

		return client
	}
}
