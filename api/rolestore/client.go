//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api"
)

type Client struct {
	client api.Connector
}

type usersResult struct {
	Count int     `json:"count"`
	Items []*User `json:"items"`
}

type rolesResult struct {
	Count int     `json:"count"`
	Items []*Role `json:"items"`
}

func NewClient(client api.Connector) *Client {
	return &Client{client: client}
}

func (store *Client) SearchUsers(keywords, source string) ([]*User, error) {
	result := usersResult{}
	err := store.client.
		Get("/role-store/api/v1/users/search").
		Send(map[string]string{
			"keywords": keywords,
			"source":   source,
		}).
		Recv(&result)

	return result.Items, err
}

func (store *Client) GetUser(id string) (user *User, err error) {
	err = store.client.
		Get("/role-store/api/v1/users/%s", url.PathEscape(id)).
		Recv(user)

	return
}

func (store *Client) GetUserRoles(id string) ([]*Role, error) {
	result := rolesResult{}
	err := store.client.
		Get("/role-store/api/v1/users/%s/roles", url.PathEscape(id)).
		Recv(&result)

	return result.Items, err
}

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
	return store.client.
		Put("/role-store/api/v1/users/%s/roles", url.PathEscape(userID)).
		RecvStatus()
}

func (store *Client) GetRoles() ([]*Role, error) {
	result := rolesResult{}

	err := store.client.
		Get("/role-store/api/v1/roles").
		Recv(&result)

	return result.Items, err
}

func (store *Client) GetRole(id string) (role Role, err error) {
	err = store.client.
		Get("/role-store/api/v1/roles/%s", url.PathEscape(id)).
		Recv(&role)

	return
}

func (store *Client) GetRoleMembers(id string) ([]*User, error) {
	result := usersResult{}

	err := store.client.
		Get("/role-store/api/v1/roles/%s/members", url.PathEscape(id)).
		Recv(&result)

	return result.Items, err
}
