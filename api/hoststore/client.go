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

// New creates a new host-store client instance
// See http://apispecs.ssh.com/#swagger-ui-4 for details about api
func New(api restapi.Connector) *HostStore {
	return &HostStore{api: api}
}

// Host returns existing single host
func (store *HostStore) Host(id string) (host *Host, err error) {
	host = new(Host)

	_, err = store.api.
		URL("/host-store/api/v1/hosts/%s", url.PathEscape(id)).
		Get(host)

	return
}

// RegisterHost append a target to PrivX
func (store *HostStore) RegisterHost(host Host) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/host-store/api/v1/hosts").
		Post(&host, &id)

	return id.ID, err
}
