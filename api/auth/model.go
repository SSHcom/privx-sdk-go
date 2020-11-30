//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package auth

//
// CA is root certificate representation
type CA struct {
	ID        string `json:"id"`
	GroupID   string `json:"group_id"`
	Type      string `json:"type"`
	Size      int    `json:"size"`
	PublicKey string `json:"public_key"`
	X509      string `json:"x509_certificate"`
}

//
//
type AccessGroup struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Default bool   `json:"default,omitempty"`
}
