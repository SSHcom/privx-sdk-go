//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package config

import (
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// ConfFileStore is a role-store client instance.
type ConfFileStore struct {
	api restapi.Connector
}

// New creates a new client instance to fetch config files from PrivX
func New(api restapi.Connector) *ConfFileStore {
	return &ConfFileStore{api: api}
}

// ConfigExtender fetches configuration file
func (store *ConfFileStore) ConfigExtender(id string) ([]byte, error) {
	var session struct {
		ID string `json:"session_id"`
	}

	_, err := store.api.
		URL("/authorizer/api/v1/extender/conf/%s", id).
		Post("", &session)

	if err != nil {
		return nil, err
	}

	return store.api.
		URL("/authorizer/api/v1/extender/conf/%s/%s", id, session.ID).
		Fetch()
}
