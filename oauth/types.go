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

// Credential is access/secret pair
type Credential struct {
	Access string
	Secret string
}

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
	client  restapi.Connector
	token   *AccessToken
	pending bool
}

//
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
