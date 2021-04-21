//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package monitor

// Params struct for pagination queries.
type Params struct {
	Offset     string `json:"offset,omitempty"`
	Limit      string `json:"limit,omitempty"`
	Sortdir    string `json:"sortdir,omitempty"`
	Sortkey    string `json:"sortkey,omitempty"`
	FuzzyCount bool   `json:"fuzzycount,omitempty"`
}

// AuditEventCodes audit event codes definitions
type AuditEventCodes struct {
	EventCode  int        `json:"key,omitempty"`
	EventValue EventValue `json:"value,omitempty"`
}

// EventValue audit event codes value definitions
type EventValue struct {
	EventID          string `json:"event_id,omitempty"`
	EventName        string `json:"event_name,omitempty"`
	EventDescription string `json:"event_desc,omitempty"`
}

// AuditEvent audit event definitions
type AuditEvent struct {
	ServiceID   string `json:"service_id,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	EventID     string `json:"event_id,omitempty"`
	EventName   string `json:"event_name,omitempty"`
	Created     string `json:"created,omitempty"`
	Message     Message
}

// Message message definitions
type Message struct {
	Message string `json:"message,omitempty"`
}
