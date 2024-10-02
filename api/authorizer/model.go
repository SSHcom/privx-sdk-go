//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package authorizer

import "time"

// CAparams ca query parameter definition.
type CAParams struct {
	AccessGroupId string `json:"access_group_id,omitempty"`
}

// PrincipalParams principal query parameter definition.
type PrincipalParams struct {
	KeyId string `json:"key_id,omitempty"`
}

// CertTemplateParams certificate template query parameter definition.
type CertTemplateParams struct {
	Service string `json:"service,omitempty"`
}

// ApiCertificate api certificate definition.
type ApiCertificate struct {
	Type              string `json:"type,omitempty"`
	Id                string `json:"id,omitempty"`
	Serial            string `json:"serial"`
	OwnerId           string `json:"owner_id,omitempty"`
	Revoked           string `json:"revoked,omitempty"`
	RevocationReason  string `json:"revocation_reason,omitempty"`
	Cert              string `json:"cert,omitempty"`
	Chain             string `json:"chain,omitempty"`
	Issuer            string `json:"issuer,omitempty"`
	Subject           string `json:"subject,omitempty"`
	NotBefore         string `json:"not_before,omitempty"`
	NotAfter          string `json:"not_after,omitempty"`
	KeyUsage          string `json:"key_usage,omitempty"`
	BasicConstraints  string `json:"basic_constraints,omitempty"`
	Extensions        string `json:"extensions,omitempty"`
	FingerPrintSHA1   string `json:"fingerprint_sha1,omitempty"`
	FingerPrintSHA256 string `json:"fingerprint_sha256,omitempty"`
	SubjectKeyId      string `json:"subject_key_id,omitempty"`
	AuthorityKeyId    string `json:"authority_key_id,omitempty"`
	Status            string `json:"status"`
}

// ApiCertificateSearch api certificate search definition.
type ApiCertificateSearch struct {
	Type           string `json:"type"`
	Id             string `json:"id,omitempty"`
	KeyId          string `json:"key_id,omitempty"`
	OwnerId        string `json:"owner_id,omitempty"`
	Subject        string `json:"subject,omitempty"`
	Issuer         string `json:"issuer,omitempty"`
	NotBefore      string `json:"not_before,omitempty"`
	NotAfter       string `json:"not_after,omitempty"`
	IncludeRevoked bool   `json:"include_revoked,omitempty"`
	IncludeExpired bool   `json:"include_expired,omitempty"`
}

// TrustAnchor trust anchor definition
type TrustAnchor struct {
	TrustAnchor       string `json:"trust_anchor"`
	TrustAnchorSHA1   string `json:"trust_anchor_sha1,omitempty"`
	TrustAnchorSHA256 string `json:"trust_anchor_sha256,omitempty"`
}

// CertTemplate certification template definition
type CertTemplate struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Service           string   `json:"service"`
	Type              string   `json:"type"`
	KeyId             string   `json:"key_id,omitempty"`
	RsaSignatureTypes []string `json:"rsa_signature_types,omitempty"`
	Principals        []string `json:"principals,omitempty"`
	Extensions        []string `json:"extensions,omitempty"`
}

// SessionIdResponse session id response definition.
type SessionIdResponse struct {
	SessionId string `json:"session_id"`
}

// Signature principal key signature response definition.
type Signature struct {
	Signature    string `json:"signature"`
	ResponseCode int    `json:"response_code,omitempty"`
	Message      string `json:"message,omitempty"`
}

// PrincipalKeySign principal key sign request definition.
type PrincipalKeySign struct {
	Id        string `json:"id"`
	GroupId   string `json:"group_id"`
	PublicKey string `json:"publicKey"`
	Type      string `json:"type"`
	Data      string `json:"data"`
}

// PrincipalKeyImport principal key import request definition.
type PrincipalKeyImport struct {
	Algorithm string `json:"algorithm"`
	Data      string `json:"data"`
}

