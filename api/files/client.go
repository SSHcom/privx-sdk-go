//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package files

import (
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Client is a role-store client instance.
type Client struct {
	api restapi.Connector
}

// New creates a new client instance to fetch config files from PrivX
func New(api restapi.Connector) *Client {
	return &Client{api: api}
}

// ConfigExtender fetches configugartion file
func (files Client) ConfigExtender(id string) ([]byte, error) {
	var session struct {
		ID string `json:"session_id"`
	}

	_, err := files.api.
		URL("/authorizer/api/v1/extender/conf/%s", id).
		Post("", &session)

	if err != nil {
		return nil, err
	}

	return files.api.
		URL("/authorizer/api/v1/extender/conf/%s/%s", id, session.ID).
		Fetch()
}
