//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package licensemanager

import (
	"github.com/SSHcom/privx-sdk-go/restapi"
)

type LicenseManager struct {
	api restapi.Connector
}

// New creates a new license manager client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *LicenseManager {
	return &LicenseManager{api: api}
}

// SetLicenseStatistics settings for SSH license statistics
func (store *LicenseManager) SetLicenseStatistics(optin bool) error {
	statistics := License{
		Optin: optin,
	}

	_, err := store.api.
		URL("/license-manager/api/v1/license/optin").
		Post(&statistics)

	return err
}

// CreateLicense post a new license to server
func (store *LicenseManager) CreateLicense(license string) error {
	_, err := store.api.
		URL("/license-manager/api/v1/license").
		Post(license)

	return err
}

// License return privx license
func (store *LicenseManager) License() (license *License, err error) {
	license = new(License)

	_, err = store.api.
		URL("/license-manager/api/v1/license").
		Get(license)

	return
}
