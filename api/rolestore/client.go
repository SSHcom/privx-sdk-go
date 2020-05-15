//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api"
)

// Client is a role-store client instance.
type Client struct {
	api *api.Client
}

type usersResult struct {
	Count int     `json:"count"`
	Items []*User `json:"items"`
}

type rolesResult struct {
	Count int     `json:"count"`
	Items []*Role `json:"items"`
}

// NewClient creates a new role-store client instance, using the
// argument SDK API client.
func NewClient(api *api.Client) (*Client, error) {
	return &Client{
		api: api,
	}, nil
}

// SearchUsers searches for users, matching the keywords and source
// criteria.
func (store *Client) SearchUsers(keywords, source string) ([]*User, error) {
	url := fmt.Sprintf("%s/role-store/api/v1/users/search",
		store.api.Endpoint())

	body, err := json.Marshal(map[string]string{
		"keywords": keywords,
		"source":   source,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := store.api.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	result := &usersResult{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, fmt.Errorf("malformed response: %s", err)
	}

	return result.Items, nil
}

// GetUser gets information about the argument user ID.
func (store *Client) GetUser(id string) (*User, error) {
	url := fmt.Sprintf("%s/role-store/api/v1/users/%s",
		store.api.Endpoint(), url.PathEscape(id))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := store.api.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	result := new(User)
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetUserRoles gets the roles of the argument user ID.
func (store *Client) GetUserRoles(id string) ([]*Role, error) {
	url := fmt.Sprintf("%s/role-store/api/v1/users/%s/roles",
		store.api.Endpoint(), url.PathEscape(id))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := store.api.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	result := &rolesResult{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

// AddUserRole adds the specified role for the user. If the user
// already has the role, this funtion does nothing.
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
// user does not have the role, this funtion does nothing.
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
	url := fmt.Sprintf("%s/role-store/api/v1/users/%s/roles",
		store.api.Endpoint(), url.PathEscape(userID))

	body, err := json.Marshal(roles)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := store.api.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}

	return nil
}

// GetRoles gets all configured roles.
func (store *Client) GetRoles() ([]*Role, error) {
	url := fmt.Sprintf("%s/role-store/api/v1/roles", store.api.Endpoint())
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := store.api.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	result := &rolesResult{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}

// GetRole gets information about the argument role ID.
func (store *Client) GetRole(id string) (*Role, error) {
	url := fmt.Sprintf("%s/role-store/api/v1/roles/%s",
		store.api.Endpoint(), url.PathEscape(id))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := store.api.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, api.ErrorFromResponse(resp, body)
	}

	result := new(Role)
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetRoleMembers gets all members (users) of the argument role ID.
func (store *Client) GetRoleMembers(id string) ([]*User, error) {
	url := fmt.Sprintf("%s/role-store/api/v1/roles/%s/members",
		store.api.Endpoint(), url.PathEscape(id))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := store.api.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, api.ErrorFromResponse(resp, body)
	}

	result := &usersResult{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, fmt.Errorf("malformed response: %s", err)
	}

	return result.Items, nil
}
