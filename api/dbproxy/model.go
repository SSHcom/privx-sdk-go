//
// Copyright (c) 2023 SSH Communications Security Inc.
//
// All rights reserved.
//

package dbproxy

// DBProxyCACertificateInfo db proxy x509 CA certificate definition.
type DBProxyCACertificateInfo struct {
	Subject           string `json:"subject,omitempty"`
	Issuer            string `json:"issuer,omitempty"`
	Serial            string `json:"serial,omitempty"`
	NotBefore         string `json:"not_before,omitempty"`
	NotAfter          string `json:"not_after,omitempty"`
	FingerPrintSHA1   string `json:"fingerprint_sha1,omitempty"`
	FingerPrintSHA256 string `json:"fingerprint_sha256,omitempty"`
}

// DBProxyAPIConf db proxy configuration definition.
type DBProxyAPIConf struct {
	CACertificate *DBProxyCACertificateInfo `json:"ca_certificate,omitempty"`
	Chain         string                    `json:"ca_certificate_chain,omitempty"`
}
