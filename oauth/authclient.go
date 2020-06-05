//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package oauth

import (
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

type tAuthPassword struct {
	tAuth
	apiClient Credential
	digest    string
}

/*

WithClientID executes OAuth2 Resource Owner Password Grant
It uses access/secret key pair to authenticate client

  auth := oauth2.WithClientID(
		// API client API
		oauth2.Credential{Access: "...", Secret: "..."},
		// OAuth client digest
		oauth2.Credential{Access: "...", Secret: "..."},
		restapi.Endpoint("https://privx.example.com"),
	)

	client := restapi.New(
		restapi.Auth(auth),
		restapi.Endpoint("https://privx.example.com"),
	)

	rolestore.New(client)
*/
func WithClientID(apiClient, digest Credential, opts ...restapi.Option) restapi.Authorizer {
	client := restapi.New(append(opts, restapi.NoRedirect())...)

	return &tAuthPassword{
		tAuth: tAuth{
			Cond:   sync.NewCond(new(sync.Mutex)),
			client: client,
		},
		apiClient: apiClient,
		digest:    base64.StdEncoding.EncodeToString([]byte(digest.Access + ":" + digest.Secret)),
	}
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
		Access:    auth.apiClient.Access,
		Secret:    auth.apiClient.Secret,
	}
	var token AccessToken

	_, err := auth.client.
		Post("/auth/api/v1/oauth/token").
		With("Content-Type", "application/x-www-form-urlencoded").
		With("Authorization", "Basic "+auth.digest).
		Send(request).
		Recv(&token)

	if err != nil {
		token.notAfter = time.Now().Add(
			time.Duration(token.ExpiresIn) * time.Second)
	}
	auth.token = &token

	return err
}
