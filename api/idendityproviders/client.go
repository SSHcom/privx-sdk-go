//
// Copyright (c) 2022 SSH Communications Security Inc.
//
// All rights reserved.
//

package idendityproviders

import (
	"net/url"
	"strings"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// IdendityProviders is an idendity providers client instance.
type IdendityProviders struct {
	api restapi.Connector
}

// New creates an idendity providers client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *IdendityProviders {
	return &IdendityProviders{api: api}
}

// List all identity providers.
func (idProviders *IdendityProviders) GetAllIdendityProviders(offset, limit int) (IdentityProviderResponse, error) {
	result := IdentityProviderResponse{}

	filters := Params{
		Offset: offset,
		Limit:  limit,
	}

	_, err := idProviders.api.
		URL("/role-store/api/v1/identity-providers").
		Query(&filters).
		Get(&result)

	return result, err
}

// Create a new Identity Provider.
func (idProviders *IdendityProviders) CreateIdendityProvider(newIP IdentityProvider) (IdentityProviderCreateResponse, error) {
	result := IdentityProviderCreateResponse{}

	_, err := idProviders.api.
		URL("/role-store/api/v1/identity-providers").
		Post(newIP, &result)

	return result, err
}

// Get Identity Provider by ID.
func (idProviders *IdendityProviders) GetIdendityProviderByID(ID string) (IdentityProvider, error) {
	result := IdentityProvider{}

	_, err := idProviders.api.
		URL("/role-store/api/v1/identity-providers/%s", url.PathEscape(ID)).
		Get(&result)

	return result, err
}

// Delete Identity Provider by ID.
func (idProviders *IdendityProviders) DeleteIdendityProviderByID(ID string) error {

	_, err := idProviders.api.
		URL("/role-store/api/v1/identity-providers/%s", url.PathEscape(ID)).
		Delete()

	return err
}

// Update a Identity Provider.
func (idProviders *IdendityProviders) UpdateIdendityProvider(UpdatedIP IdentityProvider, ID string) error {

	_, err := idProviders.api.
		URL("/role-store/api/v1/identity-providers/%s", url.PathEscape(ID)).
		Put(UpdatedIP)

	return err
}

// Search Identity Providers.
func (idProviders *IdendityProviders) SearchIdendityProviders(offset, limit int, sortkey, sortdir, keywords string) (IdentityProviderResponse, error) {
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
	_, err := idProviders.api.
		URL("/role-store/api/v1/identity-providers/search").
		Query(filters).
		Post(body, &result)

	return result, err
}
