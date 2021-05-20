//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

import "time"

// KeyValue key value definition
type KeyValue struct {
	Key   string `json:"k"`
	Value string `json:"v"`
}

// ServiceStatus auth service status definition
type ServiceStatus struct {
	Variant       string     `json:"variant,omitempty"`
	Version       string     `json:"version,omitempty"`
	APIVersion    string     `json:"api_version,omitempty"`
	Status        string     `json:"status,omitempty"`
	StatusMessage string     `json:"status_message,omitempty"`
	ApplicationID string     `json:"app_id,omitempty"`
	ServerMode    string     `json:"server-mode,omitempty"`
	StatusDetails []KeyValue `json:"status_details,omitempty"`
	StartTime     time.Time  `json:"start_time,omitempty"`
}
