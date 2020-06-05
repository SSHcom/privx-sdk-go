//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package oauth

// Config contains OAuth2 client configuration information.
type Config struct {
	ClientID        string `toml:"oauth_client_id"`
	ClientSecret    string `toml:"oauth_client_secret"`
	RedirectURI     string `toml:"oauth_redirect_uri"`
	APIClientID     string `toml:"api_client_id"`
	APIClientSecret string `toml:"api_client_secret"`
}
