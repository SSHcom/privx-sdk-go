//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package workflow

// SMTPResponse smtp server test response definition.
type SMTPResponse struct {
	Status  string      `json:"status"`
	Details interface{} `json:"details"`
}

// WorkflowSettings workflow settings definition.
type WorkflowSettings struct {
	Host               string `json:"smtp_server"`
	Port               int    `json:"smtp_server_port"`
	MaxApprovers       int    `json:"request_role_max_approvers"`
	Enabled            bool   `json:"smtp_server_enabled"`
	BackendAddress     string `json:"privx_backend_address"`
	Username           string `json:"smtp_server_username"`
	Password           string `json:"smtp_server_password"`
	Protocol           string `json:"smtp_server_protocol"`
	SenderAddress      string `json:"smtp_sender_address"`
	EmailRetryAttempts int    `json:"smtp_retry_attempts"`
	InsecureVerify     bool   `json:"smtp_server_insecure_verify"`
}

// WorkflowStepApprover workflow step approver definition.
type WorkflowStepApprover struct {
	ID   string       `json:"id,omitempty"`
	Role WorkflowRole `json:"role,omitempty"`
}

// WorkflowStep workflow step definition.
type WorkflowStep struct {
	ID        string                 `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Match     string                 `json:"match,omitempty"`
	Approvers []WorkflowStepApprover `json:"approvers,omitempty"`
}

// WorkflowRole workflow role definition.
type WorkflowRole struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted,omitempty"`
}

// WorkflowUser workflow user definition.
type WorkflowUser struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Deleted     bool   `json:"deleted"`
}

// Workflow workflow definition
type Workflow struct {
	ID                        string         `json:"id"`
	Author                    string         `json:"author"`
	Created                   string         `json:"created"`
	Updated                   string         `json:"updated"`
	UpdatedBy                 string         `json:"updated_by"`
	Name                      string         `json:"name"`
	GrantTypes                []string       `json:"grant_types,omitempty"`
	MaxTimeRestrictedDuration int64          `json:"max_time_restricted_duration,omitempty"`
	MaxFloatingDuration       int64          `json:"max_floating_duration,omitempty"`
	MaxActiveRequests         int64          `json:"max_active_requests"`
	TargetRoles               []WorkflowRole `json:"target_roles,omitempty"`
	RequestorRoles            []WorkflowRole `json:"requester_roles,omitempty"`
	Action                    string         `json:"action,omitempty"`
	CanBypassRevokeWF         bool           `json:"can_bypass_revoke_workflow"`
	Comment                   string         `json:"comment,omitempty"`
	Steps                     []WorkflowStep `json:"steps,omitempty"`
}

// Decision request decision definition.
type Decision struct {
	Step     int    `json:"step"`
	Decision string `json:"decision"`
	Comment  string `json:"comment,omitempty"`
}

// RequestStepApprover request step approver definition.
type RequestStepApprover struct {
	ID           string        `json:"id"`
	Role         WorkflowRole  `json:"role"`
	Decision     string        `json:"decision"`
	User         *WorkflowUser `json:"user,omitempty"`
	DecisionTime *string       `json:"decision_time,omitempty"`
	Comment      string        `json:"comment"`
}

// RequestStep request step definition.
type RequestStep struct {
	ID        string                `json:"id"`
	Name      string                `json:"name"`
	Match     string                `json:"match"`
	Approvers []RequestStepApprover `json:"approvers"`
}

// AccessRequest access request definition.
type AccessRequest struct {
	ID                   string        `json:"id"`
	Author               string        `json:"author"`
	Created              string        `json:"created"`
	Updated              string        `json:"updated"`
	UpdatedBy            string        `json:"updated_by"`
	Name                 string        `json:"name"`
	Requester            *WorkflowUser `json:"requester,omitempty"`
	RequestedRole        *WorkflowRole `json:"requested_role,omitempty"`
	RequestJustification string        `json:"request_justification"`
	GrantType            string        `json:"grant_type,omitempty"`
	GrantStart           string        `json:"grant_start,omitempty"`
	GrantEnd             string        `json:"grant_end,omitempty"`
	FloatingLength       int64         `json:"floating_length,omitempty"`
	TargetUser           *WorkflowUser `json:"target_user,omitempty"`
	Action               string        `json:"action,omitempty"`
	Status               string        `json:"status,omitempty"`
	Comment              string        `json:"comment,omitempty"`
	Steps                []RequestStep `json:"steps,omitempty"`
	ApproverCanRevoke    bool          `json:"approver_can_revoke"`
	TargetRoleRevoked    bool          `json:"target_role_revoked"`
	TargetRoleRevokeTime *string       `json:"target_role_revocation_time,omitempty"`
	TargetRoleRevokedBy  *WorkflowUser `json:"target_role_revoked_by,omitempty"`
}

// AccessRequestSearch access request search definition.
type AccessRequestSearch struct {
	Keywords  string `json:"keywords,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Filter    string `json:"filter,omitempty"`
}
