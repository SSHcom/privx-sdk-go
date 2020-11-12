//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package userstore

import "github.com/SSHcom/privx-sdk-go/api/rolestore"

// ClientType is a type of trusted clients
type ClientType string

// ClientType supported values
const (
	EXTENDER = ClientType("EXTENDER")
)

// TrustedClient definition
type TrustedClient struct {
	ID          string     `json:"id,omitempty"`
	Type        ClientType `json:"type,omitempty"`
	Secret      string     `json:"secret,omitempty"`
	Name        string     `json:"name,omitempty"`
	Permissions []string   `json:"permissions,omitempty"`
	Registered  bool       `json:"registered,omitempty"`
	Enabled     bool       `json:"enabled,omitempty"`
}

// Extender creates new trusted client
func Extender(name string) TrustedClient {
	return TrustedClient{
		Type:        EXTENDER,
		Permissions: []string{"privx-extender"},
		Name:        name,
	}
}

// APIClient definition
type APIClient struct {
	ID               string              `json:"id,omitempty"`
	Name             string              `json:"name,omitempty"`
	Secret           string              `json:"secret,omitempty"`
	AuthClientID     string              `json:"oauth_client_id"`
	AuthClientSecret string              `json:"oauth_client_secret"`
	Roles            []rolestore.RoleRef `json:"roles,omitempty"`
	Created          string              `json:"created,omitempty"`
	Author           string              `json:"author,omitempty"`
}
