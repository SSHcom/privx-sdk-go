//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package connectionmanager

import (
	"time"

	"github.com/SSHcom/privx-sdk-go/api/hoststore"
	"github.com/SSHcom/privx-sdk-go/api/networkaccessmanager"
	"github.com/SSHcom/privx-sdk-go/api/rolestore"
)

// DownloadTrailLogParams query parameter definition.
type DownloadTrailLogParams struct {
	Format string `url:"format,omitempty"`
}

// ConnectionTagsParams query parameter definition.
type ConnectionTagsParams struct {
	Query string `url:"query,omitempty"`
}

// UebaTrainingParams query parameter definition.
type UebaTrainingParams struct {
	SetActiveAfterTraining bool `url:"set_active_after_training"`
}

// Connection connection struct definition.
type Connection struct {
	ID                string                              `json:"id,omitempty"`
	ProxyID           string                              `json:"proxy_id,omitempty"`
	Type              string                              `json:"type,omitempty"`
	Mode              string                              `json:"mode,omitempty"`
	UserAgent         string                              `json:"user_agent,omitempty"`
	AuthMethod        []string                            `json:"authentication_method,omitempty"`
	User              ConnectionUser                      `json:"user,omitempty"`
	UserRoles         []ConnectionRole                    `json:"user_roles,omitempty"`
	UserData          *rolestore.User                     `json:"user_data,omitempty"`
	TargetHost        ConnectionHost                      `json:"target_host,omitempty"`
	TargetHostAddress string                              `json:"target_host_address,omitempty"`
	TargetHostAccount string                              `json:"target_host_account,omitempty"`
	TargetHostRoles   []ConnectionRole                    `json:"target_host_roles,omitempty"`
	TargetHostData    *hoststore.Host                     `json:"target_host_data,omitempty"`
	TargetNetworkData *networkaccessmanager.NetworkTarget `json:"target_network_data,omitempty"`
	RemoteAddress     string                              `json:"remote_address,omitempty"`
	Connected         string                              `json:"connected,omitempty"`
	Disconnected      string                              `json:"disconnected,omitempty"`
	Duration          int32                               `json:"duration,omitempty"`
	Status            string                              `json:"status,omitempty"`
	LastActivity      string                              `json:"last_activity,omitempty"`
	BytesIn           int64                               `json:"bytes_in,omitempty"`
	BytesOut          int64                               `json:"bytes_out,omitempty"`
	ForceDisconnect   string                              `json:"force_disconnect,omitempty"`
	TerminationReason string                              `json:"termination_reason,omitempty"`
	Created           string                              `json:"created,omitempty"`
	Updated           string                              `json:"updated,omitempty"`
	UpdatedBy         string                              `json:"updated_by,omitempty"`
	AuditEnabled      bool                                `json:"audit_enabled,omitempty"`
	TrailID           string                              `json:"trail_id,omitempty"`
	TrailRemoved      bool                                `json:"trail_removed,omitempty"`
	IndexStatus       string                              `json:"index_status,omitempty"`
	AccessGroupID     string                              `json:"access_group_id,omitempty"`
	Keywords          string                              `json:"keywords,omitempty"`
	SessionID         string                              `json:"session_id,omitempty"`
	AccessRoles       []AccessRoles                       `json:"access_roles,omitempty"`
	Tags              []string                            `json:"tags,omitempty"`
}

// AccessRoles access roles definition.
type AccessRoles struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Added time.Time `json:"added"`
}

// ConnectionHost connection host definition.
type ConnectionHost struct {
	ID         string `json:"id,omitempty"`
	CommonName string `json:"common_name,omitempty"`
}

// ConnectionUser connection user definition.
type ConnectionUser struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

// ConnectionRole connection role definition.
type ConnectionRole struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ConnectionSearch connection search request definition.
type ConnectionSearch struct {
	ID                   []string         `json:"id,omitempty"`
	ProxyID              []string         `json:"proxy_id,omitempty"`
	Type                 []string         `json:"type,omitempty"`
	Mode                 []string         `json:"mode,omitempty"`
	UserAgent            []string         `json:"user_agent,omitempty"`
	AuthMethod           []string         `json:"authentication_method,omitempty"`
	UserID               []string         `json:"user_id,omitempty"`
	UserDisplayName      []string         `json:"user_display_name,omitempty"`
	UserRoles            []string         `json:"user_roles,omitempty"`
	TargetHost           []string         `json:"target_host_id,omitempty"`
	TargetHostCommonName []string         `json:"target_host_common_name,omitempty"`
	TargetHostAddress    []string         `json:"target_host_address,omitempty"`
	TargetHostAccount    []string         `json:"target_host_account,omitempty"`
	TargetHostRoles      []string         `json:"target_host_roles,omitempty"`
	RemoteAddress        []string         `json:"remote_address,omitempty"`
	Connected            *TimestampSearch `json:"connected,omitempty"`
	Disconnected         *TimestampSearch `json:"disconnected,omitempty"`
	Status               []string         `json:"status,omitempty"`
	LastActivity         *TimestampSearch `json:"last_activity,omitempty"`
	ForceDisconnect      []string         `json:"force_disconnect,omitempty"`
	KeyWords             string           `json:"keywords,omitempty"`
	AccessRoles          []string         `json:"access_roles,omitempty"`
	HasAccessRoles       *bool            `json:"has_access_roles,omitempty"`
	SessionID            string           `json:"session_id,omitempty"`
	Tags                 []string         `json:"tags,omitempty"`
}

