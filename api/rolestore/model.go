//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

import (
	authmodel "github.com/SSHcom/privx-sdk-go/api/auth"
)

// Params struct for pagination queries.
type Params struct {
	Sortdir   string `json:"sortdir,omitempty"`
	Sortkey   string `json:"sortkey,omitempty"`
	Tokencode string `json:"tokencode,omitempty"`
	Refresh   bool   `json:"refresh,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	TTL       int    `json:"ttl,omitempty"`
}

// ResolveAuthorizedKey struct for resolving authorized key.
type ResolveAuthorizedKey struct {
	Username  string `json:"username,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

// PrivateKey principal privat key definition
type PrivateKey struct {
	ID         string `json:"id,omitempty"`
	PrivateKey string `json:"private_key,omitempty"`
}

// PrincipalKey principal key definition
type PrincipalKey struct {
	ID        string `json:"id,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

// LogconfCollector logconf collectors definition
type LogconfCollector struct {
	ID                       string `json:"id,omitempty"`
	Name                     string `json:"name,omitempty"`
	Type                     string `json:"type,omitempty"`
	Updated                  string `json:"updated,omitempty"`
	StatusCode               string `json:"status_code,omitempty"`
	StatusText               string `json:"status_text,omitempty"`
	AWSLogRegion             string `json:"aws_log_region,omitempty"`
	IAMAccessKeyID           string `json:"iam_access_key_id,omitempty"`
	IAMSecretAccessKey       string `json:"iam_secret_access_key,omitempty"`
	IAMSessionToken          string `json:"iam_session_token,omitempty"`
	AzureEventHubsNamespace  string `json:"azure_event_hubs_namespace,omitempty"`
	AzureResourceGroupName   string `json:"azure_resource_group_name,omitempty"`
	AzureSubscriptionID      string `json:"azure_subscription_id,omitempty"`
	AzureEventHubName        string `json:"azure_event_hub_name,omitempty"`
	AzureTenantID            string `json:"azure_tenant_id,omitempty"`
	AzureClientID            string `json:"azure_client_id,omitempty"`
	AzureClientSecret        string `json:"azure_client_secret,omitempty"`
	AzureSasConnectionString string `json:"azure_sas_connection_string,omitempty"`
	Enabled                  bool   `json:"enabled,omitempty"`
}

// AWSToken aws token definition
type AWSToken struct {
	AccessKeyID     string   `json:"access_key_id,omitempty"`
	SecretAccessKey string   `json:"secret_access_key,omitempty"`
	SessionToken    string   `json:"session_token,omitempty"`
	Expires         string   `json:"expires,omitempty"`
	Descriptions    []string `json:"descriptions,omitempty"`
}

// AWSRoleLink aws role definition.
type AWSRoleLink struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	ARN         string    `json:"arn,omitempty"`
	Updated     string    `json:"updated,omitempty"`
	Description string    `json:"description,omitempty"`
	Source      string    `json:"source,omitempty"`
	Status      string    `json:"status,omitempty"`
	Roles       []RoleRef `json:"roles,omitempty"`
}

// Connection source connection definition
type Connection struct {
	Type                   string   `json:"type,omitempty"`
	Address                string   `json:"address,omitempty"`
	AccessKeyID            string   `json:"iam_access_key_id,omitempty"`
	SecretKey              string   `json:"iam_secret_access_key,omitempty"`
	SessionToken           string   `json:"iam_session_token,omitempty"`
	FetchRolePathPrefix    string   `json:"iam_fetch_role_path_prefix,omitempty"`
	GCConfig               string   `json:"google_cloud_config_json,omitempty"`
	OpenstackVersion       string   `json:"openstack_version,omitempty"`
	OpenStackEndpoint      string   `json:"openstack_endpoint,omitempty"`
	OpenStackUsername      string   `json:"openstack_username,omitempty"`
	OpenStackUserID        string   `json:"openstack_user_id,omitempty"`
	OpenStackPassword      string   `json:"openstack_password,omitempty"`
	OpenStackAPIkey        string   `json:"openstack_apikey,omitempty"`
	OpenStackDomainName    string   `json:"openstack_domainname,omitempty"`
	OpenStackDomainID      string   `json:"openstack_domainid,omitempty"`
	OpenStackTokenID       string   `json:"openstack_token_id,omitempty"`
	AzureBaseURL           string   `json:"azure_base_url,omitempty"`
	AzureSubscriptionID    string   `json:"azure_subscription_id,omitempty"`
	AzureTenantID          string   `json:"azure_tenant_id,omitempty"`
	AzureClientID          string   `json:"azure_client_id,omitempty"`
	AzureClientSecret      string   `json:"azure_client_secret,omitempty"`
	LDAPProtocol           string   `json:"ldap_protocol,omitempty"`
	LDAPBase               string   `json:"ldap_base,omitempty"`
	LDAPUserFilter         string   `json:"ldap_user_filter,omitempty"`
	LDAPBindDN             string   `json:"ldap_bind_dn,omitempty"`
	LDAPBindPassword       string   `json:"ldap_bind_password,omitempty"`
	LDAPUserDNPattern      string   `json:"ldap_user_dn_pattern,omitempty"`
	GoogleGsuiteDomain     string   `json:"google_gsuite_domain,omitempty"`
	GoogleGsuiteAdminEmail string   `json:"google_gsuite_domain_admin_email,omitempty"`
	OIDCIssuer             string   `json:"oidc_issuer,omitempty"`
	OIDCButtonTitle        string   `json:"oidc_button_title,omitempty"`
	OIDCClientID           string   `json:"oidc_client_id,omitempty"`
	OIDCClientSecret       string   `json:"oidc_client_secret,omitempty"`
	OIDCTagsAttributeName  string   `json:"oidc_tags_attribute_name,omitempty"`
	MFAType                string   `json:"mfa_type,omitempty"`
	MFAAddress             string   `json:"mfa_address,omitempty"`
	MFABaseDN              string   `json:"mfa_base_dn,omitempty"`
	DomainControllerFQDN   string   `json:"domain_controller_fqdn,omitempty"`
	KerberosTicket         string   `json:"kerberos_ticket,omitempty"`
	DomainControllerPort   int      `json:"domain_controller_port,omitempty"`
	MFAPort                int      `json:"mfa_port,omitempty"`
	Port                   int      `json:"port,omitempty"`
	EnableMachineAuth      bool     `json:"enable_machine_authentication,omitempty"`
	EnableUserAuth         bool     `json:"enable_user_authentication,omitempty"`
	OIDCEnabled            bool     `json:"oidc_enabled,omitempty"`
	FetchRoles             bool     `json:"iam_fetch_roles,omitempty"`
	AutoUpdate             bool     `json:"service_address_auto_update,omitempty"`
	OIDCScopesSecret       []string `json:"oidc_additional_scopes_secret,omitempty"`
	GCProjectIDs           []string `json:"google_cloud_project_ids,omitempty"`
	OpenStackTenantIDs     []string `json:"openstack_tenant_ids,omitempty"`
	OpenStackTenantNames   []string `json:"openstack_tenant_names,omitempty"`
}

// EUM external user mapping definition
type EUM struct {
	SourceID           string `json:"source_id,omitempty"`
	SourceSeaerchField string `json:"source_search_field,omitempty"`
}

// Source definitions - user and host directories
type Source struct {
	ID                          string                           `json:"id,omitempty"`
	Created                     string                           `json:"created,omitempty"`
	Updated                     string                           `json:"updated,omitempty"`
	UpdatedBy                   string                           `json:"updated_by,omitempty"`
	Author                      string                           `json:"author,omitempty"`
	Name                        string                           `json:"name,omitempty"`
	StatusCode                  string                           `json:"status_code,omitempty"`
	StatusText                  string                           `json:"status_text,omitempty"`
	Comment                     string                           `json:"comment,omitempty"`
	TTL                         int                              `json:"ttl,omitempty"`
	Enabled                     bool                             `json:"enabled,omitempty"`
	Tags                        []string                         `json:"tags,omitempty"`
	UsernamePattern             []string                         `json:"username_pattern,omitempty"`
	ExternalUserMapping         []EUM                            `json:"external_user_mapping,omitempty"`
	Connection                  Connection                       `json:"connection,omitempty"`
	SessionPasswordEnabled      bool                             `json:"session_password_enabled,omitempty"`
	SessionPasswordPolicy       *authmodel.SessionPasswordPolicy `json:"session_password_policy,omitempty"`
	ChildSessionAutoLogoutDelay int                              `json:"child_session_auto_logout_delay,omitempty"`
}

// User contains PrivX user information.
type User struct {
	ID                string          `json:"id,omitempty"`
	SourceUserID      string          `json:"source_user_id,omitempty"`
	Principal         string          `json:"principal,omitempty"`
	Source            string          `json:"source,omitempty"`
	FullName          string          `json:"full_name,omitempty"`
	Email             string          `json:"email,omitempty"`
	DistinguishedName string          `json:"distinguished_name,omitempty"`
	Created           string          `json:"created,omitempty"`
	Updated           string          `json:"updated,omitempty"`
	UpdatedBy         string          `json:"updated_by,omitempty"`
	Author            string          `json:"author,omitempty"`
	Comment           string          `json:"comment,omitempty"`
	GivenName         string          `json:"given_name,omitempty"`
	Job               string          `json:"job_title,omitempty"`
	Company           string          `json:"company,omitempty"`
	Department        string          `json:"department,omitempty"`
	Telephone         string          `json:"telephone,omitempty"`
	Locale            string          `json:"locale,omitempty"`
	StaleAccessToken  bool            `json:"stale_access_token,omitempty"`
	Permissions       []string        `json:"permissions,omitempty"`
	Tags              []string        `json:"tags"`
	MFA               MFA             `json:"mfa"`
	Roles             []Role          `json:"roles"`
	AuthorizedKeys    []AuthorizedKey `json:"authorized_keys,omitempty"`
}

// AuthorizedKey authorizednal key definition
type AuthorizedKey struct {
	ID            string   `json:"id,omitempty"`
	Username      string   `json:"username,omitempty"`
	UserID        string   `json:"user_id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Comment       string   `json:"comment,omitempty"`
	PublicKey     string   `json:"public_key,omitempty"`
	NotBefore     string   `json:"not_before,omitempty"`
	NotAfter      string   `json:"not_after,omitempty"`
	SourceAddress []string `json:"source_address,omitempty"`
}

// MFA multifactor authentication definition
type MFA struct {
	Status string `json:"status,omitempty"`
	Seed   Seed   `json:"seed,omitempty"`
}

// Seed seed definition
type Seed struct {
	SeedString string `json:"seed_string,omitempty"`
	SeedQRCode string `json:"seed_qr_code,omitempty"`
}

// Role contains PrivX role information.
type Role struct {
	ID                   string           `json:"id"`
	Name                 string           `json:"name"`
	GrantType            string           `json:"grant_type"`
	Comment              string           `json:"comment"`
	AccessGroupID        string           `json:"access_group_id"`
	GrantStart           string           `json:"grant_start"`
	GrantEnd             string           `json:"grant_end"`
	GrantValidityPeriods []ValidityPeriod `json:"grant_validity_periods,omitempty"`
	Permissions          []string         `json:"permissions"`
	PublicKey            []string         `json:"principal_public_key_strings"`
	MemberCount          int              `json:"member_count"`
	FloatingLength       int              `json:"floating_length"`
	Explicit             bool             `json:"explicit" tabulate:"@userCtx"`
	Implicit             bool             `json:"implicit" tabulate:"@userCtx"`
	System               bool             `json:"system"`
	PermitAgent          bool             `json:"permit_agent"`
	Context              *Context         `json:"context"`
	SourceRule           SourceRule       `json:"source_rules"`
}

type ValidityPeriod struct {
	GrantStart string `json:"grant_start,omitempty"`
	GrantEnd   string `json:"grant_end,omitempty"`
}

// RoleRef is a reference to role object
type RoleRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Context defines the context information for a role.
type Context struct {
	Enabled   bool   `json:"enabled"`
	BlockRole bool   `json:"block_role"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Timezone  string `json:"timezone"`
}

// SourceRule defines a mapping of role to object objects in directory
type SourceRule struct {
	Type    string       `json:"type"`
	Match   string       `json:"match"`
	Source  string       `json:"source,omitempty"`
	Pattern string       `json:"search_string,omitempty"`
	Rules   []SourceRule `json:"rules"`
}

// SourceRuleNone creates an empty mapping source for the role
func SourceRuleNone() SourceRule {
	return SourceRule{
		Type:  "GROUP",
		Match: "ANY",
	}
}

// UserSearchObject user search parameters
type UserSearchObject struct {
	Keywords string   `json:"keywords,omitempty"`
	Source   string   `json:"source,omitempty"`
	UserIDs  []string `json:"user_id,omitempty"`
}
type IdentityProvider struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	TokenType string `json:"token_type"`

	JWTIssuer                     string `json:"jwt_issuer"`
	JWTAudience                   string `json:"jwt_audience"`
	JWTSubjectType                string `json:"jwt_subject_type"`
	JWTSubjectDNUsernameAttribute string `json:"jwt_subject_dn_username_attribute,omitempty"`

	CustomAttributes []CustomAttributeValidation `json:"custom_attributes,omitempty"`

	PublicKey       []PublicKey `json:"public_key,omitempty"`
	PublicKeyMethod string      `json:"public_key_method"`

	// Used for validating certs fetched from x5u urls
	X5uTrustAnchor string `json:"x5u_trust_anchor,omitempty"`
	// Optional TLS trust anchor cert used when doing x5u https requests
	X5uTLSTrustAnchor string `json:"x5u_tls_trust_anchor,omitempty"`

	X5uPrefix      string `json:"x5u_prefix,omitempty"`
	UsersDirectory string `json:"users_directory"`

	Enabled   bool   `json:"enabled"`
	Author    string `json:"author"`
	Created   string `json:"created"`
	Updated   string `json:"updated,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}

type CustomAttributeValidation struct {
	FieldName     string `json:"field_name" validate:"required"`
	Type          string `json:"type"`
	ExpectedValue string `json:"expected_value"`
	Start         string `json:"start"`
	End           string `json:"end"`
}

type PublicKey struct {
	KeyID     string `json:"key_id"`
	Comment   string `json:"comment,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

type IdentityProviderResponse struct {
	Count int                `json:"count"`
	Items []IdentityProvider `json:"items"`
}

type IdentityProviderSearch struct {
	Keywords string `json:"keywords,omitempty"`
}

type IdentityProviderResolveUserRequest struct {
	Principal string `json:"principal"`
}
type IdentityProviderCreateResponse struct {
	ID string `json:"id"`
}
