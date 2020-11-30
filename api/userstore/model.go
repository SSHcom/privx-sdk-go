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
	ClientExtender         = ClientType("EXTENDER")
	ClientHostProvisioning = ClientType("HOST_PROVISIONING")
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
		Type:        ClientExtender,
		Permissions: []string{"privx-extender"},
		Name:        name,
	}
}

// HostProvisioning creates new trusted client
func HostProvisioning(name string) TrustedClient {
	return TrustedClient{
		Type:        ClientHostProvisioning,
		Permissions: []string{"privx-host-provisioning"},
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
