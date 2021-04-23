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

type eventsResult struct {
	Count int          `json:"count"`
	Items []AuditEvent `json:"items"`
}

// New creates a new monitor service client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *Monitor {
	return &Monitor{api: api}
}

// TerminateInstances terminate PrivX instances
func (store *Monitor) TerminateInstances() error {
	_, err := store.api.
		URL("/monitor-service/api/v1/instance/exit").
		Post(nil)

	return err
}

// InstanceStatus status of the whole instance
func (store *Monitor) InstanceStatus() (status *json.RawMessage, err error) {
	_, err = store.api.
		URL("/monitor-service/api/v1/instance/status").
		Get(&status)

	return
}

// AuditEventCodes get audit event codes
func (store *Monitor) AuditEventCodes() (codes *AuditEventCodes, err error) {
	codes = new(AuditEventCodes)

	_, err = store.api.
		URL("/monitor-service/api/v1/auditevents/codes").
		Get(codes)

	return
}

// SearchAuditEvents search for audit events
func (store *Monitor) SearchAuditEvents(keywords, offset, limit, sortkey, sortdir string, fuzzycount bool) ([]AuditEvent, error) {
	result := eventsResult{}
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

	return result.Items, err
}

// AuditEvents get all audit events
func (store *Monitor) AuditEvents(offset, limit, sortkey, sortdir string, fuzzycount bool) ([]AuditEvent, error) {
	result := eventsResult{}
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

	return result.Items, err
}

// ComponentsStatus get the status of all deployed privx components
func (store *Monitor) ComponentsStatus() (status *json.RawMessage, err error) {
	_, err = store.api.
		URL("/monitor-service/api/v1/components").
		Get(&status)

	return
}

// HostComponentStatus get component status object by hostname.
func (store *Monitor) HostComponentStatus(hostname string) (status *json.RawMessage, err error) {
	_, err = store.api.
		URL("/monitor-service/api/v1/components/%s", url.PathEscape(hostname)).
		Get(&status)

	return
}
