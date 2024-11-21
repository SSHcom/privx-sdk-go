//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package monitor

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Monitor is a monitor service client instance.
type Monitor struct {
	api restapi.Connector
}

// New monitor service client constructor.
func New(api restapi.Connector) *Monitor {
	return &Monitor{api: api}
}

// MARK: Status
// Status get monitor service microservice status.
func (c *Monitor) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/monitor-service/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Audit Events
// SearchAuditEvents search audit events.
func (c *Monitor) SearchAuditEvents(search *AuditEventSearch, opts ...filters.Option) (*response.ResultSet[AuditEvent], error) {
	events := &response.ResultSet[AuditEvent]{}
	params := url.Values{}

	// Define default start time if empty (similar on how UI is handling start time).
	if search != nil && search.StartTime == "" {
		search = &AuditEventSearch{
			StartTime: time.Now().UTC().AddDate(0, 0, -8).Format("2006-01-02T15:04:05.000Z"),
		}
	}

	// Set default options, which will be overwritten by opts if defined.
	options := append([]filters.Option{
		filters.Paging(0, 25),
		filters.Sort("created", "DESC"),
		filters.FuzzyCount(true),
	}, opts...)

	for _, opt := range options {
		opt(&params)
	}

	_, err := c.api.
		URL("/monitor-service/api/v1/auditevents/search").
		Query(params).
		Post(&search, &events)

	return events, err
}

// GetAuditEvents get audit events.
func (c *Monitor) GetAuditEvents(opts ...filters.Option) (*response.ResultSet[AuditEvent], error) {
	events := &response.ResultSet[AuditEvent]{}
	params := url.Values{}

	// Set default options, which will be overwritten by opts if defined.
	options := append([]filters.Option{
		filters.Paging(0, 25),
		filters.Sort("created", "DESC"),
		filters.FuzzyCount(true),
	}, opts...)

	for _, opt := range options {
		opt(&params)
	}

	_, err := c.api.
		URL("/monitor-service/api/v1/auditevents").
		Query(params).
		Get(&events)

	return events, err
}

// GetAuditEventCodes get audit event codes.
func (c *Monitor) GetAuditEventCodes() (*AuditEventCodes, error) {
	codes := &AuditEventCodes{}

	_, err := c.api.
		URL("/monitor-service/api/v1/auditevents/codes").
		Get(&codes)

	return codes, err
}

// MARK: Components
// GetComponentsStatus get components status.
func (c *Monitor) GetComponentsStatus() (*json.RawMessage, error) {
	status := &json.RawMessage{}

	_, err := c.api.
		URL("/monitor-service/api/v1/components").
		Get(&status)

	return status, err
}

// GetComponentStatus get component status by hostname.
func (c *Monitor) GetComponentStatus(hostname string) (*json.RawMessage, error) {
	status := &json.RawMessage{}

	_, err := c.api.
		URL("/monitor-service/api/v1/components/%s", hostname).
		Get(&status)

	return status, err
}

// MARK: Instance
// GetInstanceStatus get PrivX instance status.
func (c *Monitor) GetInstanceStatus() (*json.RawMessage, error) {
	status := &json.RawMessage{}

	_, err := c.api.
		URL("/monitor-service/api/v1/instance/status").
		Get(&status)

	return status, err
}

// TerminateInstances terminate PrivX instances.
func (c *Monitor) TerminateInstances() error {
	_, err := c.api.
		URL("/monitor-service/api/v1/instance/exit").
		Post(nil)

	return err
}

// MARK: Time
// GetServerTime get current PrivX server time.
func (c *Monitor) GetServerTime() (Clock, error) {
	clock := Clock{}

	_, err := c.api.
		URL("/monitor-service/api/v1/time").
		Get(&clock)

	return clock, err
}
