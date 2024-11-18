//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package userstore

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// UserStore is a local user store client instance.
type UserStore struct {
	api restapi.Connector
}

// New local user store client constructor.
func New(api restapi.Connector) *UserStore {
	return &UserStore{api: api}
}

// MARK: Status
// Status get local user store microservice status.
func (c *UserStore) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/local-user-store/api/v1/status").
		Get(status)

	return status, err
}

// MARK: API Client
// GetAPIClients get registered api clients.
func (c *UserStore) GetAPIClients(opts ...filters.Option) (*response.ResultSet[APIClient], error) {
	clients := &response.ResultSet[APIClient]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/local-user-store/api/v1/api-clients").
		Query(params).
		Get(&clients)

	return clients, err
}

// CreateAPIClient create api client.
func (c *UserStore) CreateAPIClient(client *APIClientCreate) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/local-user-store/api/v1/api-clients").
		Post(&client, &identifier)

	return identifier, err
}

// SearchAPIClients search api clients.
func (c *UserStore) SearchAPIClients(search *APIClientSearch) (*response.ResultSet[APIClient], error) {
	clients := &response.ResultSet[APIClient]{}

	_, err := c.api.
		URL("/local-user-store/api/v1/api-clients/search").
		Post(search, &clients)

	return clients, err
}

// GetAPIClient get api client by id.
func (c *UserStore) GetAPIClient(clientID string) (*APIClient, error) {
	client := &APIClient{}

	_, err := c.api.
		URL("/local-user-store/api/v1/api-clients/%s", clientID).
		Get(client)

	return client, err
}

// UpdateAPIClient update api client.
func (c *UserStore) UpdateAPIClient(clientID string, client *APIClient) error {
	_, err := c.api.
		URL("/local-user-store/api/v1/api-clients/%s", clientID).
		Put(&client)

	return err
}

// DeleteAPIClient delete api client.
func (c *UserStore) DeleteAPIClient(clientID string) error {
	_, err := c.api.
		URL("/local-user-store/api/v1/api-clients/%s", clientID).
		Delete()

	return err
}

// MARK: Extender Clients
// GetExtenderClients get extender clients.
func (c *UserStore) GetExtenderClients() (*response.ResultSet[ExtenderClient], error) {
	clients := &response.ResultSet[ExtenderClient]{}

	_, err := c.api.
		URL("/local-user-store/api/v1/extender-clients").
		Get(&clients)

	return clients, err
}

// MARK: Users
// GetUsers get local users.
func (c *UserStore) GetUsers(opts ...filters.Option) (*response.ResultSet[LocalUser], error) {
	users := &response.ResultSet[LocalUser]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/local-user-store/api/v1/users").
		Query(params).
		Get(&users)

	return users, err
}

// CreateUser create local user.
func (c *UserStore) CreateUser(user *LocalUser) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/local-user-store/api/v1/users").
		Post(&user, &identifier)

	return identifier, err
}

// GetUser get local user by id.
func (c *UserStore) GetUser(userID string) (*LocalUser, error) {
	user := &LocalUser{}

	_, err := c.api.
		URL("/local-user-store/api/v1/users/%s", userID).
		Get(&user)

	return user, err
}

// UpdateUser update local user.
func (c *UserStore) UpdateUser(userID string, user *LocalUser) error {
	_, err := c.api.
		URL("/local-user-store/api/v1/users/%s", userID).
		Put(user)

	return err
}

// DeleteUser delete local user.
func (c *UserStore) DeleteUser(userID string) error {
	_, err := c.api.
		URL("/local-user-store/api/v1/users/%s", userID).
		Delete()

	return err
}

// UpdateUserPassword update local user password.
func (c *UserStore) UpdateUserPassword(userID string, password LocalUserPassword) error {
	_, err := c.api.
		URL("/local-user-store/api/v1/users/%s/password", userID).
		Put(password)

	return err
}

// GetUserTags get local user tags.
func (c *UserStore) GetUserTags(opts ...filters.Option) (*response.ResultSet[string], error) {
	tags := &response.ResultSet[string]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/local-user-store/api/v1/users/tags").
		Query(params).
		Get(&tags)

	return tags, err
}

// MARK: Trusted Clients
// GetTrustedClients get trusted clients.
func (c *UserStore) GetTrustedClients() (*response.ResultSet[TrustedClient], error) {
	clients := &response.ResultSet[TrustedClient]{}

	_, err := c.api.
		URL("/local-user-store/api/v1/trusted-clients").
		Get(&clients)

	return clients, err
}

// CreateTrustedClient created trusted client.
func (c *UserStore) CreateTrustedClient(client *TrustedClient) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/local-user-store/api/v1/trusted-clients").
		Post(&client, &identifier)

	return identifier, err
}

// GetTrustedClient get trusted client by id.
func (c *UserStore) GetTrustedClient(clientID string) (*TrustedClient, error) {
	client := &TrustedClient{}

	_, err := c.api.
		URL("/local-user-store/api/v1/trusted-clients/%s", clientID).
		Get(&client)

	return client, err
}

// UpdateTrustedClient update trusted client.
func (c *UserStore) UpdateTrustedClient(clientID string, client *TrustedClient) error {
	_, err := c.api.
		URL("/local-user-store/api/v1/trusted-clients/%s", clientID).
		Put(client)

	return err
}

// DeleteTrustedClient delete trusted client.
func (c *UserStore) DeleteTrustedClient(clientID string) error {
	_, err := c.api.
		URL("/local-user-store/api/v1/trusted-clients/%s", clientID).
		Delete()

	return err
}
