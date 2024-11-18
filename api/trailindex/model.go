//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package trailindex

// TrailIndexResponse trail index response definition.
type TrailIndexResponse struct {
	ConnID      string          `json:"connection_id"`
	ChanID      string          `json:"channel_id"`
	Protocol    string          `json:"protocol"`
	ChannelType string          `json:"type"`
	Extra       TrailIndexExtra `json:"extra,omitempty"`
	TimeStamp   string          `json:"timestamp"`
	Position    int             `json:"position"`
	Content     string          `json:"content"`
}

// TrailIndexExtra trail index extra definition.
type TrailIndexExtra struct {
	Command           string `json:"command,omitempty"`
	SubsystemName     string `json:"subsystem_name,omitempty"`
	OriginatorAddress string `json:"originator_address,omitempty"`
	OriginatorPort    uint32 `json:"originator_port,omitempty"`
	ListenerAddress   string `json:"listener_address,omitempty"`
	ListenerPort      uint32 `json:"listener_port,omitempty"`
	DstAddress        string `json:"dst_address,omitempty"`
	DstPort           uint32 `json:"dst_port,omitempty"`
	PTY               bool   `json:"pty,omitempty"`
}

// TranscriptSearch transcript search request object definition.
type TranscriptSearch struct {
	ConnID    string `json:"connection_id"`
	ChanID    string `json:"channel_id"`
	Protocol  string `json:"protocol"`
	Keywords  string `json:"keywords"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	StartPos  int    `json:"start_position"`
	EndPos    int    `json:"end_position"`
}

// ConnectionTranscriptStatus connection transcript status definition.
type ConnectionTranscriptStatus struct {
	ConnectionID string `json:"connection_id"`
	Status       string `json:"status"`
}
