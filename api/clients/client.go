//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package clients

import (
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Client is a role-store client instance.
type Client struct {
	api restapi.Connector
}

// New creates a new user-store client instance
func New(api restapi.Connector) *Client {
	return &Client{api: api}
}

// CreateTrustedClient registers new client to PrivX
func (clients *Client) CreateTrustedClient(client TrustedClient) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	_, err := clients.api.
		URL("/local-user-store/api/v1/trusted-clients").
		Post(client, &id)

	return id.ID, err
}

// TrustedClients fetches all known trusted clients
func (clients *Client) TrustedClients() ([]TrustedClient, error) {
	var seq struct {
		Items []TrustedClient
	}

	_, err := clients.api.
		URL("/local-user-store/api/v1/trusted-clients").
		Get(&seq)

	return seq.Items, err
}

// TrustedClient returns details about the client
func (clients *Client) TrustedClient(id string) (*TrustedClient, error) {
	client := new(TrustedClient)

	_, err := clients.api.
		URL("/local-user-store/api/v1/trusted-clients/%s", id).
		Get(client)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// DeleteTrustedClient removes the client
func (clients *Client) DeleteTrustedClient(id string) error {
	_, err := clients.api.
		URL("/local-user-store/api/v1/trusted-clients/%s", id).
		Delete()

	return err
}
