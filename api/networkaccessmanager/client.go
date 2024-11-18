//
// Copyright (c) 2022 SSH Communications Security Inc.
//
// All rights reserved.
//

package networkaccessmanager

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// NetworkAccessManager is a network access manager client instance.
type NetworkAccessManager struct {
	api restapi.Connector
}

// New network access manager client constructor.
func New(api restapi.Connector) *NetworkAccessManager {
	return &NetworkAccessManager{api: api}
}

// MARK: Status
// Status get network access manager microservice status.
func (c *NetworkAccessManager) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/network-access-manager/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Network Targets
// GetNetworkTargets get network targets.
func (c *NetworkAccessManager) GetNetworkTargets(opts ...filters.Option) (*response.ResultSet[NetworkTarget], error) {
	targets := &response.ResultSet[NetworkTarget]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/network-access-manager/api/v1/nwtargets").
		Query(params).
		Get(&targets)

	return targets, err
}

// CreateNetworkTarget create network target.
func (c *NetworkAccessManager) CreateNetworkTarget(target *NetworkTarget) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/network-access-manager/api/v1/nwtargets").
		Post(&target, &identifier)

	return identifier, err
}

// SearchNetworkTargets search network target.
func (c *NetworkAccessManager) SearchNetworkTargets(search NetworkTargetSearch, opts ...filters.Option) (*response.ResultSet[NetworkTarget], error) {
	targets := &response.ResultSet[NetworkTarget]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/network-access-manager/api/v1/nwtargets/search").
		Query(params).
		Post(&search, &targets)

	return targets, err
}

// GetNetworkTargetTags get network target tags.
func (c *NetworkAccessManager) GetNetworkTargetTags(opts ...filters.Option) (*response.ResultSet[string], error) {
	tags := &response.ResultSet[string]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/network-access-manager/api/v1/nwtargets/tags").
		Query(params).
		Get(&tags)

	return tags, err
}

// GetNetworkTarget get network target by id.
func (c *NetworkAccessManager) GetNetworkTarget(targetID string) (*NetworkTarget, error) {
	target := &NetworkTarget{}

	_, err := c.api.
		URL("/network-access-manager/api/v1/nwtargets/%s", targetID).
		Get(&target)

	return target, err
}

// UpdateNetworkTarget update network target.
func (c *NetworkAccessManager) UpdateNetworkTarget(targetID string, target *NetworkTarget) error {
	_, err := c.api.
		URL("/network-access-manager/api/v1/nwtargets/%s", targetID).
		Put(&target)

	return err
}

// DeleteNetworkTarget delete network target by id.
func (c *NetworkAccessManager) DeleteNetworkTarget(targetID string) error {
	_, err := c.api.
		URL("/network-access-manager/api/v1/nwtargets/%s", targetID).
		Delete()

	return err
}

// DisableNetworkTarget disable network target by id.
func (c *NetworkAccessManager) DisableNetworkTarget(targetID string, disable NetworkTargetDisable) error {
	_, err := c.api.
		URL("/network-access-manager/api/v1/nwtargets/%s/disabled", targetID).
		Put(&disable)

	return err
}
