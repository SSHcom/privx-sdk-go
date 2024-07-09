//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// RoleStore is a role-store client instance.
type RoleStore struct {
	api restapi.Connector
}

type usersResult struct {
	Count int    `json:"count"`
	Items []User `json:"items"`
}

type rolesResult struct {
	Count int    `json:"count"`
	Items []Role `json:"items"`
}

type sourcesResult struct {
	Count int      `json:"count"`
	Items []Source `json:"items"`
}

type awsrolesResult struct {
	Count int           `json:"count"`
	Items []AWSRoleLink `json:"items"`
}

type awsTokenResult struct {
	Count int        `json:"count"`
	Items []AWSToken `json:"items"`
}

type principalkeysResult struct {
	Count int            `json:"count"`
	Items []PrincipalKey `json:"items"`
}

type authorizedkeysResult struct {
	Count int             `json:"count"`
	Items []AuthorizedKey `json:"items"`
}

type collectorsResult struct {
	Count int                `json:"count"`
	Items []LogconfCollector `json:"items"`
}

// New creates a new role-store client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *RoleStore {
	return &RoleStore{api: api}
}

// Sources get all sources.
func (store *RoleStore) Sources() ([]Source, error) {
	result := sourcesResult{}

	_, err := store.api.
		URL("/role-store/api/v1/sources").
		Get(&result)

	return result.Items, err
}

// CreateSource create a new source
func (store *RoleStore) CreateSource(source Source) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/sources").
		Post(&source, &object)

	return object.ID, err
}

// Source returns a source
func (store *RoleStore) Source(sourceID string) (*Source, error) {
	source := &Source{}

	_, err := store.api.
		URL("/role-store/api/v1/sources/%s", url.PathEscape(sourceID)).
		Get(source)

	return source, err
}

// DeleteSource delete a source
func (store *RoleStore) DeleteSource(sourceID string) error {
	_, err := store.api.
		URL("/role-store/api/v1/sources/%s", sourceID).
		Delete()

	return err
}

// UpdateSource update existing source
func (store *RoleStore) UpdateSource(sourceID string, source *Source) error {
	_, err := store.api.
		URL("/role-store/api/v1/sources/%s", url.PathEscape(sourceID)).
		Put(source)

	return err
}

// RefreshSources refresh all host and user sources
func (store *RoleStore) RefreshSources(sourceIDs []string) error {
	_, err := store.api.
		URL("/role-store/api/v1/sources/refresh").
		Post(&sourceIDs)

	return err
}

// AWSRoleLinks returns all aws roles.
func (store *RoleStore) AWSRoleLinks(refresh bool) ([]AWSRoleLink, error) {
	result := awsrolesResult{}
	filters := Params{
		Refresh: refresh,
	}

	_, err := store.api.
		URL("/role-store/api/v1/awsroles").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// AWSRoleLink returns existing single aws role
func (store *RoleStore) AWSRoleLink(awsroleID string) (*AWSRoleLink, error) {
	role := &AWSRoleLink{}

	_, err := store.api.
		URL("/role-store/api/v1/awsroles/%s", url.PathEscape(awsroleID)).
		Get(role)

	return role, err
}

// DeleteAWSRoleLInk delete a aws role
func (store *RoleStore) DeleteAWSRoleLInk(awsroleID string) error {
	_, err := store.api.
		URL("/role-store/api/v1/awsroles/%s", awsroleID).
		Delete()

	return err
}

// UpdateAWSRoleLink update existing aws role
func (store *RoleStore) UpdateAWSRoleLink(awsRoleID string, roles []RoleRef) error {
	_, err := store.api.
		URL("/role-store/api/v1/awsroles/%s/roles", url.PathEscape(awsRoleID)).
		Put(&roles)

	return err
}

// LinkedRoles return AWS role granting PrivX roles
func (store *RoleStore) LinkedRoles(awsroleID string) ([]AWSRoleLink, error) {
	result := awsrolesResult{}

	_, err := store.api.
		URL("/role-store/api/v1/awsroles/%s/roles", url.PathEscape(awsroleID)).
		Get(&result)

	return result.Items, err
}

// Roles gets all configured roles.
func (store *RoleStore) Roles(offset, limit int, sortkey, sortdir string) ([]Role, error) {
	result := rolesResult{}

	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles").Query(filters).Get(&result)

	return result.Items, err
}

// CreateRole creates new role
func (store *RoleStore) CreateRole(role Role) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles").
		Post(&role, &object)

	return object.ID, err
}

