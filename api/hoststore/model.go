//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package hoststore

import (
	"time"

	"github.com/SSHcom/privx-sdk-go/api/rolestore"
	"github.com/SSHcom/privx-sdk-go/api/secretsmanager"
)

// HostSearch host search request definition.
type HostSearch struct {
	ID                    string   `json:"id,omitempty"`
	Keywords              string   `json:"keywords,omitempty"`
	DistinguishedName     []string `json:"distinguished_name,omitempty"`
	ExternalID            string   `json:"external_id,omitempty"`
	InstanceID            string   `json:"instance_id,omitempty"`
	SourceID              string   `json:"source_id,omitempty"`
	CommonName            []string `json:"common_name,omitempty"`
	Organization          []string `json:"organization,omitempty"`
	OrganizationalUnit    []string `json:"organizational_unit,omitempty"`
	Address               []string `json:"address,omitempty"`
	Service               []string `json:"service,omitempty"`
	Port                  []int    `json:"port,omitempty"`
	Zone                  []string `json:"zone,omitempty"`
	HostType              []string `json:"host_type,omitempty"`
	HostClassification    []string `json:"host_classification,omitempty"`
	Role                  []string `json:"role,omitempty"`
	Scope                 []string `json:"scope,omitempty"`
	IgnoreDisabledSources bool     `json:"ignore_disabled_sources,omitempty"`
	Tags                  []string `json:"tags,omitempty"`
	AccessGroupIDs        []string `json:"access_group_ids,omitempty"`
	CloudProviders        []string `json:"cloud_providers,omitempty"`
	CloudProviderRegions  []string `json:"cloud_provider_regions,omitempty"`
	Deployable            bool     `json:"deployable,omitempty"`
	Statuses              []string `json:"statuses,omitempty"`
	Disabled              string   `json:"disabled,omitempty"`
}

// SessionRecordingOptions optional host options to disable session recording per feature.
type SessionRecordingOptions struct {
	DisableClipboardRecording    bool `json:"disable_clipboard_recording"`
	DisableFileTransferRecording bool `json:"disable_file_transfer_recording"`
}

// HostResponse host response definition.
type HostResponse struct {
	ID     string `json:"id"`
	Action string `json:"action"`
}

