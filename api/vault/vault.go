//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package vault

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api"
)

// Bag contains secret data. The secret data is a JSON object and it
// can have nested values and objects.
type Bag map[string]interface{}

// Client is a Vault client instance.
type Client struct {
	api *api.Client
}

// NewClient creates a new Vault client instance, using the argument
// SDK API client.
func NewClient(api *api.Client) (*Client, error) {
	return &Client{
		api: api,
	}, nil
}

// Get gets the content of the argument secret.
func (vault *Client) Get(name string) (Bag, error) {
	secretURL := fmt.Sprintf("%s/vault/api/v1/secrets/%s",
		vault.api.Endpoint(),
		url.PathEscape(name))
	req, err := http.NewRequest(http.MethodGet, secretURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := vault.api.Do(req)
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

	bag := new(Bag)
	err = json.Unmarshal(body, bag)
	if err != nil {
		return nil, err
	}

	return *bag, nil
}
