//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package trailindex

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// TrailIndex is a trail index client instance.
type TrailIndex struct {
	api restapi.Connector
}

// New trail index client constructor.
func New(api restapi.Connector) *TrailIndex {
	return &TrailIndex{api: api}
}

// MARK: Status
// Status get trail index microservice status.
func (c *TrailIndex) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/trail-index/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Index
// GetIndexingStatus get indexing status.
func (c *TrailIndex) GetIndexingStatus(connectionID string) (ConnectionTranscriptStatus, error) {
	status := ConnectionTranscriptStatus{}

	_, err := c.api.
		URL("/trail-index/api/v1/index/%s/status", connectionID).
		Get(&status)

	return status, err
}

// GetIndexingStatuses gets the statuses of the specified connections.
func (c *TrailIndex) GetIndexingStatuses(connectionIDs []string) ([]ConnectionTranscriptStatus, error) {
	statuses := []ConnectionTranscriptStatus{}

	_, err := c.api.
		URL("/trail-index/api/v1/index/status").
		Post(connectionIDs, &statuses)

	return statuses, err
}

// StartIndexing starts indexing of the specified connections.
func (c *TrailIndex) StartIndexing(connectionIDs []string) ([]ConnectionTranscriptStatus, error) {
	statuses := []ConnectionTranscriptStatus{}

	_, err := c.api.
		URL("/trail-index/api/v1/index/start").
		Post(connectionIDs, &statuses)

	return statuses, err
}

// SearchIndexes search for the content based on the search parameters defined.
func (c *TrailIndex) SearchIndexes(search *TranscriptSearch, opts ...filters.Option) (*response.ResultSet[TrailIndexResponse], error) {
	indexes := &response.ResultSet[TrailIndexResponse]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/trail-index/api/v1/index/search").
		Query(params).
		Post(&search, &indexes)

	return indexes, err
}
