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

// Provider define additional parameters to HTTP request
type Provider interface {
	Headers() (map[string]string, error)
}
