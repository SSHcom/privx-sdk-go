//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package monitor

import (
	"encoding/json"
	"net/url"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Monitor is a monitor service client instance.
type Monitor struct {
	api restapi.Connector
}

// EventsResult list of event results
type EventsResult struct {
	Count int          `json:"count"`
	Items []AuditEvent `json:"items"`
}

// New creates a new monitor service client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *Monitor {
	return &Monitor{api: api}
}

// ComponentsStatus get the status of all deployed privx components
func (store *Monitor) ComponentsStatus() (*json.RawMessage, error) {
	status := &json.RawMessage{}

	_, err := store.api.
		URL("/monitor-service/api/v1/components").
		Get(&status)

	return status, err
}

// HostComponentStatus get component status object by hostname.
func (store *Monitor) ComponentStatus(hostname string) (*json.RawMessage, error) {
	status := &json.RawMessage{}

	_, err := store.api.
		URL("/monitor-service/api/v1/components/%s", url.PathEscape(hostname)).
		Get(&status)

	return status, err
}

// SearchAuditEvents search for audit events
func (store *Monitor) SearchAuditEvents(offset, limit int, keywords, sortkey, sortdir string, fuzzycount bool) (*EventsResult, error) {
	result := &EventsResult{}
	filters := Params{
		Offset:     offset,
		Limit:      limit,
		Sortkey:    sortkey,
		Sortdir:    sortdir,
		FuzzyCount: fuzzycount,
	}

	_, err := store.api.
		URL("/monitor-service/api/v1/auditevents/search").
		Query(&filters).
		Post(map[string]string{
			"keywords": keywords,
		}, &result)

	return result, err
}

// AuditEvents get all audit events
func (store *Monitor) AuditEvents(offset, limit int, sortkey, sortdir string, fuzzycount bool) (*EventsResult, error) {
	result := &EventsResult{}
	filters := Params{
		Offset:     offset,
		Limit:      limit,
		Sortdir:    sortdir,
		Sortkey:    sortkey,
		FuzzyCount: fuzzycount,
	}

	_, err := store.api.
		URL("/monitor-service/api/v1/auditevents").
		Query(&filters).
		Get(&result)

	return result, err
}

// AuditEventCodes get audit event codes
func (store *Monitor) AuditEventCodes() (*AuditEventCodes, error) {
	codes := &AuditEventCodes{}

	_, err := store.api.
		URL("/monitor-service/api/v1/auditevents/codes").
		Get(&codes)

	return codes, err
}

// InstanceStatus status of the whole instance
func (store *Monitor) InstanceStatus() (*json.RawMessage, error) {
	status := &json.RawMessage{}

	_, err := store.api.
		URL("/monitor-service/api/v1/instance/status").
		Get(&status)

	return status, err
}

// TerminateInstances terminate PrivX instances
func (store *Monitor) TerminateInstances() error {
	_, err := store.api.
		URL("/monitor-service/api/v1/instance/exit").
		Post(nil)

	return err
}
