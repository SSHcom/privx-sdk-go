//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package vault

import (
	"time"

	"github.com/SSHcom/privx-sdk-go/api/rolestore"
)

// Secret secret definition.
type Secret struct {
	SecretRequest
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	UpdatedBy string    `json:"updated_by"`
	Author    string    `json:"author"`
	Path      string    `json:"path"`
}

// SecretRequest secret request definition.
type SecretRequest struct {
	Name       string                  `json:"name"`
	ReadRoles  []rolestore.RoleHandle  `json:"read_roles"`
	WriteRoles []rolestore.RoleHandle  `json:"write_roles"`
	Data       *map[string]interface{} `json:"data,omitempty"`
	OwnerID    string                  `json:"owner_id,omitempty"`
}

// SecretCreate secret create create response definition.
type SecretCreate struct {
	Name string `json:"name"`
}

// SecretSearch secret search request definition.
type SecretSearch struct {
	Keywords string   `json:"keywords"`
	SortDir  string   `json:"sortdir"`
	SortKey  string   `json:"sortkey"`
	Filter   string   `json:"filter"`
	OwnerIDs []string `json:"owner_id"`
	Limit    int      `json:"limit"`
	Offset   int      `json:"offset"`
}
