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

// Params struct for pagination queries.
type Params struct {
	Offset string `json:"offset,omitempty"`
	Limit  string `json:"limit,omitempty"`
}

// FilterUser struct for local users queries.
type FilterUser struct {
	Params
	UserID   string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

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

// LocalUser definition
type LocalUser struct {
	ID         string   `json:"id,omnitempty"`
	Created    string   `json:"created,omitempty"`
	Updated    string   `json:"updated,omitempty"`
	UpdatedBy  string   `json:"updated_by,omitempty"`
	Author     string   `json:"author,omitempty"`
	Comment    string   `json:"comment,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	Username   string   `json:"username,omitempty"`
	GivenName  string   `json:"given_name,omitempty"`
	FullName   string   `json:"full_name,omitempty"`
	JobTitle   string   `json:"job_title,omitempty"`
	Company    string   `json:"company,omitempty"`
	Department string   `json:"department,omitempty"`
	Email      string   `json:"email,omitempty"`
	Telephone  string   `json:"telephone,omitempty"`
	Locale     string   `json:"locale,omitempty"`
	Password   Password `json:"password,omitempty"`
}

// Password definition
type Password struct {
	Password string `json:"password,omitempty"`
}
