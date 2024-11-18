//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package userstore

import "github.com/SSHcom/privx-sdk-go/api/rolestore"

// LocalUserParams local user query parameters definition.
type LocalUserParams struct {
	UserID   string `url:"id,omitempty"`
	Username string `url:"username,omitempty"`
}

// TrustedClient trusted client definition.
type TrustedClient struct {
	ID                            string   `json:"id"`
	Type                          string   `json:"type"`
	Secret                        string   `json:"secret"`
	Name                          string   `json:"name"`
	AccessGroupID                 string   `json:"access_group_id"`
	Created                       string   `json:"created,omitempty"`
	Updated                       string   `json:"updated,omitempty"`
	UpdatedBy                     string   `json:"updated_by,omitempty"`
	Author                        string   `json:"author,omitempty"`
	Permissions                   []string `json:"permissions"`
	Subnets                       []string `json:"subnets"`
	Enabled                       bool     `json:"enabled"`
	Registered                    bool     `json:"registered"`
	ExtenderAddress               []string `json:"extender_address"`
	OAuthClientID                 string   `json:"oauth_client_id,omitempty"`
	OAuthClientSecret             string   `json:"oauth_client_secret,omitempty"`
	GroupID                       string   `json:"group_id"`
	WebProxyAddress               string   `json:"web_proxy_address,omitempty"`
	WebProxyPort                  string   `json:"web_proxy_port,omitempty"`
	WebProxyExtenderRoutePatterns []string `json:"web_proxy_extender_route_patterns,omitempty"`
	Data                          string   `json:"data,omitempty"`
	RoutingPrefix                 string   `json:"routing_prefix"`
}

type ExtenderClient struct {
	Type                          string   `json:"type"`
	Name                          string   `json:"name"`
	RoutingPrefix                 string   `json:"routing_prefix"`
	WebProxyExtenderRoutePatterns []string `json:"web_proxy_extender_route_patterns,omitempty"`
	WebProxyExtenderRoutes        []string `json:"web_proxy_extender_routes,omitempty"`
}

// APIClient api client definition.
type APIClient struct {
	ID                string                 `json:"id"`
	Secret            string                 `json:"secret"`
	Name              string                 `json:"name"`
	Created           string                 `json:"created,omitempty"`
	Updated           string                 `json:"updated,omitempty"`
	UpdatedBy         string                 `json:"updated_by,omitempty"`
	Author            string                 `json:"author,omitempty"`
	Roles             []rolestore.RoleHandle `json:"roles"`
	OAuthClientID     string                 `json:"oauth_client_id,omitempty"`
	OAuthClientSecret string                 `json:"oauth_client_secret,omitempty"`
}

// APIClientCreate api client create request definition.
type APIClientCreate struct {
	Name  string                 `json:"name"`
	Roles []rolestore.RoleHandle `json:"roles"`
}

// APIClientSearch api client search request definition.
type APIClientSearch struct {
	Keywords string `json:"keywords"`
	SortDir  string `json:"sortdir"`
	SortKey  string `json:"sortkey"`
	Limit    *int   `json:"limit,omitempty"`
	Offset   *int   `json:"offset,omitempty"`
}

// LocalUser local user definition.
type LocalUser struct {
	ID                     string            `json:"id,omitempty"`
	Created                string            `json:"created,omitempty"`
	Updated                string            `json:"updated,omitempty"`
	UpdatedBy              string            `json:"updated_by,omitempty"`
	Author                 string            `json:"author,omitempty"`
	Comment                string            `json:"comment,omitempty"`
	Tags                   []string          `json:"tags,omitempty"`
	Principal              string            `json:"username,omitempty"`
	WindowsAccount         string            `json:"windows_account,omitempty"`
	UnixAccount            string            `json:"unix_account,omitempty"`
	FullName               string            `json:"full_name,omitempty"`
	DisplayName            string            `json:"display_name,omitempty"`
	FirstName              string            `json:"first_name,omitempty"`
	LastName               string            `json:"last_name,omitempty"`
	JobTitle               string            `json:"job_title,omitempty"`
	Company                string            `json:"company,omitempty"`
	Department             string            `json:"department,omitempty"`
	Email                  string            `json:"email,omitempty"`
	Telephone              string            `json:"telephone,omitempty"`
	Locale                 string            `json:"locale,omitempty"`
	Password               LocalUserPassword `json:"password"`
	PasswordChangeRequired bool              `json:"password_change_required"`
	Attributes             []Attributes      `json:"attributes"`
}

// Attributes user attribute definition.
type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// LocalUserPassword local user password definition.
type LocalUserPassword struct {
	Password string `json:"password,omitempty"`
	Created  string `json:"created,omitempty"`
}
