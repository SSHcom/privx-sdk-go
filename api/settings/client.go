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

// SectionSchema get schema for the section
func (store *Settings) SectionSchema(scope, section string) (schema *json.RawMessage, err error) {
	_, err = store.api.
		URL("/settings/api/v1/schema/%s/%s", url.PathEscape(scope), url.PathEscape(section)).
		Get(&schema)

	return
}

// ScopeSchema get schema for the scope
func (store *Settings) ScopeSchema(scope string) (schema *json.RawMessage, err error) {
	_, err = store.api.
		URL("/settings/api/v1/schema/%s", url.PathEscape(scope)).
		Get(&schema)

	return
}

// UpdateScopeSectionSettings update settings for a scope and section combination
func (store *Settings) UpdateScopeSectionSettings(settings *json.RawMessage, scope, section string) error {
	_, err := store.api.
		URL("/settings/api/v1/settings/%s/%s", url.PathEscape(scope), url.PathEscape(section)).
		Put(settings)

	return err
}

// ScopeSectionSettings get settings for the scope
func (store *Settings) ScopeSectionSettings(scope, section string) (settings *json.RawMessage, err error) {
	_, err = store.api.
		URL("/settings/api/v1/settings/%s/%s", url.PathEscape(scope), url.PathEscape(section)).
		Get(&settings)

	return
}

// ScopeSettings get settings for the scope
func (store *Settings) ScopeSettings(scope, merge string) (settings *json.RawMessage, err error) {
	filter := Params{
		Merge: merge,
	}

	_, err = store.api.
		URL("/settings/api/v1/settings/%s", url.PathEscape(scope)).
		Query(&filter).
		Get(&settings)

	return
}

// UpdateScopeSettings update settings for a scope.
func (store *Settings) UpdateScopeSettings(settings *json.RawMessage, scope string) error {
	_, err := store.api.
		URL("/settings/api/v1/settings/%s", url.PathEscape(scope)).
		Put(settings)

	return err
}
