package secretsmanager

import "time"

type Params struct {
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Sortdir string `json:"sortdir,omitempty"`
	Sortkey string `json:"sortkey,omitempty"`
}

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
	Created                *time.Time `json:"created"`
	CreatedBy              string     `json:"created_by"`
	Updated                *time.Time `json:"updated"`
	UpdatedBy              string     `json:"updated_by"`
}

type ScriptTemplate struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	OperatingSystem string     `json:"operating_system"`
	Script          string     `json:"script"`
	Created         *time.Time `json:"created"`
	CreatedBy       string     `json:"created_by"`
	Updated         *time.Time `json:"updated"`
	UpdatedBy       string     `json:"updated_by"`
}

type CompileScriptRequest struct {
	OperatingSystem string `json:"operating_system"`
	Script          string `json:"script"`
}

type TargetDomain struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	Enabled              bool                   `json:"enabled"`
	PeriodicScan         bool                   `json:"periodic_scan"`
	PeriodicScanInterval int                    `json:"periodic_scan_interval,omitempty"`
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

type PasswordPolicyHandle struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

type TargetDomainEndpoint struct {
	TargetDomainID          string            `json:"-"`
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

type TargetDomainHandle struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

type ScannedAccountsSearch struct {
	Keywords      string     `json:"keywords,omitempty"`
	CreatedAfter  *time.Time `json:"created_after,omitempty"`
	CreatedBefore *time.Time `json:"created_before,omitempty"`
	UpdatedAfter  *time.Time `json:"updated_after,omitempty"`
	UpdatedBefore *time.Time `json:"updated_before,omitempty"`
	State         string     `json:"state,omitempty"`
	Ignored       *bool      `json:"ignored,omitempty"`
}

type ScannedAccountChangeSet struct {
	Ignored *bool   `json:"ignored,omitempty"`
	Comment *string `json:"comment,omitempty"`
}

type ScannedAccountEditBatch struct {
	IDs       []string                `json:"ids"`
	ChangeSet ScannedAccountChangeSet `json:"changes"`
}

type ManagedAccount struct {
	ID               string                `json:"id"`
	Username         string                `json:"username"`
	Email            string                `json:"email,omitempty"`
	FullName         string                `json:"full_name,omitempty"`
	SourceID         string                `json:"source_id,omitempty"`
	SecurityID       string                `json:"security_id,omitempty"`
	AdditionalData   map[string]string     `json:"additional_data,omitempty"`
	TargetDomain     TargetDomainHandle    `json:"target_domain"`
	PasswordPolicy   *PasswordPolicyHandle `json:"password_policy,omitempty"`
	Enabled          bool                  `json:"enabled"`
	RotationEnabled  bool                  `json:"rotation_enabled"`
	ExplicitCheckout bool                  `json:"explicit_checkout"`
	State            string                `json:"state"`
	Comment          string                `json:"comment,omitempty"`
	SecretName       string                `json:"secret_name,omitempty"`
	Locked           bool                  `json:"locked"`
	LockedTimestamp  *time.Time            `json:"locked_timestamp,omitempty"`
	RotationHistory  []SecretRotationEvent `json:"rotation_history,omitempty"`
	SecretCheckouts  []SecretCheckout      `json:"checkouts,omitempty"`
	Created          time.Time             `json:"created,omitempty"`
	Author           string                `json:"author,omitempty"`
	Updated          *time.Time            `json:"updated,omitempty"`
	UpdatedBy        string                `json:"updated_by,omitempty"`
}

type SecretRotationEvent struct {
	Version int       `json:"version"`
	Rotated time.Time `json:"rotated"`
	Trigger string    `json:"trigger"`
	Status  string    `json:"status"`
}

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

type SecretVersion struct {
	Version int       `json:"version"`
	Secret  string    `json:"secret"`
	Created time.Time `json:"created"`
}

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

type ManagedAccountPasswordRequest struct {
	Password string `json:"password"`
}

type ManagedAccountCreateBatch struct {
	IDs  []string                 `json:"ids"`
	Data ManagedAccountCreateData `json:"data"`
}

type ManagedAccountCreateData struct {
	Enabled          bool                 `json:"enabled"`
	RotationEnabled  bool                 `json:"rotation_enabled"`
	Rotate           bool                 `json:"rotate"`
	ExplicitCheckout bool                 `json:"explicit_checkout"`
	PasswordPolicy   PasswordPolicyHandle `json:"password_policy,omitempty"`
	Comment          string               `json:"comment,omitempty"`
}

type ManagedAccountEditBatch struct {
	IDs       []string                `json:"ids"`
	ChangeSet ManagedAccountChangeSet `json:"changes"`
}

type ManagedAccountChangeSet struct {
	Enabled          *bool                 `json:"enabled"`
	RotationEnabled  *bool                 `json:"rotation_enabled"`
	ExplicitCheckout *bool                 `json:"explicit_checkout"`
	PasswordPolicy   *PasswordPolicyHandle `json:"password_policy,omitempty"`
	Comment          *string               `json:"comment,omitempty"`
}

type ManagedAccountBatch struct {
	IDs []string `json:"ids"`
}

type PwPolicyResult struct {
	Count int              `json:"count"`
	Items []PasswordPolicy `json:"items"`
}

type ScriptTemplateResult struct {
	Count int              `json:"count"`
	Items []ScriptTemplate `json:"items"`
}

type TdResult struct {
	Count int            `json:"count"`
	Items []TargetDomain `json:"items"`
}

type ScannedAccountResult struct {
	Count int              `json:"count"`
	Items []ScannedAccount `json:"items"`
}

type ManagedAccountResult struct {
	Count int              `json:"count"`
	Items []ManagedAccount `json:"items"`
}
