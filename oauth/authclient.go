//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package oauth

import (
	"fmt"
	"time"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

type tAuthPassword struct{ *tAuth }

/*
WithClientID executes OAuth2 Resource Owner Password Grant
It uses access/secret key pair to authenticate client

	  auth := oauth2.WithClientID(
			restapi.New(...),
			oauth2.Access(...),
			oauth2.Secret(...),
		)

		client := restapi.New(
			restapi.Auth(auth),
			restapi.Endpoint("https://privx.example.com"),
		)

		rolestore.New(client)
*/
func WithClientID(client restapi.Connector, opts ...Option) restapi.Authorizer {
	return &tAuthPassword{tAuth: newAuth(client, opts...)}
}

func (auth *tAuthPassword) AccessToken() (token string, err error) {
	if err = auth.synchronized(auth.grantPasswordCredentials); err == nil {
		token = fmt.Sprintf("Bearer %s", auth.token.AccessToken)
	}
	return
}

func (auth *tAuthPassword) grantPasswordCredentials() error {
	auth.token = nil

	request := reqAccessTokenPassword{
		GrantType: "password",
		Access:    auth.access,
		Secret:    auth.secret,
	}
	var token AccessToken

	header, err := auth.client.
		URL("/auth/api/v1/oauth/token").
		Header("Content-Type", "application/x-www-form-urlencoded").
		Header("Authorization", "Basic "+auth.digest).
		Post(request, &token)

	if auth.useCookies && auth.cookie == "" {
		auth.cookie = header.Get("Set-Cookie")
	}

	if err != nil {
		token.notAfter = time.Now().Add(
			time.Duration(token.ExpiresIn) * time.Second)
	}
	auth.token = &token

	return err
}
