//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package oauth

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// ClientID is a pair of unique client id and redirect uri
type ClientID struct {
	ID          string `json:"client_id"`
	RedirectURI string `json:"redirect_uri"`
}

var clientID = ClientID{
	ID:          "privx-ui",
	RedirectURI: "/privx/oauth-callback",
}

type tAuthCode struct {
	tAuth
	credential Credential
}

// WithCredential obtains access token using access/secret key pair
func WithCredential(credential Credential, opts ...restapi.Option) restapi.Authorizer {
	client := restapi.New(append(opts, restapi.NoRedirect())...)

	return &tAuthCode{
		tAuth: tAuth{
			Cond:   sync.NewCond(new(sync.Mutex)),
			client: client,
		},
		credential: credential,
	}
}

func (auth *tAuthCode) AccessToken() (token string, err error) {
	if err = auth.synchronized(auth.grantAuthorizationCode); err == nil {
		token = fmt.Sprintf("Bearer %s", auth.token.AccessToken)
	}
	return
}

func (auth *tAuthCode) grantAuthorizationCode() error {
	auth.token = nil

	session, err := auth.authSession()
	if err != nil {
		return err
	}

	exchange, err := auth.authCredential(session)
	if err != nil {
		return err
	}

	token, err := auth.authAccessToken(exchange)
	if err != nil {
		return err
	}

	auth.token = token
	return nil
}

func (auth *tAuthCode) authSession() (string, error) {
	request := struct {
		ClientID
		ResponseType string `json:"response_type"`
		State        string `json:"state"`
	}{
		ClientID:     clientID,
		ResponseType: "code",
		State:        "",
	}

	head, err := auth.client.
		Get("/auth/api/v1/oauth/authorize").
		Params(request).
		RecvStatus(307)

	if err != nil {
		return "", err
	}

	uri, err := url.Parse(head.Get("location"))
	if err != nil {
		return "", err
	}

	return uri.Query().Get("token"), nil
}

func (auth *tAuthCode) authCredential(session string) (string, error) {
	request := struct {
		Access string `json:"username"`
		Secret string `json:"password"`
		Token  string `json:"token"`
	}{
		Access: auth.credential.Access,
		Secret: auth.credential.Secret,
		Token:  session,
	}

	var response struct {
		Code string `json:"code"`
	}

	_, err := auth.client.
		Post("/auth/api/v1/login").
		Send(request).
		Recv(&response)

	return response.Code, err
}

func (auth *tAuth) authAccessToken(code string) (*AccessToken, error) {
	request := struct {
		ClientID
		GrantType string `json:"grant_type"`
		Code      string `json:"code"`
	}{
		ClientID:  clientID,
		GrantType: "authorization_code",
		Code:      code,
	}
	var token AccessToken

	_, err := auth.client.
		Post("/auth/api/v1/oauth/token").
		With("Content-Type", "application/x-www-form-urlencoded").
		Send(request).
		Recv(&token)

	if err != nil {
		token.notAfter = time.Now().Add(
			time.Duration(token.ExpiresIn) * time.Second)
	}

	return &token, err
}
