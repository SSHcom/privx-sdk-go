//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

import (
	"github.com/SSHcom/privx-sdk-go/api/filters"
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
func (c *Auth) CreateIdpClient(idpClient *IdpClient) (response.Identifier, error) {
	idpClientResponse := response.Identifier{}

	_, err := c.api.
		URL("/auth/api/v1/idp/clients").
		Post(&idpClient, &idpClientResponse)

	return idpClientResponse, err
}

// UpdateIdpClient updates existing identity provider client configuration definition.
func (c *Auth) UpdateIdpClient(idpClient *IdpClient, idpId string) error {
	_, err := c.api.
		URL("/auth/api/v1/idp/clients/%s", idpId).
		Put(&idpClient)

	return err
}

// GetIdpClient get existing identity provider client configuration.
func (c *Auth) GetIdpClient(idpId string) (*IdpClient, error) {
	idpClient := &IdpClient{}

	_, err := c.api.
		URL("/auth/api/v1/idp/clients/%s", idpId).
		Get(&idpClient)

	return idpClient, err
}

// DeleteIdpClient delete identity provider client configuration by Id.
func (c *Auth) DeleteIdpClient(idpId string) error {
	_, err := c.api.
		URL("/auth/api/v1/idp/clients/%s", idpId).
		Delete()

	return err
}

// RegenerateIdpClientConfig regenerates client_id and client_secret
// for OIDC identity provider client configuration.
func (c *Auth) RegenerateIdpClientConfig(idpId string) (*IdpClientConfig, error) {
	clientConfig := &IdpClientConfig{}

	_, err := c.api.
		URL("/auth/api/v1/idp/clients/%s/regenerate", idpId).
		Post(nil, &clientConfig)

	return clientConfig, err
}

// MARK: Session Storage
// GetUserSessions get valid sessions by userID.
func (c *Auth) GetUserSessions(userId string, opts ...filters.Option) (*response.ResultSet[Session], error) {
	userSessions := &response.ResultSet[Session]{}
	params := filters.Default()

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/auth/api/v1/sessionstorage/users/%s/sessions", userId).
		Query(params).
		Get(&userSessions)

	return userSessions, err
}

// GetSourceSessions get valid sessions by sourceID.
func (c *Auth) GetSourceSessions(sourceId string, opts ...filters.Option) (*response.ResultSet[Session], error) {
	sourceSessions := &response.ResultSet[Session]{}
	params := filters.Default()

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/auth/api/v1/sessionstorage/sources/%s/sessions", sourceId).
		Query(params).
		Get(&sourceSessions)

	return sourceSessions, err
}

// SearchSessions searches for sessions
func (c *Auth) SearchSessions(search *SessionSearch, opts ...filters.Option) (*response.ResultSet[Session], error) {
	sessions := &response.ResultSet[Session]{}
	params := filters.Default()

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/auth/api/v1/sessionstorage/sessions/search").
		Query(params).
		Post(search, &sessions)

	return sessions, err
}

// TerminateSession terminates single session by Id.
func (c *Auth) TerminateSession(sessionId string) error {
	_, err := c.api.
		URL("/auth/api/v1/sessionstorage/sessions/%s/terminate", sessionId).
		Post(nil)

	return err
}

// TerminateUserSessions terminates all sessions for a user.
func (store *Auth) TerminateUserSessions(userId string) error {
	_, err := store.api.
		URL("/auth/api/v1/sessionstorage/users/%s/sessions/terminate", userId).
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
func (store *Auth) GetUserPairedDevices(userId string) (*response.ResultSet[Device], error) {
	devices := &response.ResultSet[Device]{}

	_, err := store.api.
		URL("/auth/api/v1/users/%s/devices", userId).
		Get(devices)

	return devices, err
}

// UnpairUserDevice unpair users device.
func (store *Auth) UnpairUserDevice(userId, deviceId string) error {
	_, err := store.api.
		URL("/auth/api/v1/users/%s/devices/%s", userId, deviceId).
		Delete()

	return err
}
