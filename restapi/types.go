//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package restapi

// Connector is HTTP connector for api
type Connector interface {
	URL(string, string) CURL
	Get(string, ...interface{}) CURL
	Put(string, ...interface{}) CURL
}

// CURL is HTTP request
type CURL interface {
	Send(data interface{}) CURL
	Recv(data interface{}) error
	RecvStatus() error
}

// Provider define additional parameters to HTTP request
type Provider interface {
	Headers() (map[string]string, error)
}
