//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package restapi

import "net/http"

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
}

// Authorizer provides access token for REST API client
type Authorizer interface {
	AccessToken() (string, error)
}

const (
	// UserAgent specifies the HTTP user-agent string for the SDK
	// clients.
	UserAgent = "privx-sdk-go"
)
