//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package vault

import (
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

// New creates a new Vault client instance, using the argument
// SDK API client.
func New(api restapi.Connector) *Vault {
	return &Vault{api: api}
}

// Secrets returns secrets client has access to
func (vault *Vault) Secrets(offset, limit string) ([]Secret, error) {
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
func (vault *Vault) Secret(name string) (bag *Secret, err error) {
	bag = new(Secret)
	_, err = vault.api.
		URL("/vault/api/v1/secrets/%s", url.PathEscape(name)).
		Get(bag)

	return
}

type tVaultReq struct {
	Name       string              `json:"name,omitempty"`
	Data       interface{}         `json:"data"`
	AllowRead  []rolestore.RoleRef `json:"read_roles,omitempty"`
	AllowWrite []rolestore.RoleRef `json:"write_roles,omitempty"`
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

// Create new secret to PrivX Vault
func (vault *Vault) Create(
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

// Update existing secret at PrivX Vault
func (vault *Vault) Update(
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

// Remove existing secret from PrivX vault
func (vault *Vault) Remove(name string) error {
	_, err := vault.api.
		URL("/vault/api/v1/secrets/%s", name).
		Delete()

	return err
}
