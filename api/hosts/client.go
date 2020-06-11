//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package hosts

import (
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Client is a role-store client instance.
type Client struct {
	api restapi.Connector
}

// New creates a new host-store client instance
// See http://apispecs.ssh.com/#swagger-ui-4 for details about api
func New(api restapi.Connector) *Client {
	return &Client{api: api}
}

// Register target host to PrivX
func (hosts Client) Register(host Host) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	_, err := hosts.api.
		URL("/host-store/api/v1/hosts").
		Post(&host, &id)

	return id.ID, err
}
