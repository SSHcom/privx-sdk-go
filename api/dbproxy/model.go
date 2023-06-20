//
// Copyright (c) 2023 SSH Communications Security Inc.
//
// All rights reserved.
//

package dbproxy

import "time"

// KeyValue key value definition
type KeyValue struct {
	Key   string `json:"k"`
	Value string `json:"v"`
}

// ServiceStatus db proxy service status definition
type ServiceStatus struct {
	Version       string     `json:"version"`
	APIVersion    string     `json:"api_version"`
	Status        string     `json:"status"`
	StatusMessage string     `json:"status_message,omitempty"`
	ApplicationID string     `json:"app_id,omitempty"`
	StatusDetails []KeyValue `json:"status_details,omitempty"`
	StartTime     time.Time  `json:"start_time"`
}

// DBProxyCACertificateInfo DB proxy x509 CA certificate information
type DBProxyCACertificateInfo struct {
	Subject           string `json:"subject,omitempty"`
	Issuer            string `json:"issuer,omitempty"`
	Serial            string `json:"serial,omitempty"`
	NotBefore         string `json:"not_before,omitempty"`
	NotAfter          string `json:"not_after,omitempty"`
	FingerPrintSHA1   string `json:"fingerprint_sha1,omitempty"`
	FingerPrintSHA256 string `json:"fingerprint_sha256,omitempty"`
}

// DBProxyConf DB proxy config definition
type DBProxyConf struct {
	CACertificate *DBProxyCACertificateInfo `json:"ca_certificate,omitempty"`
	Chain         string                    `json:"ca_certificate_chain,omitempty"`
}
