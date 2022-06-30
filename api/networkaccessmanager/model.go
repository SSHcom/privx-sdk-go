//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package networkaccessmanager

type Nat struct {
	Addr string `json:"addr,omitempty"`
	Port int    `json:"port,omitempty"`
}
type Port struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}
type Ip struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}
type Selector struct {
	IP    Ip     `json:"ip,omitempty"`
	Port  Port   `json:"port,omitempty"`
	Proto string `json:"proto,omitempty"`
}
type Dst struct {
	Selector Selector `json:"selector,omitempty"`
	Nat      Nat      `json:"nat,omitempty"`
}
type Role struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type Item struct {
	ID               string `json:"id,omitempty"`
	Created          string `json:"created,omitempty"`
	Updated          string `json:"updated,omitempty"`
	UpdatedBy        string `json:"updated_by,omitempty"`
	Author           string `json:"author,omitempty"`
	Comment          string `json:"comment,omitempty"`
	Name             string `json:"name,omitempty"`
	UserInstructions string `json:"user_instructions,omitempty"`
	SrcNat           bool   `json:"src_nat,omitempty"`
	Roles            []Role `json:"roles,omitempty"`
	Dst              []Dst  `json:"dst,omitempty"`
	ExclusiveAccess  bool   `json:"exclusive_access,omitempty"`
	Disabled         bool   `json:"disabled,omitempty"`
}
type ApiNwtargetsResponse struct {
	Count int    `json:"count,omitempty"`
	Items []Item `json:"items,omitempty"`
}
type ApiNwtargetsResponsePost struct {
	ID string `json:"id,omitempty"`
}
type Params struct {
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Sortkey string `json:"sortkey,omitempty"`
	Sortdir string `json:"sortdir,omitempty"`
	Name    string `json:"name,omitempty"`
	ID      string `json:"id,omitempty"`
	Filter  string `json:"filter,omitempty"`
}
