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
func (c *Authorizer) GetCACertificates(opts ...filters.Option) ([]CA, error) {
	cas := []CA{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/cas").
		Query(params).
		Get(&cas)

	return cas, err
}

// DownloadCACertificate fetch authorizers root certificate as a download object.
func (c *Authorizer) DownloadCACertificate(caId, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/cas/%s", caId).
		Download(filename)

	return err
}

// DownloadCertificateRevocationList fetch authorizer CA certificate revocation list as a download object.
func (c *Authorizer) DownloadCertificateRevocationList(caId, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/cas/%s/crl", caId).
		Download(filename)

	return err
}

// RenewCAKey renew a CA key.
func (c *Authorizer) RenewCAKey(accessGroupId string) (response.Identifier, error) {
	caResponse := response.Identifier{}

	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s/cas", accessGroupId).
		Post(nil, &caResponse)

	return caResponse, err
}

// RevokeCAKey revoke a CA key.
func (c *Authorizer) RevokeCAKey(accessGroupId, caId string) error {
	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s/cas/%s", accessGroupId, caId).
		Delete()

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
func (c *Authorizer) GetPrincipals() (*Principal, error) {
	principals := &Principal{}

	_, err := c.api.
		URL("/authorizer/api/v1/principals").
		Get(&principals)

	return principals, err
}

// GetPrincipal get principal by its group id.
func (c *Authorizer) GetPrincipal(groupId string, opts ...filters.Option) (*Principal, error) {
	principal := &Principal{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/principals/%s", groupId).
		Query(params).
		Get(&principal)

	return principal, err
}

// DeletePrincipalKey delete the principal key by its group Id.
func (c *Authorizer) DeletePrincipalKey(groupId string, opts ...filters.Option) error {
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/principals/%s", groupId).
		Query(params).
		Delete()

	return err
}

// CreatePrincipalKey create a principal key pair.
func (c *Authorizer) CreatePrincipalKey(groupId string) (*Principal, error) {
	principal := &Principal{}

	_, err := c.api.
		URL("/authorizer/api/v1/principals/%s/create", groupId).
		Post(nil, &principal)

	return principal, err
}

// ImportPrincipalKey import a principal key pair.
func (c *Authorizer) ImportPrincipalKey(groupId string, key *PrincipalKeyImport) (*Principal, error) {
	principal := &Principal{}

	_, err := c.api.
		URL("/authorizer/api/v1/principals/%s/import", groupId).
		Post(&key, &principal)

	return principal, err
}

// SignPrincipalKey get a principal key signature.
func (c *Authorizer) SignPrincipalKey(groupId string, sign *PrincipalKeySign, opts ...filters.Option) (*Signature, error) {
	signature := &Signature{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/principals/%s/sign", groupId).
		Query(params).
		Post(&sign, &signature)

	return signature, err
}

// MARK: Extender
// GetExtenderCACertificates gets authorizers extender CA certificates.
func (c *Authorizer) GetExtenderCACertificates(opts ...filters.Option) ([]CA, error) {
	certificates := []CA{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/extender/cas").
		Query(params).
		Get(&certificates)

	return certificates, err
}

// GetExtenderCACertificate gets authorizers extender CA certificate by id.
func (c *Authorizer) GetExtenderCACertificate(id string) (*CA, error) {
	certificate := &CA{}

	_, err := c.api.
		URL("/authorizer/api/v1/extender/cas/%s", id).
		Get(&certificate)

	return certificate, err
}

// DownloadExtenderCertificateCRL fetch authorizer CA certificate revocation list as a download object.
func (c *Authorizer) DownloadExtenderCertificateCRL(filename, id string) error {
	err := c.api.
		URL("/authorizer/api/v1/extender/cas/%s/crl", id).
		Download(filename)

	return err
}

// EnrollExtenderCertificate enroll certificate from extender CA.
func (c *Authorizer) EnrollExtenderCertificate(request *CertificateEnroll) (*CertificateEnrollResponse, error) {
	enroll := &CertificateEnrollResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/extender/enroll").
		Post(&request, &enroll)

	return enroll, err
}

// RevokeExtenderCertificate revoke certificate.
func (c *Authorizer) RevokeExtenderCertificate(request *CertificateRevocation) (*CertificateRevocationResponse, error) {
	revoke := &CertificateRevocationResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/extender/revoke").
		Post(&request, &revoke)

	return revoke, err
}

// GetExtenderConfigSessions get extenders config session ids.
func (c *Authorizer) GetExtenderConfigSessions(trustedClientId string) (*SessionIdResponse, error) {
	sessionIds := &SessionIdResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/extender/conf/%s", trustedClientId).
		Post(nil, &sessionIds)

	return sessionIds, err
}

// DownloadExtenderConfig fetch a pre-configured extender config as a download object.
func (c *Authorizer) DownloadExtenderConfig(trustedClientId, sessionId, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/extender/conf/%s/%s", trustedClientId, sessionId).
		Download(filename)

	return err
}

// MARK: Deploy
// GetDeployScriptSessions get deploy script session ids.
func (c *Authorizer) GetDeployScriptSessions(trustedClientId string) (*SessionIdResponse, error) {
	sessionIds := &SessionIdResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/deploy/%s", trustedClientId).
		Post(nil, &sessionIds)

	return sessionIds, err
}

// DownloadDeployScript fetch a pre-configured deployment script.
func (c *Authorizer) DownloadDeployScript(trustedClientId, sessionId, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/deploy/%s/%s", trustedClientId, sessionId).
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
func (c *Authorizer) GetCarrierConfigSessions(trustedClientId string) (*SessionIdResponse, error) {
	sessionIds := &SessionIdResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/carrier/conf/%s", trustedClientId).
		Post(nil, &sessionIds)

	return sessionIds, err
}

// DownloadCarrierConfig fetch a pre-configured carrier config.
func (c *Authorizer) DownloadCarrierConfig(trustedClientId, sessionId, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/carrier/conf/%s/%s", trustedClientId, sessionId).
		Download(filename)

	return err
}

// MARK: Web-Proxy
// GetWebProxyCACertificates gets authorizer's web proxy CA certificates.
func (c *Authorizer) GetWebProxyCACertificates(opts ...filters.Option) ([]CA, error) {
	certificates := []CA{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/authorizer/api/v1/icap/cas").
		Query(params).
		Get(&certificates)

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

// EnrollWebProxyCertificate enroll certificate from web proxy CA.
func (c *Authorizer) EnrollWebProxyCertificate(request *CertificateEnroll) (*CertificateEnrollResponse, error) {
	enroll := &CertificateEnrollResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/icap/enroll").
		Post(&request, &enroll)

	return enroll, err
}

// RevokeWebProxyCertificate revoke certificate.
func (c *Authorizer) RevokeWebProxyCertificate(request *CertificateRevocation) (*CertificateRevocationResponse, error) {
	revoke := &CertificateRevocationResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/icap/revoke").
		Post(&request, &revoke)

	return revoke, err
}

// GetWebProxyConfigSessions get web proxy config session ids.
func (c *Authorizer) GetWebProxyConfigSessions(trustedClientId string) (*SessionIdResponse, error) {
	sessionIds := &SessionIdResponse{}

	_, err := c.api.
		URL("/authorizer/api/v1/icap/conf/%s", trustedClientId).
		Post(nil, &sessionIds)

	return sessionIds, err
}

// DownloadWebProxyConfig fetch a pre-configured web proxy config as a download object.
func (c *Authorizer) DownloadWebProxyConfig(trustedClientId, sessionId, filename string) error {
	err := c.api.
		URL("/authorizer/api/v1/icap/conf/%s/%s", trustedClientId, sessionId).
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
func (c *Authorizer) GetAccessGroup(accessGroupId string) (*AccessGroup, error) {
	accessGroup := &AccessGroup{}

	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s", accessGroupId).
		Get(&accessGroup)

	return accessGroup, err
}

// UpdateAccessGroup update access group by id.
func (c *Authorizer) UpdateAccessGroup(accessGroupId string, update *AccessGroup) error {
	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s", accessGroupId).
		Put(update)

	return err
}

// DeleteAccessGroup delete access group by id.
func (c *Authorizer) DeleteAccessGroup(accessGroupId string) error {
	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s", accessGroupId).
		Delete()

	return err
}

// RenewAccessGroupCAKey renew access group CA key.
func (c *Authorizer) RenewAccessGroupCAKey(accessGroupId string) (string, error) {
	var keyId string

	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s/cas", accessGroupId).
		Post(nil, &keyId)

	return keyId, err
}

// RevokeAccessGroupCAKey revoke access group CA key.
func (c *Authorizer) RevokeAccessGroupCAKey(accessGroupId string, caId string) error {
	_, err := c.api.
		URL("/authorizer/api/v1/accessgroups/%s/cas/%s", accessGroupId, caId).
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
func (c *Authorizer) GetCert(certId string) (*ApiCertificate, error) {
	cert := &ApiCertificate{}

	_, err := c.api.
		URL("/authorizer/api/v1/cert/%s", certId).
		Get(&cert)

	return cert, err
}

// MARK: Secrets
// GetAccountSecrets get all account secrets.
func (c *Authorizer) GetAccountSecrets(opts ...filters.Option) (*response.ResultSet[AccountSecret], error) {
	secrets := &response.ResultSet[AccountSecret]{}
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
func (c *Authorizer) SearchAccountSecrets(search *AccountSecretSearch, opts ...filters.Option) (*response.ResultSet[AccountSecret], error) {
	secrets := &response.ResultSet[AccountSecret]{}
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
func (c *Authorizer) GetSecretCheckout(checkoutId string) (*Checkout, error) {
	checkout := &Checkout{}

	_, err := c.api.
		URL("/authorizer/api/v1/secrets/checkouts/%s", checkoutId).
		Get(&checkout)

	return checkout, err
}

// ReleaseSecretCheckout release secret checkout.
func (c *Authorizer) ReleaseSecretCheckout(checkoutId string) error {
	_, err := c.api.
		URL("/authorizer/api/v1/secrets/checkouts/%s/release", checkoutId).
		Post(nil)

	return err
}
