//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package trailindex

// TrailIndexResponse trail index response definition.
type TrailIndexResponse struct {
	ConnectionType string          `json:"connection_type"`
	ConnID         string          `json:"connection_id"`
	ChanID         string          `json:"channel_id"`
	Protocol       string          `json:"protocol"`
	ChannelType    string          `json:"type"`
	Extra          TrailIndexExtra `json:"extra,omitempty"`
	TimeStamp      string          `json:"timestamp"`
	Position       int             `json:"position"`
	Content        string          `json:"content"`
}

// TrailIndexExtra trail index extra definition.
type TrailIndexExtra struct {
	// exec
	Command string `json:"command,omitempty"`
	PTY     bool   `json:"pty,omitempty"`
	// api-proxy
	HTTPRequest           string `json:"http_request,omitempty"`
	HTTPRequestID         string `json:"http_request_id,omitempty"`
	HTTPResponse          string `json:"http_response,omitempty"`
	HTTPResponseTimeStamp string `json:"http_response_timestamp,omitempty"`
	HTTPTransport         string `json:"http_transport,omitempty"`
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
