//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package rolestore

// These functions are defined to maintain backward compatibility with versions before 1.35.0.

// Deprecated: See CreateIdentityProvider.
func (store *RoleStore) CreateIdendityProvider(newIP IdentityProvider) (IdentityProviderCreateResponse, error) {
	return store.CreateIdentityProvider(newIP)
}

// Deprecated: See DeleteIdentityProviderByID.
func (store *RoleStore) DeleteIdendityProviderByID(ID string) error {
	return store.DeleteIdentityProviderByID(ID)
}

// Deprecated: See GetAllIdentityProviders.
func (store *RoleStore) GetAllIdendityProviders(offset, limit int) (IdentityProviderResponse, error) {
	return store.GetAllIdentityProviders(offset, limit)
}

// Deprecated: See GetIdentityProviderByID.
func (store *RoleStore) GetIdendityProviderByID(ID string) (IdentityProvider, error) {
	return store.GetIdentityProviderByID(ID)
}

// Deprecated: See SearchIdentityProviders.
func (store *RoleStore) SearchIdendityProviders(offset, limit int, sortkey, sortdir, keywords string) (IdentityProviderResponse, error) {
	return store.SearchIdentityProviders(offset, limit, sortkey, sortdir, keywords)
}

// Deprecated: See UpdateIdentityProvider.
func (store *RoleStore) UpdateIdendityProvider(UpdatedIP IdentityProvider, ID string) error {
	return store.UpdateIdentityProvider(UpdatedIP, ID)
}
