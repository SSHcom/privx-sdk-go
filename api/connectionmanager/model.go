//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package connectionmanager

import "time"

// Params query params definition
type Params struct {
	Offset     int    `json:"offset,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Sortdir    string `json:"sortdir,omitempty"`
	Sortkey    string `json:"sortkey,omitempty"`
	Format     string `json:"format,omitempty"`
	Filter     string `json:"filter,omitempty"`
	FuzzyCount bool   `json:"fuzzycount,omitempty"`
	Query      string `json:"query,omitempty"`
}

// ConnectionHost connection host struct definition
type ConnectionHost struct {
	ID         string `json:"id,omitempty"`
	CommonName string `json:"common_name,omitempty"`
}

// ConnectionRole connection role struct definition
type ConnectionRole struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// UserData user data struct definition
type UserData struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"display_name,omitempty"`
}

// AccessRoles access roles struct definition
type AccessRoles struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Added string `json:"added"`
}

// Connection connection struct definition
type Connection struct {
	ID                string           `json:"id,omitempty"`
	ProxyID           string           `json:"proxy_id,omitempty"`
	Type              string           `json:"type,omitempty"`
	UserAgent         string           `json:"user_agent,omitempty"`
	TargetHostAddress string           `json:"target_host_address,omitempty"`
	TargetHostAccount string           `json:"target_host_account,omitempty"`
	RemoteAddress     string           `json:"remote_address,omitempty"`
	Connected         string           `json:"connected,omitempty"`
	Disconnected      string           `json:"disconnected,omitempty"`
	Status            string           `json:"status,omitempty"`
	LastActivity      string           `json:"last_activity,omitempty"`
	ForceDisconnect   string           `json:"force_disconnect,omitempty"`
	TerminationReason string           `json:"termination_reason,omitempty"`
	Created           string           `json:"created,omitempty"`
	Updated           string           `json:"updated,omitempty"`
	UpdatedBy         string           `json:"updated_by,omitempty"`
	TrailID           string           `json:"trail_id,omitempty"`
	IndexStatus       string           `json:"index_status,omitempty"`
	AccessGroupID     string           `json:"access_group_id,omitempty"`
	AuthMethod        []string         `json:"authentication_method,omitempty"`
	BytesIn           int              `json:"bytes_in,omitempty"`
	BytesOut          int              `json:"bytes_out,omitempty"`
	Duration          int              `json:"duration,omitempty"`
	TrailRemoved      bool             `json:"trail_removed,omitempty"`
	AuditEnabled      bool             `json:"audit_enabled,omitempty"`
	TargetHostData    ConnectionHost   `json:"target_host_data,omitempty"`
	UserData          UserData         `json:"user,omitempty"`
	UserRoles         []ConnectionRole `json:"user_roles,omitempty"`
	TargetHostRoles   []ConnectionRole `json:"target_host_roles,omitempty"`
	AccessRoles       []AccessRoles    `json:"access_roles,omitempty"`
	Tags              []string         `json:"tags,omitempty"`
}

type connectionsTagResult struct {
	Count int      `json:"count"`
	Items []string `json:"items"`
}

// TimestampSearch timestamp search struct definition
type TimestampSearch struct {
	Start string
	End   string
}

// ConnectionSearch connection search struct definition
type ConnectionSearch struct {
	ID                   []string        `json:"id,omitempty"`
	ProxyID              []string        `json:"proxy_id,omitempty"`
	Type                 []string        `json:"type,omitempty"`
	Mode                 []string        `json:"mode,omitempty"`
	UserAgent            []string        `json:"user_agent,omitempty"`
	AuthMethod           []string        `json:"authentication_method,omitempty"`
	UserID               []string        `json:"user_id,omitempty"`
	UserDisplayName      []string        `json:"user_display_name,omitempty"`
	UserRoles            []string        `json:"user_roles,omitempty"`
	TargetHost           []string        `json:"target_host_id,omitempty"`
	TargetHostCommonName []string        `json:"target_host_common_name,omitempty"`
	TargetHostAddress    []string        `json:"target_host_address,omitempty"`
	TargetHostAccount    []string        `json:"target_host_account,omitempty"`
	TargetHostRoles      []string        `json:"target_host_roles,omitempty"`
	RemoteAddress        []string        `json:"remote_address,omitempty"`
	Status               []string        `json:"status,omitempty"`
	ForceDisconnect      []string        `json:"force_disconnect,omitempty"`
	AccessRoles          []string        `json:"access_roles,omitempty"`
	KeyWords             string          `json:"keywords,omitempty"`
	HasAccessRoles       bool            `json:"has_access_roles,omitempty"`
	Connected            TimestampSearch `json:"connected,omitempty"`
	Disconnected         TimestampSearch `json:"disconnected,omitempty"`
	LastActivity         TimestampSearch `json:"last_activity,omitempty"`
	Tags                 []string        `json:"tags,omitempty"`
}

