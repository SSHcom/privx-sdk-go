//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

// User contains PrivX user information.
type User struct {
	ID                string   `json:"id"`
	SourceUserID      string   `json:"source_user_id"`
	Tags              []string `json:"tags"`
	Principal         string   `json:"principal"`
	Source            string   `json:"source"`
	FullName          string   `json:"full_name"`
	Email             string   `json:"email"`
	DistinguishedName string   `json:"distinguished_name"`
	Roles             []Role   `json:"roles"`
}

// Role contains PrivX role information.
type Role struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Explicit    bool       `json:"explicit" tabulate:"@userCtx"`
	Implicit    bool       `json:"implicit" tabulate:"@userCtx"`
	System      bool       `json:"system"`
	GrantType   string     `json:"grant_type"`
	Comment     string     `json:"comment"`
	SourceRule  SourceRule `json:"source_rules"`
	Permissions []string   `json:"permissions"`
	Context     *Context   `json:"context"`
	MemberCount int        `json:"member_count"`
	PublicKey   []string   `json:"principal_public_key_strings,omitempty"`
}

// RoleRef is a reference to role object
type RoleRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Context defines the context information for a role.
type Context struct {
	Enabled   bool   `json:"enabled"`
	BlockRole bool   `json:"block_role"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Timezone  string `json:"timezone"`
}

// SourceRule defines a mapping of role to object objects in directory
type SourceRule struct {
	Type    string       `json:"type"`
	Match   string       `json:"match"`
	Source  string       `json:"source,omitempty"`
	Pattern string       `json:"search_string,omitempty"`
	Rules   []SourceRule `json:"rules"`
}

// SourceRuleNone creates an empty mapping source for the role
func SourceRuleNone() SourceRule {
	return SourceRule{
		Type:  "GROUP",
		Match: "ANY",
	}
}
