//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package hoststore

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// HostStore is a host store client instance.
type HostStore struct {
	api restapi.Connector
}

// New host store client constructor.
func New(api restapi.Connector) *HostStore {
	return &HostStore{api: api}
}

// MARK: Status
// Status get host store microservice status.
func (c *HostStore) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/host-store/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Hosts
// SearchHosts search hosts.
func (c *HostStore) SearchHosts(search *HostSearch, opts ...filters.Option) (*response.ResultSet[Host], error) {
	hosts := &response.ResultSet[Host]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/host-store/api/v1/hosts/search").
		Query(params).
		Post(&search, &hosts)

	return hosts, err
}

// GetHosts get hosts.
func (c *HostStore) GetHosts(opts ...filters.Option) (*response.ResultSet[Host], error) {
	hosts := &response.ResultSet[Host]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/host-store/api/v1/hosts").
		Query(params).
		Get(&hosts)

	return hosts, err
}

// CreateHost create a host.
func (c *HostStore) CreateHost(host *Host) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/host-store/api/v1/hosts").
		Post(&host, &identifier)

	return identifier, err
}

// ResolveHost resolve service to a single host.
func (c *HostStore) ResolveHost(resolve HostResolve) (*Host, error) {
	host := &Host{}

	_, err := c.api.
		URL("/host-store/api/v1/hosts/resolve").
		Post(&resolve, &host)

	return host, err
}

// GetHost get host by id.
func (c *HostStore) GetHost(hostID string) (*Host, error) {
	host := &Host{}

	_, err := c.api.
		URL("/host-store/api/v1/hosts/%s", hostID).
		Get(&host)

	return host, err
}

// UpdateHost update host.
func (c *HostStore) UpdateHost(hostID string, host *Host) error {
	_, err := c.api.
		URL("/host-store/api/v1/hosts/%s", hostID).
		Put(&host)

	return err
}

// DeleteHost delete host.
func (c *HostStore) DeleteHost(hostID string) error {
	_, err := c.api.
		URL("/host-store/api/v1/hosts/%s", hostID).
		Delete()

	return err
}

// DeployHost deploy host.
func (c *HostStore) DeployHost(host *Host) (HostResponse, error) {
	response := HostResponse{}

	_, err := c.api.
		URL("/host-store/api/v1/hosts/deploy").
		Post(&host, &response)

	return response, err
}

// UpdateDeployStatus update host to be deployable or undeployable.
func (c *HostStore) UpdateDeployStatus(hostID string, deployable bool) error {
	d := HostDeployable{
		Deployable: deployable,
	}

	_, err := c.api.
		URL("/host-store/api/v1/hosts/%s/deployable", hostID).
		Put(&d)

	return err
}

// GetHostTags get host tags.
func (c *HostStore) GetHostTags(opts ...filters.Option) (*response.ResultSet[string], error) {
	tags := &response.ResultSet[string]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/host-store/api/v1/hosts/tags").
		Query(params).
		Get(&tags)

	return tags, err
}

// UpdateHostStatus enable/disable host.
func (c *HostStore) UpdateHostStatus(hostID string, disabled bool) error {
	d := HostDisabled{
		Disabled: disabled,
	}

	_, err := c.api.
		URL("/host-store/api/v1/hosts/%s/disabled", hostID).
		Put(&d)

	return err
}

// MARK: Settings
// GetServiceOptions get default service options.
func (c *HostStore) GetServiceOptions() (*HostServiceOptions, error) {
	options := &HostServiceOptions{}

	_, err := c.api.
		URL("/host-store/api/v1/settings/default_service_options").
		Get(&options)

	return options, err
}

// MARK: WhiteLists
// GetWhitelists get whitelists.
func (c *HostStore) GetWhitelists(opts ...filters.Option) (*response.ResultSet[Whitelist], error) {
	result := &response.ResultSet[Whitelist]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}
	_, err := c.api.
		URL("/host-store/api/v1/whitelists").
		Query(params).
		Get(&result)

	return result, err
}

// CreateWhitelist create whitelist.
func (c *HostStore) CreateWhitelist(whitelist *Whitelist) (response.Identifier, error) {
	identifier := response.Identifier{}
	_, err := c.api.
		URL("/host-store/api/v1/whitelists").
		Post(&whitelist, &identifier)

	return identifier, err
}

// GetWhitelist get whitelist by id.
func (c *HostStore) GetWhitelist(whitelistID string) (*Whitelist, error) {
	whitelist := &Whitelist{}
	_, err := c.api.
		URL("/host-store/api/v1/whitelists/%s", whitelistID).
		Get(&whitelist)

	return whitelist, err
}

// UpdateWhitelist update whitelist.
func (c *HostStore) UpdateWhitelist(whitelistID string, whitelist Whitelist) error {
	_, err := c.api.
		URL("/host-store/api/v1/whitelists/%s", whitelistID).
		Put(&whitelist)

	return err
}

// DeleteWhitelist delete whitelist.
func (c *HostStore) DeleteWhitelist(whitelistID string) error {
	_, err := c.api.
		URL("/host-store/api/v1/whitelists/%s", whitelistID).
		Delete()

	return err
}

// SearchWhitelists search whitelists.
func (c *HostStore) SearchWhitelists(search WhitelistSearch, opts ...filters.Option) (*response.ResultSet[Whitelist], error) {
	whitelists := &response.ResultSet[Whitelist]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/host-store/api/v1/whitelists/search").
		Query(params).
		Post(&search, &whitelists)

	return whitelists, err
}

// EvaluateWhitelist evaluate commands against whitelist patterns.
func (c *HostStore) EvaluateWhitelist(evaluate *WhitelistEvaluate) (*WhitelistEvaluateResponse, error) {
	result := &WhitelistEvaluateResponse{}
	_, err := c.api.
		URL("/host-store/api/v1/whitelists/evaluate").
		Post(&evaluate, &result)

	return result, err
}
