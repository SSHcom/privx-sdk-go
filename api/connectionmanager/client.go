//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package connectionmanager

import (
	"fmt"
	"net/url"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// ConnectionManager is a connection manager client instance.
type ConnectionManager struct {
	api restapi.Connector
}

type connectionsResult struct {
	Count int          `json:"count"`
	Items []Connection `json:"items"`
}

// New creates a new connection manager client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *ConnectionManager {
	return &ConnectionManager{api: api}
}

// Connections get all connections
func (store *ConnectionManager) Connections(offset, limit int, sortkey, sortdir string) ([]Connection, error) {
	result := connectionsResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/connections").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// SearchConnections search for connections
func (store *ConnectionManager) SearchConnections(offset, limit int, sortdir, sortkey string, searchObject ConnectionSearch) ([]Connection, error) {
	result := connectionsResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortdir: sortdir,
		Sortkey: sortkey,
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/connections/search").
		Query(&filters).
		Post(&searchObject, &result)

	return result.Items, err
}

// Connection get a single connection
func (store *ConnectionManager) Connection(connID string) (*Connection, error) {
	conn := &Connection{}

	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s", url.PathEscape(connID)).
		Get(conn)

	return conn, err
}

// CreateSessionIDFileDownload create session ID for trail stored file download
func (store *ConnectionManager) CreateSessionIDFileDownload(connID, chanID, fileID string) (string, error) {
	var object struct {
		SessionID string `json:"session_id"`
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/file/%s",
			url.PathEscape(connID), url.PathEscape(chanID), url.PathEscape(fileID)).
		Post(nil, &object)

	return object.SessionID, err
}

// DownloadStoredFile download trail stored file transferred within audited connection channel
func (store *ConnectionManager) DownloadStoredFile(connID, chanID, fileID, sessionID, filename string) error {
	err := store.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/file/%s/%s",
			url.PathEscape(connID), url.PathEscape(chanID), url.PathEscape(fileID), url.PathEscape(sessionID)).
		Download(filename)

	return err
}

// CreateSessionIDTrailLog create session ID for trail log download
func (store *ConnectionManager) CreateSessionIDTrailLog(connID, chanID string) (string, error) {
	var object struct {
		SessionID string `json:"session_id"`
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/log",
			url.PathEscape(connID), url.PathEscape(chanID)).
		Post(nil, &object)

	return object.SessionID, err
}

// DownloadTrailLog download trail log of audited connection channel
func (store *ConnectionManager) DownloadTrailLog(connID, chanID, sessionID, format, filter, filename string) error {
	filters := Params{
		Format: format,
		Filter: filter,
	}

	err := store.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/log/%s",
			url.PathEscape(connID), url.PathEscape(chanID), url.PathEscape(sessionID)).
		Query(&filters).
		Download(filename)

	return err
}

// AccessRoles get saved access roles for a connection
func (store *ConnectionManager) AccessRoles(connID string) ([]AccessRoles, error) {
	var result []AccessRoles

	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s/access_roles", url.PathEscape(connID)).
		Get(&result)

	return result, err
}

// GrantAccessRoleToConnection grant a role permission for a connection
func (store *ConnectionManager) GrantAccessRoleToConnection(connID, roleID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s/access_roles/%s",
			url.PathEscape(connID), url.PathEscape(roleID)).
		Post(nil)

	return err
}

// RevokeAccessRoleFromConnection revoke a permission for a role from a connection
func (store *ConnectionManager) RevokeAccessRoleFromConnection(connID, roleID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/connections/%s/access_roles/%s",
			url.PathEscape(connID), url.PathEscape(roleID)).
		Delete()

	return err
}

// RevokeAccessRoleFromAllConnections revoke permissions for a role from connections
func (store *ConnectionManager) RevokeAccessRoleFromAllConnections(roleID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/connections/access_roles/%s",
			url.PathEscape(roleID)).
		Delete()

	return err
}

// TerminateConnection terminate connection by ID.
func (store *ConnectionManager) TerminateConnection(connID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/terminate/connection/%s", url.PathEscape(connID)).
		Post(nil)

	return err
}

// TerminateConnectionsByTargetHost terminate connection(s) from host
func (store *ConnectionManager) TerminateConnectionsByTargetHost(hostID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/terminate/host/%s", url.PathEscape(hostID)).
		Post(nil)

	return err
}

// TerminateConnectionsByUser terminate connection(s) of a user
func (store *ConnectionManager) TerminateConnectionsByUser(userID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/terminate/user/%s", url.PathEscape(userID)).
		Post(nil)

	return err
}

// UEBA

// UebaConfigurations get ueba configurations
func (store *ConnectionManager) UebaConfigurations() (UebaConfigurations, error) {
	configurations := UebaConfigurations{}
	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/configure").
		Get(&configurations)

	return configurations, err
}

