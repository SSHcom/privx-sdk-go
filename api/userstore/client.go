//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package userstore

import (
	"net/url"

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

type tagsResult struct {
	Count int      `json:"count"`
	Items []string `json:"items"`
}

type clientsResult struct {
	Count int             `json:"count"`
	Items []TrustedClient `json:"items"`
}

// New creates a new user-store client instance
func New(api restapi.Connector) *UserStore {
	return &UserStore{api: api}
}

// LocalUserTags returns local user tags
func (store *UserStore) LocalUserTags(offset, limit, sortdir, query string) ([]string, error) {
	result := tagsResult{}
	filters := FilterUser{
		Params: Params{
			Offset:  offset,
			Limit:   limit,
			Sortdir: sortdir,
			Query:   query,
		},
	}

	_, err := store.api.
		URL("/local-user-store/api/v1/users/tags").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// UpdateLocalUserPassword update existing local user password
func (store *UserStore) UpdateLocalUserPassword(id string, password *Password) error {
	_, err := store.api.
		URL("/local-user-store/api/v1/users/%s/password", url.PathEscape(id)).
		Put(password)

	return err
}

// UpdateLocalUser update existing local user
func (store *UserStore) UpdateLocalUser(id string, localUser *LocalUser) error {
	_, err := store.api.
		URL("/local-user-store/api/v1/users/%s", url.PathEscape(id)).
		Put(localUser)

	return err
}

// DeleteLocalUser delete a local user
func (store *UserStore) DeleteLocalUser(id string) error {
	_, err := store.api.
		URL("/local-user-store/api/v1/users/%s", id).
		Delete()

	return err
}

// LocalUser returns details about the local user
func (store *UserStore) LocalUser(id string) (user *LocalUser, err error) {
	user = new(LocalUser)

	_, err = store.api.
		URL("/local-user-store/api/v1/users/%s", url.PathEscape(id)).
		Get(user)

	return user, err
}

// CreateLocalUser create a new local PrivX user
func (store *UserStore) CreateLocalUser(newUser LocalUser) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/local-user-store/api/v1/users").
		Post(newUser, &id)

	return id.ID, err
}

// LocalUsers returns user details from all known local users
func (store *UserStore) LocalUsers(offset, limit, userID, username string) ([]LocalUser, error) {
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

// UpdateTrustedClient update existing trusted client
func (store *UserStore) UpdateTrustedClient(id string, client *TrustedClient) error {
	_, err := store.api.
		URL("/local-user-store/api/v1/trusted-clients/%s", url.PathEscape(id)).
		Put(client)

	return err
}

// DeleteTrustedClient removes the client
func (store *UserStore) DeleteTrustedClient(id string) error {
	_, err := store.api.
		URL("/local-user-store/api/v1/trusted-clients/%s", id).
		Delete()

	return err
}

// ExtenderClients returns a list of extender client names and types
func (store *UserStore) ExtenderClients() ([]TrustedClient, error) {
	result := clientsResult{}

	_, err := store.api.
		URL("/local-user-store/api/v1/extender-clients").
		Get(&result)

	return result.Items, err
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
