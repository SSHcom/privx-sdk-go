//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package oauth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/SSHcom/privx-sdk-go/pkce"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

var clientID = tClientID{
	ID:          "privx-ui",
	RedirectURI: "/privx/oauth-callback",
}

type tAuthCode struct{ *tAuth }

/*

WithCredential executes OAuth2 Authorization Code Grant
It uses access/secret key pair to authenticate client

  auth := oauth2.WithCredential(
		oauth2.Access(...),
		oauth2.Secret(...),
		oauth2.Transport(...),
	)

	client := restapi.New(
		restapi.Auth(auth),
		restapi.Endpoint("https://privx.example.com"),
	)

	rolestore.New(client)

*/
func WithCredential(client restapi.Connector, opts ...Option) restapi.Authorizer {
	return &tAuthCode{tAuth: newAuth(client, opts...)}
}

func (auth *tAuthCode) AccessToken() (token string, err error) {
	if err = auth.synchronized(auth.grantAuthorizationCode); err == nil {
		token = fmt.Sprintf("Bearer %s", auth.token.AccessToken)
	}
	return
}

func (auth *tAuthCode) grantAuthorizationCode() error {
	auth.token = nil

	cv, err := pkce.NewCodeVerifier()
	if err != nil {
		return err
	}

	challenge, method := cv.ChallengeS256()
	state, err := auth.random()
	if err != nil {
		return err
	}

	session, err := auth.authSession(challenge, method, state)
	if err != nil {
		return err
	}

	exchange, err := auth.authCredential(session, state)
	if err != nil {
		return err
	}

	token, err := auth.authAccessToken(exchange, cv)
	if err != nil {
		return err
	}

	auth.token = token
	return nil
}

//
func (auth *tAuthCode) authSession(challenge, method, state string) (string, error) {
	request := reqAuthSession{
		tClientID:     clientID,
		ResponseType:  "code",
		State:         state,
		UserAgent:     restapi.UserAgent,
		CodeChallenge: challenge,
		CodeMethod:    method,
	}

	head, err := auth.client.
		URL("/auth/api/v1/oauth/authorize").
		Query(request).
		Status(307)

	if err != nil {
		return "", err
	}

	uri, err := url.Parse(head.Get("location"))
	if err != nil {
		return "", err
	}

	return uri.Query().Get("token"), nil
}

//
func (auth *tAuthCode) authCredential(session, state string) (string, error) {
	request := reqExchangeCode{
		Access: auth.access,
		Secret: auth.secret,
		Token:  session,
	}

	var response struct {
		Code  string `json:"code"`
		State string `json:"state"`
	}

	_, err := auth.client.
		URL("/auth/api/v1/login").
		Post(request, &response)

	if response.State != state {
		return "", errors.New("invalid response state")
	}

	return response.Code, err
}

//
func (auth *tAuth) authAccessToken(code string, cv pkce.CodeVerifier) (*AccessToken, error) {
	request := reqAccessToken{
		tClientID:  clientID,
		GrantType:  "authorization_code",
		Code:       code,
		CodeVerify: cv.String(),
	}
	var token AccessToken

	_, err := auth.client.
		URL("/auth/api/v1/oauth/token").
		Header("Content-Type", "application/x-www-form-urlencoded").
		Post(request, &token)

	if err != nil {
		token.notAfter = time.Now().Add(
			time.Duration(token.ExpiresIn) * time.Second)
	}

	return &token, err
}

func (auth *tAuth) random() (string, error) {
	var buf [32]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.
		WithPadding(base64.NoPadding).
		EncodeToString(buf[:]), nil
}
