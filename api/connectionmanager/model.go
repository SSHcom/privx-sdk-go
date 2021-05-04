//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package connectionmanager

// Params query paramas definition
type Params struct {
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Sortdir string `json:"sortdir,omitempty"`
	Sortkey string `json:"sortkey,omitempty"`
	Format  string `json:"format,omitempty"`
	Filter  string `json:"filter,omitempty"`
}

// ConnectionHost connection host struct definition
type ConnectionHost struct {
	ID         string `json:"id,omitempty"`
	CommonName string `json:"common_name,omitempty"`
}

// ConnectionRole connection role struct definition
type ConnectionRole struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// UserData user data struct definition
type UserData struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"display_name,omitempty"`
}

// AccessRoles access roles struct definition
type AccessRoles struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Added string `json:"added"`
}

// Connection connection struct definition
type Connection struct {
	ID                string           `json:"id,omitempty"`
	ProxyID           string           `json:"proxy_id,omitempty"`
	Type              string           `json:"type,omitempty"`
	UserAgent         string           `json:"user_agent,omitempty"`
	TargetHostAddress string           `json:"target_host_address,omitempty"`
	TargetHostAccount string           `json:"target_host_account,omitempty"`
	RemoteAddress     string           `json:"remote_address,omitempty"`
	Connected         string           `json:"connected,omitempty"`
	Disconnected      string           `json:"disconnected,omitempty"`
	Status            string           `json:"status,omitempty"`
	LastActivity      string           `json:"last_activity,omitempty"`
	ForceDisconnect   string           `json:"force_disconnect,omitempty"`
	TerminationReason string           `json:"termination_reason,omitempty"`
	Created           string           `json:"created,omitempty"`
	Updated           string           `json:"updated,omitempty"`
	UpdatedBy         string           `json:"updated_by,omitempty"`
	TrailID           string           `json:"trail_id,omitempty"`
	IndexStatus       string           `json:"index_status,omitempty"`
	AccessGroupID     string           `json:"access_group_id,omitempty"`
	AuthMethod        []string         `json:"authentication_method,omitempty"`
	BytesIn           int              `json:"bytes_in,omitempty"`
	BytesOut          int              `json:"bytes_out,omitempty"`
	Duration          int              `json:"duration,omitempty"`
	TrailRemoved      bool             `json:"trail_removed,omitempty"`
	AuditEnabled      bool             `json:"audit_enabled,omitempty"`
	TargetHostData    ConnectionHost   `json:"target_host_data,omitempty"`
	UserData          UserData         `json:"user,omitempty"`
	UserRoles         []ConnectionRole `json:"user_roles,omitempty"`
	TargetHostRoles   []ConnectionRole `json:"target_host_roles,omitempty"`
	AccessRoles       []AccessRoles    `json:"access_roles,omitempty"`
}

// TimestampSearch timestamp search struct definition
type TimestampSearch struct {
	Start string
	End   string
}

// ConnectionSearch connection search struct definition
type ConnectionSearch struct {
	ID                   []string        `json:"id,omitempty"`
	ProxyID              []string        `json:"proxy_id,omitempty"`
	Type                 []string        `json:"type,omitempty"`
	Mode                 []string        `json:"mode,omitempty"`
	UserAgent            []string        `json:"user_agent,omitempty"`
	AuthMethod           []string        `json:"authentication_method,omitempty"`
	UserID               []string        `json:"user_id,omitempty"`
	UserDisplayName      []string        `json:"user_display_name,omitempty"`
	UserRoles            []string        `json:"user_roles,omitempty"`
	TargetHost           []string        `json:"target_host_id,omitempty"`
	TargetHostCommonName []string        `json:"target_host_common_name,omitempty"`
	TargetHostAddress    []string        `json:"target_host_address,omitempty"`
	TargetHostAccount    []string        `json:"target_host_account,omitempty"`
	TargetHostRoles      []string        `json:"target_host_roles,omitempty"`
	RemoteAddress        []string        `json:"remote_address,omitempty"`
	Status               []string        `json:"status,omitempty"`
	ForceDisconnect      []string        `json:"force_disconnect,omitempty"`
	AccessRoles          []string        `json:"access_roles,omitempty"`
	KeyWords             string          `json:"keywords,omitempty"`
	HasAccessRoles       bool            `json:"has_access_roles,omitempty"`
	Connected            TimestampSearch `json:"connected,omitempty"`
	Disconnected         TimestampSearch `json:"disconnected,omitempty"`
	LastActivity         TimestampSearch `json:"last_activity,omitempty"`
}
