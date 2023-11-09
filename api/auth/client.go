//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/common"
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
func (store *Auth) AuthStatus() (*common.ServiceStatus, error) {
	status := &common.ServiceStatus{}

	_, err := store.api.
		URL("/auth/api/v1/status").
		Get(status)

	return status, err
}

// CreateIdpClient creates a new identity provider client configuration.
func (store *Auth) CreateIdpClient(idpClient *IDPClient) (IDstruct, error) {
	idpClientId := IDstruct{}

	_, err := store.api.
		URL("/auth/api/v1/idp/clients").
		Post(&idpClient, &idpClientId)

	return idpClientId, err
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

// UserSessions fetches valid sessions by userID.
func (store *Auth) UserSessions(offset, limit int, sortkey, sortdir, userID string) (*sessionsResult, error) {
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}
	userSessions := &sessionsResult{}

	_, err := store.api.
		URL("/auth/api/v1/sessionstorage/users/%s/sessions", userID).
		Query(&filters).
		Get(&userSessions)

	return userSessions, err
}

// SourceSessions fetches valid sessions by sourceID.
func (store *Auth) SourceSessions(offset, limit int, sortkey, sortdir, sourceID string) (*sessionsResult, error) {
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}
	sourceSessions := &sessionsResult{}

	_, err := store.api.
		URL("/auth/api/v1/sessionstorage/sources/%s/sessions", sourceID).
		Query(&filters).
		Get(&sourceSessions)

	return sourceSessions, err
}

// SearchSessions searches for sessions
func (store *Auth) SearchSessions(offset, limit int, sortkey, sortdir string, search *SearchParams) (*sessionsResult, error) {
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}
	sessions := &sessionsResult{}

	_, err := store.api.
		URL("/auth/api/v1/sessionstorage/sessions/search").
		Query(&filters).
		Post(search, &sessions)

	return sessions, err
}

// TerminateSession terminates single session by ID.
func (store *Auth) TerminateSession(sessionID string) error {

	_, err := store.api.
		URL("/auth/api/v1/sessionstorage/sessions/%s/terminate", sessionID).
		Post(nil)
	return err
}

// TerminateUserSessions terminates all sessions for a user.
func (store *Auth) TerminateUserSessions(userID string) error {

	_, err := store.api.
		URL("/auth/api/v1/sessionstorage/users/%s/sessions/terminate", userID).
		Post(nil)
	return err
}

// Logout logs out user.
func (store *Auth) Logout() error {

	_, err := store.api.
		URL("/auth/api/v1/logout").
		Post(nil)
	return err
}

func (store *Auth) GetUserPairedDevices(userID string) (*PairedDevices, error) {
	devices := &PairedDevices{}

	_, err := store.api.
		URL("/auth/api/v1/users/%s/devices", userID).
		Get(devices)

	return devices, err
}

func (store *Auth) UnpairUserDevice(userID, deviceID string) error {
	_, err := store.api.
		URL("/auth/api/v1/users/%s/devices/%s", userID, deviceID).
		Delete()

	return err
}
