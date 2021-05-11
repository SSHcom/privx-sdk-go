//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package monitor

// Params struct for pagination queries.
type Params struct {
	Offset     int    `json:"offset,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Sortdir    string `json:"sortdir,omitempty"`
	Sortkey    string `json:"sortkey,omitempty"`
	FuzzyCount bool   `json:"fuzzycount,omitempty"`
}

type AuditEventSearchObject struct {
	Keywords      string `json:"keywords"`
	UserID        string `json:"user_id"`
	ConnectionID  string `json:"connection_id"`
	HostID        string `json:"host_id"`
	SourceID      string `json:"source_id"`
	SessionID     string `json:"session_id"`
	AccessGroupID string `json:"access_group_id"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
}

// AuditEventCodes audit event codes definitions
type AuditEventCodes map[string]AuditEventInfo

// AuditEventInfo audit event codes value definitions
type AuditEventInfo struct {
	EventID          int    `json:"event_id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_desc"`
}

// AuditEvent audit event definitions
type AuditEvent struct {
	ServiceID   string            `json:"service_id,omitempty"`
	ServiceName string            `json:"service_name,omitempty"`
	EventID     string            `json:"event_id,omitempty"`
	EventName   string            `json:"event_name,omitempty"`
	Created     string            `json:"created,omitempty"`
	Message     map[string]string `json:"message,omitempty"`
}
