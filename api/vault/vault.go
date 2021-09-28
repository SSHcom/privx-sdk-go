//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package vault

import (
	"encoding/json"
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/rolestore"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Vault is client instance.
type Vault struct {
	api restapi.Connector
}

type secretResult struct {
	Count int      `json:"count"`
	Items []Secret `json:"items"`
}

type secretID struct {
	ownerID string
	name    string
}

// tVaultReq t vault request definition
type tVaultReq struct {
	Name       string              `json:"name,omitempty"`
	Data       interface{}         `json:"data"`
	AllowRead  []rolestore.RoleRef `json:"read_roles,omitempty"`
	AllowWrite []rolestore.RoleRef `json:"write_roles,omitempty"`
	OwnerID    string              `json:"owner_id,omitempty"`
}

// New creates a new Vault client instance, using the argument
// SDK API client.
func New(api restapi.Connector) *Vault {
	return &Vault{api: api}
}

// CreateSecret create new secret to PrivX Vault
func (vault *Vault) CreateSecret(
	name string,
	allowReadBy []string,
	allowWriteBy []string,
	secret interface{},
) error {
	req := vault.mkVaultReq(allowReadBy, allowWriteBy, secret)
	req.Name = name

	_, err := vault.api.
		URL("/vault/api/v1/secrets").
		Post(req)

	return err
}

//CreateUserSecret creates a user secret
func (vault *Vault) CreateUserSecret(
	secretID secretID,
	allowReadBy []string,
	allowWriteBy []string,
	secret interface{},
) error {
	req := vault.mkVaultReq(allowReadBy, allowWriteBy, secret)
	req.Name = secretID.name
	req.OwnerID = url.PathEscape(secretID.ownerID)

	_, err := vault.api.
		URL("/vault/api/v1/user/%s/secrets", req.OwnerID).
		Post(req)

	return err
}

// Secrets returns secrets client has access to
func (vault *Vault) Secrets(offset, limit int) ([]Secret, error) {
	result := secretResult{}
	filters := Params{
		Offset: offset,
		Limit:  limit,
	}

	_, err := vault.api.
		URL("/vault/api/v1/secrets").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// UserSecrets returns user secrets client has access to
func (vault *Vault) UserSecrets(secretID secretID, offset, limit int) ([]Secret, error) {
	result := secretResult{}
	filters := Params{
		Offset: offset,
		Limit:  limit,
	}

	_, err := vault.api.
		URL("/vault/api/v1/user/%s/secrets", url.PathEscape(secretID.ownerID)).
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// Secret gets the content of the argument secret.
func (vault *Vault) Secret(name string) (*Secret, error) {
	bag := &Secret{}
	_, err := vault.api.
		URL("/vault/api/v1/secrets/%s", url.PathEscape(name)).
		Get(&bag)

	return bag, err
}

// UserSecret gets the content of the argument user secret.
func (vault *Vault) UserSecret(secretID secretID) (*Secret, error) {
	bag := &Secret{}
	_, err := vault.api.
		URL("/vault/api/v1/user/%s/secrets/%s", url.PathEscape(secretID.ownerID), url.PathEscape(secretID.name)).
		Get(&bag)

	return bag, err
}

// UpdateSecret existing secret at PrivX Vault
func (vault *Vault) UpdateSecret(
	name string,
	allowReadTo []string,
	allowWriteTo []string,
	secret interface{},
) error {
	req := vault.mkVaultReq(allowReadTo, allowWriteTo, secret)

	_, err := vault.api.
		URL("/vault/api/v1/secrets/%s", name).
		Put(req)

	return err
}

// UpdateUserSecret existing secret at PrivX Vault
func (vault *Vault) UpdateUserSecret(
	secretID secretID,
	allowReadTo []string,
	allowWriteTo []string,
	secret interface{},
) error {
	req := vault.mkVaultReq(allowReadTo, allowWriteTo, secret)
	req.Name = url.PathEscape(secretID.name)
	req.OwnerID = url.PathEscape(secretID.ownerID)
	_, err := vault.api.
		URL("/vault/api/v1/user/%s/secrets/%s", req.OwnerID, req.Name).
		Put(req)

	return err
}

// DeleteSecret delete existing secret from PrivX vault
func (vault *Vault) DeleteSecret(name string) error {
	_, err := vault.api.
		URL("/vault/api/v1/secrets/%s", name).
		Delete()

	return err
}

// DeleteSecret delete existing secret from PrivX vault
func (vault *Vault) DeleteUserSecret(secretID secretID) error {
	ownerID := url.PathEscape(secretID.ownerID)
	name := url.PathEscape(secretID.name)

	_, err := vault.api.
		URL("/vault/api/v1/user/%s/secrets/%s", ownerID, name).
		Delete()

	return err
}

// SecretMetadata returns secret metadata
func (vault *Vault) SecretMetadata(name string) (*Secret, error) {
	metadata := &Secret{}

	_, err := vault.api.
		URL("/vault/api/v1/metadata/secrets/%s", url.PathEscape(name)).
		Get(&metadata)

	return metadata, err
}

// SecretMetadata returns secret metadata
func (vault *Vault) UserSecretMetadata(secretID secretID) (*Secret, error) {
	metadata := &Secret{}
	ownerID := url.PathEscape(secretID.ownerID)
	name := url.PathEscape(secretID.name)

	_, err := vault.api.
		URL("/vault/api/v1/user/%s/metadata/secrets/%s", ownerID, name).
		Get(&metadata)

	return metadata, err
}

// SearchSecrets search for existing secrets
func (vault *Vault) SearchSecrets(offset, limit int, keywords, sortkey, sortdir string) ([]Secret, error) {
	result := secretResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := vault.api.
		URL("/vault/api/v1/search/secrets").
		Query(&filters).
		Post(map[string]string{
			"keywords": keywords,
		}, &result)

	return result.Items, err
}

// VaultSchemas returns the defined schemas
func (vault *Vault) VaultSchemas() (*json.RawMessage, error) {
	schemas := &json.RawMessage{}

	_, err := vault.api.
		URL("/vault/api/v1/schemas").
		Get(&schemas)

	return schemas, err
}

func (vault *Vault) mkVaultReq(
	allowReadBy []string,
	allowWriteBy []string,
	secret interface{},
) tVaultReq {
	allow := func(ids []string) []rolestore.RoleRef {
		seq := []rolestore.RoleRef{}
		for _, id := range ids {
			seq = append(seq, rolestore.RoleRef{ID: id})
		}
		return seq
	}

	return tVaultReq{
		Data:       secret,
		AllowRead:  allow(allowReadBy),
		AllowWrite: allow(allowWriteBy),
		OwnerID:    "",
	}
}
