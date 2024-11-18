//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package networkaccessmanager

// NetworkTargetTagsParams network target tags query parameter definition.
type NetworkTargetTagsParams struct {
	Query int `url:"query,omitempty"`
}

// NetworkTargetDisable network target disable request definition.
type NetworkTargetDisable struct {
	Disabled bool `json:"disabled"`
}

// NetworkTarget network target definition.
type NetworkTarget struct {
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Dst              []Destination `json:"dst"`
	SrcNAT           bool          `json:"src_nat,omitempty"`
	Roles            []RoleHandle  `json:"roles"`
	Tags             []string      `json:"tags"`
	Comment          string        `json:"comment,omitempty"`
	UserInstructions string        `json:"user_instructions,omitempty"`
	ExclusiveAccess  bool          `json:"exclusive_access,omitempty"`
	Disabled         string        `json:"disabled,omitempty"`
	Created          string        `json:"created"`
	Author           string        `json:"author"`
	Updated          string        `json:"updated"`
	UpdatedBy        string        `json:"updated_by"`
}

// RoleHandle is a handle to network target role definition.
type RoleHandle struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

// Destination network target destination definition.
type Destination struct {
	Sel Selector       `json:"selector"`
	NAT *NATParameters `json:"nat,omitempty"`
}

// Selector network target selector definition.
type Selector struct {
	IP       IPRange    `json:"ip"`
	Protocol string     `json:"proto,omitempty"`
	Port     *PortRange `json:"port,omitempty"`
}

// PortRange port range definition
type PortRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// IPRange ip range definition.
type IPRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// NATParameters network target parameters.
type NATParameters struct {
	Addr string `json:"addr"`
	Port int    `json:"port,omitempty"`
}

// NetworkTargetSearch network target search request definition.
type NetworkTargetSearch struct {
	Keywords string   `json:"keywords,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}
