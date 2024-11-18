//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

import (
	"encoding/json"

	"github.com/SSHcom/privx-sdk-go/api/auth"
)

const (
	// Enumerated values for MFA actions.
	MFAActionEnable  MFAAction = "enable"
	MFAActionDisable MFAAction = "disable"
	MFAActionReset   MFAAction = "reset"
)

// MFAAction definition for possible actions related to MFA.
type MFAAction string

// AWSRoleParams aws role query parameter definition.
type AWSRoleParams struct {
	Refresh bool `url:"refresh,omitempty"`
}

// AuthorizedKeyResolve authorized key resolve request definition.
type AuthorizedKeyResolve struct {
	Username  string `json:"username,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

// RolePrincipalKeyImport role principal key import request definition.
type RolePrincipalKeyImport struct {
	PrivateKey string `json:"private_key,omitempty"`
}

// RolePrincipalKey role principal key definition.
type RolePrincipalKey struct {
	ID        string `json:"id,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

// LogConfCollector logconf collectors definition.
type LogConfCollector struct {
	ID                  string `json:"id,omitempty"`
	Name                string `json:"name,omitempty"`
	Type                string `json:"type,omitempty"`
	Enabled             bool   `json:"enabled"`
	Updated             string `json:"updated,omitempty"`
	StatusCode          string `json:"status_code,omitempty"`
	StatusText          string `json:"status_text,omitempty"`
	AWSLogRegion        string `json:"aws_log_region,omitempty"`
	IAMAccessKeyID      string `json:"iam_access_key_id,omitempty"`
	IAMSecretAccessKey  string `json:"iam_secret_access_key,omitempty"`
	IAMSessionToken     string `json:"iam_session_token,omitempty"`
	Namespace           string `json:"azure_event_hubs_namespace,omitempty"`
	EventHubName        string `json:"azure_event_hub_name,omitempty"`
	TenantID            string `json:"azure_tenant_id,omitempty"`
	ClientID            string `json:"azure_client_id,omitempty"`
	ClientSecret        string `json:"azure_client_secret,omitempty"`
	SASConnectionString string `json:"azure_sas_connection_string,omitempty"`
}

// AWSRole aws role definition.
type AWSRole struct {
	ID          string            `json:"id,omitempty"`
	DisplayName string            `json:"name,omitempty"`
	ARN         string            `json:"arn,omitempty"`
	Description string            `json:"description,omitempty"`
	Source      string            `json:"source,omitempty"`
	Status      string            `json:"status,omitempty"`
	Roles       []LinkedPrivXRole `json:"roles,omitempty"`
	Updated     string            `json:"updated,omitempty"`
}

type LinkedPrivXRole struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"name,omitempty"`
}

// Source source definitions.
type Source struct {
	ID                          string                      `json:"id,omitempty"`
	Name                        string                      `json:"name,omitempty"`
	Created                     string                      `json:"created,omitempty"`
	Updated                     string                      `json:"updated,omitempty"`
	Author                      string                      `json:"author,omitempty"`
	UpdatedBy                   string                      `json:"updatedby,omitempty"`
	TTL                         int                         `json:"ttl,omitempty"`
	RegionFilter                []string                    `json:"region_filter,omitempty"`
	Comment                     string                      `json:"comment,omitempty"`
	Tags                        []string                    `json:"tags,omitempty"`
	StatusCode                  string                      `json:"status_code"`
	StatusText                  string                      `json:"status_text"`
	Connection                  SourceConnection            `json:"connection,omitempty"`
	UsernamePattern             []string                    `json:"username_pattern,omitempty"`
	ExternalUserMapping         []UserMapping               `json:"external_user_mapping,omitempty"`
	Enabled                     bool                        `json:"enabled"`
	SessionPasswordEnabled      bool                        `json:"session_password_enabled,omitempty"`
	SessionPasswordPolicy       *auth.SessionPasswordPolicy `json:"session_password_policy,omitempty"`
	ChildSessionAutoLogoutDelay int                         `json:"child_session_auto_logout_delay,omitempty"`
}

// UserMapping user mapping definition
type UserMapping struct {
	SourceID          string `json:"source_id,omitempty"`
	SourceSearchField string `json:"source_search_field,omitempty"`
}

// SourceConnection source connection definition.
type SourceConnection struct {
	Type                        string            `json:"type,omitempty"`
	AttributeMapping            map[string]string `json:"attribute_mapping"`
	MFAType                     string            `json:"mfa_type"`
	MFAAddress                  string            `json:"mfa_address"`
	MFAPort                     int               `json:"mfa_port"`
	ServiceAddressAutoUpdate    bool              `json:"service_address_auto_update"`
	UseInstanceTags             bool              `json:"use_instance_tags"`
	HostFilterTag               string            `json:"host_filter_tag,omitempty"`
	AWSRoleFilterName           string            `json:"aws_role_filter_name,omitempty"`
	EnableUserAuthentication    bool              `json:"enable_user_authentication,omitempty"`
	EnableMachineAuthentication bool              `json:"enable_machine_authentication,omitempty"`

	// AWS
	IAMAccessKeyID     string `json:"iam_access_key_id"`
	IAMSecretAccessKey string `json:"iam_secret_access_key"`
	IAMSessionToken    string `json:"iam_session_token"`
	IAMRoleARN         string `json:"iam_role_arn"`

	IAMFetchRoles          bool   `json:"iam_fetch_roles"`
	IAMFetchRolePathPrefix string `json:"iam_fetch_role_path_prefix"`
	IAMExternalID          string `json:"iam_external_id"`

	// Google Cloud
	GoogleCloudProjectIDs []string `json:"google_cloud_project_ids"`
	GoogleCloudConfigJSON string   `json:"google_cloud_config_json"`

	// Google GSuite
	GoogleGSuiteDomainAdminEmail string `json:"google_gsuite_domain_admin_email"`
	GoogleGSuiteDomain           string `json:"google_gsuite_domain"`

	// OpenStack
	OpenStackIdentityEndpoint string   `json:"openstack_endpoint"`
	OpenStackUsername         string   `json:"openstack_username"`
	OpenStackUserID           string   `json:"openstack_user_id"`
	OpenStackPassword         string   `json:"openstack_password"`
	OpenStackDomainName       string   `json:"openstack_domainname"`
	OpenStackDomainID         string   `json:"openstack_domain_id"`
	OpenStackTokenID          string   `json:"openstack_token_id"`
	OpenStackTenantIDs        []string `json:"openstack_tenant_ids"`
	OpenStackTenantNames      []string `json:"openstack_tenant_names"`
	OpenStackVersion          string   `json:"openstack_version"`
	OpenStackRegion           string   `json:"openstack_region"`

	// Azure
	AzureEndpoint       string `json:"azure_base_url"`
	AzureSubscriptionID string `json:"azure_subscription_id"`
	AzureTenantID       string `json:"azure_tenant_id"`
	AzureClientID       string `json:"azure_client_id"`
	AzureClientSecret   string `json:"azure_client_secret"`

	// Microsoft Graph API
	MsGraphEndpoint     string `json:"msgraph_base_url"`
	MsGraphTenantID     string `json:"msgraph_tenant_id"`
	MsGraphClientID     string `json:"msgraph_client_id"`
	MsGraphClientSecret string `json:"msgraph_client_secret"`
	MSGraphBatchSize    int    `json:"msgraph_batch_size"`
	MSGraphPageSize     int    `json:"msgraph_page_size"`

	// AD/LDAP
	Address                                 string `json:"address"`
	Port                                    int    `json:"port"`
	LdapBaseDN                              string `json:"ldap_base_dn"`
	LdapUserDNPattern                       string `json:"ldap_user_dn_pattern"`
	LdapUserFilter                          string `json:"ldap_user_filter"`
	LdapBindDN                              string `json:"ldap_bind_dn"`
	LdapBindPassword                        string `json:"ldap_bind_password"`
	LdapProtocol                            string `json:"ldap_protocol"`
	Certificates                            string `json:"root_certificates"`
	PasswordChangeEnabled                   bool   `json:"password_change_enabled"`
	ClientCertificateAuthenticationEnabled  bool   `json:"client_certificate_authentication_enabled"`
	ClientCAPEM                             string `json:"client_ca_pem"`
	ClientCertificateAuthenticationRequired bool   `json:"client_certificate_authentication_required"`

	// AD/LDAP/UKM
	SkipStrictCertCheck bool `json:"skip_strict_cert_check"`

	// OpenID Connect
	OIDCEnabled             bool     `json:"oidc_enabled"`
	OIDCIssuer              string   `json:"oidc_issuer"`
	OIDCClientID            string   `json:"oidc_client_id"`
	OIDCClientSecret        string   `json:"oidc_client_secret"`
	OIDCButtonTitle         string   `json:"oidc_button_title"`
	OIDCTagsAttributeName   string   `json:"oidc_tags_attribute_name"`
	OIDCAllowLoginTag       string   `json:"oidc_allow_login_tag"`
	OIDCAdditionalScopes    []string `json:"oidc_additional_scopes"`
	OIDCUseUserInfoEndpoint bool     `json:"oidc_use_userinfo_endpoint"`

	// User directory group filter
	GroupFilter []string `json:"group_filter"`

	// SCIM
	SCIMAuthenticationType string `json:"scim_authentication_type"`
	SCIMCreateRoles        bool   `json:"scim_create_roles"`
	SCIMUpdateSSHHostKeys  bool   `json:"scim_update_ssh_host_keys"`
	SCIMBasicUsername      string `json:"scim_username"`
	SCIMBasicPassword      string `json:"scim_password"`
	SCIMUserFilter         string `json:"scim_user_filter"`

	// VMWare ESXi/vCenter
	VMWareEndpoint   string `json:"vmware_url"`
	VMWareUsername   string `json:"vmware_username"`
	VMWarePassword   string `json:"vmware_password"`
	VMWareDataCenter string `json:"vmware_datacenter"`

	// UKM
	UKMEndpoint    string `json:"ukm_endpoint"`
	UKMToken       string `json:"ukm_token"`
	UKMTrustAnchor string `json:"um_ca_pem"`
}

// User user definition.
type User struct {
	ID                     string          `json:"id,omitempty"`
	SourceUserID           string          `json:"source_user_id,omitempty"`
	Created                string          `json:"created,omitempty"`
	Updated                string          `json:"updated,omitempty"`
	Author                 string          `json:"author,omitempty"`
	UpdatedBy              string          `json:"updatedby,omitempty"`
	Principal              string          `json:"principal,omitempty"`
	Source                 string          `json:"source,omitempty"`
	SourceType             string          `json:"source_type,omitempty"`
	Comment                string          `json:"comment,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	Roles                  []Role          `json:"roles"`
	Attributes             []UserAttribute `json:"attributes"`
	Permissions            []string        `json:"permissions"`
	FirstName              string          `json:"first_name,omitempty"`
	LastName               string          `json:"last_name,omitempty"`
	FullName               string          `json:"full_name,omitempty"`
	JobTitle               string          `json:"job_title,omitempty"`
	Company                string          `json:"company,omitempty"`
	Department             string          `json:"department,omitempty"`
	Email                  string          `json:"email,omitempty"`
	Telephone              string          `json:"telephone,omitempty"`
	DistinguishedName      string          `json:"distinguished_name,omitempty"`
	Locale                 string          `json:"locale,omitempty"`
	SamAccountName         string          `json:"samaccountname,omitempty"`
	WindowsAccount         string          `json:"windows_account,omitempty"`
	UnixAccount            string          `json:"unix_account,omitempty"`
	Password               string          `json:"password,omitempty"`
	MFA                    MFAStatus       `json:"mfa,omitempty"`
	Settings               json.RawMessage `json:"settings,omitempty"`
	ExternalID             string          `json:"external_id,omitempty"`
	AuthorizedKeys         []AuthorizedKey `json:"authorized_keys,omitempty"`
	WebAuthnCredentials    []Credential    `json:"webauthn_credentials,omitempty"`
	SessionPasswordEnabled bool            `json:"session_password_enabled,omitempty"`
	StaleAccessToken       bool            `json:"stale_access_token,omitempty"`
	CurrentSessionID       string          `json:"current_session_id,omitempty"`
	RefreshTimestamp       string
}

