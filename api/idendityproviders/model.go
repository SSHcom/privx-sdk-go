//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package idendityproviders

type IdentityProvider struct {
	ID        string `json:"id"`
	Name      string `json:"name" validate:"required,min=2,max=100"`
	TokenType string `json:"token_type" validate:"required,eq=JWT"`

	JWTIssuer                     string `json:"jwt_issuer" validate:"required"`
	JWTAudience                   string `json:"jwt_audience"`
	JWTSubjectType                string `json:"jwt_subject_type" validate:"required,eq=plain|eq=dn"`
	JWTSubjectDNUsernameAttribute string `json:"jwt_subject_dn_username_attribute,omitempty" validate:"required_if=JWTSubjectType dn"`

	CustomAttributes []CustomAttributeValidation `json:"custom_attributes,omitempty" validate:"dive"`

	PublicKey       []PublicKey `json:"public_key,omitempty" validate:"dive"`
	PublicKeyMethod string      `json:"public_key_method" validate:"required,eq=static|eq=x5u|eq=x5u-publickey"`

	// Used for validating certs fetched from x5u urls
	X5uTrustAnchor string `json:"x5u_trust_anchor,omitempty" validate:"required_if=PublicKeyMethod x5u"`
	// Optional TLS trust anchor cert used when doing x5u https requests
	X5uTLSTrustAnchor string `json:"x5u_tls_trust_anchor,omitempty" validate:"required_if=PublicKeyMethod x5u-publickey"`

	X5uPrefix      string `json:"x5u_prefix,omitempty" validate:"required_if=PublicKeyMethod x5u-publickey"`
	UsersDirectory string `json:"users_directory" validate:"required"`

	Enabled   bool   `json:"enabled"`
	Author    string `json:"author"`
	Created   string `json:"created"`
	Updated   string `json:"updated,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}

type CustomAttributeValidation struct {
	FieldName     string `json:"field_name" validate:"required"`
	Type          string `json:"type" validate:"required,eq=string_pattern|eq=numeric_range|eq=ip_range|eq=ip_client"`
	ExpectedValue string `json:"expected_value" validate:"required_if=Type string_pattern,max=2042"`
	Start         string `json:"start" validate:"required_if=Type numeric_range,required_if=Type ip_range"`
	End           string `json:"end" validate:"required_if=Type numeric_range,required_if=Type ip_range"`
}

type PublicKey struct {
	KeyID     string `json:"key_id" validate:"required"`
	Comment   string `json:"comment,omitempty" validate:"max=2042"`
	PublicKey string `json:"public_key,omitempty" validate:"required"`
}

type IdentityProviderResponse struct {
	Count int                `json:"count"`
	Items []IdentityProvider `json:"items"`
}

type IdentityProviderSearch struct {
	Keywords string `json:"keywords,omitempty"`
}

type IdentityProviderResolveUserRequest struct {
	Principal string `json:"principal" validate:"required"`
}
type IdentityProviderCreateResponse struct {
	ID string `json:"id"`
}
type Params struct {
	Offset  int    `json:"offset,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Sortkey string `json:"sortkey,omitempty"`
	Sortdir string `json:"sortdir,omitempty"`
}
