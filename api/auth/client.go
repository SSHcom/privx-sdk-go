//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

import (
	"net/url"

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

// CreateIdpClient creates a new identity provider client configuration.
func (store *Auth) CreateIdpClient(idpClient *IDPClient) (map[string]string, error) {
	idpClientIdMap := make(map[string]string)

	_, err := store.api.
		URL("/auth/api/v1/idp/clients").
		Post(&idpClient, &idpClientIdMap)

	return idpClientIdMap, err
}

// UpdateIdpClient updates existing identity provider client configuration definition.
func (store *Auth) UpdateIdpClient(idpClient *IDPClient, idpID string) error {

	_, err := store.api.
		URL("/auth/api/v1/idp/clients/%s", idpID).
		Put(&idpClient)

	return err
}

// IdpClient fetches existing identity provider client configuration.
func (store *Auth) IdpClient(idpID string) (*IDPClient, error) {
	idpClient := &IDPClient{}

	_, err := store.api.
		URL("/auth/api/v1/idp/clients/%s", idpID).
		Get(&idpClient)

	return idpClient, err
}

// DeleteIdpClient delete identity provider client configuration by ID.
func (store *Auth) DeleteIdpClient(idpID string) error {

	_, err := store.api.
		URL("/auth/api/v1/idp/clients/%s", idpID).
		Delete()

	return err
}

// RegenerateIdpClientConfig regenerates client_id and client_secret
// for OIDC identity provider client configuration.
func (store *Auth) RegenerateIdpClientConfig(idpID string) (*IdpClientConfig, error) {
	clientConfig := &IdpClientConfig{}

	_, err := store.api.
		URL("/auth/api/v1/idp/clients/%s/regenerate", url.PathEscape(idpID)).
		Post(nil, &clientConfig)

	return clientConfig, err
}
