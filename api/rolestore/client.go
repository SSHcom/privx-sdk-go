//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

import (
	"net/url"

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
	Count int       `json:"count"`
	Items []AWSRole `json:"items"`
}

type principalkeysResult struct {
	Count int            `json:"count"`
	Items []PrincipalKey `json:"items"`
}

// New creates a new role-store client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *RoleStore {
	return &RoleStore{api: api}
}

// ResolveUserRole resolve users role
func (store *RoleStore) ResolveUserRole(id string) (user *User, err error) {
	user = new(User)

	_, err = store.api.
		URL("/role-store/api/v1/users/%s/resolve", url.PathEscape(id)).
		Get(user)

	return
}

// UsersAWSRoles returns current user's AWS roles
func (store *RoleStore) UsersAWSRoles() ([]AWSRole, error) {
	result := awsrolesResult{}
	_, err := store.api.
		URL("/role-store/api/v1/users/current/awsroles").
		Get(&result)

	return result.Items, err
}

// DeletePrincipalKey delete a role's principal key
func (store *RoleStore) DeletePrincipalKey(roleID, keyID string) error {
	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/%s", roleID, keyID).
		Delete()

	return err
}

// PrincipalKey returns a role's principal key object.
func (store *RoleStore) PrincipalKey(roleID, keyID string) (key *PrincipalKey, err error) {
	key = new(PrincipalKey)

	_, err = store.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/%s", url.PathEscape(roleID), url.PathEscape(keyID)).
		Get(key)

	return
}

// ImportPrincipalKey import new principal key for existing role
func (store *RoleStore) ImportPrincipalKey(key PrivateKey, id string) (string, error) {
	var keyID struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/import", url.PathEscape(id)).
		Post(&key, &keyID)

	return keyID.ID, err
}

// GeneratePrincipalKey generate new principal key for existing role
func (store *RoleStore) GeneratePrincipalKey(id string) (string, error) {
	var keyID struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/principalkeys/generate", url.PathEscape(id)).
		Post(id, &keyID)

	return keyID.ID, err
}

// PrincipalKeys returns all principal keys
func (store *RoleStore) PrincipalKeys(id string) ([]PrincipalKey, error) {
	result := principalkeysResult{}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/principalkeys", url.PathEscape(id)).
		Get(&result)

	return result.Items, err
}

// UpdateRole update existing role
func (store *RoleStore) UpdateRole(id string, role *Role) error {
	_, err := store.api.
		URL("/role-store/api/v1/roles/%s", url.PathEscape(id)).
		Put(role)

	return err
}

// DeleteRole delete a role
func (store *RoleStore) DeleteRole(id string) error {
	_, err := store.api.
		URL("/role-store/api/v1/roles/%s", id).
		Delete()

	return err
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

// AWSGrantedRoles return AWS role granting PrivX roles
func (store *RoleStore) AWSGrantedRoles(id string) ([]AWSRole, error) {
	result := awsrolesResult{}

	_, err := store.api.
		URL("/role-store/api/v1/awsroles/%s/roles", url.PathEscape(id)).
		Get(&result)

	return result.Items, err
}

// DeleteAWSRole delete a aws role
func (store *RoleStore) DeleteAWSRole(id string) error {
	_, err := store.api.
		URL("/role-store/api/v1/awsroles/%s", id).
		Delete()

	return err
}

// AWSRole returns existing single aws role
func (store *RoleStore) AWSRole(id string) (role *AWSRole, err error) {
	role = new(AWSRole)

	_, err = store.api.
		URL("/role-store/api/v1/awsroles/%s", url.PathEscape(id)).
		Get(role)

	return
}

// AWSRoles returns all aws roles.
func (store *RoleStore) AWSRoles(refresh string) ([]AWSRole, error) {
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

// RefreshSources refresh all host and user sources
func (store *RoleStore) RefreshSources(id []string) error {
	_, err := store.api.
		URL("/role-store/api/v1/sources/refresh").
		Post(&id)

	return err
}

// UpdateSource update existing source
func (store *RoleStore) UpdateSource(id string, source *Source) error {
	_, err := store.api.
		URL("/role-store/api/v1/sources/%s", url.PathEscape(id)).
		Put(source)

	return err
}

// DeleteSource delete a source
func (store *RoleStore) DeleteSource(id string) error {
	_, err := store.api.
		URL("/role-store/api/v1/sources/%s", id).
		Delete()

	return err
}

// Source returns a source
func (store *RoleStore) Source(id string) (source *Source, err error) {
	source = new(Source)

	_, err = store.api.
		URL("/role-store/api/v1/sources/%s", url.PathEscape(id)).
		Get(source)

	return
}

// CreateSource create a new source
func (store *RoleStore) CreateSource(source Source) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/sources").
		Post(&source, &id)

	return id.ID, err
}

// Sources get all sources.
func (store *RoleStore) Sources() ([]Source, error) {
	result := sourcesResult{}

	_, err := store.api.
		URL("/role-store/api/v1/sources").
		Get(&result)

	return result.Items, err
}

// SearchUsers searches for users, matching the keywords and source
// criteria.
func (store *RoleStore) SearchUsers(keywords, source string) ([]User, error) {
	result := usersResult{}
	_, err := store.api.
		URL("/role-store/api/v1/users/search").
		Post(map[string]string{
			"keywords": keywords,
			"source":   source,
		}, &result)

	return result.Items, err
}

// User gets information about the argument user ID.
func (store *RoleStore) User(id string) (user *User, err error) {
	user = new(User)

	_, err = store.api.
		URL("/role-store/api/v1/users/%s", url.PathEscape(id)).
		Get(user)

	return
}

// UserRoles gets the roles of the argument user ID.
func (store *RoleStore) UserRoles(id string) ([]Role, error) {
	result := rolesResult{}
	_, err := store.api.
		URL("/role-store/api/v1/users/%s/roles", url.PathEscape(id)).
		Get(&result)

	return result.Items, err
}

// AddUserRole adds the specified role for the user. If the user
// already has the role, this function does nothing.
func (store *RoleStore) AddUserRole(userID, roleID string) error {
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

// RemoveUserRole removes the specified role from the user. If the
// user does not have the role, this function does nothing.
func (store *RoleStore) RemoveUserRole(userID, roleID string) error {
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

// Roles gets all configured roles.
func (store *RoleStore) Roles() ([]Role, error) {
	result := rolesResult{}

	_, err := store.api.
		URL("/role-store/api/v1/roles").
		Get(&result)

	return result.Items, err
}

// Role gets information about the argument role ID.
func (store *RoleStore) Role(id string) (role *Role, err error) {
	role = new(Role)

	_, err = store.api.
		URL("/role-store/api/v1/roles/%s", url.PathEscape(id)).
		Get(role)

	return
}

// GetRoleMembers gets all members (users) of the argument role ID.
func (store *RoleStore) GetRoleMembers(id string) ([]User, error) {
	result := usersResult{}

	_, err := store.api.
		URL("/role-store/api/v1/roles/%s/members", url.PathEscape(id)).
		Get(&result)

	return result.Items, err
}

// CreateRole creates new role
func (store *RoleStore) CreateRole(role Role) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/role-store/api/v1/roles").
		Post(&role, &id)

	return id.ID, err
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