//UEBA

// UebaConfigurations uebaconfigurations struct definition
type UebaConfigurations struct {
	Address      string `json:"address"`
	TrustAnchors string `json:"trust_anchors"`
}

// UebaAnomalySettings ueba anomaly settings struct definition
type UebaAnomalySettings struct {
	Action    string  `json:"action"`
	Threshold float32 `json:"threshold"`
}

// UebaDatasetQueryParams query params definition for Ueba DataSet
type UebaDatasetQueryParams struct {
	Logs     bool `json:"logs,omitempty"`
	BinCount int  `json:"bin_count,omitempty"`
}

// TimeRange time range struct definition
type TimeRange struct {
	Start   *time.Time         `json:"start,omitempty"`
	End     *time.Time         `json:"end,omitempty"`
	Exclude []ExcludeTimeRange `json:"exclude,omitempty"`
}

type ExcludeTimeRange struct {
	Start time.Time `json:"start" validate:"required"`
	End   time.Time `json:"end" validate:"required"`
}

// Dataset dataset struct definition for Ueba
type Dataset struct {
	ID                         string               `db:"id" json:"id" validate:"omitempty,uuid"`
	LastTraining               *time.Time           `db:"last_training" json:"last_training"`
	FeatureConfigName          string               `db:"feature_config_name" json:"-"`
	IsActive                   bool                 `db:"is_active" json:"is_active"`
	UseForInferenceOnceTrained bool                 `db:"use_for_inference_once_trained" json:"use_for_inference_once_trained"`
	Quantile99                 float32              `db:"quantile_99" json:"-"`
	Quantile999                float32              `db:"quantile_999" json:"-"`
	Std                        float32              `db:"std" json:"-"`
	TimeRangeSettings          *TimeRange           `json:"time_range_settings" validate:"required"`
	DBTimeRangeSettings        string               `db:"time_range_settings" json:"-"`
	TrainingResults            []UebaTrainingResult `json:"training_results"`
	Created                    *time.Time           `db:"created" json:"created,omitempty"`
	CreatedBy                  string               `db:"created_by" json:"created_by,omitempty"`
	Updated                    *time.Time           `db:"updated" json:"updated,omitempty"`
	UpdatedBy                  string               `db:"updated_by" json:"updated_by,omitempty"`
	Comment                    string               `db:"comment" json:"comment,omitempty"`
}

// DatasetBodyParam struct definition for body params in ueba dataset api calls
type DatasetBodyParam struct {
	ID                string     `db:"id" json:"id" validate:"omitempty"`
	TimeRangeSettings *TimeRange `json:"time_range_settings" validate:"required"`
	Created           *time.Time `db:"created" json:"created,omitempty"`
	CreatedBy         string     `db:"created_by" json:"created_by,omitempty"`
	Updated           *time.Time `db:"updated" json:"updated,omitempty"`
	UpdatedBy         string     `db:"updated_by" json:"updated_by,omitempty"`
	Comment           string     `db:"comment" json:"comment,omitempty"`
}

type uebaDatasetsResult struct {
	Items []Dataset `json:"items"`
	Count int       `json:"count"`
}

// UebaTrainingResult ueba training result struct definition
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

type Histogram struct {
	Hist     []float32 `json:"hist"`
	BinEdges []float32 `json:"bin_edges"`
}

// trainingQueryParams struct definition for ueba training query params
type trainingQueryParams struct {
	SetActiveAfterTraining bool `json:"set_active_after_training"`
}

type ConnectionCount struct {
	Count int `json:"count"`
}

type IDstruct struct {
	ID string `json:"id"`
}

type UebaInternalModelInstance struct {
	ID                string `json:"id" validate:"uuid"`
	FeatureConfigName string `json:"feature_config_name"`
	Status            string `json:"status"`
	Created           string `json:"created"`
}

type UebaInternalStatus struct {
	TrainingStatus      string                      `json:"training_status"`
	InferenceStatus     string                      `json:"inference_status"`
	DatasetID           string                      `json:"dataset_id" validate:"uuid,omitempty"`
	ModelInstanceStatus []UebaInternalModelInstance `json:"model_instance_status"`
}
