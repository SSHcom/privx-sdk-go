//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

// Params struct for pagination queries.
type Params struct {
	Refresh string `json:"refresh,omitempty"`
}

// AWSRole aws role definition.
type AWSRole struct {
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
	ID                  string     `json:"id,omitempty"`
	Created             string     `json:"created,omitempty"`
	Updated             string     `json:"updated,omitempty"`
	UpdatedBy           string     `json:"updated_by,omitempty"`
	Author              string     `json:"author,omitempty"`
	Name                string     `json:"name,omitempty"`
	StatusCode          string     `json:"status_code,omitempty"`
	StatusText          string     `json:"status_text,omitempty"`
	Comment             string     `json:"comment,omitempty"`
	TTL                 int        `json:"ttl,omitempty"`
	Enabled             bool       `json:"enabled,omitempty"`
	Tags                []string   `json:"tags,omitempty"`
	UsernamePattern     []string   `json:"username_pattern,omitempty"`
	ExternalUserMapping []EUM      `json:"external_user_mapping,omitempty"`
	Connection          Connection `json:"connection,omitempty"`
}

// User contains PrivX user information.
type User struct {
	ID                string   `json:"id"`
	SourceUserID      string   `json:"source_user_id"`
	Tags              []string `json:"tags"`
	Principal         string   `json:"principal"`
	Source            string   `json:"source"`
	FullName          string   `json:"full_name"`
	Email             string   `json:"email"`
	DistinguishedName string   `json:"distinguished_name"`
	Roles             []Role   `json:"roles"`
}

// Role contains PrivX role information.
type Role struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Explicit    bool       `json:"explicit" tabulate:"@userCtx"`
	Implicit    bool       `json:"implicit" tabulate:"@userCtx"`
	System      bool       `json:"system"`
	GrantType   string     `json:"grant_type"`
	Comment     string     `json:"comment"`
	SourceRule  SourceRule `json:"source_rules"`
	Permissions []string   `json:"permissions"`
	Context     *Context   `json:"context"`
	MemberCount int        `json:"member_count"`
	PublicKey   []string   `json:"principal_public_key_strings,omitempty"`
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
