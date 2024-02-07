//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package oauth

import (
	"sync"
	"time"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// AccessToken contains OAuth2 access token information.
type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    uint   `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	notAfter     time.Time
}

// isInvalid checks if token is valid
func (token *AccessToken) isInvalid() bool {
	return token == nil || time.Now().After(token.notAfter)
}

// tAuth authorizer client
type tAuth struct {
	*sync.Cond
	access     string
	secret     string
	digest     string
	client     restapi.Connector
	token      *AccessToken
	useCookies bool
	cookie     string
	pending    bool
}

func newAuth(client restapi.Connector, opts ...Option) *tAuth {
	auth := &tAuth{
		Cond:   sync.NewCond(new(sync.Mutex)),
		client: client,
	}

	for _, opt := range opts {
		auth = opt(auth)
	}
	return auth
}

// synchronized closure execution in the context of authorizer
func (auth *tAuth) synchronized(f func() error) (err error) {
	auth.L.Lock()
	for auth.pending {
		auth.Wait()
	}
	defer auth.L.Unlock()
	if auth.token.isInvalid() {
		auth.pending = true
		auth.L.Unlock()

		err = f()

		auth.L.Lock()
		auth.pending = false
		auth.Broadcast()
	}

	return
}

func (auth *tAuth) Cookie() string {
	return auth.cookie
}

// tClientID is a pair of unique client id and redirect uri
type tClientID struct {
	ID          string `json:"client_id"`
	RedirectURI string `json:"redirect_uri"`
}

// reqAuthSession establishes new auth session
type reqAuthSession struct {
	tClientID
	ResponseType  string `json:"response_type"`
	State         string `json:"state"`
	UserAgent     string `json:"user_agent"`
	CodeChallenge string `json:"code_challenge"`
	CodeMethod    string `json:"code_challenge_method"`
}

// reqExchangeCode fetches the code from authorizer
type reqExchangeCode struct {
	Access string `json:"username"`
	Secret string `json:"password"`
	Token  string `json:"token"`
}

// reqAccessToken exchanges the code for access token
type reqAccessToken struct {
	tClientID
	GrantType  string `json:"grant_type"`
	Code       string `json:"code"`
	CodeVerify string `json:"code_verifier"`
}

// reqAccessToken
type reqAccessTokenPassword struct {
	GrantType string `json:"grant_type"`
	Access    string `json:"username"`
	Secret    string `json:"password"`
}