// SetUebaConfigurations set ueba configurations
func (store *ConnectionManager) SetUebaConfigurations(configurations *UebaConfigurations) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/configure").
		Post(&configurations)

	return err
}

// UebaAnomalySettings get ueba anomaly settings
func (store *ConnectionManager) UebaAnomalySettings() (UebaAnomalySettings, error) {
	settings := UebaAnomalySettings{}
	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/anomaly-settings").
		Get(&settings)

	return settings, err
}

// CreateAnomalySettings create a host to host store
func (store *ConnectionManager) CreateAnomalySettings(settings UebaAnomalySettings) error {

	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/anomaly-settings").
		Post(&settings)

	return err
}

// StartAnalyzing start ueba analysis
func (store *ConnectionManager) StartAnalyzing(datasetID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/start-analyzing/%s", url.PathEscape(datasetID)).
		Post(nil)

	return err
}

// StopAnalyzing stop ueba analysis
func (store *ConnectionManager) StopAnalyzing() error {
	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/stop-analyzing").
		Post(nil)

	return err
}

// CreateIdForUebaScript create session ID for Ueba setup script
func (store *ConnectionManager) CreateIdForUebaScript() (IDstruct, error) {
	sessionId := IDstruct{}
	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/setup-script").
		Post(nil, &sessionId)

	return sessionId, err
}

// DownloadUebaScript download ueba setup script.
func (store *ConnectionManager) DownloadUebaScript(sessionID string) error {
	filename := fmt.Sprintf("ueba-%s-startup.sh", sessionID)
	err := store.api.
		URL("/connection-manager/api/v1/ueba/setup-script/%s", url.PathEscape(sessionID)).
		Download(filename)
	return err
}

// UebaDatasets get dataset object list for ueba.
func (store *ConnectionManager) UebaDatasets(logs bool, bin_count int) (uebaDatasetsResult, error) {
	result := uebaDatasetsResult{}
	filters := UebaDatasetQueryParams{
		Logs:     logs,
		BinCount: bin_count,
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/datasets").
		Query(&filters).
		Get(&result)

	return result, err
}

// CreateUebaDataset Save new dataset definition.
func (store *ConnectionManager) CreateUebaDataset(uebaDatasetParam DatasetBodyParam) (IDstruct, error) {
	datasetID := IDstruct{}

	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/datasets").
		Post(&uebaDatasetParam, &datasetID)

	return datasetID, err
}

// UebaDataset Get dataset by id, possibility to filter training history.
func (store *ConnectionManager) UebaDataset(logs bool, bin_count int, datasetID string) (Dataset, error) {
	result := Dataset{}
	filters := UebaDatasetQueryParams{
		Logs:     logs,
		BinCount: bin_count,
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/datasets/%s", datasetID).
		Query(&filters).
		Get(&result)

	return result, err
}

// UpdateUebaDataset Update dataset.
func (store *ConnectionManager) UpdateUebaDataset(uebaDatasetParam DatasetBodyParam, datasetID string) error {

	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/datasets/%s", datasetID).
		Put(&uebaDatasetParam)

	return err
}

// DeleteUebaDataset Delete dataset.
func (store *ConnectionManager) DeleteUebaDataset(datasetID string) error {
	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/datasets/%s", datasetID).
		Delete()

	return err
}

// TrainUebaDataset Train or retrain saved dataset.
func (store *ConnectionManager) TrainUebaDataset(datasetID string, set_active_after_training bool) (ConnectionCount, error) {
	count := ConnectionCount{}
	filters := trainingQueryParams{
		SetActiveAfterTraining: set_active_after_training,
	}

	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/train/%s", url.PathEscape(datasetID)).
		Query(&filters).
		Post(nil, &count)

	return count, err
}

// ConnectionCounts Get number of connections for dataset with given parameters.
// All connections, if json empty in body.
func (store *ConnectionManager) ConnectionCounts(timerange TimeRange) (ConnectionCount, error) {
	count := ConnectionCount{}

	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/query-connection-count").
		Post(&timerange, &count)

	return count, err
}

// UebaStatus Get Ueba service status
func (store *ConnectionManager) UebaStatus() (ServiceStatus, error) {
	uebaStatus := ServiceStatus{}

	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/status").
		Get(&uebaStatus)

	return uebaStatus, err
}

// UebaInternalStatus Get Ueba microservice internal status
func (store *ConnectionManager) UebaInternalStatus() (UebaInternalStatus, error) {
	uebaInternalStatus := UebaInternalStatus{}

	_, err := store.api.
		URL("/connection-manager/api/v1/ueba/status/internal").
		Get(&uebaInternalStatus)

	return uebaInternalStatus, err
}
