//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package vault

import (
	"encoding/json"
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Vault is a vault client instance.
type Vault struct {
	api restapi.Connector
}

// New vault client constructor.
func New(api restapi.Connector) *Vault {
	return &Vault{api: api}
}

// MARK: Status
// Status get role store microservice status.
func (c *Vault) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/vault/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Metadata
// GetSecretsMetadata get secrets metadata.
func (c *Vault) GetSecretsMetadata(name string) (*Secret, error) {
	metadata := &Secret{}

	_, err := c.api.
		URL("/vault/api/v1/metadata/secrets/%s", name).
		Get(&metadata)

	return metadata, err
}

// GetUsersSecretsMetadata get users secrets metadata.
func (c *Vault) GetUsersSecretsMetadata(userID, name string) (*Secret, error) {
	metadata := &Secret{}

	_, err := c.api.
		URL("/vault/api/v1/user/%s/metadata/secrets/%s", userID, name).
		Get(&metadata)

	return metadata, err
}

// MARK: Personal Secrets
// GetUserSecrets get user secrets.
func (c *Vault) GetUserSecrets(userID string, opts ...filters.Option) (*response.ResultSet[Secret], error) {
	secrets := &response.ResultSet[Secret]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/vault/api/v1/user/%s/secrets", userID).
		Query(params).
		Get(&secrets)

	return secrets, err
}

// CreateUserSecret create user secret.
func (c *Vault) CreateUserSecret(userID string, secret *SecretRequest) (SecretCreate, error) {
	created := SecretCreate{}

	_, err := c.api.
		URL("/vault/api/v1/user/%s/secrets", userID).
		Post(&secret, &created)

	return created, err
}

// UserSecret get user secret by secret name.
func (c *Vault) GetUserSecret(userID, secretName string) (*Secret, error) {
	secret := &Secret{}

	_, err := c.api.
		URL("/vault/api/v1/user/%s/secrets/%s", userID, secretName).
		Get(&secret)

	return secret, err
}

// UpdateUserSecret update user secret.
func (c *Vault) UpdateUserSecret(userID, secretName string, secret *SecretRequest) error {
	_, err := c.api.
		URL("/vault/api/v1/user/%s/secrets/%s", userID, secretName).
		Put(&secret)

	return err
}

// DeleteUserSecret delete user secret.
func (c *Vault) DeleteUserSecret(userID, secretName string) error {
	_, err := c.api.
		URL("/vault/api/v1/user/%s/secrets/%s", userID, secretName).
		Delete()

	return err
}

// MARK: Schemas
// GetSchemas get the defined vault schemas.
func (c *Vault) GetSchemas() (*json.RawMessage, error) {
	schemas := &json.RawMessage{}

	_, err := c.api.
		URL("/vault/api/v1/schemas").
		Get(&schemas)

	return schemas, err
}

// MARK: Secrets
// GetSecrets get secrets.
func (c *Vault) GetSecrets(opts ...filters.Option) (*response.ResultSet[Secret], error) {
	secrets := &response.ResultSet[Secret]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/vault/api/v1/secrets").
		Query(params).
		Get(&secrets)

	return secrets, err
}

// CreateSecret create secret.
func (c *Vault) CreateSecret(secret *SecretRequest) (SecretCreate, error) {
	created := SecretCreate{}

	_, err := c.api.
		URL("/vault/api/v1/secrets").
		Post(&secret, &created)

	return created, err
}

// GetSecret get secret by secret name.
func (c *Vault) GetSecret(secretName string) (*Secret, error) {
	secret := &Secret{}

	_, err := c.api.
		URL("/vault/api/v1/secrets/%s", secretName).
		Get(&secret)

	return secret, err
}

// UpdateSecret update secret.
func (c *Vault) UpdateSecret(secretName string, secret *SecretRequest) error {
	_, err := c.api.
		URL("/vault/api/v1/secrets/%s", secretName).
		Put(&secret)

	return err
}

// DeleteSecret delete secret.
func (c *Vault) DeleteSecret(secretName string) error {
	_, err := c.api.
		URL("/vault/api/v1/secrets/%s", secretName).
		Delete()

	return err
}

// SearchSecrets search secrets.
func (c *Vault) SearchSecrets(search SecretSearch, opts ...filters.Option) (*response.ResultSet[Secret], error) {
	secrets := &response.ResultSet[Secret]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/vault/api/v1/search/secrets").
		Query(params).
		Post(&search, &secrets)

	return secrets, err
}
