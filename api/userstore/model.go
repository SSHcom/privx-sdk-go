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
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Sortdir string `json:"sortdir,omitempty"`
	Query   string `json:"query,omitempty"`
}

// FilterUser struct for local users queries.
type FilterUser struct {
	Params
	UserID   string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

// TrustedClient definition
type TrustedClient struct {
	ID                            string     `json:"id,omitempty"`
	Secret                        string     `json:"secret,omitempty"`
	Name                          string     `json:"name,omitempty"`
	WebProxyAddress               string     `json:"web_proxy_address,omitempty"`
	WebProxyPort                  string     `json:"web_proxy_port,omitempty"`
	Registered                    bool       `json:"registered,omitempty"`
	Enabled                       bool       `json:"enabled,omitempty"`
	Type                          ClientType `json:"type,omitempty"`
	Permissions                   []string   `json:"permissions,omitempty"`
	WebProxyExtenderRoutePatterns []string   `json:"web_proxy_extender_route_patterns,omitempty"`
	ExtenderAddress               []string   `json:"extender_address,omitempty"`
	Subnets                       []string   `json:"subnets,omitempty"`
	RoutingPrefix                 string     `json:"routing_prefix,omitempty"`
	AccessGroupId                 string     `json:"access_group_id,omitempty"`
	GroupId                       string     `json:"group_id,omitempty"`
	OAuthClientID                 string     `json:"oauth_client_id,omitempty"`
	OAuthClientSecret             string     `json:"oauth_client_secret,omitempty"`
	Data                          string     `json:"data,omitempty"`
	Created                       string     `json:"created,omitempty"`
	Updated                       string     `json:"updated,omitempty"`
	UpdatedBy                     string     `json:"updated_by,omitempty"`
	Author                        string     `json:"author,omitempty"`
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
	ID         string      `json:"id,omitempty"`
	Created    string      `json:"created,omitempty"`
	Updated    string      `json:"updated,omitempty"`
	UpdatedBy  string      `json:"updated_by,omitempty"`
	Author     string      `json:"author,omitempty"`
	Comment    string      `json:"comment,omitempty"`
	Tags       []string    `json:"tags,omitempty"`
	Username   string      `json:"username,omitempty"`
	GivenName  string      `json:"given_name,omitempty"`
	FullName   string      `json:"full_name,omitempty"`
	JobTitle   string      `json:"job_title,omitempty"`
	Company    string      `json:"company,omitempty"`
	Department string      `json:"department,omitempty"`
	Email      string      `json:"email,omitempty"`
	Telephone  string      `json:"telephone,omitempty"`
	Locale     string      `json:"locale,omitempty"`
	Password   Password    `json:"password,omitempty"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

// User attribute
type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Password definition
type Password struct {
	Password string `json:"password,omitempty"`
}