// HostResolve host resolve request definition.
type HostResolve struct {
	Service string `json:"service"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// HostDeployable host deployable request definition.
type HostDeployable struct {
	Deployable bool `json:"deployable"`
}

// HostDisabled host disabled request definition.
type HostDisabled struct {
	Disabled bool `json:"disabled"`
}

// Whitelist whitelist definition.
type Whitelist struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Comment           string   `json:"comment,omitempty"`
	Type              string   `json:"type"`
	WhiteListPatterns []string `json:"whitelist_patterns,omitempty"`
	Author            string   `json:"author"`
	Created           string   `json:"created"`
	UpdatedBy         string   `json:"updated_by,omitempty"`
	Updated           string   `json:"updated,omitempty"`
}

// WhitelistSearch whitelist search request definition.
type WhitelistSearch struct {
	Keywords string `json:"keywords"`
}

// WhitelistEvaluate whitelist evaluate request definition.
type WhitelistEvaluate struct {
	WhiteList     Whitelist `json:"whitelist"`
	RShellVariant string    `json:"rshell_variant"`
	Commands      []string  `json:"commands"`
}

// WhitelistEvaluateResponse white list evaluate response definition.
type WhitelistEvaluateResponse struct {
	WhiteListPatternResults []WhitelistPatternResult `json:"whitelist_pattern_results"`
	CommandResults          []CommandResult          `json:"command_results"`
}

// CommandResult command result definition.
type CommandResult struct {
	Command string `json:"command"`
	Allowed bool   `json:"allowed"`
}

// WhitelistPatternResult whitelist pattern result definition.
type WhitelistPatternResult struct {
	WhiteListPattern string   `json:"whitelist_pattern"`
	Status           []string `json:"status"`
}

// Principal of the target host
type Principal struct {
	ID             string                 `json:"principal"`
	Roles          []rolestore.RoleHandle `json:"roles"`
	Source         string                 `json:"source"`
	UseUserAccount bool                   `json:"use_user_account"`
	Passphrase     string                 `json:"passphrase"`
	Applications   []string               `json:"applications"`
}

// Host defines PrivX target
type Host struct {
	ID                      string                   `json:"id"`
	Deployable              *bool                    `json:"deployable,omitempty"`
	Tofu                    *bool                    `json:"tofu,omitempty"`
	StandAloneHost          bool                     `json:"stand_alone_host"`
	ExternalID              string                   `json:"external_id"`
	InstanceID              string                   `json:"instance_id"`
	SSHHostPubKeys          []HostSSHPubKeys         `json:"ssh_host_public_keys"`
	HostCertificateRaw      string                   `json:"host_certificate_raw"`
	HostCertificate         *HostCertificateInfo     `json:"host_certificate,omitempty"`
	ContactAddress          string                   `json:"contact_address"`
	PasswordRotationEnabled bool                     `json:"password_rotation_enabled"`
	Services                []HostService            `json:"services"`
	Principals              []HostPrincipals         `json:"principals"`
	PasswordRotation        *RotationMetadata        `json:"password_rotation,omitempty"`
	SourceID                string                   `json:"source_id"`
	AccessGroupID           string                   `json:"access_group_id"`
	CloudProvider           string                   `json:"cloud_provider"`
	CloudProviderRegion     string                   `json:"cloud_provider_region"`
	Status                  []HostStatus             `json:"status"`
	Created                 string                   `json:"created"`
	Updated                 string                   `json:"updated" diff:"-"`
	UpdatedBy               string                   `json:"updated_by"`
	DistinguishedName       string                   `json:"distinguished_name"`
	CommonName              string                   `json:"common_name"`
	Organization            string                   `json:"organization"`
	OrganizationalUnit      string                   `json:"organizational_unit"`
	Zone                    string                   `json:"zone"`
	Scope                   []string                 `json:"scope"`
	HostType                string                   `json:"host_type"`
	HostClassification      string                   `json:"host_classification"`
	Comment                 string                   `json:"comment"`
	Addresses               []string                 `json:"addresses"`
	AuditEnabled            *bool                    `json:"audit_enabled,omitempty"`
	Tags                    []string                 `json:"tags"`
	UserMessage             string                   `json:"user_message"`
	Disabled                string                   `json:"disabled"`
	SessionRecordingOptions *SessionRecordingOptions `json:"session_recording_options,omitempty"`
	Deleted                 bool                     `json:"deleted,omitempty"`
}

type HostStatus struct {
	Key   string `json:"k"`
	Value string `json:"v"`
}

type RotationMetadata struct {
	AccessGroupID                    string               `json:"access_group_id"`
	UseMainAccount                   bool                 `json:"use_main_account"`
	OperatingSystem                  string               `json:"operating_system"`
	WinrmAddress                     string               `json:"winrm_address"`
	WinrmPort                        int                  `json:"winrm_port,omitempty"`
	CertificateValidationOptions     string               `json:"certificate_validation_options"`
	WinRMHostCertificateTrustAnchors string               `json:"winrm_host_certificate_trust_anchors"`
	Protocol                         string               `json:"protocol"`
	RotationStatus                   []RotationStatusItem `json:"rotation_status,omitempty"`
	PasswordPolicyId                 string               `json:"password_policy_id"`
	ScriptTemplateId                 string               `json:"script_template_id"`
	Created                          *time.Time           `json:"created,omitempty"`
	Updated                          *time.Time           `json:"updated,omitempty"`
	CreatedBy                        string               `json:"created_by,omitempty"`
	UpdatedBy                        string               `json:"updated_by,omitempty"`
}

type RotationStatusItem struct {
	Account          string     `json:"principal"`
	LastRotated      *time.Time `json:"last_rotated,omitempty"`
	LastError        *time.Time `json:"last_error,omitempty"`
	LastErrorDetails string     `json:"last_error_details"`
}

type HostPrincipals struct {
	Principal              string                             `json:"principal"`
	TargetDomain           *secretsmanager.TargetDomainHandle `json:"target_domain,omitempty"`
	Passphrase             string                             `json:"passphrase"`
	Rotate                 bool                               `json:"rotate"`
	UseForPasswordRotation bool                               `json:"use_for_password_rotation"`
	UsernameAttribute      string                             `json:"username_attribute"`
	UseUserAccount         bool                               `json:"use_user_account"`
	Source                 string                             `json:"source"`
	Roles                  []HostRole                         `json:"roles"`
	Applications           []HostPrincipalApplications        `json:"applications"`
	ServiceOptions         *HostServiceOptions                `json:"service_options,omitempty"`
	CommandRestrictions    HostCommandRestrictions            `json:"command_restrictions,omitempty"`
}

type HostCommandRestrictions struct {
	Enabled          bool             `json:"enabled"`
	RShellVariant    string           `json:"rshell_variant,omitempty"`
	DefaultWhiteList WhiteListHandle  `json:"default_whitelist"`
	WhiteLists       []WhiteListGrant `json:"whitelists"`
	AllowNoMatch     bool             `json:"allow_no_match,omitempty"`
	AuditMatch       bool             `json:"audit_match,omitempty"`
	AuditNoMatch     bool             `json:"audit_no_match,omitempty"`
	Banner           string           `json:"banner,omitempty"`
}

type WhiteListGrant struct {
	WhiteList WhiteListHandle `json:"whitelist"`
	Roles     []HostRole      `json:"roles"`
}

type HostRole struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted,omitempty"`
}

type WhiteListHandle struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

type HostServiceOptions struct {
	SSHServiceOptions *SSHServiceOptions `json:"ssh,omitempty"`
	RDPServiceOptions *RDPServiceOptions `json:"rdp,omitempty"`
	WebServiceOptions *WebServiceOptions `json:"web,omitempty"`
	VNCServiceOptions *VNCServiceOptions `json:"vnc,omitempty"`
	DBServiceOptions  *DBServiceOptions  `json:"db,omitempty"`
}

