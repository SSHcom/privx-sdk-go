//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package licensemanager

import "github.com/SSHcom/privx-sdk-go/api/rolestore"

// License license definition
type License struct {
	LicenseStatus           string   `json:"license_status,omitempty"`
	Version                 string   `json:"version,omitempty"`
	CreationDate            string   `json:"creation_date,omitempty"`
	ExpiryDate              string   `json:"expiry_date,omitempty"`
	LastRefreshedDate       string   `json:"last_refreshed_date,omitempty"`
	Customer                string   `json:"customer,omitempty"`
	SerialNumber            string   `json:"serial_number,omitempty"`
	Product                 string   `json:"product,omitempty"`
	LicenseCode             string   `json:"license_code,omitempty"`
	LicenseMessage          string   `json:"license_message,omitempty"`
	Status                  int      `json:"status,omitempty"`
	Message                 int      `json:"message,omitempty"`
	MaxHosts                int      `json:"max_hosts,omitempty"`
	MaxAuditedHosts         int      `json:"max_audited_hosts,omitempty"`
	MaxConcurrentSSHConns   int      `json:"max_concurrent_ssh_conns,omitempty"`
	MaxConcurrentRDPConns   int      `json:"max_concurrent_rdp_conns,omitempty"`
	MaxConcurrentHTTPSConns int      `json:"max_concurrent_https_conns,omitempty"`
	MaxConcurrentVNCConns   int      `json:"max_concurrent_vnc_conns,omitempty"`
	MaxUsers                int      `json:"max_users,omitempty"`
	AnalyticsEnabled        bool     `json:"analytics_enabled,omitempty"`
	IsOffline               bool     `json:"isoffline,omitempty"`
	Optin                   bool     `json:"optin,omitempty"`
	Features                []string `json:"features,omitempty"`
	HostsInUse              int      `json:"hosts_in_use,omitempty"`
	AuditHostsInUse         int      `json:"audit_hosts_in_use,omitempty"`
	UsersInUse              int      `json:"users_in_use,omitempty"`
}

type RegistrationStatus struct {
	Status      string             `json:"status"`
	UsedSources []rolestore.Source `json:"used_sources"`
	ProductId   string             `json:"product_id"`
}
