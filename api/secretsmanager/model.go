package secretsmanager

import "time"

// HostSecretMetadata host secret metadata definition.
type HostSecretMetadata struct {
	Metadata HostSecret      `json:"metadata"`
	Accounts []AccountSecret `json:"accounts"`
}

// AccountSecret account secret definition.
type AccountSecret struct {
	Account           string     `json:"account"`
	HostID            string     `json:"host_id"`
	RotationInitiated *time.Time `json:"rotation_initiate,omitempty"`
	LastRotated       *time.Time `json:"last_rotated,omitempty"`
	LastError         *time.Time `json:"last_error,omitempty"`
	LastErrorDetails  string     `json:"last_error_details"`
	InitialPassword   string     `json:"initial_password,omitempty"`
	Created           *time.Time `json:"created,omitempty"`
	CreatedBy         string     `json:"created_by"`
}

// HostSecret host secret definition.
type HostSecret struct {
	HostID                           string     `json:"host_id"`
	AccessGroupID                    string     `json:"access_group_id"`
	Address                          string     `json:"address"`
	Port                             int        `json:"port"`
	OperatingSystem                  string     `json:"operating_system"`
	Protocol                         string     `json:"protocol"`
	CertificateValidationOptions     string     `json:"certificate_validation_options"`
	WinRMHostCertificateTrustAnchors string     `json:"winrm_host_certificate_trust_anchors"`
	UseMainAccount                   bool       `json:"use_main_account"`
	MainAccount                      string     `json:"main_account"`
	PolicyID                         string     `json:"policy_id"`
	ScriptTemplateID                 string     `json:"script_template_id"`
	Created                          *time.Time `json:"created,omitempty"`
	CreatedBy                        string     `json:"created_by"`
	Updated                          *time.Time `json:"updated,omitempty"`
	UpdatedBy                        string     `json:"updated_by"`
}

// PasswordPolicy password policy definition.
type PasswordPolicy struct {
	ID                     string     `json:"id"`
	Name                   string     `json:"name"`
	RotationInterval       string     `json:"rotation_interval"`
	PasswordMinLength      int        `json:"password_min_length"`
	PasswordMaxLength      int        `json:"password_max_length" `
	UseSpecialCharacters   bool       `json:"use_special_characters"`
	UseLowercase           bool       `json:"use_lower_case"`
	UseUppercase           bool       `json:"use_upper_case"`
	UseNumbers             bool       `json:"use_numbers"`
	MaxVersions            int        `json:"max_versions"`
	NumberOfRetries        int        `json:"number_of_retries"`
	RetryInterval          string     `json:"retry_interval"`
	MaxConcurrentCheckouts int        `json:"max_concurrent_checkouts"`
	MaxCheckoutDuration    string     `json:"max_checkout_duration"`
	RotateOnRelease        bool       `json:"rotate_on_release"`
	VerifyAfterRotation    bool       `json:"verify_after_rotation"`
	Created                *time.Time `json:"created,omitempty"`
	CreatedBy              string     `json:"created_by"`
	Updated                *time.Time `json:"updated,omitempty"`
	UpdatedBy              string     `json:"updated_by"`
}

// ScriptTemplate script template definition.
type ScriptTemplate struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	OperatingSystem string     `json:"operating_system"`
	Script          string     `json:"script"`
	Created         *time.Time `json:"created,omitempty"`
	CreatedBy       string     `json:"created_by"`
	Updated         *time.Time `json:"updated,omitempty"`
	UpdatedBy       string     `json:"updated_by"`
}

// CompileScript compile script request definition.
type CompileScript struct {
	OperatingSystem string `json:"operating_system"`
	Script          string `json:"script"`
}

// CompileScriptResponse compile script response definition.
type CompileScriptResponse struct {
	Script string `json:"script"`
}

// TargetDomain target domain definition.
type TargetDomain struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	DomainName           string                 `json:"domain_name"`
	Enabled              bool                   `json:"enabled"`
	PeriodicScan         bool                   `json:"periodic_scan"`
	PeriodicScanInterval int                    `json:"periodic_scan_interval"`
	ScanStatus           string                 `json:"scan_status,omitempty"`
	ScanMessage          string                 `json:"scan_message,omitempty"`
	LastScanned          *time.Time             `json:"last_scanned,omitempty"`
	AutoOnboarding       bool                   `json:"auto_onboarding"`
	AutoOnboardingPolicy *PasswordPolicyHandle  `json:"auto_onboarding_policy,omitempty"`
	EndPoints            []TargetDomainEndpoint `json:"endpoints"`
	Comment              string                 `json:"comment"`
	Created              time.Time              `json:"created,omitempty"`
	Author               string                 `json:"author,omitempty"`
	Updated              time.Time              `json:"updated,omitempty"`
	UpdatedBy            string                 `json:"updated_by,omitempty"`
}

// PasswordPolicyHandle password policy handle definition.
type PasswordPolicyHandle struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

