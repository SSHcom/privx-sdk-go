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
