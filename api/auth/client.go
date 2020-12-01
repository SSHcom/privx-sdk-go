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

// AccessGroups lists all access group
func (auth *Client) AccessGroups() ([]AccessGroup, error) {
	var result struct {
		Count int           `json:"count"`
		Items []AccessGroup `json:"items"`
	}

	_, err := auth.api.
		URL("/authorizer/api/v1/accessgroups").
		Get(&result)

	return result.Items, err
}