// TargetDomainEndpoint target domain endpoint definition.
type TargetDomainEndpoint struct {
	Type                    string            `json:"type"`
	ScanPriority            int               `json:"scan_priority"`
	RotationPriority        int               `json:"rotation_priority"`
	AttributeMapping        map[string]string `json:"attribute_mapping,omitempty"`
	LdapProtocol            string            `json:"ldap_protocol,omitempty"`
	LdapAddress             string            `json:"ldap_address,omitempty"`
	LdapPort                int               `json:"ldap_port,omitempty"`
	LdapBaseDN              string            `json:"ldap_base_dn,omitempty"`
	LdapBindDN              string            `json:"ldap_bind_dn,omitempty"`
	LdapBindPassword        string            `json:"ldap_bind_password,omitempty"`
	LdapUserFilter          string            `json:"ldap_user_filter,omitempty"`
	LdapRootCertificates    string            `json:"ldap_root_certificates,omitempty"`
	LdapSkipStrictCertCheck bool              `json:"ldap_skip_strict_cert_check,omitempty"`
	EntraBaseUrl            string            `json:"entra_base_url,omitempty"`
	EntraTenantID           string            `json:"entra_tenant_id,omitempty"`
	EntraClientID           string            `json:"entra_client_id,omitempty"`
	EntraClientSecret       string            `json:"entra_client_secret,omitempty"`
	EntraBatchSize          int               `json:"entra_batch_size,omitempty"`
	EntraPageSize           int               `json:"entra_page_size,omitempty"`
	EntraGroupFilter        []string          `json:"entra_group_filter,omitempty"`
}

type TargetDomainsSearch struct {
	Keywords       string `json:"keywords,omitempty"`
	Enabled        *bool  `json:"enabled,omitempty"`
	PeriodicScan   *bool  `json:"periodic_scan,omitempty"`
	AutoOnboarding *bool  `json:"auto_onboarding,omitempty"`
}

// ScannedAccount scanned account definition.
type ScannedAccount struct {
	ID             string             `json:"id"`
	Username       string             `json:"username"`
	Email          string             `json:"email"`
	FullName       string             `json:"full_name"`
	SourceID       string             `json:"source_id"`
	SecurityID     string             `json:"security_id"`
	AdditionalData map[string]string  `json:"additional_data"`
	TargetDomain   TargetDomainHandle `json:"target_domain"`
	State          string             `json:"state"`
	Ignored        bool               `json:"ignored"`
	Comment        string             `json:"comment"`
	Created        time.Time          `json:"created,omitempty"`
	Updated        time.Time          `json:"updated,omitempty"`
	UpdatedBy      string             `json:"updated_by,omitempty"`
}

// TargetDomainHandle target domain handle definition.
type TargetDomainHandle struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

// ScannedAccountsSearch scanned account search request definition.
type ScannedAccountsSearch struct {
	Keywords      string     `json:"keywords,omitempty"`
	CreatedAfter  *time.Time `json:"created_after,omitempty"`
	CreatedBefore *time.Time `json:"created_before,omitempty"`
	UpdatedAfter  *time.Time `json:"updated_after,omitempty"`
	UpdatedBefore *time.Time `json:"updated_before,omitempty"`
	State         string     `json:"state,omitempty"`
	Ignored       *bool      `json:"ignored,omitempty"`
}

// ScannedAccountChangeSet scanned account change set definition.
type ScannedAccountChangeSet struct {
	Ignored *bool   `json:"ignored,omitempty"`
	Comment *string `json:"comment,omitempty"`
}

// ScannedAccountEditBatch scanned account edit batch request definition.
type ScannedAccountEditBatch struct {
	IDs       []string                `json:"ids"`
	ChangeSet ScannedAccountChangeSet `json:"changes"`
}

// ManagedAccount managed account definition.
type ManagedAccount struct {
	ID                 string                `json:"id"`
	Username           string                `json:"username"`
	Email              string                `json:"email,omitempty"`
	FullName           string                `json:"full_name,omitempty"`
	SamAccountName     string                `json:"sam_account_name,omitempty"`
	SourceID           string                `json:"source_id,omitempty"`
	SecurityID         string                `json:"security_id,omitempty"`
	AdditionalData     map[string]string     `json:"additional_data,omitempty"`
	TargetDomain       TargetDomainHandle    `json:"target_domain"`
	PasswordPolicy     *PasswordPolicyHandle `json:"password_policy,omitempty"`
	Enabled            bool                  `json:"enabled"`
	RotationEnabled    bool                  `json:"rotation_enabled"`
	ExplicitCheckout   bool                  `json:"explicit_checkout"`
	State              string                `json:"state"`
	Comment            string                `json:"comment,omitempty"`
	SecretName         string                `json:"secret_name,omitempty"`
	Locked             bool                  `json:"locked"`
	LockedTimestamp    *time.Time            `json:"locked_timestamp,omitempty"`
	RotationHistory    []SecretRotationEvent `json:"rotation_history,omitempty"`
	SecretCheckouts    []SecretCheckout      `json:"checkouts,omitempty"`
	Created            time.Time             `json:"created,omitempty"`
	Author             string                `json:"author,omitempty"`
	Updated            *time.Time            `json:"updated,omitempty"`
	UpdatedBy          string                `json:"updated_by,omitempty"`
	DisableRDPCertAuth bool                  `json:"disable_rdp_cert_auth"`
}

