package apiproxy

// ApiTarget object specifies all information necessary for performing access
// control for requests to an API target
type ApiTarget struct {
	// ID is the unique UUID for the API target, if set must be a valid UUID
	ID string `json:"id"`

	// Name is a unique human readable name for the API target. Required; Not empty; Trimmed.
	Name string `json:"name"`

	// Comment is an optional human readable comment. Trimmed.
	Comment string `json:"comment,omitempty"`

	// Tags are optional tags for the API target. Trimmed entries; empty tags rejected.
	Tags []string `json:"tags"`

	// AccessGroupID specifies the access group the API target is associated to. if set must be a valid UUID
	AccessGroupID string `json:"access_group_id"`

	// Roles are the required roles which grant access to the API target
	Roles []RoleHandle `json:"roles"`

	// AuthorizedEndpoints specify the address, scheme, method and path patterns
	// for matching requests. A request must match at least one authorized
	// endpoint to be authorized.
	AuthorizedEndpoints []ApiTargetEndpoint `json:"authorized_endpoints"`

	// UnauthorizedEndpoints specify the address, scheme, method and path
	// patterns for matching requests. An authorized request must not match any
	// unauthorized endpoint.
	UnauthorizedEndpoints []ApiTargetEndpoint `json:"unauthorized_endpoints"`

	// TLSTrustAnchors specify optional X.509 certificate trust anchors for
	// validating the api target TLS server certificates. These trust anchors
	// are used together with trust anchors configured in the host OS, in
	// api-proxy settings and in the access group.
	TLSTrustAnchors string `json:"tls_trust_anchors,omitempty"`

	// TLSInsecureSkipVerify turns off the api target TLS server certificate
	// validation. It should be used with great caution and only when strictly needed.
	TLSInsecureSkipVerify bool `json:"tls_insecure_skip_verify,omitempty"`

	// TargetCredential specifies the credentials used for authenticating to the api target.
	TargetCredential TargetCredential `json:"target_credential"` // Credential validated according to selected type.

	// Disabled specifies whether this api target is enabled or not. All request
	// to disabled api targets are rejected.
	// Can be one of NOT_DISABLED, BY_ADMIN, BY_LICENSE.
	Disabled string `json:"disabled"`

	// AuditEnabled specifies whether to session record requests to this target api.
	AuditEnabled bool `json:"audit_enabled,omitempty"`

	Created   string `json:"created"`
	Author    string `json:"author"`
	Updated   string `json:"updated,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}

// TargetCredential object contains the credentials for authenticating to the api target.
type TargetCredential struct {
	// Type is the api target credential type. Accepted values are "basicauth" and "token".
	Type string `json:"type"`

	// BasicAuthUsername is the username for credentials of type "basicauth"
	BasicAuthUsername string `json:"basic_auth_username,omitempty"`

	// BasicAuthPassword is the password for credentials of type "basicauth"
	BasicAuthPassword string `json:"basic_auth_password,omitempty"`

	// BearerToken is the static bearer token for credentials of type "token"
	BearerToken string `json:"bearer_token,omitempty"`

	// Certificate is the certificate for api target of type "certificate"
	Certificate string `json:"certificate,omitempty"`

	// PrivateKey is the private key for api target of type "certificate"
	PrivateKey string `json:"private_key,omitempty"`
}

// ApiTargetEndpoint API endpoint patterns that are matched against request URLs.
type ApiTargetEndpoint struct {
	// Host is matched against the host part in the request URL. Port may be
	// omitted from host if it matches the default port 80 for http or 443 for
	// https. Matching is done using case-insensitive exact match.
	Host string `json:"host"`

	// Protocols are matched against request URL scheme. Accepted values are
	// "http" ,"https" and "*" (for http or https).
	Protocols []string `json:"protocols"`

	// Methods are matched against request http method. Accepted values are
	// "GET", "PUT", "POST", "DELETE", "HEAD", "PATCH", "OPTIONS", "TRACE", "*"
	Methods []string `json:"methods"`

	// Paths are matched against request URL path. Following wildcards can be used:
	//   "*"   matches one path segment
	//   "**"  matches rest of the path
	Paths []string `json:"paths"`

	// AllowUnauthenticated specifies if unauthenticated requests should be allowed.
	AllowUnauthenticated bool `json:"allow_unauthenticated"`

	// NATTargetHost specifies optional api target host address. It is used as
	// the api target host address when forwarding requests to api targets over an extender.
	NATTargetHost string `json:"nat_target_host,omitempty"`
}

type RoleHandle struct {
	ID      string `json:"id"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

// ApiProxyCACertificateInfo api proxy x509 CA certificate information
type ApiProxyCACertificateInfo struct {
	Subject           string `json:"subject,omitempty"`
	Issuer            string `json:"issuer,omitempty"`
	Serial            string `json:"serial,omitempty"`
	NotBefore         string `json:"not_before,omitempty"`
	NotAfter          string `json:"not_after,omitempty"`
	FingerPrintSHA1   string `json:"fingerprint_sha1,omitempty"`
	FingerPrintSHA256 string `json:"fingerprint_sha256,omitempty"`
}

// ApiProxyAPIConf response for GetApiProxyConfig()
type ApiProxyAPIConf struct {
	Addresses     []string                   `json:"addresses,omitempty"`
	CACertificate *ApiProxyCACertificateInfo `json:"ca_certificate,omitempty"`
	Chain         string                     `json:"ca_certificate_chain,omitempty"`
}

type ApiTargetSearchRequest struct {
	// Keywords is a comma or space limited string of search keywords
	Keywords string `json:"keywords"`

	// Name is the search constraint for api target name
	Name string `json:"name"`

	// AccessGroupID is the search constraint for api target access group
	AccessGroupID string `json:"access_group_id"`

	// Tags is a search constraint for api target tags
	Tags []string `json:"tags"`
}
