//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

import (
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Auth is a auth client instance.
type Auth struct {
	api restapi.Connector
}

// New creates a new auth client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *Auth {
	return &Auth{api: api}
}

// AuthStatus get microservice status
func (store *Auth) AuthStatus() (*ServiceStatus, error) {
	status := &ServiceStatus{}

	_, err := store.api.
		URL("/auth/api/v1/status").
		Get(status)

	return status, err
}
