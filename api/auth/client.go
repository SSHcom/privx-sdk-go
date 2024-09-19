//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Auth is a auth client instance.
type Auth struct {
	api restapi.Connector
}

// New auth client constructor.
func New(api restapi.Connector) *Auth {
	return &Auth{api: api}
}

// MARK: Status
// Status get auth microservice status.
func (c *Auth) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/auth/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Identity Provider
// CreateIdpClient creates a new identity provider client configuration.
func (c *Auth) CreateIdpClient(idpClient *IdpClient) (*response.Id, error) {
	idpClientResponse := &response.Id{}

	_, err := c.api.
		URL("/auth/api/v1/idp/clients").
		Post(&idpClient, &idpClientResponse)

	return idpClientResponse, err
}

// UpdateIdpClient updates existing identity provider client configuration definition.
func (c *Auth) UpdateIdpClient(idpClient *IdpClient, idpId string) error {
	_, err := c.api.
		URL("/auth/api/v1/idp/clients/%s", url.PathEscape(idpId)).
		Put(&idpClient)

	return err
}

// GetIdpClient get existing identity provider client configuration.
func (c *Auth) GetIdpClient(idpId string) (*IdpClient, error) {
	idpClient := &IdpClient{}

	_, err := c.api.
		URL("/auth/api/v1/idp/clients/%s", url.PathEscape(idpId)).
		Get(&idpClient)

	return idpClient, err
}

// DeleteIdpClient delete identity provider client configuration by Id.
func (c *Auth) DeleteIdpClient(idpId string) error {
	_, err := c.api.
		URL("/auth/api/v1/idp/clients/%s", url.PathEscape(idpId)).
		Delete()

	return err
}

// RegenerateIdpClientConfig regenerates client_id and client_secret
// for OIDC identity provider client configuration.
func (c *Auth) RegenerateIdpClientConfig(idpId string) (*IdpClientConfig, error) {
	clientConfig := &IdpClientConfig{}

	_, err := c.api.
		URL("/auth/api/v1/idp/clients/%s/regenerate", url.PathEscape(idpId)).
		Post(nil, &clientConfig)

	return clientConfig, err
}

// MARK: Session Storage
// GetUserSessions get valid sessions by userID.
func (c *Auth) GetUserSessions(filters *response.FilterParams, userId string) (*response.Page[Session], error) {
	userSessions := &response.Page[Session]{}

	_, err := c.api.
		URL("/auth/api/v1/sessionstorage/users/%s/sessions", url.PathEscape(userId)).
		Query(&filters).
		Get(&userSessions)

	return userSessions, err
}

// GetSourceSessions get valid sessions by sourceID.
func (c *Auth) GetSourceSessions(filters *response.FilterParams, sourceId string) (*response.Page[Session], error) {
	sourceSessions := &response.Page[Session]{}

	_, err := c.api.
		URL("/auth/api/v1/sessionstorage/sources/%s/sessions", url.PathEscape(sourceId)).
		Query(&filters).
		Get(&sourceSessions)

	return sourceSessions, err
}

// SearchSessions searches for sessions
func (c *Auth) SearchSessions(filters *response.FilterParams, search *SessionSearchRequest) (*response.Page[Session], error) {
	sessions := &response.Page[Session]{}

	_, err := c.api.
		URL("/auth/api/v1/sessionstorage/sessions/search").
		Query(&filters).
		Post(search, &sessions)

	return sessions, err
}

// TerminateSession terminates single session by Id.
func (c *Auth) TerminateSession(sessionId string) error {
	_, err := c.api.
		URL("/auth/api/v1/sessionstorage/sessions/%s/terminate", url.PathEscape(sessionId)).
		Post(nil)

	return err
}

// TerminateUserSessions terminates all sessions for a user.
func (store *Auth) TerminateUserSessions(userId string) error {
	_, err := store.api.
		URL("/auth/api/v1/sessionstorage/users/%s/sessions/terminate", url.PathEscape(userId)).
		Post(nil)

	return err
}

// MARK: Users
// Logout log out user.
func (store *Auth) Logout() error {
	_, err := store.api.
		URL("/auth/api/v1/logout").
		Post(nil)

	return err
}

// MARK: Mobile Gateway
// GetUserPairedDevices get users paired devices.
func (store *Auth) GetUserPairedDevices(userId string) (*response.Page[Device], error) {
	devices := &response.Page[Device]{}

	_, err := store.api.
		URL("/auth/api/v1/users/%s/devices", url.PathEscape(userId)).
		Get(devices)

	return devices, err
}

// UnpairUserDevice unpair users device.
func (store *Auth) UnpairUserDevice(userId, deviceId string) error {
	_, err := store.api.
		URL("/auth/api/v1/users/%s/devices/%s", url.PathEscape(userId), url.PathEscape(deviceId)).
		Delete()

	return err
}
