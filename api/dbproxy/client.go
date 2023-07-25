//
// Copyright (c) 2023 SSH Communications Security Inc.
//
// All rights reserved.
//

package dbproxy

import (
	"github.com/SSHcom/privx-sdk-go/common"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// DbProxy is a db proxy instance.
type DbProxy struct {
	api restapi.Connector
}

// New creates a new db proxy client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *DbProxy {
	return &DbProxy{api: api}
}

// DbProxyStatus get microservice status
func (store *DbProxy) DbProxyStatus() (*common.ServiceStatus, error) {
	status := &common.ServiceStatus{}

	_, err := store.api.
		URL("/db-proxy/api/v1/status").
		Get(status)

	return status, err
}

// DbProxyConf get db proxy configuration
func (store *DbProxy) DbProxyConf() (*DBProxyConf, error) {
	config := &DBProxyConf{}

	_, err := store.api.
		URL("/db-proxy/api/v1/conf").
		Get(config)

	return config, err
}
