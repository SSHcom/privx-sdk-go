//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package config

import (
	"github.com/SSHcom/privx-sdk-go/v2/restapi"
)

// ConfFileStore is a role-store client instance.
type ConfFileStore struct {
	api restapi.Connector
}

type tSession struct {
	ID string `json:"session_id"`
}

// New creates a new client instance to fetch config files from PrivX
func New(api restapi.Connector) *ConfFileStore {
	return &ConfFileStore{api: api}
}

// ConfigExtender fetches configuration file
func (store *ConfFileStore) ConfigExtender(id string) ([]byte, error) {
	return store.config("extender/conf", id)
}

// ConfigDeploy fetches deployment script
func (store *ConfFileStore) ConfigDeploy(id string) ([]byte, error) {
	return store.config("deploy", id)
}

func (store *ConfFileStore) config(config, id string) ([]byte, error) {
	var session tSession

	_, err := store.api.
		URL("/authorizer/api/v1/%s/%s", config, id).
		Post("", &session)

	if err != nil {
		return nil, err
	}

	return store.api.
		URL("/authorizer/api/v1/%s/%s/%s", config, id, session.ID).
		Fetch()
}