// Credential webauthn credential definition.
type Credential struct {
	ID           string `json:"id"`
	CredentialID string `json:"credential_id"`
	Name         string `json:"name,omitempty" diff:"name"`
	Comment      string `json:"comment,omitempty" diff:"comment"`
	LastUsed     string `json:"last_used,omitempty"`
	Created      string `json:"created,omitempty"`
	Author       string `json:"author,omitempty"`
	Updated      string `json:"updated,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
}

// MFAStatus mfa status definition.
type MFAStatus struct {
	Status          string  `json:"user_mfa_status,omitempty"`
	TotpStatus      string  `json:"status,omitempty"`
	TotpMFASeed     MFASeed `json:"seed,omitempty"`
	MobileMfaStatus string  `json:"mobile_mfa_status"`
}

type MFASeed struct {
	Seed_string  string `json:"seed_string,omitempty"`
	Seed_qr_code string `json:"seed_qr_code,omitempty"`
}

// UserAttribute user attribute definition.
type UserAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// AuthorizedKey authorized key definition.
type AuthorizedKey struct {
	ID            string   `json:"id"`
	UserID        string   `json:"user_id,omitempty"`
	Username      string   `json:"username,omitempty"`
	Source        string   `json:"source,omitempty"`
	PublicKey     string   `json:"public_key,omitempty"`
	NotBefore     string   `json:"not_before,omitempty"`
	NotAfter      string   `json:"not_after,omitempty"`
	ExpiresIn     int64    `json:"expires_in,omitempty"`
	SourceAddress []string `json:"source_address"`
	Fingerprints  []string `json:"fingerprints,omitempty"`
	Name          string   `json:"name,omitempty"`
	Comment       string   `json:"comment,omitempty"`
	Created       string   `json:"created,omitempty"`
	Updated       string   `json:"updated,omitempty"`
	UpdatedBy     string   `json:"updated_by,omitempty"`
	Author        string   `json:"author,omitempty"`
}

// UserSettings user settings update request definition.
type UserSettings struct {
	Locale            UserLocale              `json:"locale,omitempty"`
	RDPClient         UserRDPClient           `json:"rdpClient,omitempty"`
	SSHClient         UserSSHClient           `json:"sshClient,omitempty"`
	ConnectionHistory []UserConnectionHistory `json:"connectionHistory,omitempty"`
	Bookmarks         UserBookmarks           `json:"bookmarks,omitempty"`
}

// Bookmark bookmark user settings definition.
type Bookmark struct {
	Id    string `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

// UserBookmarks user bookmarks settings definition.
type UserBookmarks struct {
	HostsBookmarks                   []Bookmark `json:"available-hosts,omitempty"`
	NetworkTargetsBookmarks          []Bookmark `json:"available-network-targets,omitempty"`
	SecretsBookmarks                 []Bookmark `json:"secrets,omitempty"`
	UserRequestsBookmarks            []Bookmark `json:"requests,omitempty"`
	ApprovalsBookmarks               []Bookmark `json:"approvals,omitempty"`
	ConnectionsBookmarks             []Bookmark `json:"connections,omitempty"`
	AuditEventBookmarks              []Bookmark `json:"audit-events,omitempty"`
	CertificatesBookmarks            []Bookmark `json:"certificates,omitempty"`
	AdminRolesBookmarks              []Bookmark `json:"roles,omitempty"`
	AdminUsersBookmarks              []Bookmark `json:"users,omitempty"`
	AdminHostsBookmarks              []Bookmark `json:"hosts,omitempty"`
	AdminNetworkTargetsBookmarks     []Bookmark `json:"network-targets,omitempty"`
	AdminDirectoriesBookmarks        []Bookmark `json:"sources,omitempty"`
	AdminAccessGroupsBookmarks       []Bookmark `json:"access-groups,omitempty"`
	AdminWorkflowsBookmarks          []Bookmark `json:"workflows,omitempty"`
	SessionsBookmarks                []Bookmark `json:"sessions,omitempty"`
	IdentityProviderClientsBookmarks []Bookmark `json:"identity-provider-clients,omitempty"`
	ExternalTokenProvidersBookmarks  []Bookmark `json:"external-token-providers,omitempty"`
	AWSRolesBookmarks                []Bookmark `json:"aws-roles,omitempty"`
	APIClientsBookmarks              []Bookmark `json:"api-clients,omitempty"`
	CommandWhitelistsBookmarks       []Bookmark `json:"command-whitelists,omitempty"`
}

// UserConnectionHistory user connection history settings definition.
type UserConnectionHistory struct {
	Id          string `json:"id"`
	Time        string `json:"time"`
	Type        string `json:"type"`
	Target      string `json:"target"`
	TargetID    string `json:"targetId,omitempty"`
	Account     string `json:"account,omitempty"`
	Name        string `json:"name,omitempty"`
	Application string `json:"application,omitempty"`
}

// UserSSHClient user ssh client settings definition.
type UserSSHClient struct {
	FontSize          int    `json:"fontSize,omitempty"`
	Encoding          string `json:"encoding,omitempty"`
	Locale            string `json:"locale,omitempty"`
	Theme             string `json:"theme,omitempty"`
	CopyOnSelect      bool   `json:"copyOnSelect,omitempty"`
	PasteOnRightClick bool   `json:"pasteOnRightClick,omitempty"`
	SendCtrlV         bool   `json:"sendCtrlV,omitempty"`
	AltAsMeta         bool   `json:"altAsMeta,omitempty"`
	ClickableLinks    bool   `json:"clickableLinks,omitempty"`
	ScrollbackLength  int    `json:"scrollbackLength,omitempty"`
}

// UserLocale user locale settings definition.
type UserLocale struct {
	Locale string `json:"locale,omitempty"`
}

// UserRDPClient user rdp client settings definition.
type UserRDPClient struct {
	KeyboardLayout            string  `json:"keyboardLayout,omitempty"`
	Scale                     float32 `json:"scale,omitempty"`
	ScalingMode               string  `json:"scalingMode,omitempty"`
	ImageScalingAlgorithm     string  `json:"imageScalingAlgorithm,omitempty"`
	ClipboardSync             bool    `json:"clipboardSync,omitempty"`
	ClipboardSyncStartsPaused bool    `json:"clipboardSyncStartsPaused,omitempty"`
}

// RoleSearch role search request definition.
type RoleSearch struct {
	Name []string `json:"name"`
}

// Role PrivX role definition.
type Role struct {
	ID                   string           `json:"id"`
	Name                 string           `json:"name,omitempty"`
	Explicit             bool             `json:"explicit"`
	Implicit             bool             `json:"implicit"`
	System               bool             `json:"system"`
	GrantType            string           `json:"grant_type,omitempty"`
	GrantStart           string           `json:"grant_start,omitempty"`
	GrantEnd             string           `json:"grant_end,omitempty"`
	GrantValidityPeriods []ValidityPeriod `json:"grant_validity_periods,omitempty"`
	FloatingLength       int64            `json:"floating_length,omitempty"`
	Comment              string           `json:"comment,omitempty"`
	Permissions          []string         `json:"permissions"`
	AccessGroupID        string           `json:"access_group_id"`
	PrincipalPublicKeys  []string         `json:"principal_public_key_strings,omitempty"`
	PermitAgent          bool             `json:"permit_agent,omitempty"`
	Context              ContextualLimit  `json:"context"`
}

type ContextualLimit struct {
	Enabled   bool     `json:"enabled"`
	BlockRole bool     `json:"block_role,omitempty"`
	Validity  []string `json:"validity"`
	StartTime string   `json:"start_time"`
	EndTime   string   `json:"end_time"`
	TimeZone  string   `json:"timezone"`
	IPMasks   []string `json:"ip_masks"`
}

// ValidityPeriod validity period definition.
type ValidityPeriod struct {
	GrantStart string `json:"grant_start,omitempty"`
	GrantEnd   string `json:"grant_end,omitempty"`
}

// RoleHandle role handle definition.
type RoleHandle struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted,omitempty"`
}