// SecretRotationEvent secret rotation event definition.
type SecretRotationEvent struct {
	Version int       `json:"version"`
	Rotated time.Time `json:"rotated"`
	Trigger string    `json:"trigger"`
	Status  string    `json:"status"`
}

// SecretCheckout secret checkout definition.
type SecretCheckout struct {
	ID               string          `json:"id"`
	Type             string          `json:"type"`
	UserID           string          `json:"user_id"`
	Expires          time.Time       `json:"expires"`
	Created          time.Time       `json:"created"`
	ExplicitCheckout bool            `json:"explicit_checkout"`
	Secrets          []SecretVersion `json:"secrets,omitempty"`
	Username         string          `json:"username"`
	Email            string          `json:"email,omitempty"`
	FullName         string          `json:"full_name,omitempty"`
	TargetDomainID   string          `json:"target_domain_id,omitempty"`
	ManagedAccountID string          `json:"managed_account_id,omitempty"`
	HostID           string          `json:"host_id,omitempty"`
	SecretName       string          `json:"secret_name,omitempty"`
	Meta             string          `json:"meta,omitempty"`
}

// SecretVersion secret version definition.
type SecretVersion struct {
	Version int       `json:"version"`
	Secret  string    `json:"secret"`
	Created time.Time `json:"created"`
}

// ManagedAccountsSearch managed account search request definition.
type ManagedAccountsSearch struct {
	Keywords         string     `json:"keywords,omitempty"`
	Enabled          *bool      `json:"enabled,omitempty"`
	CreatedAfter     *time.Time `json:"created_after,omitempty"`
	CreatedBefore    *time.Time `json:"created_before,omitempty"`
	UpdatedAfter     *time.Time `json:"updated_after,omitempty"`
	UpdatedBefore    *time.Time `json:"updated_before,omitempty"`
	State            string     `json:"state,omitempty"`
	RotationEnabled  *bool      `json:"rotation_enabled,omitempty"`
	ExplicitCheckout *bool      `json:"explicit_checkout,omitempty"`
}

// ManagedAccountPasswordSet manage account password set request definition.
type ManagedAccountPasswordSet struct {
	Password string `json:"password"`
}

// IDList id list response definition.
type IDList struct {
	IDs []string `json:"ids"`
}

// ManagedAccountCreateBatch managed account create batch definition.
type ManagedAccountCreateBatch struct {
	IDs  []string                 `json:"ids"`
	Data ManagedAccountCreateData `json:"data"`
}

// ManagedAccountCreateData managed account batch create data definition.
type ManagedAccountCreateData struct {
	Enabled            bool                 `json:"enabled"`
	RotationEnabled    bool                 `json:"rotation_enabled"`
	Rotate             bool                 `json:"rotate"`
	ExplicitCheckout   bool                 `json:"explicit_checkout"`
	DisableRDPCertAuth bool                 `json:"disable_rdp_cert_auth"`
	PasswordPolicy     PasswordPolicyHandle `json:"password_policy,omitempty"`
	Comment            string               `json:"comment,omitempty"`
}

type ManagedAccountEditBatch struct {
	IDs       []string                `json:"ids"`
	ChangeSet ManagedAccountChangeSet `json:"changes"`
}

// ManagedAccountChangeSet manage account change set request definition.
type ManagedAccountChangeSet struct {
	Enabled            *bool                 `json:"enabled,omitempty"`
	RotationEnabled    *bool                 `json:"rotation_enabled,omitempty"`
	ExplicitCheckout   *bool                 `json:"explicit_checkout,omitempty"`
	DisableRDPCertAuth *bool                 `json:"disable_rdp_cert_auth,omitempty"`
	PasswordPolicy     *PasswordPolicyHandle `json:"password_policy,omitempty"`
	Comment            *string               `json:"comment,omitempty"`
}

// ManagedAccountDeleteBatch manage account batch delete request definition.
type ManagedAccountDeleteBatch struct {
	IDs []string `json:"ids"`
}

// ManagedAccountRotateBatch manage account batch rotate request definition.

type ManagedAccountRotateBatch struct {
	IDs []string `json:"ids"`
}

type TargetDomainsResolveResponse struct {
	Count int                    `json:"count"`
	Items []TargetDomainsResolve `json:"items"`
}

type TargetDomainsResolve struct {
	ID               string `json:"id"`
	TargetDomainName string `json:"target_domain_name"`
}