// TimestampSearch timestamp search request definition.
type TimestampSearch struct {
	Start string
	End   string
}

// UebaConfigurations ueba configurations definition.
type UebaConfigurations struct {
	Address          string                          `json:"address"`
	TrustAnchors     string                          `json:"trust_anchors"`
	TrustAnchorsInfo []hoststore.HostCertificateInfo `json:"trust_anchors_info,omitempty"`
}

// UebaAnomalySettings ueba anomaly settings definition.
type UebaAnomalySettings struct {
	Action    string  `json:"action"`
	Threshold float32 `json:"threshold"`
}

// Dataset ueba dataset definition.
type Dataset struct {
	ID                         string               `json:"id"`
	LastTraining               *time.Time           `json:"last_training,omitempty"`
	IsActive                   bool                 `json:"is_active"`
	UseForInferenceOnceTrained bool                 `json:"use_for_inference_once_trained"`
	TimeRangeSettings          *TimeRange           `json:"time_range_settings,omitempty"`
	TrainingResults            []UebaTrainingResult `json:"training_results"`
	Created                    *time.Time           `json:"created,omitempty"`
	CreatedBy                  string               `json:"created_by,omitempty"`
	Updated                    *time.Time           `json:"updated,omitempty"`
	UpdatedBy                  string               `json:"updated_by,omitempty"`
	Comment                    string               `json:"comment,omitempty"`
}

// TimeRange time range definition.
type TimeRange struct {
	Start   *time.Time         `json:"start,omitempty"`
	End     *time.Time         `json:"end,omitempty"`
	Exclude []ExcludeTimeRange `json:"exclude,omitempty"`
}

// ExcludeTimeRange exclude time range definition.
type ExcludeTimeRange struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// UebaTrainingResult ueba training result definition.
type UebaTrainingResult struct {
	DatasetID                  string    `json:"dataset_id"`
	Created                    time.Time `json:"created"`
	FeatureConfigName          string    `json:"feature_config_name"`
	Status                     string    `json:"status"`
	ErrorCode                  string    `json:"error_code"`
	ErrorDetails               string    `json:"error_details"`
	NumConnections             int       `json:"num_connections"`
	Mean                       float32   `json:"mean"`
	Std                        float32   `json:"std"`
	Quantile99                 float32   `json:"quantile_99"`
	Quantile999                float32   `json:"quantile_999"`
	TrainingLog                string    `json:"training_log"`
	TrainingDatasetLoss        []float32 `json:"training_dataset_loss"`
	ValidationDatasetLoss      []float32 `json:"validation_dataset_loss"`
	ValidationDatasetHistogram Histogram `json:"validation_dataset_histogram"`
}

// Histogram ueba histogram definition.
type Histogram struct {
	Hist     []float32 `json:"hist"`
	BinEdges []float32 `json:"bin_edges"`
}

// ConnectionCount ueba connection count response definition.
type ConnectionCount struct {
	Count int `json:"count"`
}

// UebaModelInstance ueba model instance definition.
type UebaModelInstance struct {
	ID                string `json:"id"`
	FeatureConfigName string `json:"feature_config_name"`
	Status            string `json:"status"`
	Created           string `json:"created"`
}

// UebaInternalStatus ueba internal status definition.
type UebaInternalStatus struct {
	TrainingStatus  string              `json:"training_status"`
	InferenceStatus string              `json:"inference_status"`
	DatasetID       string              `json:"dataset_id"`
	Instances       []UebaModelInstance `json:"instances"`
}

// DownloadSessionID download sessions id response definition.
type DownloadSessionID struct {
	SessionID string `json:"session_id"`
}

// ConnectionPermission connection access permission definition.
type ConnectionPermission struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Added time.Time `json:"added"`
}
