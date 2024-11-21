//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package licensemanager

import (
	"github.com/SSHcom/privx-sdk-go/api/rolestore"
)

// LicenseStatistics license statistics request definition.
type LicenseStatistics struct {
	Optin bool `json:"optin"`
}

// RegistrationStatus mobile gw registration status response definition.
type RegistrationStatus struct {
	Status      string             `json:"status"`
	UsedSources []rolestore.Source `json:"used_sources"`
	ProductId   string             `json:"product_id"`
}
