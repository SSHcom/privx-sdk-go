//
// Copyright (c) 2022 SSH Communications Security Inc.
//
// All rights reserved.
//

package networkaccessmanager

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// NetworkAccessManager is a network access manager client instance.
type NetworkAccessManager struct {
	api restapi.Connector
}

// New creates a new network access manager client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *NetworkAccessManager {
	return &NetworkAccessManager{api: api}
}

// nwtargets Get network targets
func (nam *NetworkAccessManager) GetNetworkTargets(offset, limit int, sortkey, sortdir, name, id string) (ApiNwtargetsResponse, error) {
	result := ApiNwtargetsResponse{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
		Name:    name,
		ID:      id,
	}

	_, err := nam.api.
		URL("/network-access-manager/api/v1/nwtargets").
		Query(&filters).
		Get(&result)

	return result, err
}

// nwtargets Create network target
func (nam *NetworkAccessManager) CreateNetworkTargets(network Item) (ApiNwtargetsResponsePost, error) {
	result := ApiNwtargetsResponsePost{}

	_, err := nam.api.
		URL("/network-access-manager/api/v1/nwtargets").
		Post(&network, &result)

	return result, err
}

// nwtargets Search network target
func (nam *NetworkAccessManager) SearchNetworkTargets(offset, limit int, sortkey, sortdir, filter, keywords string) (ApiNwtargetsResponse, error) {
	result := ApiNwtargetsResponse{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
		Filter:  filter,
	}
	body := KeywordsStruct{
		Keywords: keywords,
	}
	_, err := nam.api.
		URL("/network-access-manager/api/v1/nwtargets/search").
		Query(&filters).
		Post(body, &result)

	return result, err
}

// Get microservice status
func (nam *NetworkAccessManager) NetworkAccessManagerStatus() (ApiNAMstatus, error) {
	result := ApiNAMstatus{}

	_, err := nam.api.
		URL("/network-access-manager/api/v1/status").
		Get(&result)

	return result, err
}

// nwtarget Get network targets by ID
func (nam *NetworkAccessManager) GetNetworkTargetByID(NetworkTargetID string) (Item, error) {
	result := Item{}

	_, err := nam.api.
		URL("/network-access-manager/api/v1/nwtargets/%s", url.PathEscape(NetworkTargetID)).
		Get(&result)

	return result, err
}

//nwtarget Update a network target
func (nam *NetworkAccessManager) UpdateNetworkTarget(networkTarget *Item, NetworkTargetID string) error {

	_, err := nam.api.
		URL("/network-access-manager/api/v1/nwtargets/%s", url.PathEscape(NetworkTargetID)).
		Put(networkTarget)

	return err
}

// nwtarget Delete network target by ID
func (nam *NetworkAccessManager) DeleteNetworkTargetByID(NetworkTargetID string) error {

	_, err := nam.api.
		URL("/network-access-manager/api/v1/nwtargets/%s", url.PathEscape(NetworkTargetID)).
		Delete()

	return err
}

//nwtarget disable a network target
func (nam *NetworkAccessManager) DisableNetworkTargetByID(DisabledVal bool, NetworkTargetID string) error {
	dis := DisabledStruct{}
	dis.Disabled = DisabledVal

	_, err := nam.api.
		URL("/network-access-manager/api/v1/nwtargets/%s/disabled", url.PathEscape(NetworkTargetID)).
		Put(dis)

	return err
}
