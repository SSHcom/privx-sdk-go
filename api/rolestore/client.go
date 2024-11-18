//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

import (
	"encoding/json"
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// RoleStore is a role-store client instance.
type RoleStore struct {
	api restapi.Connector
}

// New role store client constructor.
func New(api restapi.Connector) *RoleStore {
	return &RoleStore{api: api}
}

// MARK: Status
// Status get role store microservice status.
func (c *RoleStore) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/role-store/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Sources
// GetSources get sources.
func (c *RoleStore) GetSources() (*response.ResultSet[Source], error) {
	sources := &response.ResultSet[Source]{}

	_, err := c.api.
		URL("/role-store/api/v1/sources").
		Get(&sources)

	return sources, err
}

// CreateSource create source.
func (c *RoleStore) CreateSource(source *Source) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/role-store/api/v1/sources").
		Post(&source, &identifier)

	return identifier, err
}

// GetSource get source by id.
func (c *RoleStore) GetSource(sourceID string) (*Source, error) {
	source := &Source{}

	_, err := c.api.
		URL("/role-store/api/v1/sources/%s", sourceID).
		Get(&source)

	return source, err
}

// UpdateSource update source.
func (c *RoleStore) UpdateSource(sourceID string, source *Source) error {
	_, err := c.api.
		URL("/role-store/api/v1/sources/%s", sourceID).
		Put(&source)

	return err
}

// DeleteSource delete source.
func (c *RoleStore) DeleteSource(sourceID string) error {
	_, err := c.api.
		URL("/role-store/api/v1/sources/%s", sourceID).
		Delete()

	return err
}

// RefreshSources refresh sources.
func (c *RoleStore) RefreshSources(sourceIDs []string) error {
	_, err := c.api.
		URL("/role-store/api/v1/sources/refresh").
		Post(&sourceIDs)

	return err
}

// MARK: AWS Roles
// GetAWSRoles get AWS roles.
func (c *RoleStore) GetAWSRoles(opts ...filters.Option) (*response.ResultSet[AWSRole], error) {
	roles := &response.ResultSet[AWSRole]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/awsroles").
		Query(params).
		Get(&roles)

	return roles, err
}

// GetAWSRole get AWS role by id.
func (c *RoleStore) GetAWSRole(awsRoleID string) (*AWSRole, error) {
	role := &AWSRole{}

	_, err := c.api.
		URL("/role-store/api/v1/awsroles/%s", awsRoleID).
		Get(role)

	return role, err
}

// DeleteAWSRole delete AWS role.
func (c *RoleStore) DeleteAWSRole(awsRoleID string) error {
	_, err := c.api.
		URL("/role-store/api/v1/awsroles/%s", awsRoleID).
		Delete()

	return err
}

// GetLinkedRoles get AWS role granting PrivX roles.
func (c *RoleStore) GetLinkedRoles(awsRoleID string) (*response.ResultSet[LinkedPrivXRole], error) {
	roles := &response.ResultSet[LinkedPrivXRole]{}

	_, err := c.api.
		URL("/role-store/api/v1/awsroles/%s/roles", awsRoleID).
		Get(&roles)

	return roles, err
}

// UpdateAWSRole update AWS role granting PrivX roles.
func (c *RoleStore) UpdateAWSRole(awsRoleID string, roles []LinkedPrivXRole) error {
	_, err := c.api.
		URL("/role-store/api/v1/awsroles/%s/roles", awsRoleID).
		Put(&roles)

	return err
}

// MARK: Users
// GetUser get user by id.
func (c *RoleStore) GetUser(userID string) (*User, error) {
	user := &User{}

	_, err := c.api.
		URL("/role-store/api/v1/users/%s", userID).
		Get(user)

	return user, err
}

// GetUserSettings get user settings.
func (c *RoleStore) GetUserSettings(userID string) (*json.RawMessage, error) {
	settings := &json.RawMessage{}

	_, err := c.api.
		URL("/role-store/api/v1/users/%s/settings", userID).
		Get(&settings)

	return settings, err
}

// UpdateUserSettings update specific user's settings
func (c *RoleStore) UpdateUserSettings(userID string, settings *UserSettings) error {
	_, err := c.api.
		URL("/role-store/api/v1/users/%s/settings", userID).
		Put(&settings)

	return err
}

// GetUserRoles get roles of user by id.
func (c *RoleStore) GetUserRoles(userID string) (*response.ResultSet[Role], error) {
	roles := &response.ResultSet[Role]{}
	_, err := c.api.
		URL("/role-store/api/v1/users/%s/roles", userID).
		Get(&roles)

	return roles, err
}