// UserSearch user search request definition.
type UserSearch struct {
	Keywords string   `json:"keywords,omitempty"`
	Source   string   `json:"source,omitempty"`
	UserIDs  []string `json:"user_id,omitempty"`
}

// IdentityProviderSearch identity provider search request definition.
type IdentityProviderSearch struct {
	Keywords string `json:"keywords,omitempty"`
}

// IdentityProvider identity provider definition.
type IdentityProvider struct {
	ID                            string                      `json:"id"`
	Name                          string                      `json:"name"`
	TokenType                     string                      `json:"token_type"`
	JWTIssuer                     string                      `json:"jwt_issuer"`
	JWTAudience                   string                      `json:"jwt_audience"`
	JWTSubjectType                string                      `json:"jwt_subject_type"`
	JWTSubjectDNUsernameAttribute string                      `json:"jwt_subject_dn_username_attribute,omitempty"`
	CustomAttributes              []CustomAttributeValidation `json:"custom_attributes,omitempty"`
	PublicKeys                    []PublicKey                 `json:"public_keys,omitempty"`
	PublicKeyMethod               string                      `json:"public_key_method"`
	X5uTrustAnchor                string                      `json:"x5u_trust_anchor,omitempty"`
	X5uTLSTrustAnchor             string                      `json:"x5u_tls_trust_anchor,omitempty"`
	X5uPrefix                     string                      `json:"x5u_prefix,omitempty"`
	UsersDirectory                string                      `json:"users_directory"`
	Enabled                       bool                        `json:"enabled"`
	Author                        string                      `json:"author"`
	Created                       string                      `json:"created"`
	Updated                       string                      `json:"updated,omitempty"`
	UpdatedBy                     string                      `json:"updated_by,omitempty"`
}

// CustomAttributeValidation identity provider custom attribute definition.
type CustomAttributeValidation struct {
	FieldName     string `json:"field_name"`
	Type          string `json:"type"`
	ExpectedValue string `json:"expected_value"`
	Start         string `json:"start"`
	End           string `json:"end"`
}

// PublicKey identity provider public key definition.
type PublicKey struct {
	KeyID     string `json:"key_id"`
	Comment   string `json:"comment,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}
