//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package userstore

import (
	"github.com/SSHcom/privx-sdk-go/api/rolestore"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// UserStore is a role-store client instance.
type UserStore struct {
	api restapi.Connector
}

type usersResult struct {
	Count int         `json:"count"`
	Items []LocalUser `json:"items"`
}

// New creates a new user-store client instance
func New(api restapi.Connector) *UserStore {
	return &UserStore{api: api}
}

// GetLocalUsers get a local user with details
func (store *UserStore) GetLocalUsers(offset, limit, userID, username string) ([]LocalUser, error) {
	result := usersResult{}
	filters := FilterUser{
		Params: Params{
			Offset: offset,
			Limit:  limit,
		},
		UserID:   userID,
		Username: username,
	}

	_, err := store.api.
		URL("/local-user-store/api/v1/users").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// CreateTrustedClient registers new client to PrivX
func (store *UserStore) CreateTrustedClient(client TrustedClient) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/local-user-store/api/v1/trusted-clients").
		Post(client, &id)

	return id.ID, err
}

// TrustedClients fetches all known trusted clients
func (store *UserStore) TrustedClients() ([]TrustedClient, error) {
	var seq struct {
		Items []TrustedClient
	}

	_, err := store.api.
		URL("/local-user-store/api/v1/trusted-clients").
		Get(&seq)

	return seq.Items, err
}

// TrustedClient returns details about the client
func (store *UserStore) TrustedClient(id string) (*TrustedClient, error) {
	client := new(TrustedClient)

	_, err := store.api.
		URL("/local-user-store/api/v1/trusted-clients/%s", id).
		Get(client)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// DeleteTrustedClient removes the client
func (store *UserStore) DeleteTrustedClient(id string) error {
	_, err := store.api.
		URL("/local-user-store/api/v1/trusted-clients/%s", id).
		Delete()

	return err
}

// APIClients returns list of all registered api clients
func (store *UserStore) APIClients() ([]APIClient, error) {
	var seq struct {
		Items []APIClient
	}

	_, err := store.api.
		URL("/local-user-store/api/v1/api-clients").
		Get(&seq)

	return seq.Items, err
}

// APIClient returns details about API client
func (store *UserStore) APIClient(id string) (*APIClient, error) {
	client := new(APIClient)

	_, err := store.api.
		URL("/local-user-store/api/v1/api-clients/%s", id).
		Get(client)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// CreateAPIClient creates new API client
func (store *UserStore) CreateAPIClient(name string, roles []string) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	req := struct {
		Name  string              `json:"name"`
		Roles []rolestore.RoleRef `json:"roles"`
	}{Name: name, Roles: []rolestore.RoleRef{}}

	for _, role := range roles {
		req.Roles = append(req.Roles, rolestore.RoleRef{ID: role})
	}

	_, err := store.api.
		URL("/local-user-store/api/v1/api-clients").
		Post(req, &id)

	return id.ID, err
}

// DeleteAPIClient removes existing API client
func (store *UserStore) DeleteAPIClient(id string) error {
	_, err := store.api.
		URL("/local-user-store/api/v1/api-clients/%s", id).
		Delete()

	return err
}