// ResolveRoles searches give role name and returns corresponding ids
func (store *RoleStore) ResolveRoles(names []string) ([]RoleRef, error) {
	var result struct {
		Count int       `json:"count"`
		Items []RoleRef `json:"items"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles/resolve").
		Post(&names, &result)

	return result.Items, err
}

// EvaluateRole evaluate a new role definition
func (store *RoleStore) EvaluateRole(role *Role) ([]User, error) {
	var result struct {
		Count int    `json:"count"`
		Items []User `json:"items"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles/evaluate").
		Post(role, &result)

	return result.Items, err
}

// Role gets information about the argument role ID.
func (store *RoleStore) Role(roleID string) (*Role, error) {
	role := &Role{}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s", url.PathEscape(roleID)).
		Get(role)

	return role, err
}

// DeleteRole delete a role
func (store *RoleStore) DeleteRole(roleID string) error {
	_, err := store.api.
		URL("/role-store/api/v1/roles/%s", roleID).
		Delete()

	return err
}

// UpdateRole update existing role
func (store *RoleStore) UpdateRole(roleID string, role *Role) error {
	_, err := store.api.
		URL("/role-store/api/v1/roles/%s", url.PathEscape(roleID)).
		Put(role)

	return err
}

// GetRoleMembers gets all members (users) of the argument role ID.
func (store *RoleStore) GetRoleMembers(roleID string, offset, limit int, sortkey, sortdir string) ([]User, error) {
	result := usersResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/members", url.PathEscape(roleID)).
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// AWSToken returns AWS token for a specified role
func (store *RoleStore) AWSToken(roleID, tokencode string, ttl int) ([]AWSToken, error) {
	result := awsTokenResult{}
	filters := Params{
		Tokencode: tokencode,
		TTL:       ttl,
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/awstoken", url.PathEscape(roleID)).
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// PrincipalKeys returns all principal keys
func (store *RoleStore) PrincipalKeys(roleID string) ([]PrincipalKey, error) {
	result := principalkeysResult{}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/principalkeys", url.PathEscape(roleID)).
		Get(&result)

	return result.Items, err
}

// GeneratePrincipalKey generate new principal key for existing role
func (store *RoleStore) GeneratePrincipalKey(roleID string) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/generate", url.PathEscape(roleID)).
		Post(nil, &object)

	return object.ID, err
}

// ImportPrincipalKey import new principal key for existing role
func (store *RoleStore) ImportPrincipalKey(key PrivateKey, roleID string) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/import", url.PathEscape(roleID)).
		Post(&key, &object)

	return object.ID, err
}

// PrincipalKey returns a role's principal key object.
func (store *RoleStore) PrincipalKey(roleID, keyID string) (*PrincipalKey, error) {
	key := &PrincipalKey{}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/%s", url.PathEscape(roleID), url.PathEscape(keyID)).
		Get(key)

	return key, err
}

// DeletePrincipalKey delete a role's principal key
func (store *RoleStore) DeletePrincipalKey(roleID, keyID string) error {
	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/%s", roleID, keyID).
		Delete()

	return err
}

// User gets information about the argument user ID.
func (store *RoleStore) User(userID string) (*User, error) {
	user := &User{}

	_, err := store.api.
		URL("/role-store/api/v1/users/%s", url.PathEscape(userID)).
		Get(user)

	return user, err
}

// UserSettings get specific user settings
func (store *RoleStore) UserSettings(userID string) (*json.RawMessage, error) {
	settings := &json.RawMessage{}

	_, err := store.api.
		URL("/role-store/api/v1/users/%s/settings", url.PathEscape(userID)).
		Get(&settings)

	return settings, err
}

// UpdateUserSettings update specific user's settings
func (store *RoleStore) UpdateUserSettings(settings *json.RawMessage, userID string) error {
	_, err := store.api.
		URL("/role-store/api/v1/users/%s/settings", url.PathEscape(userID)).
		Put(settings)

	return err
}

// UserRoles gets the roles of the argument user ID.
func (store *RoleStore) UserRoles(userID string) ([]Role, error) {
	result := rolesResult{}
	_, err := store.api.
		URL("/role-store/api/v1/users/%s/roles", url.PathEscape(userID)).
		Get(&result)

	return result.Items, err
}

// GrantUserRole adds the specified role for the user. If the user
// already has the role, this function does nothing.
func (store *RoleStore) GrantUserRole(userID, roleID string) error {
	// Get user's current roles.
	roles, err := store.UserRoles(userID)
	if err != nil {
		return err
	}
	// Does user already have the specified role?
	for _, role := range roles {
		if role.ID == roleID {
			// Already granted.
			return nil
		}
	}

	// Get new role.
	role, err := store.Role(roleID)
	if err != nil {
		return err
	}

	// Add an explicit role grant request.
	roles = append(roles, Role{
		ID:       role.ID,
		Explicit: true,
	})

	return store.setUserRoles(userID, roles)
}

// RevokeUserRole removes the specified role from the user. If the
// user does not have the role, this function does nothing.
func (store *RoleStore) RevokeUserRole(userID, roleID string) error {
	// Get user's current roles.
	roles, err := store.UserRoles(userID)
	if err != nil {
		return err
	}
	// Remove role from user's roles.
	var newRoles []Role
	for _, role := range roles {
		if role.ID != roleID {
			newRoles = append(newRoles, role)
		}
	}
	if len(newRoles) == len(roles) {
		// User did not have the specified role.
		return nil
	}

	// Set new roles.
	return store.setUserRoles(userID, newRoles)
}

func (store *RoleStore) setUserRoles(userID string, roles []Role) error {
	_, err := store.api.
		URL("/role-store/api/v1/users/%s/roles", url.PathEscape(userID)).
		Put(roles)

	return err
}

// EnableMFA enable multifactor authentication
func (store *RoleStore) EnableMFA(userIDs []string) error {
	_, err := store.api.
		URL("/role-store/api/v1/users/mfa/enable").
		Post(&userIDs)

	return err
}

// DisableMFA disable multifactor authentication
func (store *RoleStore) DisableMFA(userIDs []string) error {
	_, err := store.api.
		URL("/role-store/api/v1/users/mfa/disable").
		Post(&userIDs)

	return err
}

// ResetMFA reset multifactor authentication
func (store *RoleStore) ResetMFA(userIDs []string) error {
	_, err := store.api.
		URL("/role-store/api/v1/users/mfa/reset").
		Post(&userIDs)

	return err
}

// ResolveUser resolve users role
func (store *RoleStore) ResolveUser(userID string) (*User, error) {
	user := &User{}

	_, err := store.api.
		URL("/role-store/api/v1/users/%s/resolve", url.PathEscape(userID)).
		Get(user)

	return user, err
}

// SearchUsers searches for users, matching the keywords and source
// criteria.
func (store *RoleStore) SearchUsers(offset, limit int, sortkey, sortdir string, searchBody UserSearchObject) ([]User, error) {
	result := usersResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}
	_, err := store.api.
		URL("/role-store/api/v1/users/search").
		Query(&filters).
		Post(searchBody, &result)

	return result.Items, err
}

// SearchUsersExternal searche users with user search parameters.
func (store *RoleStore) SearchUsersExternal(keywords, sourceID string) ([]User, error) {
	result := usersResult{}
	_, err := store.api.
		URL("/role-store/api/v1/users/search/external").
		Post(map[string]string{
			"keywords": keywords,
			"source":   sourceID,
		}, &result)

	return result.Items, err
}

// AuthorizedKeys return user's authorized keys
func (store *RoleStore) AuthorizedKeys(userID string) ([]AuthorizedKey, error) {
	result := authorizedkeysResult{}

	_, err := store.api.
		URL("/role-store/api/v1/users/%s/authorizedkeys", url.PathEscape(userID)).
		Get(&result)

	return result.Items, err
}

// CreateAuthorizedKey register an authorized key for user
func (store *RoleStore) CreateAuthorizedKey(key AuthorizedKey, userID string) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/users/%s/authorizedkeys", url.PathEscape(userID)).
		Post(&key, &object)

	return object.ID, err
}

