//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package restapi

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
)

// Connector is HTTP connector for api
type Connector interface {
	// URL creates a request/response session
	URL(string, ...interface{}) CURL
}

// CURL is HTTP request
type CURL interface {
	// Query defines URI parameters of the request
	Query(interface{}) CURL
	// Header defines request header
	Header(string, string) CURL
	// Status evalutes the request
	Status(...int) (http.Header, error)
	Get(interface{}) (http.Header, error)
	Put(interface{}, ...interface{}) (http.Header, error)
	Post(interface{}, ...interface{}) (http.Header, error)
	Delete(...interface{}) (http.Header, error)
	Fetch() ([]byte, error)
	Download(string) error
}

// Authorizer provides access token for REST API client
type Authorizer interface {
	AccessToken() (string, error)
	Cookie() string
}

const (
	// UserAgent specifies the HTTP user-agent string for the SDK
	// clients.
	UserAgent = "privx-sdk-go"
)

// Certificate specifies a trusted CA certificate for the REST endpoint.
type Certificate struct {
	X509 *x509.Certificate
}

// MarshalText implements the encoding.TextMarshaler interface.
func (cert Certificate) MarshalText() (text []byte, err error) {
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.X509.Raw,
	}
	return pem.EncodeToMemory(block), nil
}

// UnmarshalText unmarshals certificate from a configuration file PEM
// block.
func (cert *Certificate) UnmarshalText(text []byte) error {
	block, _ := pem.Decode(text)
	if block == nil {
		return fmt.Errorf("could not decode certificate PEM data")
	}
	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}
	cert.X509 = c
	return nil
}
