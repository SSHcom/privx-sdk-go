//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package vault

import (
	"encoding/json"

	"github.com/SSHcom/privx-sdk-go/api/rolestore"
)

// Params struct for pagination queries.
type Params struct {
	Offset string `json:"offset,omitempty"`
	Limit  string `json:"limit,omitempty"`
}

// Secret contains PrivX metadata about secret and its vault
type Secret struct {
	ID         string              `json:"name"`
	Author     string              `json:"author,omitempty"`
	Editor     string              `json:"updated_by,omitempty"`
	Created    string              `json:"created,omitempty"`
	Updated    string              `json:"updated,omitempty"`
	AllowRead  []rolestore.RoleRef `json:"read_roles,omitempty"`
	AllowWrite []rolestore.RoleRef `json:"write_roles,omitempty"`
	Data       json.RawMessage     `json:"data,omitempty"`
}
