//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package settings

import (
	"encoding/json"
	"net/url"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Settings is a settings client instance.
type Settings struct {
	api restapi.Connector
}

// New creates a new settings client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *Settings {
	return &Settings{api: api}
}

// ScopeSettings get settings for the scope
func (store *Settings) ScopeSettings(scope, merge string) (*json.RawMessage, error) {
	settings := &json.RawMessage{}
	filter := Params{
		Merge: merge,
	}

	_, err := store.api.
		URL("/settings/api/v1/settings/%s", url.PathEscape(scope)).
		Query(&filter).
		Get(&settings)

	return settings, err
}

// UpdateScopeSettings update settings for a scope.
func (store *Settings) UpdateScopeSettings(settings *json.RawMessage, scope string) error {
	_, err := store.api.
		URL("/settings/api/v1/settings/%s", url.PathEscape(scope)).
		Put(settings)

	return err
}

// ScopeSectionSettings get settings for the scope
func (store *Settings) ScopeSectionSettings(scope, section string) (*json.RawMessage, error) {
	settings := &json.RawMessage{}
	_, err := store.api.
		URL("/settings/api/v1/settings/%s/%s", url.PathEscape(scope), url.PathEscape(section)).
		Get(&settings)

	return settings, err
}

// UpdateScopeSectionSettings update settings for a scope and section combination
func (store *Settings) UpdateScopeSectionSettings(settings *json.RawMessage, scope, section string) error {
	_, err := store.api.
		URL("/settings/api/v1/settings/%s/%s", url.PathEscape(scope), url.PathEscape(section)).
		Put(settings)

	return err
}

// ScopeSchema get schema for the scope
func (store *Settings) ScopeSchema(scope string) (*json.RawMessage, error) {
	schema := &json.RawMessage{}
	_, err := store.api.
		URL("/settings/api/v1/schema/%s", url.PathEscape(scope)).
		Get(&schema)

	return schema, err
}

// SectionSchema get schema for the section
func (store *Settings) SectionSchema(scope, section string) (*json.RawMessage, error) {
	schema := &json.RawMessage{}
	_, err := store.api.
		URL("/settings/api/v1/schema/%s/%s", url.PathEscape(scope), url.PathEscape(section)).
		Get(&schema)

	return schema, err
}

// RestartRequired get restart required information for given settings and scope
func (store *Settings) RestartRequired(settings *json.RawMessage, scope string) (*json.RawMessage, error) {
	isRestartRequiredResponse := &json.RawMessage{}

	_, err := store.api.
		URL("/settings/api/v1/restart_required/%s", url.PathEscape(scope)).
		Post(settings, isRestartRequiredResponse)

	return isRestartRequiredResponse, err
}
