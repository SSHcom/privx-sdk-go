//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package trailindex

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// TrailIndex is a trail index client instance.
type TrailIndex struct {
	api restapi.Connector
}

type trailIndexResult struct {
	Count int             `json:"count"`
	Items []IndexResponse `json:"items"`
}

// New creates a new trail index client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *TrailIndex {
	return &TrailIndex{api: api}
}

// IndexingStatus get indexing status of the connection
func (store *TrailIndex) IndexingStatus(connectionID string) (*Connection, error) {
	status := &Connection{}

	_, err := store.api.
		URL("/trail-index/api/v1/index/%s/status", url.PathEscape(connectionID)).
		Get(status)

	return status, err
}

// IndexingStatuses gets the statuses of the specified connections
func (store *TrailIndex) IndexingStatuses(connectionIDs []string) ([]Connection, error) {
	var connections []Connection

	_, err := store.api.
		URL("/trail-index/api/v1/index/status").
		Post(connectionIDs, &connections)

	return connections, err
}

// StartIndexing starts indexing of the specified connections
func (store *TrailIndex) StartIndexing(connectionIDs []string) ([]Connection, error) {
	var connections []Connection

	_, err := store.api.
		URL("/trail-index/api/v1/index/start").
		Post(connectionIDs, &connections)

	return connections, err
}

// SearchContent search for the content based on the search parameters defined
func (store *TrailIndex) SearchContent(offset, limit int, sortdir string, searchObject SearchRequestObject) ([]IndexResponse, error) {
	result := trailIndexResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortdir: sortdir,
	}

	_, err := store.api.
		URL("/trail-index/api/v1/index/search").
		Query(&filters).
		Post(&searchObject, &result)

	return result.Items, err
}
