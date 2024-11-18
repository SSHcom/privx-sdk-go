//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package settings

import (
	"encoding/json"
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Settings is a settings client instance.
type Settings struct {
	api restapi.Connector
}

// New settings client constructor.
func New(api restapi.Connector) *Settings {
	return &Settings{api: api}
}

// MARK: Status
// Status get settings microservice status.
func (c *Settings) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/settings/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Schema
// GetScopeSchema get schema for the scope.
func (c *Settings) GetScopeSchema(scope string) (*json.RawMessage, error) {
	schema := &json.RawMessage{}
	_, err := c.api.
		URL("/settings/api/v1/schema/%s", scope).
		Get(&schema)

	return schema, err
}

// GetSectionSchema get schema for the section
func (c *Settings) GetSectionSchema(scope, section string) (*json.RawMessage, error) {
	schema := &json.RawMessage{}
	_, err := c.api.
		URL("/settings/api/v1/schema/%s/%s", scope, section).
		Get(&schema)

	return schema, err
}

// MARK: Settings
// GetScopeSettings get settings for the scope
func (c *Settings) GetScopeSettings(scope string, opts ...filters.Option) (*json.RawMessage, error) {
	settings := &json.RawMessage{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}
	_, err := c.api.
		URL("/settings/api/v1/settings/%s", scope).
		Query(params).
		Get(&settings)

	return settings, err
}

// UpdateScopeSettings update settings for a scope.
func (c *Settings) UpdateScopeSettings(settings map[string]interface{}, scope string) error {
	_, err := c.api.
		URL("/settings/api/v1/settings/%s", scope).
		Put(&settings)

	return err
}

// GetSectionSettings get settings for the section.
func (c *Settings) GetSectionSettings(scope, section string) (*json.RawMessage, error) {
	settings := &json.RawMessage{}

	_, err := c.api.
		URL("/settings/api/v1/settings/%s/%s", scope, section).
		Get(&settings)

	return settings, err
}

// UpdateSectionSettings update settings for a scope and section combination.
func (c *Settings) UpdateSectionSettings(scope, section string, settings map[string]interface{}) error {
	_, err := c.api.
		URL("/settings/api/v1/settings/%s/%s", scope, section).
		Put(&settings)

	return err
}

// MARK: Restart Required
// VerifyRestartRequired verify if restart is required for given settings scope.
func (c *Settings) VerifyRestartRequired(scope string, settings map[string]interface{}) (*map[string]interface{}, error) {
	verification := &map[string]interface{}{}

	_, err := c.api.
		URL("/settings/api/v1/restart_required/%s", scope).
		Post(&settings, &verification)

	return verification, err
}
