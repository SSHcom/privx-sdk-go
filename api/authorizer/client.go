//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package authorizer

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Authorizer is a authorizer client instance.
type Authorizer struct {
	api restapi.Connector
}

// New authorizer client constructor.
func New(api restapi.Connector) *Authorizer {
	return &Authorizer{api: api}
}

// MARK: Status
// Status get authorizer microservice status.
func (c *Authorizer) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/authorizer/api/v1/status").
		Get(status)

	return status, err
}

// MARK: CAS
// GetCACertificates get authorizers root certificates.
// Note, the v1 endpoint doesn't return the count as part of the response body,
// this will change with v2. Until then, we will handle it internally within the SDK.
func (c *Authorizer) GetCACertificates(opts ...filters.Option) (*response.ResultSet[CA], error) {
	cas := []CA{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/cas").
		Query(params).
		Get(&cas)

	// v1 endpoint does not return count,
	// return count internally in sdk until v2 is introduced
	certs := &response.ResultSet[CA]{
		Items: cas,
		Count: len(cas),
	}

	return certs, err
}

// DownloadCACertificate fetch authorizers root certificate as a download object.
func (c *Authorizer) DownloadCACertificate(caID, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/cas/%s", caID).
		Download(filename)

	return err
}

// CAConfig get authorizers root certificate config by ca type.
func (c *Authorizer) CAConfig(caType string) (ComponentCaConfig, error) {
	caConf := ComponentCaConfig{}

	_, err := c.api.
		URL("/authorizer/api/v1/%s/cas/config", caType).
		Get(&caConf)

	return caConf, err
}

// UpdateCAConfig update authorizers root certificate config by ca type.
func (c *Authorizer) UpdateCAConfig(caType string, caConf ComponentCaConfig) error {
	_, err := c.api.
		URL("/authorizer/api/v1/%s/cas/config", caType).
		Put(&caConf)

	return err
}

// DownloadCertificateRevocationList fetch authorizer CA certificate revocation list as a download object.
func (c *Authorizer) DownloadCertificateRevocationList(caID, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/cas/%s/crl", caID).
		Download(filename)

	return err
}

// GetTargetHostCredentials get target host credentials for the user.
func (c *Authorizer) GetTargetHostCredentials(request *ApiIdentities) (*ApiIdentitiesResponse, error) {
	principal := &ApiIdentitiesResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/ca/authorize").
		Post(&request, &principal)

	return principal, err
}

// MARK: Principals
// GetPrincipals get defined principals.
// Note, the v1 endpoint doesn't return the count as part of the response body,
// this will change with v2. Until then, we will handle it internally within the SDK.
func (c *Authorizer) GetPrincipals() (*response.ResultSet[Principal], error) {
	p := []Principal{}

	_, err := c.api.
		URL("/authorizer/api/v1/principals").
		Get(&p)

	// v1 endpoint does not return count,
	// return count internally in sdk until v2 is introduced
	principals := &response.ResultSet[Principal]{
		Count: len(p),
		Items: p,
	}

	return principals, err
}

// GetPrincipal get principal by its group id.
func (c *Authorizer) GetPrincipal(groupID string, opts ...filters.Option) (*Principal, error) {
	principal := &Principal{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/principals/%s", groupID).
		Query(params).
		Get(&principal)

	return principal, err
}

// DeletePrincipalKey delete the principal key by its group id.
func (c *Authorizer) DeletePrincipalKey(groupID string, opts ...filters.Option) error {
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/principals/%s", groupID).
		Query(params).
		Delete()

	return err
}

// CreatePrincipalKey create a principal key pair.
func (c *Authorizer) CreatePrincipalKey(groupID string) (*Principal, error) {
	principal := &Principal{}

	_, err := c.api.
		URL("/authorizer/api/v1/principals/%s/create", groupID).
		Post(nil, &principal)

	return principal, err
}

// ImportPrincipalKey import a principal key pair.
func (c *Authorizer) ImportPrincipalKey(groupID string, key *PrincipalKeyImport) (*Principal, error) {
	principal := &Principal{}

	_, err := c.api.
		URL("/authorizer/api/v1/principals/%s/import", groupID).
		Post(&key, &principal)

	return principal, err
}

// SignPrincipalKey get a principal key signature.
func (c *Authorizer) SignPrincipalKey(groupID string, sign *PrincipalKeySign, opts ...filters.Option) (*Signature, error) {
	signature := &Signature{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/principals/%s/sign", groupID).
		Query(params).
		Post(&sign, &signature)

	return signature, err
}

// MARK: Extender
// GetExtenderCACertificates gets authorizers extender CA certificates.
// Note, the v1 endpoint doesn't return the count as part of the response body,
// this will change with v2. Until then, we will handle it internally within the SDK.
func (c *Authorizer) GetExtenderCACertificates(opts ...filters.Option) (*response.ResultSet[CA], error) {
	cs := []CA{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/extender/cas").
		Query(params).
		Get(&cs)

	// v1 endpoint does not return count,
	// return count internally in sdk until v2 is introduced
	certificates := &response.ResultSet[CA]{
		Count: len(cs),
		Items: cs,
	}

	return certificates, err
}

// DownloadExtenderCACertificate fetch authorizers extender CA certificate by id as a download object.
func (c *Authorizer) DownloadExtenderCACertificate(filename, id string) error {
	err := c.api.
		URL("/authorizer/api/v1/extender/cas/%s", id).
		Download(filename)

	return err
}

// DownloadExtenderCertificateCRL fetch authorizer CA certificate revocation list as a download object.
func (c *Authorizer) DownloadExtenderCertificateCRL(filename, id string) error {
	err := c.api.
		URL("/authorizer/api/v1/extender/cas/%s/crl", id).
		Download(filename)

	return err
}

// GetExtenderConfigSessions get extenders config session ids.
func (c *Authorizer) GetExtenderConfigSessions(trustedClientID string) (*SessionIDResponse, error) {
	sessionIDs := &SessionIDResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/extender/conf/%s", trustedClientID).
		Post(nil, &sessionIDs)

	return sessionIDs, err
}

// DownloadExtenderConfig fetch a pre-configured extender config as a download object.
func (c *Authorizer) DownloadExtenderConfig(trustedClientID, sessionID, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/extender/conf/%s/%s", trustedClientID, sessionID).
		Download(filename)

	return err
}

// MARK: Deploy
// GetDeployScriptSessions get deploy script session ids.
func (c *Authorizer) GetDeployScriptSessions(trustedClientID string) (*SessionIDResponse, error) {
	sessionIDs := &SessionIDResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/deploy/%s", trustedClientID).
		Post(nil, &sessionIDs)

	return sessionIDs, err
}

// DownloadDeployScript fetch a pre-configured deployment script.
func (c *Authorizer) DownloadDeployScript(trustedClientID, sessionID, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/deploy/%s/%s", trustedClientID, sessionID).
		Download(filename)

	return err
}

// DownloadPrincipalCommandScript fetch the principals_command.sh script.
func (c *Authorizer) DownloadPrincipalCommandScript(filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/deploy/principals_command.sh").
		Download(filename)

	return err
}

// MARK: Carrier
// // GetCarrierConfigSessions get carrier config session ids.
func (c *Authorizer) GetCarrierConfigSessions(trustedClientID string) (*SessionIDResponse, error) {
	sessionIDs := &SessionIDResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/carrier/conf/%s", trustedClientID).
		Post(nil, &sessionIDs)

	return sessionIDs, err
}

// DownloadCarrierConfig fetch a pre-configured carrier config.
func (c *Authorizer) DownloadCarrierConfig(trustedClientID, sessionID, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/carrier/conf/%s/%s", trustedClientID, sessionID).
		Download(filename)

	return err
}

// MARK: Web-Proxy
// GetWebProxyCACertificates gets authorizer's web proxy CA certificates.
// Note, the v1 endpoint doesn't return the count as part of the response body,
// this will change with v2. Until then, we will handle it internally within the SDK.
func (c *Authorizer) GetWebProxyCACertificates(opts ...filters.Option) (*response.ResultSet[CA], error) {
	cs := []CA{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/icap/cas").
		Query(params).
		Get(&cs)

	// v1 endpoint does not return count,
	// return count internally in sdk until v2 is introduced
	certificates := &response.ResultSet[CA]{
		Count: len(cs),
		Items: cs,
	}

	return certificates, err
}

// GetWebProxyCACertificate gets authorizer's web proxy CA certificate by id.
func (c *Authorizer) GetWebProxyCACertificate(id string) (*CA, error) {
	certificate := &CA{}

	_, err := c.api.
		URL("/authorizer/api/v1/icap/cas/%s", id).
		Get(&certificate)

	return certificate, err
}

// DownloadWebProxyCertificateCRL fetch authorizer CA certificate revocation list as a download object.
func (c *Authorizer) DownloadWebProxyCertificateCRL(filename, id string) error {
	err := c.api.
		URL("/authorizer/api/v1/icap/cas/%s/crl", id).
		Download(filename)

	return err
}

// GetWebProxyConfigSessions get web proxy config session ids.
func (c *Authorizer) GetWebProxyConfigSessions(trustedClientID string) (*SessionIDResponse, error) {
	sessionIDs := &SessionIDResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/icap/conf/%s", trustedClientID).
		Post(nil, &sessionIDs)

	return sessionIDs, err
}

// DownloadWebProxyConfig fetch a pre-configured web proxy config as a download object.
func (c *Authorizer) DownloadWebProxyConfig(trustedClientID, sessionID, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/icap/conf/%s/%s", trustedClientID, sessionID).
		Download(filename)

	return err
}

// MARK: Templates
// GetCertTemplates returns the certificate authentication templates.
func (c *Authorizer) GetCertTemplates(opts ...filters.Option) (*response.ResultSet[CertTemplate], error) {
	templates := &response.ResultSet[CertTemplate]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/cert/templates").
		Query(params).
		Get(&templates)

	return templates, err
}

// MARK: Trust Anchors
// GetSSLTrustAnchor returns the SSL trust anchor.
func (c *Authorizer) GetSSLTrustAnchor() (*TrustAnchor, error) {
	trustAnchor := &TrustAnchor{}

	_, err := c.api.
		URL("/authorizer/api/v1/ssl-trust-anchor").
		Get(&trustAnchor)

	return trustAnchor, err
}

// GetExtenderTrustAnchor returns the extender trust anchor.
func (c *Authorizer) GetExtenderTrustAnchor() (*TrustAnchor, error) {
	trustAnchor := &TrustAnchor{}

	_, err := c.api.
		URL("/authorizer/api/v1/extender-trust-anchor").
		Get(&trustAnchor)

	return trustAnchor, err
}

// MARK: Access Groups
// GetAccessGroups get all access group.
func (c *Authorizer) GetAccessGroups(opts ...filters.Option) (*response.ResultSet[AccessGroup], error) {
	accessGroups := &response.ResultSet[AccessGroup]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups").
		Query(params).
		Get(&accessGroups)

	return accessGroups, err
}

// CreateAccessGroup create access group.
func (c *Authorizer) CreateAccessGroup(accessGroup *AccessGroup) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups").
		Post(&accessGroup, &identifier)

	return identifier, err
}

// SearchAccessGroups search for access groups.
func (c *Authorizer) SearchAccessGroups(search *AccessGroupSearch, opts ...filters.Option) (*response.ResultSet[AccessGroup], error) {
	accessGroups := &response.ResultSet[AccessGroup]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/search").
		Query(params).
		Post(search, &accessGroups)

	return accessGroups, err
}

// GetAccessGroup get access group by id.
func (c *Authorizer) GetAccessGroup(accessGroupID string) (*AccessGroup, error) {
	accessGroup := &AccessGroup{}

	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s", accessGroupID).
		Get(&accessGroup)

	return accessGroup, err
}

// UpdateAccessGroup update access group by id.
func (c *Authorizer) UpdateAccessGroup(accessGroupID string, update *AccessGroup) error {
	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s", accessGroupID).
		Put(update)

	return err
}

// DeleteAccessGroup delete access group by id.
func (c *Authorizer) DeleteAccessGroup(accessGroupID string) error {
	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s", accessGroupID).
		Delete()

	return err
}

// RenewAccessGroupCAKey renew access group CA key.
func (c *Authorizer) RenewAccessGroupCAKey(accessGroupID string) (string, error) {
	var keyID string

	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s/cas", accessGroupID).
		Post(nil, &keyID)

	return keyID, err
}

// RevokeAccessGroupCAKey revoke access group CA key.
func (c *Authorizer) RevokeAccessGroupCAKey(accessGroupID string, caID string) error {
	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s/cas/%s", accessGroupID, caID).
		Delete()

	return err
}

// MARK: Certs
// SearchCerts search certificates.
func (c *Authorizer) SearchCerts(search *ApiCertificateSearch, opts ...filters.Option) (*response.ResultSet[ApiCertificate], error) {
	certs := &response.ResultSet[ApiCertificate]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/cert/search").
		Query(params).
		Post(search, &certs)

	return certs, err
}

// GetAllCertificates get all certificates.
func (c *Authorizer) GetAllCertificates() (*response.ResultSet[ApiCertificate], error) {
	certs := &response.ResultSet[ApiCertificate]{}

	_, err := c.api.
		URL("/authorizer/api/v1/cert").
		Get(&certs)

	return certs, err
}

// GetCert get certificate by id.
func (c *Authorizer) GetCert(certID string) (*ApiCertificate, error) {
	cert := &ApiCertificate{}

	_, err := c.api.
		URL("/authorizer/api/v1/cert/%s", certID).
		Get(&cert)

	return cert, err
}

// MARK: Secrets
// GetAccountSecrets get all account secrets.
func (c *Authorizer) GetAccountSecrets(opts ...filters.Option) (*response.ResultSet[HostAccountSecret], error) {
	secrets := &response.ResultSet[HostAccountSecret]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/secrets").
		Query(params).
		Get(&secrets)

	return secrets, err
}

// SearchAccountSecrets search for account secrets.
func (c *Authorizer) SearchAccountSecrets(search *AccountSecretSearch, opts ...filters.Option) (*response.ResultSet[HostAccountSecret], error) {
	secrets := &response.ResultSet[HostAccountSecret]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/secrets/search").
		Query(params).
		Post(search, &secrets)

	return secrets, err
}

// GetSecretCheckouts get secret checkouts.
func (c *Authorizer) GetSecretCheckouts(opts ...filters.Option) (*response.ResultSet[Checkout], error) {
	checkouts := &response.ResultSet[Checkout]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/secrets/checkouts").
		Query(params).
		Get(&checkouts)

	return checkouts, err
}

// CheckoutAccountSecret checkout account secret.
func (c *Authorizer) CheckoutAccountSecret(checkout CheckoutRequest) (*response.ResultSet[Checkout], error) {
	checkoutResp := &response.ResultSet[Checkout]{}

	_, err := c.api.
		URL("/authorizer/api/v1/secrets/checkouts").
		Post(checkout, &checkoutResp)

	return checkoutResp, err
}

// GetSecretCheckout get secret checkout by id.
func (c *Authorizer) GetSecretCheckout(checkoutID string) (*Checkout, error) {
	checkout := &Checkout{}

	_, err := c.api.
		URL("/authorizer/api/v1/secrets/checkouts/%s", checkoutID).
		Get(&checkout)

	return checkout, err
}

// ReleaseSecretCheckout release secret checkout.
func (c *Authorizer) ReleaseSecretCheckout(checkoutID string) error {
	_, err := c.api.
		URL("/authorizer/api/v1/secrets/checkouts/%s/release", checkoutID).
		Post(nil)

	return err
}
