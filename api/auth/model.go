//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

import "time"

// query params definition for auth service.
type params struct {
	Sortkey string `json:"sortkey"`
	Sortdir string `json:"sortdir"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
}

// IdpClient identity provider client definition.
type IdpClient struct {
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

// IdpClientConfig identity provider config definition.
type IdpClientConfig struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// IdpResponse identity provider create response definition.
type IdpResponse struct {
	ID string `json:"id"`
}

// Session user session definition.
type Session struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id,omitempty"`
	SourceID        string    `json:"source_id,omitempty"`
	Domain          string    `json:"domain,omitempty"`
	Username        string    `json:"username,omitempty"`
	RemoteAddr      string    `json:"remote_addr,omitempty"`
	UserAgent       string    `json:"user_agent,omitempty"`
	Type            string    `json:"type,omitempty"`
	ParentSessionId string    `json:"parent_session_id,omitempty"`
	Created         time.Time `json:"created,omitempty"`
	Updated         time.Time `json:"updated,omitempty"`
	Expires         time.Time `json:"expires,omitempty"`
	TokenExpires    time.Time `json:"token_expires,omitempty"`
	LoggedOut       bool      `json:"logged_out,omitempty"`
	Current         bool      `json:"current,omitempty"`
}

// SessionPasswordPolicy session password policy definition.
type SessionPasswordPolicy struct {
	PasswordMinLength    int    `json:"password_min_length,omitempty"`
	PasswordMaxLength    int    `json:"password_max_length,omitempty"`
	UseSpecialCharacters bool   `json:"use_special_characters,omitempty"`
	UseLowercase         bool   `json:"use_lower_case,omitempty"`
	UseUppercase         bool   `json:"use_upper_case,omitempty"`
	UseNumbers           bool   `json:"use_numbers,omitempty"`
	PasswordEntropy      int    `json:"password_entropy,omitempty"`
	PasswordStrength     string `json:"password_strength,omitempty"`
}

// SessionSearchRequest session search request parameter definition.
type SessionSearchRequest struct {
	Keywords string `json:"keywords,omitempty"`
	UserID   string `json:"user_id,omitempty"`
	Type     string `json:"type,omitempty"`
}

// Device mobile gateway device definition.
type Device struct {
	ID        string `json:"id"`
	OS        string `json:"os,omitempty"`
	Name      string `json:"name,omitempty"`
	Activated string `json:"activated,omitempty"`
	Updated   string `json:"updated,omitempty"`
	LastUsed  string `json:"lastUsed,omitempty"`
}
