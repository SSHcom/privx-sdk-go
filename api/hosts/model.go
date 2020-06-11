//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package hosts

import "github.com/SSHcom/privx-sdk-go/api/rolestore"

// Source of host objects
type Source string

// Source constants
const (
	UI   = Source("UI")
	SCAN = Source("SCAN")
)

// Address is fully qualified domain names, IPv4 or IPv6 addresses of the host
type Address string

// Scheme of protocols allowed by the host
type Scheme string

// Scheme constants, all supported protocols
const (
	SSH   = Scheme("SSH")
	RDP   = Scheme("RDP")
	VNC   = Scheme("VNC")
	HTTP  = Scheme("HTTP")
	HTTPS = Scheme("HTTPS")
)

// Service specify the service available on target host
type Service struct {
	Scheme  Scheme  `json:"service"`
	Address Address `json:"address"`
	Port    int     `json:"port"`
	Source  Source  `json:"source"`
}

// Principal of the target host
type Principal struct {
	ID     string              `json:"principal"`
	Roles  []rolestore.RoleRef `json:"roles"`
	Source Source              `json:"source"`
}

// Host defines PrivX target
type Host struct {
	ID         string      `json:"id,omitempty"`
	Name       string      `json:"common_name,omitempty"`
	Addresses  []Address   `json:"addresses,omitempty"`
	Services   []Service   `json:"services,omitempty"`
	Principals []Principal `json:"principals,omitempty"`
}

// Service creates a corresponding service definition
//   hosts.SSH.Service(...)
func (scheme Scheme) Service(addr Address, port int) Service {
	return Service{
		Scheme:  scheme,
		Address: addr,
		Port:    port,
		Source:  UI,
	}
}

// NewPrincipal creates a corresponding definition from roles
func NewPrincipal(id string, role ...rolestore.RoleRef) Principal {
	return Principal{
		ID:     id,
		Roles:  role,
		Source: UI,
	}
}
