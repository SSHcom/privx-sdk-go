//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package licensemanager

import (
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// LicenseManager is a license manager client instance.
type LicenseManager struct {
	api restapi.Connector
}

// New creates a new license manager client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *LicenseManager {
	return &LicenseManager{api: api}
}

// RefreshLicense refresh the license info
func (store *LicenseManager) RefreshLicense() (*License, error) {
	license := &License{}

	_, err := store.api.
		URL("/license-manager/api/v1/license/refresh").
		Post(nil, license)

	return license, err
}

// DeactivateLicense deactivate license
func (store *LicenseManager) DeactivateLicense() error {
	_, err := store.api.
		URL("/license-manager/api/v1/license/deactivate").
		Post(nil)

	return err
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

// SetLicense post a new license to server
func (store *LicenseManager) SetLicense(licenseCode string) error {
	_, err := store.api.
		URL("/license-manager/api/v1/license").
		Post(licenseCode)

	return err
}

// License return privx license
func (store *LicenseManager) License() (*License, error) {
	license := &License{}

	_, err := store.api.
		URL("/license-manager/api/v1/license").
		Get(license)

	return license, err
}

// Register PrivX instance to mobilegw
func (store *LicenseManager) RegisterToMobileGW() error {
	_, err := store.api.
		URL("/license-manager/api/v1/mobilegw/register").
		Post(nil)

	return err
}

// Unregister PrivX instance from mobilegw
func (store *LicenseManager) UnregisterToMobileGW() error {
	_, err := store.api.
		URL("/license-manager/api/v1/mobilegw/unregister").
		Delete(nil)

	return err
}

// Get PrivX registration status to mobilegw
func (store *LicenseManager) GetMobileGwRegistration() (*RegistrationStatus, error) {
	status := &RegistrationStatus{}

	_, err := store.api.
		URL("/license-manager/api/v1/mobilegw/status").
		Get(status)

	return status, err
}
