//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package connectionmanager

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// ConnectionManager is a connection manager client instance.
type ConnectionManager struct {
	api restapi.Connector
}

type connectionsResult struct {
	Count int          `json:"count"`
	Items []Connection `json:"items"`
}

// New creates a new connection manager client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *ConnectionManager {
	return &ConnectionManager{api: api}
}

// Connections get all connections
func (store *ConnectionManager) Connections(offset, limit int, sortkey, sortdir string) ([]Connection, error) {
	result := connectionsResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/connections").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// SearchConnections search for connections
func (store *ConnectionManager) SearchConnections(offset, limit int, sortdir, sortkey string, searchObject ConnectionSearch) ([]Connection, error) {
	result := connectionsResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortdir: sortdir,
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/connections/search").
		Query(&filters).
		Post(&searchObject, &result)

	return result.Items, err
}

// Connection get a single connection
func (store *ConnectionManager) Connection(connID string) (*Connection, error) {
	conn := &Connection{}

	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s", url.PathEscape(connID)).
		Get(conn)

	return conn, err
}

// CreateSessionIDFileDownload create session ID for trail stored file download
func (store *ConnectionManager) CreateSessionIDFileDownload(connID, chanID, fileID string) (string, error) {
	var object struct {
		SessionID string `json:"session_id"`
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/file/%s",
			url.PathEscape(connID), url.PathEscape(chanID), url.PathEscape(fileID)).
		Post(nil, &object)

	return object.SessionID, err
}

// DownloadStoredFile download trail stored file transferred within audited connection channel
func (store *ConnectionManager) DownloadStoredFile(connID, chanID, fileID, sessionID, filename string) error {
	err := store.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/file/%s/%s",
			url.PathEscape(connID), url.PathEscape(chanID), url.PathEscape(fileID), url.PathEscape(sessionID)).
		Download(filename)

	return err
}

// CreateSessionIDTrailLog create session ID for trail log download
func (store *ConnectionManager) CreateSessionIDTrailLog(connID, chanID string) (string, error) {
	var object struct {
		SessionID string `json:"session_id"`
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/log",
			url.PathEscape(connID), url.PathEscape(chanID)).
		Post(nil, &object)

	return object.SessionID, err
}

// DownloadTrailLog download trail log of audited connection channel
func (store *ConnectionManager) DownloadTrailLog(connID, chanID, sessionID, format, filter, filename string) error {
	filters := Params{
		Format: format,
		Filter: filter,
	}

	err := store.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/log/%s",
			url.PathEscape(connID), url.PathEscape(chanID), url.PathEscape(sessionID)).
		Query(&filters).
		Download(filename)

	return err
}

// SavedAccessRoles get saved access roles for a connection
func (store *ConnectionManager) SavedAccessRoles(connID string) ([]AccessRoles, error) {
	var result []AccessRoles

	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s/access_roles", url.PathEscape(connID)).
		Get(&result)

	return result, err
}

// GrantRolePermission grant a role permission for a connection
func (store *ConnectionManager) GrantRolePermission(connID, roleID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s/access_roles/%s",
			url.PathEscape(connID), url.PathEscape(roleID)).
		Post(nil)

	return err
}

// RevokeRolePermission revoke a permission for a role from a connection
func (store *ConnectionManager) RevokeRolePermission(connID, roleID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s/access_roles/%s",
			url.PathEscape(connID), url.PathEscape(roleID)).
		Delete()

	return err
}

// RevokeRoleFromConnection revoke permissions for a role from connections
func (store *ConnectionManager) RevokeRoleFromConnection(roleID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/connections/access_roles/%s",
			url.PathEscape(roleID)).
		Delete()

	return err
}

// TerminateConnection terminate connection by ID.
func (store *ConnectionManager) TerminateConnection(connID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/terminate/connection/%s", url.PathEscape(connID)).
		Post(nil)

	return err
}

// TerminateConnectionFromHost terminate connection(s) from host
func (store *ConnectionManager) TerminateConnectionFromHost(hostID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/terminate/host/%s", url.PathEscape(hostID)).
		Post(nil)

	return err
}

// TerminateConnectionFromUser terminate connection(s) of a user
func (store *ConnectionManager) TerminateConnectionFromUser(userID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/terminate/user/%s", url.PathEscape(userID)).
		Post(nil)

	return err
}