// UpdateUserRoles update user roles by id.
func (c *RoleStore) UpdateUserRoles(userID string, roles []Role) error {
	_, err := c.api.
		URL("/role-store/api/v1/users/%s/roles", userID).
		Put(roles)

	return err
}

// SetMFA enable, disable or reset mfa authentication.
func (c *RoleStore) SetMFA(userIDs []string, action MFAAction) error {
	_, err := c.api.
		URL("/role-store/api/v1/users/mfa/%s", action).
		Post(&userIDs)

	return err
}

// GetCurrentUserInfo get current user and user settings.
func (c *RoleStore) GetCurrentUserInfo() (*json.RawMessage, error) {
	current := &json.RawMessage{}

	_, err := c.api.
		URL("/role-store/api/v1/users/current").
		Get(&current)

	return current, err
}

// GetCurrentUserSettings get current user AWS roles.
func (c *RoleStore) GetCurrentAWSRoles() (*response.ResultSet[AWSRole], error) {
	roles := &response.ResultSet[AWSRole]{}

	_, err := c.api.
		URL("/role-store/api/v1/users/current/awsroles").
		Get(&roles)

	return roles, err
}

// GetCurrentUserAndSettings get current user settings.
func (c *RoleStore) GetCurrentUserSettings() (*json.RawMessage, error) {
	settings := &json.RawMessage{}

	_, err := c.api.
		URL("/role-store/api/v1/users/current/settings").
		Get(&settings)

	return settings, err
}

// UpdateCurrentUserSettings update current user settings.
func (c *RoleStore) UpdateCurrentUserSettings(settings *UserSettings) error {
	_, err := c.api.
		URL("/role-store/api/v1/users/current/settings").
		Put(&settings)

	return err
}

// ResolveUserRoles resolve user roles.
func (c *RoleStore) ResolveUserRoles(userID string) (*User, error) {
	user := &User{}

	_, err := c.api.
		URL("/role-store/api/v1/users/%s/resolve", userID).
		Get(&user)

	return user, err
}

// SearchUsers search users.
func (c *RoleStore) SearchUsers(search UserSearch, opts ...filters.Option) (*response.ResultSet[User], error) {
	users := &response.ResultSet[User]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/users/search").
		Query(params).
		Post(&search, &users)

	return users, err
}

// SearchExternalUsers search external users.
func (c *RoleStore) SearchExternalUsers(search UserSearch) (*response.ResultSet[User], error) {
	users := &response.ResultSet[User]{}

	_, err := c.api.
		URL("/role-store/api/v1/users/search/external").
		Post(&search, &users)

	return users, err
}

// GetUsersAuthorizedKeys get users authorized keys.
func (c *RoleStore) GetUsersAuthorizedKeys(userID string, opts ...filters.Option) (*response.ResultSet[AuthorizedKey], error) {
	keys := &response.ResultSet[AuthorizedKey]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/users/%s/authorizedkeys", userID).
		Query(params).
		Get(&keys)

	return keys, err
}

// CreateUserAuthorizedKey create authorized key for user.
func (c *RoleStore) CreateUserAuthorizedKey(userID string, key *AuthorizedKey) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/role-store/api/v1/users/%s/authorizedkeys", userID).
		Post(&key, &identifier)

	return identifier, err
}

// GetUserAuthorizedKey get user authorized key by id.
func (c *RoleStore) GetUserAuthorizedKey(userID, keyID string) (*AuthorizedKey, error) {
	key := &AuthorizedKey{}

	_, err := c.api.
		URL("/role-store/api/v1/users/%s/authorizedkeys/%s", userID, keyID).
		Get(&key)

	return key, err
}

// UpdateUserAuthorizedKey update user authorized key.
func (c *RoleStore) UpdateUserAuthorizedKey(userID, keyID string, key *AuthorizedKey) error {
	_, err := c.api.
		URL("/role-store/api/v1/users/%s/authorizedkeys/%s", userID, keyID).
		Put(&key)

	return err
}

// DeleteUserAuthorizedKey delete a user authorized key.
func (c *RoleStore) DeleteUserAuthorizedKey(userID, keyID string) error {
	_, err := c.api.
		URL("/role-store/api/v1/users/%s/authorizedkeys/%s", userID, keyID).
		Delete()

	return err
}

// GetCurrentUserAuthorizedKeys get current user authorized keys.
func (c *RoleStore) GetCurrentUserAuthorizedKeys(opts ...filters.Option) (*response.ResultSet[AuthorizedKey], error) {
	keys := &response.ResultSet[AuthorizedKey]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/users/current/authorizedkeys").
		Query(params).
		Get(&keys)

	return keys, err
}

