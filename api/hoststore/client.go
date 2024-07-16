//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package hoststore

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// HostStore is a role-store client instance.
type HostStore struct {
	api restapi.Connector
}

type hostResult struct {
	Count int    `json:"count"`
	Items []Host `json:"items"`
}

type tagsResult struct {
	Count int      `json:"count"`
	Items []string `json:"items"`
}

// New creates a new host-store client instance
// See http://apispecs.ssh.com/#swagger-ui-4 for details about api
func New(api restapi.Connector) *HostStore {
	return &HostStore{api: api}
}

// SearchHost search for existing hosts
func (store *HostStore) SearchHost(sortkey, sortdir, filter string, offset, limit int, searchObject *HostSearchObject) ([]Host, error) {
	result := hostResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
		Filter:  filter,
	}

	_, err := store.api.
		URL("/host-store/api/v1/hosts/search").
		Query(&filters).
		Post(&searchObject, &result)

	return result.Items, err
}

// Hosts returns existing hosts
func (store *HostStore) Hosts(offset, limit int, sortkey, sortdir, filter string) ([]Host, error) {
	result := hostResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
		Filter:  filter,
	}

	_, err := store.api.
		URL("/host-store/api/v1/hosts").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// CreateHost create a host to host store
func (store *HostStore) CreateHost(host Host) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/host-store/api/v1/hosts").
		Post(&host, &object)

	return object.ID, err
}

// ResolveHost resolve service and address to a single host in host store
func (store *HostStore) ResolveHost(service Service) (*Host, error) {
	host := &Host{}

	_, err := store.api.
		URL("/host-store/api/v1/hosts/resolve").
		Post(&service, &host)

	return host, err
}

// Host returns existing single host
func (store *HostStore) Host(hostID string) (*Host, error) {
	host := &Host{}

	_, err := store.api.
		URL("/host-store/api/v1/hosts/%s", url.PathEscape(hostID)).
		Get(&host)

	return host, err
}

// UpdateHost update existing host
func (store *HostStore) UpdateHost(hostID string, host *Host) error {
	_, err := store.api.
		URL("/host-store/api/v1/hosts/%s", url.PathEscape(hostID)).
		Put(host)

	return err
}

// DeleteHost delete a host
func (store *HostStore) DeleteHost(hostID string) error {
	_, err := store.api.
		URL("/host-store/api/v1/hosts/%s", hostID).
		Delete()

	return err
}

// UpdateDeployStatus update host to be deployable or undeployable
func (store *HostStore) UpdateDeployStatus(hostID string, status bool) error {
	deployStatus := Host{
		Deployable: status,
	}

	_, err := store.api.
		URL("/host-store/api/v1/hosts/%s/deployable", url.PathEscape(hostID)).
		Put(deployStatus)

	return err
}

// HostTags returns host tags
func (store *HostStore) HostTags(offset, limit int, sortdir, query string) ([]string, error) {
	result := tagsResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortdir: sortdir,
		Query:   query,
	}

	_, err := store.api.
		URL("/host-store/api/v1/hosts/tags").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// UpdateDisabledHostStatus enable/disable host
func (store *HostStore) UpdateDisabledHostStatus(hostID string, status bool) error {
	disabledStatus := HostDisabledRequest{
		Disabled: status,
	}

	_, err := store.api.
		URL("/host-store/api/v1/hosts/%s/disabled", url.PathEscape(hostID)).
		Put(disabledStatus)

	return err
}

// ServiceOptions returns default service options
func (store *HostStore) ServiceOptions() (*DefaultServiceOptions, error) {
	options := &DefaultServiceOptions{}

	_, err := store.api.
		URL("/host-store/api/v1/settings/default_service_options").
		Get(&options)

	return options, err
}
