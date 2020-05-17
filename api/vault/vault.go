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

// Bag contains secret data. The secret data is a JSON object and it
// can have nested values and objects.
type Bag map[string]interface{}

// Client is a Vault client instance.
type Client struct {
	api api.Connector
}

// NewClient creates a new Vault client instance, using the argument
// SDK API client.
func NewClient(api api.Connector) (*Client, error) {
	return &Client{
		api: api,
	}, nil
}

// Get gets the content of the argument secret.
func (vault *Client) Get(name string) (bag Bag, err error) {
	err = vault.api.
		Get("/vault/api/v1/secrets/%s", url.PathEscape(name)).
		Recv(&bag)

	return
}