// CreateCurrentUserAuthorizedKey create authorized key for current user.
func (c *RoleStore) CreateCurrentUserAuthorizedKey(key *AuthorizedKey) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/role-store/api/v1/users/current/authorizedkeys").
		Post(&key, &identifier)

	return identifier, err
}

// GetCurrentUserAuthorizedKey get current user authorized key by id.
func (c *RoleStore) GetCurrentUserAuthorizedKey(keyID string) (*AuthorizedKey, error) {
	key := &AuthorizedKey{}

	_, err := c.api.
		URL("/role-store/api/v1/users/current/authorizedkeys/%s", keyID).
		Get(&key)

	return key, err
}

// UpdateCurrentUserAuthorizedKey update current user authorized key.
func (c *RoleStore) UpdateCurrentUserAuthorizedKey(keyID string, key *AuthorizedKey) error {
	_, err := c.api.
		URL("/role-store/api/v1/users/current/authorizedkeys/%s", keyID).
		Put(&key)

	return err
}

// DeleteCurrentUserAuthorizedKey delete current a user authorized key.
func (c *RoleStore) DeleteCurrentUserAuthorizedKey(keyID string) error {
	_, err := c.api.
		URL("/role-store/api/v1/users/current/authorizedkeys/%s", keyID).
		Delete()

	return err
}

// GetRoles get roles.
func (c *RoleStore) GetRoles(opts ...filters.Option) (*response.ResultSet[Role], error) {
	roles := &response.ResultSet[Role]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/roles").
		Query(params).
		Get(&roles)

	return roles, err
}

// CreateRole creates role.
func (c *RoleStore) CreateRole(role *Role) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/role-store/api/v1/roles").
		Post(&role, &identifier)

	return identifier, err
}

// ResolveRoles resolve role names to role.
func (c *RoleStore) ResolveRoles(names []string) (*response.ResultSet[Role], error) {
	roles := &response.ResultSet[Role]{}

	_, err := c.api.
		URL("/role-store/api/v1/roles/resolve").
		Post(&names, &roles)

	return roles, err
}

// SearchRoles search roles.
func (c *RoleStore) SearchRoles(search RoleSearch, opts ...filters.Option) (*response.ResultSet[Role], error) {
	roles := &response.ResultSet[Role]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/roles/search").
		Query(params).
		Post(&search, &roles)

	return roles, err
}

// EvaluateRole evaluate role definition.
func (c *RoleStore) EvaluateRole(role *Role) (*response.ResultSet[User], error) {
	users := &response.ResultSet[User]{}

	_, err := c.api.
		URL("/role-store/api/v1/roles/evaluate").
		Post(role, &users)

	return users, err
}

// GetRole get role by id.
func (c *RoleStore) GetRole(roleID string) (*Role, error) {
	role := &Role{}

	_, err := c.api.
		URL("/role-store/api/v1/roles/%s", roleID).
		Get(&role)

	return role, err
}

// UpdateRole update role.
func (c *RoleStore) UpdateRole(roleID string, role *Role) error {
	_, err := c.api.
		URL("/role-store/api/v1/roles/%s", roleID).
		Put(&role)

	return err
}

// DeleteRole delete role.
func (c *RoleStore) DeleteRole(roleID string) error {
	_, err := c.api.
		URL("/role-store/api/v1/roles/%s", roleID).
		Delete()

	return err
}

// GetRoleMembers gets users of the role.
func (c *RoleStore) GetRoleMembers(roleID string, opts ...filters.Option) (*response.ResultSet[User], error) {
	users := &response.ResultSet[User]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/roles/%s/members", roleID).
		Query(params).
		Get(&users)

	return users, err
}

// GetAWSToken get AWS token for role.
func (c *RoleStore) GetAWSToken(roleID string, opts ...filters.Option) (*json.RawMessage, error) {
	token := &json.RawMessage{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/roles/%s/awstoken", roleID).
		Query(params).
		Get(&token)

	return token, err
}

// GetPrincipalKeys get roles principal keys.
func (c *RoleStore) GetPrincipalKeys(roleID string) (*response.ResultSet[RolePrincipalKey], error) {
	keys := &response.ResultSet[RolePrincipalKey]{}

	_, err := c.api.
		URL("/role-store/api/v1/roles/%s/principalkeys", roleID).
		Get(&keys)

	return keys, err
}

// CreatePrincipalKey create principal key for role.
func (c *RoleStore) CreatePrincipalKey(roleID string) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/generate", roleID).
		Post(nil, &identifier)

	return identifier, err
}

// ImportPrincipalKey import principal key for role.
func (c *RoleStore) ImportPrincipalKey(roleID string, key RolePrincipalKeyImport) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/import", roleID).
		Post(&key, &identifier)

	return identifier, err
}