// AuthorizedKey return user's authorized key
func (store *RoleStore) AuthorizedKey(userID, keyID string) (*AuthorizedKey, error) {
	key := &AuthorizedKey{}

	_, err := store.api.
		URL("/role-store/api/v1/users/%s/authorizedkeys/%s", url.PathEscape(userID), url.PathEscape(keyID)).
		Get(key)

	return key, err
}

// UpdateAuthorizedKey update authorized key for user
func (store *RoleStore) UpdateAuthorizedKey(key *AuthorizedKey, userID, keyID string) error {
	_, err := store.api.
		URL("/role-store/api/v1/users/%s/authorizedkeys/%s", url.PathEscape(userID), url.PathEscape(keyID)).
		Put(key)

	return err
}

// DeleteAuthorizedKey delete a user's authorized key
func (store *RoleStore) DeleteAuthorizedKey(userID, keyID string) error {
	_, err := store.api.
		URL("/role-store/api/v1/users/%s/authorizedkeys/%s", userID, keyID).
		Delete()

	return err
}

// LogconfCollectors returns all logconf collectors
func (store *RoleStore) LogconfCollectors() ([]LogconfCollector, error) {
	result := collectorsResult{}

	_, err := store.api.
		URL("/role-store/api/v1/logconf/collectors").
		Get(&result)

	return result.Items, err
}