// ApiIdentities end user authorization request definition.
type ApiIdentities struct {
	PublicKey string `json:"public_key,omitempty"`
	HostId    string `json:"host_id,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	Username  string `json:"username,omitempty"`
	Service   string `json:"service,omitempty"`
	RoleId    string `json:"role_id,omitempty"`
}

// ApiIdentitiesResponse api identities response definition.
type ApiIdentitiesResponse struct {
	Certificates  []ApiSshCertificate `json:"certificates"`
	PrincipalKeys []ApiSshKey         `json:"principal_keys"`
	Passphrase    string              `json:"passphrase,omitempty"`
	ResponseCode  int                 `json:"response_code"`
	Message       string              `json:"message"`
}

// ApiSshCertificate api ssh certificate definition.
type ApiSshCertificate struct {
	Type       string   `json:"type"`
	Data       string   `json:"data"`
	DataString string   `json:"data_string"`
	Chain      []string `json:"chain"`
}

// ApiSshKey api ssh key definition.
type ApiSshKey struct {
	Id              string `json:"id"`
	GroupId         string `json:"group_id,omitempty"`
	Type            string `json:"type,omitempty"`
	Comment         string `json:"comment,omitempty"`
	PublicKey       string `json:"public_key,omitempty"`
	PublicKeyString string `json:"public_key_string,omitempty"`
	Size            int    `json:"size,omitempty"`
}

// CA root certificate definition.
type CA struct {
	Id                string `json:"id"`
	GroupId           string `json:"group_id"`
	AccessGroupId     string `json:"access_group_id,omitempty"`
	Type              string `json:"type"`
	Size              int    `json:"size"`
	PublicKey         string `json:"public_key"`
	Comment           string `json:"comment,omitempty"`
	PublicKeyString   string `json:"public_key_string"`
	X509Certificate   string `json:"x509_certificate,omitempty"`
	Subject           string `json:"subject,omitempty"`
	Issuer            string `json:"issuer,omitempty"`
	SerialNumber      string `json:"serial,omitempty"`
	NotBefore         string `json:"not_before,omitempty"`
	NotAfter          string `json:"not_after,omitempty"`
	FingerPrintSHA1   string `json:"fingerprint_sha1,omitempty"`
	FingerPrintSHA256 string `json:"fingerprint_sha256,omitempty"`
}

// Principal principal definition.
type Principal struct {
	Type            string `json:"type"`
	Id              string `json:"id"`
	GroupId         string `json:"group_id"`
	Comment         string `json:"comment,omitempty"`
	PublicKey       string `json:"public_key"`
	PublicKeyString string `json:"public_key_string"`
	Size            int    `json:"size"`
}

// CertificateEnroll certificate enroll request definition.
type CertificateEnroll struct {
	CAId  string `json:"ca_id,omitempty"`
	CSR   string `json:"csr"`
	Owner string `json:"owner"`
}

// CertificateEnrollResponse certificate enroll response definition.
type CertificateEnrollResponse struct {
	Id   string `json:"id"`
	Cert string `json:"cert"`
	CA   string `json:"ca"`
}

// CertificateRevocation certificate revocation request definition.
type CertificateRevocation struct {
	Reason string `json:"reason,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Cert   string `json:"cert,omitempty"`
}

// CertificateRevocationResponse certificate revocation response definition.
type CertificateRevocationResponse struct {
	Ids []string `json:"ids"`
}

// AccessGroup access group definition.
type AccessGroup struct {
	ID                               string `json:"id,omitempty"`
	Name                             string `json:"name,omitempty"`
	Comment                          string `json:"comment,omitempty"`
	HostCertificateTrustAnchors      string `json:"host_certificate_trust_anchors"`
	WinRMHostCertificateTrustAnchors string `json:"winrm_host_certificate_trust_anchors"`
	DBHostCertificateTrustAnchors    string `json:"db_host_certificate_trust_anchors"`
	CAId                             string `json:"ca_id,omitempty"`
	PrimaryCAId                      string `json:"primary_ca_id"`
	Author                           string `json:"author,omitempty"`
	Created                          string `json:"created,omitempty"`
	Updated                          string `json:"updated,omitempty"`
	UpdatedBy                        string `json:"updated_by,omitempty"`
	Default                          bool   `json:"default,omitempty"`
}

// AccessGroupSearch access group request search body definition.
type AccessGroupSearch struct {
	Keywords string `json:"keywords,omitempty"`
}

// AccountSecret account secret definition.
type AccountSecret struct {
	Path         string             `json:"path"`
	Type         string             `json:"type"`
	Username     string             `json:"username"`
	Email        string             `json:"email,omitempty"`
	FullName     string             `json:"full_name,omitempty"`
	TargetDomain TargetDomainHandle `json:"target_domain,omitempty"`
	Host         HostPrincipals     `json:"host,omitempty"`
	Created      string             `json:"created,omitempty"`
	Updated      string             `json:"updated,omitempty"`
}

// TargetDomainHandle target domain handle definition.
type TargetDomainHandle struct {
	Id      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

// HostPrincipals host principals definition.
type HostPrincipals struct {
	Id         string   `json:"id"`
	Addresses  []string `json:"addresses"`
	CommonName string   `json:"common_name,omitempty"`
	ExternalId string   `json:"external_id,omitempty"`
	InstanceId string   `json:"instance_id,omitempty"`
}

// AccountSecretSearch account secret search request definition.
type AccountSecretSearch struct {
	Keywords string `json:"keywords"`
	HostId   string `json:"host_id,omitempty"`
	Username string `json:"username,omitempty"`
}

// Checkout checkout definition.
type Checkout struct {
	Id               string         `json:"id"`
	Path             string         `json:"path"`
	Type             string         `json:"type"`
	Expires          string         `json:"expires"`
	Created          string         `json:"created"`
	ExplicitCheckout bool           `json:"explicit_checkout"`
	Secrets          []Secrets      `json:"secrets"`
	Username         string         `json:"username"`
	Email            string         `json:"email,omitempty"`
	FullName         string         `json:"full_name,omitempty"`
	Host             HostPrincipals `json:"host,omitempty"`
	TargetDomain     TargetDomain   `json:"target_domain,omitempty"`
	ManagedAccountId string         `json:"managed_account_id,omitempty"`
	UserId           string         `json:"user_id"`
}

// CheckoutRequest checkout request definition.
type CheckoutRequest struct {
	Path string `json:"path"`
}

// Secrets secrets definition.
type Secrets struct {
	Version int       `json:"version"`
	Secret  string    `json:"secret"`
	Created time.Time `json:"created"`
}

// TargetDomain target domain definition.
type TargetDomain struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}
