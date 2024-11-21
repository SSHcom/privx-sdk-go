//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package connectionmanager

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// ConnectionManager is a connection manager client instance.
type ConnectionManager struct {
	api restapi.Connector
}

// New connection manager client constructor.
func New(api restapi.Connector) *ConnectionManager {
	return &ConnectionManager{api: api}
}

// MARK: Status
// Status get connection manager microservice status.
func (c *ConnectionManager) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/connection-manager/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Connections
// GetConnections get connections.
func (c *ConnectionManager) GetConnections(opts ...filters.Option) (*response.ResultSet[Connection], error) {
	connections := &response.ResultSet[Connection]{}
	params := url.Values{}

	// Set default options, which will be overwritten by opts if defined.
	options := append([]filters.Option{
		filters.Paging(0, 25),
		filters.Sort("connected", "DESC"),
		filters.FuzzyCount(true),
	}, opts...)

	for _, opt := range options {
		opt(&params)
	}

	_, err := c.api.
		URL("/connection-manager/api/v1/connections").
		Query(params).
		Get(&connections)

	return connections, err
}

// SearchConnections search for connections.
func (c *ConnectionManager) SearchConnections(search *ConnectionSearch, opts ...filters.Option) (*response.ResultSet[Connection], error) {
	connections := &response.ResultSet[Connection]{}
	params := url.Values{}

	// Set default options, which will be overwritten by opts if defined.
	options := append([]filters.Option{
		filters.Paging(0, 25),
		filters.Sort("connected", "DESC"),
		filters.FuzzyCount(true),
	}, opts...)

	for _, opt := range options {
		opt(&params)
	}

	_, err := c.api.
		URL("/connection-manager/api/v1/connections/search").
		Query(params).
		Post(&search, &connections)

	return connections, err
}

// GetConnection get connection by id.
func (c *ConnectionManager) GetConnection(connID string) (*Connection, error) {
	connection := &Connection{}

	_, err := c.api.
		URL("/connection-manager/api/v1/connections/%s", connID).
		Get(&connection)

	return connection, err
}

// CreateSessionForFileDownload create session id for trail stored file download.
func (c *ConnectionManager) CreateSessionForFileDownload(connID, chanID, fileID string) (DownloadSessionID, error) {
	sessionID := DownloadSessionID{}

	_, err := c.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/file/%s", connID, chanID, fileID).
		Post(nil, &sessionID)

	return sessionID, err
}

// DownloadTrailStoredFile download trail stored file transferred within audited connection channel,
func (c *ConnectionManager) DownloadTrailStoredFile(connID, chanID, fileID, sessionID, filename string) error {
	err := c.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/file/%s/%s",
			connID, chanID, fileID, sessionID).
		Download(filename)

	return err
}

// CreateSessionForTrailLogDownload create session id for trail log download.
func (c *ConnectionManager) CreateSessionForTrailLogDownload(connID, chanID string) (DownloadSessionID, error) {
	sessionID := DownloadSessionID{}

	_, err := c.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/log", connID, chanID).
		Post(nil, &sessionID)

	return sessionID, err
}

// DownloadTrailLog download trail log of audited connection channel.
func (c *ConnectionManager) DownloadTrailLog(connID, chanID, sessionID, filename string, opts ...filters.Option) error {
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	err := c.api.
		URL("/connection-manager/api/v1/connections/%s/channel/%s/log/%s", connID, chanID, sessionID).
		Query(params).
		Download(filename)

	return err
}

// GetAccessRoles get access roles for connection by id.
// Note, the v1 endpoint doesn't return the count as part of the response body,
// this will change with v2. Until then, we will handle it internally within the SDK.
func (c *ConnectionManager) GetAccessRoles(connID string) (*response.ResultSet[ConnectionPermission], error) {
	p := []ConnectionPermission{}

	_, err := c.api.
		URL("/connection-manager/api/v1/connections/%s/access_roles", connID).
		Get(&p)

	// v1 endpoint does not return count,
	// return count internally in sdk until v2 is introduced
	perms := &response.ResultSet[ConnectionPermission]{
		Count: len(p),
		Items: p,
	}

	return perms, err
}

// GrantAccessRole grant a role permission for a connection.
func (c *ConnectionManager) GrantAccessRole(connID, roleID string) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/connections/%s/access_roles/%s", connID, roleID).
		Post(nil)

	return err
}

// RevokeAccessRole revoke a permission for a role from a connection.
func (c *ConnectionManager) RevokeAccessRole(connID, roleID string) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/connections/%s/access_roles/%s", connID, roleID).
		Delete()

	return err
}

// RevokeAccessRoleFromAllConnections revoke permissions for a role from all connections.
func (c *ConnectionManager) RevokeAccessRoleFromAllConnections(roleID string) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/connections/access_roles/%s", roleID).
		Delete()

	return err
}

// GetConnectionTags get connection tags.
func (c *ConnectionManager) GetConnectionTags(opts ...filters.Option) (*response.ResultSet[string], error) {
	tags := &response.ResultSet[string]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/connection-manager/api/v1/connections/tags").
		Query(params).
		Get(&tags)

	return tags, err
}

