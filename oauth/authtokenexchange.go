//
// Copyright (c) 2026 SSH Communications Security Inc.
//
// All rights reserved.
//

package oauth

import (
	"fmt"
	"time"

	"github.com/SSHcom/privx-sdk-go/v2/restapi"
)

type tAuthTokenExchange struct{ *tAuth }

/*
WithExchangeToken authenticate users using externally created JWT Token in exchange to a PrivX access token

	auth := oauth.WithExchangeToken(
		restapi.New(
		restapi.BaseURL(url),
	),
		oauth.ExchangeToken(token), # required
		oauth.ExchangeScope("privx-user"), # optional
		oauth.AuthClientId("privx-ui"), # optional
	)

	return restapi.New(
		restapi.Auth(auth()),
		restapi.BaseURL(url),
	)
*/
func WithExchangeToken(client restapi.Connector, opts ...Option) restapi.Authorizer {
	return &tAuthTokenExchange{tAuth: newAuth(client, opts...)}
}

func (auth *tAuthTokenExchange) AccessToken() (string, error) {
	if err := auth.synchronized(auth.getAccessToken); err != nil {
		return "", err
	}

	return fmt.Sprintf("Bearer %s", auth.token.AccessToken), nil
}

func (auth *tAuthTokenExchange) getAccessToken() error {
	if auth.token != nil && auth.token.RefreshToken != "" {
		if auth.authRefreshToken() == nil {
			return nil
		}
	}
	return auth.exchangeToken()
}

func (auth *tAuthTokenExchange) exchangeToken() error {
	auth.token = nil

	token, err := auth.authExchangeToken()
	if err != nil {
		return err
	}

	auth.token = token
	return nil
}

func (auth *tAuth) authExchangeToken() (*AccessToken, error) {
	var token AccessToken

	request := Token{
		Token:    auth.exchangeToken,
		Scope:    auth.scope,
		ClientId: auth.clientId,
	}

	_, err := auth.client.
		URL("/auth/api/v1/token/login").
		CookieJar(auth.cookieJar).
		Post(&request, &token)

	if err == nil {
		token.notAfter = time.Now().Add(
			time.Duration(token.ExpiresIn) * time.Second)
	}

	return &token, err
}
