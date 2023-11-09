//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

import "time"

// Params query params definition
type Params struct {
	Sortkey string `json:"sortkey,omitempty"`
	Sortdir string `json:"sortdir,omitempty"`
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
}

// IDPClient idp client definition
type IDPClient struct {
	ID                             string            `json:"id"`
	Name                           string            `json:"name"`
	Created                        time.Time         `json:"created,omitempty"`
	Updated                        time.Time         `json:"updated,omitempty"`
	IDPType                        string            `json:"idp_type"`
	OIDCIssuer                     string            `json:"oidc_issuer,omitempty"`
	OIDCAudience                   []string          `json:"oidc_audience"`
	OIDCClientID                   string            `json:"oidc_client_id,omitempty"`
	OIDCClientSecret               string            `json:"oidc_client_secret,omitempty"`
	OIDCScopesEnabled              []string          `json:"oidc_scopes_enabled"`
	OIDCResponseTypesSupported     []string          `json:"oidc_response_types_supported,omitempty"`
	OIDCGrantTypesSupported        []string          `json:"oidc_grant_types_supported,omitempty"`
	OIDCEnablePKCE                 bool              `json:"oidc_code_challenge_method_enabled,omitempty"`
	OIDCEnabledAuthMethod          string            `json:"oidc_auth_method_enabled,omitempty"`
	OIDCAuthMethodPost             bool              `json:"oidc_auth_method_post,omitempty"`
	OIDCGrantTypeRefreshToken      bool              `json:"oidc_grant_type_refresh_token,omitempty"`
	OIDCAllowedRedirectURIs        []string          `json:"oidc_allowed_redirect_uris,omitempty"`
	OIDCDefaultLogoutRedirectURI   string            `json:"oidc_default_logout_redirect_uri,omitempty"`
	OIDCAllowedLogoutRedirectURIs  []string          `json:"oidc_allowed_logout_redirect_uris,omitempty"`
	OIDCAttributeMapping           map[string]string `json:"oidc_attribute_mapping,omitempty"`
	OIDCSignatureAlgorithm         string            `json:"oidc_signature_algorithm,omitempty"`
	OIDCAccessTokenValidInMinutes  int               `json:"oidc_access_token_valid_in_minutes,omitempty"`
	OIDCRefreshTokenValidInMinutes int               `json:"oidc_refresh_token_valid_in_minutes,omitempty"`
	UserFilter                     string            `json:"user_filter,omitempty"`
	Enabled                        bool              `json:"enabled"`
}

// IDPClientsResult idp client list result definition
type IDPClientsResult struct {
	Count int         `json:"count"`
	Items []IDPClient `json:"items"`
}

// IdpClientConfig config definition with client_id and client_secret
type IdpClientConfig struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type IDstruct struct {
	ID string `json:"id"`
}

type Session struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	SourceID     string    `json:"source_id"`
	Domain       string    `json:"domain"`
	Username     string    `json:"username"`
	RemoteAddr   string    `json:"remote_addr"`
	UserAgent    string    `json:"user_agent"`
	Type         string    `json:"type"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
	Expires      time.Time `json:"expires"`
	TokenExpires time.Time `json:"token_expires"`
	LoggedOut    bool      `json:"logged_out"`
	Current      bool      `json:"current,omitempty"`
}

type sessionsResult struct {
	Items []Session `json:"items"`
	Count int       `json:"count"`
}

type SessionPasswordPolicy struct {
	PasswordMinLength    int    `json:"password_min_length"`
	PasswordMaxLength    int    `json:"password_max_length"`
	UseSpecialCharacters bool   `json:"use_special_characters"`
	UseLowercase         bool   `json:"use_lower_case"`
	UseUppercase         bool   `json:"use_upper_case"`
	UseNumbers           bool   `json:"use_numbers"`
	PasswordEntropy      int    `json:"password_entropy,omitempty"`
	PasswordStrength     string `json:"password_strength,omitempty"`
}

// SearchParams search params definition
type SearchParams struct {
	Keywords string `json:"keywords,omitempty"`
	UserID   string `json:"user_id,omitempty"`
	Type     string `json:"type,omitempty"`
}

type Device struct {
	ID        string `json:"id"`
	OS        string `json:"os"`
	Name      string `json:"name"`
	Activated string `json:"activated"`
	Updated   string `json:"updated"`
	LastUsed  string `json:"lastUsed"`
}

type PairedDevices struct {
	Count int      `json:"count"`
	Items []Device `json:"items"`
}
