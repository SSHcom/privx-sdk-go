//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

import "time"

// KeyValue key value definition
type KeyValue struct {
	Key   string `json:"k"`
	Value string `json:"v"`
}

// ServiceStatus auth service status definition
type ServiceStatus struct {
	Variant       string     `json:"variant,omitempty"`
	Version       string     `json:"version,omitempty"`
	APIVersion    string     `json:"api_version,omitempty"`
	Status        string     `json:"status,omitempty"`
	StatusMessage string     `json:"status_message,omitempty"`
	ApplicationID string     `json:"app_id,omitempty"`
	ServerMode    string     `json:"server-mode,omitempty"`
	StatusDetails []KeyValue `json:"status_details,omitempty"`
	StartTime     time.Time  `json:"start_time,omitempty"`
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
