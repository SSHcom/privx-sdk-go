//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package licensemanager

import (
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// LicenseManager is a license manager client instance.
type LicenseManager struct {
	api restapi.Connector
}

// New license manager client constructor.
func New(api restapi.Connector) *LicenseManager {
	return &LicenseManager{api: api}
}

// MARK: Status
// Status get license manager microservice status.
func (c *LicenseManager) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/license-manager/api/v1/status").
		Get(status)

	return status, err
}

// MARK: License
// GetLicense get license.
func (c *LicenseManager) GetLicense() (map[string]interface{}, error) {
	license := map[string]interface{}{}

	_, err := c.api.
		URL("/license-manager/api/v1/license").
		Get(&license)

	return license, err
}

// SetLicense set new license.
func (c *LicenseManager) SetLicense(licenseCode string) error {
	_, err := c.api.
		URL("/license-manager/api/v1/license").
		Post(licenseCode)

	return err
}

// RefreshLicense refresh license info.
func (c *LicenseManager) RefreshLicense() (map[string]interface{}, error) {
	license := map[string]interface{}{}

	_, err := c.api.
		URL("/license-manager/api/v1/license/refresh").
		Post(nil, license)

	return license, err
}

// SetLicenseStatistics set settings for SSH license statistics.
func (c *LicenseManager) SetLicenseStatistics(optin LicenseStatistics) error {
	_, err := c.api.
		URL("/license-manager/api/v1/license/optin").
		Post(&optin)

	return err
}

// DeactivateLicense deactivate license.
func (c *LicenseManager) DeactivateLicense() error {
	_, err := c.api.
		URL("/license-manager/api/v1/license/deactivate").
		Post(nil)

	return err
}

// GetLicenseJSSnippet get PrivX license javascript snippet.
func (c *LicenseManager) GetLicenseJSSnippet() (string, error) {
	snippet := ""

	_, err := c.api.
		URL("/license-manager/api/v1/license.js").
		Get(&snippet)

	return snippet, err
}

// MARK: Mobile Gateway
// Get PrivX registration status to mobile gateway.
func (c *LicenseManager) GetMobileGwRegistration() (*RegistrationStatus, error) {
	status := &RegistrationStatus{}

	_, err := c.api.
		URL("/license-manager/api/v1/mobilegw/status").
		Get(status)

	return status, err
}

// RegisterToMobileGw register PrivX instance to mobile gateway.
func (c *LicenseManager) RegisterToMobileGw() error {
	_, err := c.api.
		URL("/license-manager/api/v1/mobilegw/register").
		Post(nil)

	return err
}

// UnregisterFromMobileGw unregister PrivX instance from mobile gateway.
func (c *LicenseManager) UnregisterFromMobileGw() error {
	_, err := c.api.
		URL("/license-manager/api/v1/mobilegw/unregister").
		Delete(nil)

	return err
}
