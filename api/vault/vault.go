//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package vault

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api"
)

type Bag map[string]interface{}

type Client struct {
	api api.Connector
}

func NewClient(api api.Connector) (*Client, error) {
	return &Client{
		api: api,
	}, nil
}

func (vault *Client) Get(name string) (bag Bag, err error) {
	err = vault.api.
		Get("/vault/api/v1/secrets/%s", url.PathEscape(name)).
		Recv(&bag)

	return
}