// CreateLogconfCollector create a logconf collector
func (store *RoleStore) CreateLogconfCollector(conf LogconfCollector) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/logconf/collectors").
		Post(&conf, &object)

	return object.ID, err
}

// LogconfCollector returns existing single logconf collector
func (store *RoleStore) LogconfCollector(collectorID string) (*LogconfCollector, error) {
	conf := &LogconfCollector{}

	_, err := store.api.
		URL("/role-store/api/v1/logconf/collectors/%s", url.PathEscape(collectorID)).
		Get(conf)

	return conf, err
}

// UpdateLogconfCollector update existing logconf collector
func (store *RoleStore) UpdateLogconfCollector(collectorID string, conf *LogconfCollector) error {
	_, err := store.api.
		URL("/role-store/api/v1/logconf/collectors/%s", url.PathEscape(collectorID)).
		Put(conf)

	return err
}

// DeleteLogconfCollector delete a logconf collector
func (store *RoleStore) DeleteLogconfCollector(collectorID string) error {
	_, err := store.api.
		URL("/role-store/api/v1/logconf/collectors/%s", collectorID).
		Delete()

	return err
}

// AllAuthorizedKeys returns all authorized keys
func (store *RoleStore) AllAuthorizedKeys(offset, limit int, sortdir, sortkey string) ([]AuthorizedKey, error) {
	result := authorizedkeysResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortdir: sortdir,
		Sortkey: sortkey,
	}

	_, err := store.api.
		URL("/role-store/api/v1/authorizedkeys").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// ResolveAuthorizedKey resolve authorized keys
func (store *RoleStore) ResolveAuthorizedKey(resolve ResolveAuthorizedKey) ([]AuthorizedKey, error) {
	result := authorizedkeysResult{}

	_, err := store.api.
		URL("/role-store/api/v1/authorizedkeys/resolve").
		Post(&resolve, &result)

	return result.Items, err
}

/////////////////////////////
//// Idendity providers ////
///////////////////////////

// List all identity providers.
func (store *RoleStore) GetAllIdendityProviders(offset, limit int) (IdentityProviderResponse, error) {
	result := IdentityProviderResponse{}

	filters := Params{
		Offset: offset,
		Limit:  limit,
	}

	_, err := store.api.
		URL("/role-store/api/v1/identity-providers").
		Query(&filters).
		Get(&result)
	return result, err
}

// Create a new Identity Provider.
func (store *RoleStore) CreateIdendityProvider(newIP IdentityProvider) (IdentityProviderCreateResponse, error) {
	result := IdentityProviderCreateResponse{}

	_, err := store.api.
		URL("/role-store/api/v1/identity-providers").
		Post(newIP, &result)

	return result, err
}

// Get Identity Provider by ID.
func (store *RoleStore) GetIdendityProviderByID(ID string) (IdentityProvider, error) {
	result := IdentityProvider{}

	_, err := store.api.
		URL("/role-store/api/v1/identity-providers/%s", url.PathEscape(ID)).
		Get(&result)

	return result, err
}

// Delete Identity Provider by ID.
func (store *RoleStore) DeleteIdendityProviderByID(ID string) error {

	_, err := store.api.
		URL("/role-store/api/v1/identity-providers/%s", url.PathEscape(ID)).
		Delete()

	return err
}

// Update a Identity Provider.
func (store *RoleStore) UpdateIdendityProvider(UpdatedIP IdentityProvider, ID string) error {

	_, err := store.api.
		URL("/role-store/api/v1/identity-providers/%s", url.PathEscape(ID)).
		Put(UpdatedIP)

	return err
}

// Search Identity Providers.
func (store *RoleStore) SearchIdendityProviders(offset, limit int, sortkey, sortdir, keywords string) (IdentityProviderResponse, error) {
	result := IdentityProviderResponse{}

	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: strings.ToUpper(sortdir),
	}
	body := IdentityProviderSearch{
		Keywords: keywords,
	}
	_, err := store.api.
		URL("/role-store/api/v1/identity-providers/search").
		Query(filters).
		Post(body, &result)

	return result, err
}
