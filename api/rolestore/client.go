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

// Client is a role-store client instance.
type Client struct {
	api restapi.Connector
}

type usersResult struct {
	Count int     `json:"count"`
	Items []*User `json:"items"`
}

type rolesResult struct {
	Count int     `json:"count"`
	Items []*Role `json:"items"`
}

// New creates a new role-store client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *Client {
	return &Client{api: api}
}

// SearchUsers searches for users, matching the keywords and source
// criteria.
func (store *Client) SearchUsers(keywords, source string) ([]*User, error) {
	result := usersResult{}
	err := store.api.
		Post("/role-store/api/v1/users/search").
		Send(map[string]string{
			"keywords": keywords,
			"source":   source,
		}).
		Recv(&result)

	return result.Items, err
}

// GetUser gets information about the argument user ID.
func (store *Client) GetUser(id string) (user *User, err error) {
	err = store.api.
		Get("/role-store/api/v1/users/%s", url.PathEscape(id)).
		Recv(user)

	return
}

// GetUserRoles gets the roles of the argument user ID.
func (store *Client) GetUserRoles(id string) ([]*Role, error) {
	result := rolesResult{}
	err := store.api.
		Get("/role-store/api/v1/users/%s/roles", url.PathEscape(id)).
		Recv(&result)

	return result.Items, err
}

// AddUserRole adds the specified role for the user. If the user
// already has the role, this function does nothing.
func (store *Client) AddUserRole(userID, roleID string) error {
	// Get user's current roles.
	roles, err := store.GetUserRoles(userID)
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
	role, err := store.GetRole(roleID)
	if err != nil {
		return err
	}

	// Add an explicit role grant request.
	roles = append(roles, &Role{
		ID:       role.ID,
		Explicit: true,
	})

	return store.setUserRoles(userID, roles)
}

// RemoveUserRole removes the specified role from the user. If the
// user does not have the role, this function does nothing.
func (store *Client) RemoveUserRole(userID, roleID string) error {
	// Get user's current roles.
	roles, err := store.GetUserRoles(userID)
	if err != nil {
		return err
	}
	// Remove role from user's roles.
	var newRoles []*Role
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

func (store *Client) setUserRoles(userID string, roles []*Role) error {
	return store.api.
		Put("/role-store/api/v1/users/%s/roles", url.PathEscape(userID)).
		RecvStatus()
}

// GetRoles gets all configured roles.
func (store *Client) GetRoles() ([]*Role, error) {
	result := rolesResult{}

	err := store.api.
		Get("/role-store/api/v1/roles").
		Recv(&result)

	return result.Items, err
}

// GetRole gets information about the argument role ID.
func (store *Client) GetRole(id string) (role Role, err error) {
	err = store.api.
		Get("/role-store/api/v1/roles/%s", url.PathEscape(id)).
		Recv(&role)

	return
}

// GetRoleMembers gets all members (users) of the argument role ID.
func (store *Client) GetRoleMembers(id string) ([]*User, error) {
	result := usersResult{}

	err := store.api.
		Get("/role-store/api/v1/roles/%s/members", url.PathEscape(id)).
		Recv(&result)

	return result.Items, err
}
