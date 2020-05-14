//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

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

type Role struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Explicit    bool     `json:"explicit" tabulate:"@userCtx"`
	Implicit    bool     `json:"implicit" tabulate:"@userCtx"`
	System      bool     `json:"system"`
	GrantType   string   `json:"grant_type"`
	Comment     string   `json:"comment"`
	Permissions []string `json:"permissions"`
	Context     *Context `json:"context"`
	MemberCount int      `json:"member_count"`
}

type Context struct {
	Enabled   bool   `json:"enabled"`
	BlockRole bool   `json:"block_role"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Timezone  string `json:"timezone"`
}