// UpdateConnectionTags update connection tags.
func (c *ConnectionManager) UpdateConnectionTags(tags []string, connectionID string) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/connections/%s/tags", connectionID).
		Put(&tags)

	return err
}

// TerminateConnection terminate connection by id.
func (c *ConnectionManager) TerminateConnection(connID string) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/terminate/connection/%s", connID).
		Post(nil)

	return err
}

// MARK: Terminate
// TerminateConnectionsByHost terminate connections from host.
func (c *ConnectionManager) TerminateConnectionsByHost(hostID string) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/terminate/host/%s", hostID).
		Post(nil)

	return err
}

// TerminateConnectionsByUser terminate connection(s) of a user
func (c *ConnectionManager) TerminateConnectionsByUser(userID string) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/terminate/user/%s", userID).
		Post(nil)

	return err
}

// MARK: UEBA Management
// GetUebaConfigurations get ueba configurations.
func (c *ConnectionManager) GetUebaConfigurations() (*UebaConfigurations, error) {
	configurations := &UebaConfigurations{}
	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/configure").
		Get(&configurations)

	return configurations, err
}

// SetUebaConfigurations set ueba configurations.
func (c *ConnectionManager) SetUebaConfigurations(configurations *UebaConfigurations) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/configure").
		Post(&configurations)

	return err
}

// GetUebaAnomalySettings get ueba anomaly settings.
func (c *ConnectionManager) GetUebaAnomalySettings() (UebaAnomalySettings, error) {
	settings := UebaAnomalySettings{}
	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/anomaly-settings").
		Get(&settings)

	return settings, err
}

// CreateUebaAnomalySettings create Ueba anomaly settings.
func (c *ConnectionManager) CreateUebaAnomalySettings(settings UebaAnomalySettings) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/anomaly-settings").
		Post(&settings)

	return err
}

// StartUebaAnalyzing start ueba analyzing connections with a saved dataset.
func (c *ConnectionManager) StartUebaAnalyzing(datasetID string) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/start-analyzing/%s", datasetID).
		Post(nil)

	return err
}

// StopUebaAnalyzing stop ueba analyzing connection anomalies.
func (c *ConnectionManager) StopUebaAnalyzing() error {
	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/stop-analyzing").
		Post(nil)

	return err
}

// MARK: UEBA Train
// GetUebaDatasets get dataset list for ueba.
func (c *ConnectionManager) GetUebaDatasets() (*response.ResultSet[Dataset], error) {
	datasets := &response.ResultSet[Dataset]{}

	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/datasets").
		Get(&datasets)

	return datasets, err
}

// CreateUebaDataset create a new dataset.
func (c *ConnectionManager) CreateUebaDataset(dataset *Dataset) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/datasets").
		Post(&dataset, &identifier)

	return identifier, err
}

// GetUebaDataset get ueba dataset by id.
func (c *ConnectionManager) GetUebaDataset(datasetID string) (*Dataset, error) {
	dataset := &Dataset{}

	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/datasets/%s", datasetID).
		Get(&dataset)

	return dataset, err
}

// UpdateUebaDataset update ueba dataset.
func (c *ConnectionManager) UpdateUebaDataset(dataset *Dataset, datasetID string) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/datasets/%s", datasetID).
		Put(&dataset)

	return err
}

// DeleteUebaDataset delete ueba dataset.
func (c *ConnectionManager) DeleteUebaDataset(datasetID string) error {
	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/datasets/%s", datasetID).
		Delete()

	return err
}

// TrainUebaDataset train or retrain ueba dataset.
func (c *ConnectionManager) TrainUebaDataset(datasetID string, opts ...filters.Option) (ConnectionCount, error) {
	count := ConnectionCount{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/train/%s", datasetID).
		Query(params).
		Post(nil, &count)

	return count, err
}

// GetUebaConnectionCounts get number of connections for dataset.
func (c *ConnectionManager) GetUebaConnectionCounts(timeRange TimeRange) (ConnectionCount, error) {
	count := ConnectionCount{}

	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/query-connection-count").
		Post(&timeRange, &count)

	return count, err
}

// MARK: UEBA Setup
// CreateSessionForUebaScriptDownload create session id for ueba setup script download.
func (c *ConnectionManager) CreateSessionForUebaScriptDownload() (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/setup-script").
		Post(nil, &identifier)

	return identifier, err
}

// DownloadUebaScript download ueba setup script.
func (c *ConnectionManager) DownloadUebaScript(sessionID, filename string) error {
	err := c.api.
		URL("/connection-manager/api/v1/ueba/setup-script/%s", sessionID).
		Download(filename)

	return err
}

// MARK: UEBA Status
// GetUebaStatus get ueba service status.
func (c *ConnectionManager) GetUebaStatus() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/status").
		Get(&status)

	return status, err
}

// GetUebaInternalStatus get ueba internal status.
func (c *ConnectionManager) GetUebaInternalStatus() (UebaInternalStatus, error) {
	uebaInternalStatus := UebaInternalStatus{}

	_, err := c.api.
		URL("/connection-manager/api/v1/ueba/status/internal").
		Get(&uebaInternalStatus)

	return uebaInternalStatus, err
}