type DBServiceOptions struct {
	MaxBytesUpload   int64 `json:"max_bytes_upload,omitempty"`
	MaxBytesDownload int64 `json:"max_bytes_download,omitempty"`
}

type VNCServiceOptions struct {
	FileTransfer bool `json:"file_transfer"`
	Clipboard    bool `json:"clipboard"`
}

type WebServiceOptions struct {
	FileTransfer bool `json:"file_transfer"`
	Audio        bool `json:"audio"`
	Clipboard    bool `json:"clipboard"`
}

type RDPServiceOptions struct {
	FileTransfer bool `json:"file_transfer"`
	Audio        bool `json:"audio"`
	Clipboard    bool `json:"clipboard"`
}

type SSHServiceOptions struct {
	Shell        bool `json:"shell"`
	FileTransfer bool `json:"file_transfer"`
	Exec         bool `json:"exec"`
	Tunnels      bool `json:"tunnels"`
	X11          bool `json:"x11"`
	Other        bool `json:"other"`
}

type HostPrincipalApplications struct {
	Name             string `json:"name,omitempty"`
	Application      string `json:"application,omitempty"`
	Arguments        string `json:"arguments,omitempty"`
	WorkingDirectory string `json:"working_directory,omitempty"`
}

type HostService struct {
	Service                      string                  `json:"service"`
	Address                      string                  `json:"address"`
	Port                         int                     `json:"port"`
	UseForPasswordRotation       bool                    `json:"use_for_password_rotation"`
	TunnelPort                   int                     `json:"ssh_tunnel_port"`
	UsePlainTextVNC              bool                    `json:"use_plaintext_vnc"`
	Source                       string                  `json:"source"`
	Realm                        string                  `json:"realm,omitempty"`
	LoginPageURL                 string                  `json:"login_page_url"`
	UsernameFieldName            string                  `json:"username_field_name"`
	PasswordFieldName            string                  `json:"password_field_name"`
	LoginRequestUrl              string                  `json:"login_request_url"`
	LoginRequestPasswordProperty string                  `json:"login_request_password_property"`
	AuthType                     string                  `json:"auth_type"`
	HealthCheckStatus            string                  `json:"status"`
	HealthCheckStatusUpdated     string                  `json:"status_updated"`
	AllowedDomains               []string                `json:"allowed_domains"`
	Browser                      string                  `json:"browser"`
	BrowserKioskMode             bool                    `json:"kiosk_mode"`
	BrowserUrlBar                bool                    `json:"enable_urlbar"`
	BrowserNaviBar               bool                    `json:"enable_navibar"`
	BrowserNaviBarAutoHide       bool                    `json:"autohide_navibar"`
	BrowserDevTools              bool                    `json:"enable_devtools"`
	BrowserPopups                bool                    `json:"enable_popups"`
	BrowserWebCompatibleMode     bool                    `json:"enable_web_compatibility_mode"`
	BrowserTimeZone              string                  `json:"timezone"`
	WebIdleTimeLimit             int                     `json:"idle_time_limit"`
	ServiceVersion               string                  `json:"service_version"`
	Created                      time.Time               `json:"created"`
	Updated                      time.Time               `json:"updated"`
	CertificateTemplate          string                  `json:"certificate_template"`
	AllowModifiedWebParams       bool                    `json:"allow_modified_web_params"`
	ProtocolVersion              string                  `json:"protocol_version,omitempty"`
	Latency                      int                     `json:"latency_in_microseconds,omitempty"`
	DB                           HostServiceDBParameters `json:"db"`
	UseLegacyCipherSuites        bool                    `json:"use_legacy_cipher_suites"`
	TLSMinVersion                string                  `json:"tls_min_version"`
	TLSMaxVersion                string                  `json:"tls_max_version"`
}

type HostSSHPubKeys struct {
	Key         string `json:"key"`
	FingerPrint string `json:"fingerprint"`
}

type HostServiceDBParameters struct {
	Protocol                   string `json:"protocol"`
	TLSCertificateValidation   string `json:"tls_certificate_validation"`
	TLSCertificateTrustAnchors string `json:"tls_certificate_trust_anchors"`
	AuditSkipBytes             int64  `json:"audit_skip_bytes"`
}

type HostCertificateInfo struct {
	Subject           string   `json:"subject,omitempty"`
	Issuer            string   `json:"issuer,omitempty"`
	Serial            string   `json:"serial,omitempty"`
	NotBefore         string   `json:"not_before,omitempty"`
	NotAfter          string   `json:"not_after,omitempty"`
	DNSNames          []string `json:"dns_names,omitempty"`
	EmailAddresses    []string `json:"email_addresses,omitempty"`
	IPAddresses       []string `json:"ip_addresses,omitempty"`
	URIs              []string `json:"uris,omitempty"`
	FingerPrintSHA1   string   `json:"fingerprint_sha1,omitempty"`
	FingerPrintSHA256 string   `json:"fingerprint_sha256,omitempty"`
}
