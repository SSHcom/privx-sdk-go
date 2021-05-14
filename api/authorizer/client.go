//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package authorizer

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Client is a authorizer client instance.
type Client struct {
	api restapi.Connector
}

// New creates a new authorizer client instance
func New(api restapi.Connector) *Client {
	return &Client{api: api}
}

// CACertificates gets authorizer's root certificates
func (auth *Client) CACertificates(accessGroupID string) ([]CA, error) {
	ca := []CA{}
	filters := Params{
		AccessGroupID: accessGroupID,
	}

	_, err := auth.api.
		URL("/authorizer/api/v1/cas").
		Query(&filters).
		Get(&ca)

	return ca, err
}

// CACertificate gets authorizer's root certificate
func (auth *Client) CACertificate(caID, filename string) error {
	err := auth.api.
		URL("/authorizer/api/v1/cas/%s", url.PathEscape(caID)).
		Download(filename)

	return err
}

// CertificateRevocationList gets authorizer CA's certificate revocation list.
func (auth *Client) CertificateRevocationList(caID, filename string) error {
	err := auth.api.
		URL("/authorizer/api/v1/cas/%s/crl", url.PathEscape(caID)).
		Download(filename)

	return err
}

// TargetHostCredentials get target host credentials for the user
func (auth *Client) TargetHostCredentials(authorizer *AuthorizationRequest) (*Principal, error) {
	principal := &Principal{}

	_, err := auth.api.
		URL("/authorizer/api/v1/ca/authorize").
		Post(&authorizer, &principal)

	return principal, err
}

// Principals gets defined principals from the authorizer
func (auth *Client) Principals() ([]Principal, error) {
	principals := []Principal{}

	_, err := auth.api.
		URL("/authorizer/api/v1/cas").
		Get(&principals)

	return principals, err
}

// Principal gets the principal key by its group ID
func (auth *Client) Principal(groupID, keyID, filter string) (*Principal, error) {
	principal := &Principal{}
	filters := Params{
		KeyID:  keyID,
		Filter: filter,
	}

	_, err := auth.api.
		URL("/authorizer/api/v1/principals/%s", url.PathEscape(groupID)).
		Query(&filters).
		Get(&principal)

	return principal, err
}

// DeletePrincipalKey delete the principal key by its group ID
func (auth *Client) DeletePrincipalKey(groupID, keyID string) error {
	filters := Params{
		KeyID: keyID,
	}

	_, err := auth.api.
		URL("/authorizer/api/v1/principals/%s", url.PathEscape(groupID)).
		Query(filters).
		Delete()

	return err
}

// CreatePrincipalKey create a principal key pair
func (auth *Client) CreatePrincipalKey(groupID string) (*Principal, error) {
	principal := &Principal{}

	_, err := auth.api.
		URL("/authorizer/api/v1/principals/%s/create", url.PathEscape(groupID)).
		Post(nil, &principal)

	return principal, err
}

// ImportPrincipalKey mport a principal key pair
func (auth *Client) ImportPrincipalKey(groupID string, key *PrincipalKeyImportRequest) (*Principal, error) {
	principal := &Principal{}

	_, err := auth.api.
		URL("/authorizer/api/v1/principals/%s/import", url.PathEscape(groupID)).
		Post(&key, &principal)

	return principal, err
}

// SignPrincipalKey sign a principal key and get a signature
func (auth *Client) SignPrincipalKey(groupID, keyID string, credential *Credential) (*Signature, error) {
	signature := &Signature{}
	filters := Params{
		KeyID: keyID,
	}

	_, err := auth.api.
		URL("/authorizer/api/v1/principals/%s/sign", url.PathEscape(groupID)).
		Query(&filters).
		Post(&credential, &signature)

	return signature, err
}

// ExtenderCACertificates gets authorizer's extender CA certificates
func (auth *Client) ExtenderCACertificates(accessGroupID string) ([]CA, error) {
	certificates := []CA{}
	filters := Params{
		AccessGroupID: accessGroupID,
	}

	_, err := auth.api.
		URL("/authorizer/api/v1/extender/cas").
		Query(&filters).
		Get(&certificates)

	return certificates, err
}

// ExtenderCACertificate gets authorizer's extender CA certificate
func (auth *Client) ExtenderCACertificate(id string) (*CA, error) {
	certificate := &CA{}

	_, err := auth.api.
		URL("/authorizer/api/v1/extender/cas/%s", url.PathEscape(id)).
		Get(&certificate)

	return certificate, err
}

// DownloadExtenderCertificateCRL gets authorizer CA's certificate revocation list
func (auth *Client) DownloadExtenderCertificateCRL(filename, id string) error {
	err := auth.api.
		URL("/authorizer/api/v1/extender/cas/%s/crl", url.PathEscape(id)).
		Download(filename)

	return err
}

// ExtenderConfigDownloadHandle get a session id
func (auth *Client) ExtenderConfigDownloadHandle(trustedClientID string) (*DownloadHandle, error) {
	sessionID := &DownloadHandle{}

	_, err := auth.api.
		URL("/authorizer/api/v1/extender/conf/%s", url.PathEscape(trustedClientID)).
		Post(nil, &sessionID)

	return sessionID, err
}

// ExtenderConfig gets a pre-configured extender config
func (auth *Client) ExtenderConfig(trustedClientID, sessionID, filename string) error {
	err := auth.api.
		URL("/authorizer/api/v1/extender/conf/%s/%s", url.PathEscape(trustedClientID), url.PathEscape(sessionID)).
		Download(filename)

	return err
}

// TrustedClientHealthCheck test the health of the trusted client connection with the server
func (auth *Client) TrustedClientHealthCheck(time *HealthCheckParams) (*HealthCheckStatus, error) {
	configErrors := &HealthCheckStatus{}

	_, err := auth.api.
		URL("/authorizer/api/v1/trusted-clients/health-check").
		Post(&time, &configErrors)

	return configErrors, err
}

// DeployScriptSessionID get a session id for a deployment script
func (auth *Client) DeployScriptSessionID(trustedClientID string) (*DownloadHandle, error) {
	sessionID := &DownloadHandle{}

	_, err := auth.api.
		URL("/authorizer/api/v1/deploy/%s", url.PathEscape(trustedClientID)).
		Post(nil, &sessionID)

	return sessionID, err
}

// DownloadDeployScript gets a pre-configured deployment script
func (auth *Client) DownloadDeployScript(trustedClientID, sessionID, filename string) error {
	err := auth.api.
		URL("/authorizer/api/v1/deploy/%s/%s", url.PathEscape(trustedClientID), url.PathEscape(sessionID)).
		Download(filename)

	return err
}

// DownloadPrincipalCommandScript gets the principals_command.sh script
func (auth *Client) DownloadPrincipalCommandScript(filename string) error {
	err := auth.api.
		URL("/authorizer/api/v1/deploy/principals_command.sh").
		Download(filename)

	return err
}
