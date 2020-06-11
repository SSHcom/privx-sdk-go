//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

import (
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Client is a authorizer client instance.
type Client struct {
	api restapi.Connector
}

// New creates a new authorizer client instance
func New(api restapi.Connector) *Client {
	return &Client{api: api}
}

// RootCertificate fetches
func (auth *Client) RootCertificate() (ca []CA, err error) {
	_, err = auth.api.
		URL("/authorizer/api/v1/cas").
		Get(&ca)

	return
}
