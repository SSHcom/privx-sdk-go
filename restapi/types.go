//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package restapi

import "net/http"

// Connector is HTTP connector for api
type Connector interface {
	URL(string, string) CURL
	Get(string, ...interface{}) CURL
	Put(string, ...interface{}) CURL
	Post(string, ...interface{}) CURL
}

// CURL is HTTP request
type CURL interface {
	// Params defines query parameters
	Params(data interface{}) CURL
	// With defines HTTP header
	With(string, string) CURL
	// Send payload to HTTP endpoint
	Send(data interface{}) CURL
	// Recv payload(s) from HTTP endpoint
	Recv(data interface{}) (http.Header, error)
	RecvStatus(...int) (http.Header, error)
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
