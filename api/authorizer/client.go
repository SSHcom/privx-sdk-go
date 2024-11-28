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

// CAConfig get authorizers root certificate config by ca type.
func (auth *Client) CAConfig(caType string) (ComponentCaConfig, error) {
	caConf := ComponentCaConfig{}

	_, err := auth.api.
		URL("/authorizer/api/v1/%s/cas/config", caType).
		Get(&caConf)

	return caConf, err
}

// UpdateCAConfig update authorizers root certificate config by ca type.
func (auth *Client) UpdateCAConfig(caType string, caConf ComponentCaConfig) error {
	_, err := auth.api.
		URL("/authorizer/api/v1/%s/cas/config", caType).
		Put(&caConf)

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
func (auth *Client) TargetHostCredentials(authorizer *AuthorizationRequest) (*ApiIdentitiesResponse, error) {
	principal := &ApiIdentitiesResponse{}

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

// DownloadExtenderConfig gets a pre-configured extender config
func (auth *Client) DownloadExtenderConfig(trustedClientID, sessionID, filename string) error {
	err := auth.api.
		URL("/authorizer/api/v1/extender/conf/%s/%s", url.PathEscape(trustedClientID), url.PathEscape(sessionID)).
		Download(filename)

	return err
}

// DeployScriptDownloadHandle get a session id for a deployment script
func (auth *Client) DeployScriptDownloadHandle(trustedClientID string) (*DownloadHandle, error) {
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

// CarrierConfigDownloadHandle get a session id for a carrier config
func (auth *Client) CarrierConfigDownloadHandle(trustedClientID string) (*DownloadHandle, error) {
	sessionID := &DownloadHandle{}

	_, err := auth.api.
		URL("/authorizer/api/v1/carrier/conf/%s", url.PathEscape(trustedClientID)).
		Post(nil, &sessionID)

	return sessionID, err
}

// DownloadCarrierConfig gets a pre-configured carrier config
func (auth *Client) DownloadCarrierConfig(trustedClientID, sessionID, filename string) error {
	err := auth.api.
		URL("/authorizer/api/v1/carrier/conf/%s/%s", url.PathEscape(trustedClientID), url.PathEscape(sessionID)).
		Download(filename)

	return err
}

// WebProxyCACertificates gets authorizer's web proxy CA certificates
func (auth *Client) WebProxyCACertificates(accessGroupID string) ([]CA, error) {
	certificates := []CA{}
	filters := Params{
		AccessGroupID: accessGroupID,
	}

	_, err := auth.api.
		URL("/authorizer/api/v1/icap/cas").
		Query(&filters).
		Get(&certificates)

	return certificates, err
}

// WebProxyCACertificate gets authorizer's web proxy CA certificate
func (auth *Client) WebProxyCACertificate(trustedClientID string) (*CA, error) {
	certificate := &CA{}

	_, err := auth.api.
		URL("/authorizer/api/v1/icap/cas/%s", url.PathEscape(trustedClientID)).
		Get(&certificate)

	return certificate, err
}

// DownloadWebProxyCertificateCRL gets authorizer CA's certificate revocation list
func (auth *Client) DownloadWebProxyCertificateCRL(filename, trustedClientID string) error {
	err := auth.api.
		URL("/authorizer/api/v1/icap/cas/%s/crl", url.PathEscape(trustedClientID)).
		Download(filename)

	return err
}

// WebProxySessionDownloadHandle get a session id for a web proxy config
func (auth *Client) WebProxySessionDownloadHandle(trustedClientID string) (*DownloadHandle, error) {
	sessionID := &DownloadHandle{}

	_, err := auth.api.
		URL("/authorizer/api/v1/icap/conf/%s", url.PathEscape(trustedClientID)).
		Post(nil, &sessionID)

	return sessionID, err
}

// DownloadWebProxyConfig gets a pre-configured web proxy config
func (auth *Client) DownloadWebProxyConfig(trustedClientID, sessionID, filename string) error {
	err := auth.api.
		URL("/authorizer/api/v1/icap/conf/%s/%s", url.PathEscape(trustedClientID), url.PathEscape(sessionID)).
		Download(filename)

	return err
}

// CertTemplates returns the certificate authentication templates for the service
func (auth *Client) CertTemplates(service string) ([]CertTemplate, error) {
	result := templatesResult{}
	filters := Params{
		Service: service,
	}

	_, err := auth.api.
		URL("/authorizer/api/v1/cert/templates").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// SSLTrustAnchor returns the SSL trust anchor (PrivX TLS CA certificate)
func (auth *Client) SSLTrustAnchor() (*TrustAnchor, error) {
	anchor := &TrustAnchor{}

	_, err := auth.api.
		URL("/authorizer/api/v1/ssl-trust-anchor").
		Get(&anchor)

	return anchor, err
}

// ExtenderTrustAnchor returns the extender trust anchor (PrivX TLS CA certificate)
func (auth *Client) ExtenderTrustAnchor() (*TrustAnchor, error) {
	anchor := &TrustAnchor{}

	_, err := auth.api.
		URL("/authorizer/api/v1/extender-trust-anchor").
		Get(&anchor)

	return anchor, err
}

// MARK: Access Groups
// AccessGroups lists all access group
func (auth *Client) AccessGroups(offset, limit int, sortkey, sortdir string) ([]AccessGroup, error) {
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}
	result := accessGroupResult{}

	_, err := auth.api.
		URL("/authorizer/api/v1/accessgroups").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// CreateAccessGroup create a access group
func (auth *Client) CreateAccessGroup(accessGroup *AccessGroup) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := auth.api.
		URL("/authorizer/api/v1/accessgroups").
		Post(&accessGroup, &object)

	return object.ID, err
}

// SearchAccessGroup search for access groups
func (auth *Client) SearchAccessGroup(offset, limit int, sortkey, sortdir string, search *SearchParams) ([]AccessGroup, error) {
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}
	result := accessGroupResult{}

	_, err := auth.api.
		URL("/authorizer/api/v1/accessgroups/search").
		Query(&filters).
		Post(search, &result)

	return result.Items, err
}

// AccessGroup get access group
func (auth *Client) AccessGroup(accessGroupID string) (*AccessGroup, error) {
	accessGroup := &AccessGroup{}

	_, err := auth.api.
		URL("/authorizer/api/v1/accessgroups/%s", url.PathEscape(accessGroupID)).
		Get(&accessGroup)

	return accessGroup, err
}

// UpdateAccessGroup update access group
func (auth *Client) UpdateAccessGroup(accessGroupID string, accessGroup *AccessGroup) error {
	_, err := auth.api.
		URL("/authorizer/api/v1/accessgroups/%s", url.PathEscape(accessGroupID)).
		Put(accessGroup)

	return err
}

// DeleteAccessGroup delete a access group
func (auth *Client) DeleteAccessGroup(accessGroupID string) error {
	_, err := auth.api.
		URL("/authorizer/api/v1/accessgroups/%s", accessGroupID).
		Delete()

	return err
}

// CreateAccessGroupsIdCas create CA Key to an access group
func (auth *Client) CreateAccessGroupsIdCas(accessGroupID string) (string, error) {
	var result string

	_, err := auth.api.
		URL("/authorizer/api/v1/accessgroups/%s/cas", accessGroupID).
		Post(nil, &result)

	return result, err
}

// DeleteAccessGroup delete a CA Key to an access group
func (auth *Client) DeleteAccessGroupsIdCas(accessGroupID string, caID string) error {
	_, err := auth.api.
		URL("/authorizer/api/v1/accessgroups/%s/cas/%s", accessGroupID, caID).
		Delete()

	return err
}

// MARK: Certs
// SearchCert search for certificates
func (auth *Client) SearchCert(offset, limit int, sortkey, sortdir string, cert *APICertificateSearch) ([]APICertificate, error) {
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}
	result := apiCertificateResult{}

	_, err := auth.api.
		URL("/authorizer/api/v1/cert/search").
		Query(&filters).
		Post(cert, &result)

	return result.Items, err
}

// Get all Certificates
func (auth *Client) GetAllCertificates() (apiCertificateResult, error) {
	certificates := apiCertificateResult{}

	_, err := auth.api.
		URL("/authorizer/api/v1/cert").
		Get(&certificates)

	return certificates, err
}

// Get Certificate by ID
func (auth *Client) GetCertByID(ID string) (ApiCertificateObject, error) {
	cert := ApiCertificateObject{}

	_, err := auth.api.
		URL("/authorizer/api/v1/cert/%s", url.PathEscape(ID)).
		Get(&cert)

	return cert, err
}

// MARK: Secrets
// AccountSecrets lists all account secrets
func (auth *Client) AccountSecrets(limit int, sortdir string) (AccountSecretsResult, error) {
	filters := Params{
		Limit:   limit,
		Sortdir: sortdir,
	}
	result := AccountSecretsResult{}

	_, err := auth.api.
		URL("/authorizer/api/v1/secrets").
		Query(&filters).
		Get(&result)

	return result, err
}

// SearchAccountSecrets search for account secrets
func (auth *Client) SearchAccountSecrets(limit int, sortdir string, search *AccountSecretsSearchRequest) (AccountSecretsResult, error) {
	filters := Params{
		Limit:   limit,
		Sortdir: sortdir,
	}
	result := AccountSecretsResult{}

	_, err := auth.api.
		URL("/authorizer/api/v1/secrets/search").
		Query(&filters).
		Post(search, &result)

	return result, err
}

// CheckoutAccountSecret checkout account secret
func (auth *Client) CheckoutAccountSecret(path string) (CheckoutResult, error) {
	checkoutReq := CheckoutRequest{
		Path: path,
	}
	result := CheckoutResult{}

	_, err := auth.api.
		URL("/authorizer/api/v1/secrets/checkouts").
		Post(checkoutReq, &result)

	return result, err
}

// Checkouts lists secret checkouts
func (auth *Client) Checkouts(limit int, sortdir string) (CheckoutResult, error) {
	filters := Params{
		Limit:   limit,
		Sortdir: sortdir,
	}
	result := CheckoutResult{}

	_, err := auth.api.
		URL("/authorizer/api/v1/secrets/checkouts").
		Query(&filters).
		Get(&result)

	return result, err
}

// Checkout get checkout by id
func (auth *Client) Checkout(checkoutId string) (*Checkout, error) {
	checkout := &Checkout{}

	_, err := auth.api.
		URL("/authorizer/api/v1/secrets/checkouts/%s", url.PathEscape(checkoutId)).
		Get(&checkout)

	return checkout, err
}

// ReleaseCheckout release secret checkout
func (auth *Client) ReleaseCheckout(checkoutId string) error {
	_, err := auth.api.
		URL("/authorizer/api/v1/secrets/checkouts/%s/release", url.PathEscape(checkoutId)).
		Post(nil)

	return err
}
