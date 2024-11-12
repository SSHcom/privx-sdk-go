//
// Copyright (c) 2023 SSH Communications Security Inc.
//
// All rights reserved.
//

package dbproxy

import (
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// DbProxy is a db proxy client instance.
type DbProxy struct {
	api restapi.Connector
}

// New db proxy client constructor.
func New(api restapi.Connector) *DbProxy {
	return &DbProxy{api: api}
}

// MARK: Status
// Status get db proxy microservice status.
func (c *DbProxy) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/db-proxy/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Config
// GetDbProxyConfig get db proxy configuration.
func (c *DbProxy) GetDbProxyConfig() (*DBProxyAPIConf, error) {
	config := &DBProxyAPIConf{}

	_, err := c.api.
		URL("/db-proxy/api/v1/conf").
		Get(config)

	return config, err
}
