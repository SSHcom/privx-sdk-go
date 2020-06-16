//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package hoststore

import (
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// HostStore is a role-store client instance.
type HostStore struct {
	api restapi.Connector
}

// New creates a new host-store client instance
// See http://apispecs.ssh.com/#swagger-ui-4 for details about api
func New(api restapi.Connector) *HostStore {
	return &HostStore{api: api}
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