// GetPrincipalKey get roles principal key.
func (c *RoleStore) GetPrincipalKey(roleID, keyID string) (RolePrincipalKey, error) {
	key := RolePrincipalKey{}

	_, err := c.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/%s", roleID, keyID).
		Get(&key)

	return key, err
}

// DeletePrincipalKey delete roles principal key.
func (c *RoleStore) DeletePrincipalKey(roleID, keyID string) error {
	_, err := c.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/%s", roleID, keyID).
		Delete()

	return err
}

// MARK: Identity Providers
// GetIdentityProviders get identity providers.
func (c *RoleStore) GetIdentityProviders(opts ...filters.Option) (*response.ResultSet[IdentityProvider], error) {
	providers := &response.ResultSet[IdentityProvider]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/identity-providers").
		Query(params).
		Get(&providers)

	return providers, err
}

// CreateIdentityProvider create a identity provider.
func (c *RoleStore) CreateIdentityProvider(provider *IdentityProvider) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/role-store/api/v1/identity-providers").
		Post(&provider, &identifier)

	return identifier, err
}

// GetIdentityProvider get identity provider by id.
func (c *RoleStore) GetIdentityProvider(providerID string) (*IdentityProvider, error) {
	provider := &IdentityProvider{}

	_, err := c.api.
		URL("/role-store/api/v1/identity-providers/%s", providerID).
		Get(&provider)

	return provider, err
}

// UpdateIdentityProvider update identity provider.
func (c *RoleStore) UpdateIdentityProvider(providerID string, provider *IdentityProvider) error {
	_, err := c.api.
		URL("/role-store/api/v1/identity-providers/%s", providerID).
		Put(&provider)

	return err
}

// DeleteIdentityProvider delete identity provider by id.
func (c *RoleStore) DeleteIdentityProvider(providerID string) error {
	_, err := c.api.
		URL("/role-store/api/v1/identity-providers/%s", providerID).
		Delete()

	return err
}

// SearchIdentityProviders search identity providers.
func (c *RoleStore) SearchIdentityProviders(search IdentityProviderSearch, opts ...filters.Option) (*response.ResultSet[IdentityProvider], error) {
	providers := &response.ResultSet[IdentityProvider]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/identity-providers/search").
		Query(params).
		Post(&search, &providers)

	return providers, err
}

// MARK: Authorized Keys
// GetAuthorizedKeys get authorized keys.
func (c *RoleStore) GetAuthorizedKeys(opts ...filters.Option) (*response.ResultSet[AuthorizedKey], error) {
	keys := &response.ResultSet[AuthorizedKey]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/role-store/api/v1/authorizedkeys").
		Query(params).
		Get(&keys)

	return keys, err
}

// ResolveAuthorizedKey resolve authorized key.
func (c *RoleStore) ResolveAuthorizedKey(resolve AuthorizedKeyResolve) (*AuthorizedKey, error) {
	key := &AuthorizedKey{}

	_, err := c.api.
		URL("/role-store/api/v1/authorizedkeys/resolve").
		Post(&resolve, &key)

	return key, err
}

// MARK: Logconf
// GetLogConfCollectors get logconf collectors.
func (c *RoleStore) GetLogConfCollectors() (*response.ResultSet[LogConfCollector], error) {
	collectors := &response.ResultSet[LogConfCollector]{}

	_, err := c.api.
		URL("/role-store/api/v1/logconf/collectors").
		Get(&collectors)

	return collectors, err
}

// CreateLogConfCollector create logconf collector.
func (c *RoleStore) CreateLogConfCollector(collector *LogConfCollector) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/role-store/api/v1/logconf/collectors").
		Post(&collector, &identifier)

	return identifier, err
}

// GetLogConfCollector get logconf collector by id.
func (c *RoleStore) GetLogConfCollector(collectorID string) (*LogConfCollector, error) {
	collector := &LogConfCollector{}

	_, err := c.api.
		URL("/role-store/api/v1/logconf/collectors/%s", collectorID).
		Get(&collector)

	return collector, err
}

// UpdateLogConfCollector update logconf collector.
func (c *RoleStore) UpdateLogConfCollector(collectorID string, collector *LogConfCollector) error {
	_, err := c.api.
		URL("/role-store/api/v1/logconf/collectors/%s", collectorID).
		Put(&collector)

	return err
}

// DeleteLogConfCollector delete logconf collector.
func (c *RoleStore) DeleteLogConfCollector(collectorID string) error {
	_, err := c.api.
		URL("/role-store/api/v1/logconf/collectors/%s", collectorID).
		Delete()

	return err
}
