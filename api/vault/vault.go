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

// tVaultReq t vault request definition
type tVaultReq struct {
	Name       string              `json:"name,omitempty"`
	Data       interface{}         `json:"data"`
	AllowRead  []rolestore.RoleRef `json:"read_roles,omitempty"`
	AllowWrite []rolestore.RoleRef `json:"write_roles,omitempty"`
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

// Secret gets the content of the argument secret.
func (vault *Vault) Secret(name string) (*Secret, error) {
	bag := &Secret{}
	_, err := vault.api.
		URL("/vault/api/v1/secrets/%s", url.PathEscape(name)).
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

// DeleteSecret delete existing secret from PrivX vault
func (vault *Vault) DeleteSecret(name string) error {
	_, err := vault.api.
		URL("/vault/api/v1/secrets/%s", name).
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
	}
}
