//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package authorizer

import "time"

// Params query params definition
type Params struct {
	ResponseType  string `json:"response_type,omitempty"`
	ClientID      string `json:"client_id,omitempty"`
	State         string `json:"state,omitempty"`
	RedirectURI   string `json:"redirect_uri,omitempty"`
	UserAgent     string `json:"user_agent,omitempty"`
	OidcID        string `json:"oidc_id,omitempty"`
	AccessGroupID string `json:"access_group_id,omitempty"`
	KeyID         string `json:"key_id,omitempty"`
	Filter        string `json:"filter,omitempty"`
	Service       string `json:"service,omitempty"`
	Sortkey       string `json:"sortkey,omitempty"`
	Sortdir       string `json:"sortdir,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty"`
}

// SearchParams search params definition
type SearchParams struct {
	Keywords string `json:"keywords,omitempty"`
}

// APICertificate api certificate definition
type APICertificate struct {
	ID               string `json:"id,omitempty"`
	Type             string `json:"type,omitempty"`
	OwnerID          string `json:"owner_id,omitempty"`
	Revoked          string `json:"revoked,omitempty"`
	RevocationReason string `json:"revocation_reason,omitempty"`
	Cert             string `json:"cert,omitempty"`
	Chain            string `json:"chain,omitempty"`
}

// APICertificateSearch api certificate search definition
type APICertificateSearch struct {
	ID             string `json:"id,omitempty"`
	Type           string `json:"type,omitempty"`
	KeyID          string `json:"key_id,omitempty"`
	OwnerID        string `json:"owner_id,omitempty"`
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
	KeyID             string   `json:"key_id,omitempty"`
	RsaSignatureTypes []string `json:"rsa_signature_types,omitempty"`
	Principals        []string `json:"principals,omitempty"`
	Extensions        []string `json:"extensions,omitempty"`
}

// DownloadHandle download handle definition
type DownloadHandle struct {
	SessionID string `json:"session_id"`
}

// Signature signature  definition
type Signature struct {
	Signature string `json:"signature"`
}

// Credential end user authentication credentials definition
type Credential struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// PrincipalKeyImportRequest principal key import definition
type PrincipalKeyImportRequest struct {
	Algorithm string `json:"algorithm"`
	Data      string `json:"data"`
}

// AuthorizationRequest end user authorization request definition
type AuthorizationRequest struct {
	PublicKey string `json:"public_key,omitempty"`
	HostID    string `json:"host_id,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	Username  string `json:"username,omitempty"`
	Service   string `json:"service,omitempty"`
	RoleID    string `json:"role_id,omitempty"`
}

// Principal principal definition
type Principal struct {
	ID              string `json:"id"`
	GroupID         string `json:"group_id,omitempty"`
	Type            string `json:"type,omitempty"`
	Comment         string `json:"comment,omitempty"`
	PublicKey       string `json:"public_key,omitempty"`
	PublicKeyString string `json:"public_key_string,omitempty"`
	Size            int    `json:"size,omitempty"`
}

type ApiSshCertificate struct {
	Type       string   `json:"type"`
	Data       string   `json:"data"`
	DataString string   `json:"data_string"`
	Chain      []string `json:"chain"`
}
type ApiIdentitiesResponse struct {
	Certificates  []ApiSshCertificate `json:"certificates"`
	PrincipalKeys []Principal         `json:"principal_keys"`
	Passphrase    string              `json:"passphrase,omitempty"`
	ResponseCode  int                 `json:"response_code"`
	Message       string              `json:"message"`
}

// CA is root certificate representation
type CA struct {
	ID        string `json:"id"`
	GroupID   string `json:"group_id"`
	Type      string `json:"type"`
	Size      int    `json:"size"`
	PublicKey string `json:"public_key"`
	X509      string `json:"x509_certificate"`
}

// AccessGroup access group definition
type AccessGroup struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Comment   string `json:"comment,omitempty"`
	CAID      string `json:"ca_id,omitempty"`
	Author    string `json:"author,omitempty"`
	Created   string `json:"created,omitempty"`
	Updated   string `json:"updated,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	Default   bool   `json:"default,omitempty"`
}
type ApiCertificateSearchResponse struct {
	Count int                    `json:"count"`
	Items []ApiCertificateObject `json:"items"`
}

type ApiCertificateObject struct {
	Type              string `json:"type"`
	ID                string `json:"id"`
	Serial            string `json:"serial"`
	OwnerID           string `json:"owner_id,omitempty"`
	Revoked           string `json:"revoked,omitempty"`
	RevocationReason  string `json:"revocation_reason,omitempty"`
	Cert              string `json:"cert"`
	Chain             string `json:"chain"`
	Issuer            string `json:"issuer,omitempty"`
	Subject           string `json:"subject,omitempty"`
	NotBefore         string `json:"not_before,omitempty"`
	NotAfter          string `json:"not_after,omitempty"`
	KeyUsage          string `json:"key_usage,omitempty"`
	BasicConstraints  string `json:"basic_constraints,omitempty"`
	Extensions        string `json:"extensions,omitempty"`
	FingerPrintSHA1   string `json:"fingerprint_sha1,omitempty"`
	FingerPrintSHA256 string `json:"fingerprint_sha256,omitempty"`
	SubjectKeyID      string `json:"subject_key_id,omitempty"`
	AuthorityKeyID    string `json:"authority_key_id,omitempty"`
	ExpiryStatus      string `json:"expiry_status,omitempty"`
}

type AccountSecrets struct {
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

type TargetDomainHandle struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

type HostPrincipals struct {
	ID         string   `json:"id"`
	Addresses  []string `json:"addresses"`
	CommonName string   `json:"common_name,omitempty"`
	ExternalID string   `json:"external_id,omitempty"`
	InstanceID string   `json:"instance_id,omitempty"`
}

type AccountSecretsSearchRequest struct {
	Keywords string `json:"keywords"`
	HostID   string `json:"host_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type Checkout struct {
	ID               string         `json:"id"`
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
	ManagedAccountID string         `json:"managed_account_id,omitempty"`
	UserID           string         `json:"user_id"`
}

type CheckoutRequest struct {
	Path string `json:"path"`
}

type Secrets struct {
	Version int       `json:"version"`
	Secret  string    `json:"secret"`
	Created time.Time `json:"created"`
}

type TargetDomain struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}
