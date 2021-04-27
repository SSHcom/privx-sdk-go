//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package workflow

// Params struct for pagination queries
type Params struct {
	Offset  string `json:"offset,omitempty"`
	Limit   string `json:"limit,omitempty"`
	Sortkey string `json:"sortkey,omitempty"`
	Sortdir string `json:"sortdir,omitempty"`
	Filter  string `json:"filter,omitempty"`
}

// WorkflowStepApprover workflow step approver defintion
type WorkflowStepApprover struct {
	ID   string       `json:"id,omitempty"`
	Role WorkflowRole `json:"role,omitempty"`
}

// WorkflowStep workflow step definition
type WorkflowStep struct {
	ID        string                 `json:"id,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Match     string                 `json:"match,omitempty"`
	Approvers []WorkflowStepApprover `json:"approvers,omitempty"`
}

// WorkflowRole workflow frole definition
type WorkflowRole struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

// WorkflowUser workflow user definition
type WorkflowUser struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

// Workflow workflow response definition
type Workflow struct {
	ID                   string         `json:"id,omitempty"`
	Author               string         `json:"author,omitempty"`
	Created              string         `json:"created,omitempty"`
	Updated              string         `json:"updated,omitempty"`
	UpdatedBy            string         `json:"updated_by,omitempty"`
	Name                 string         `json:"name,omitempty"`
	RequestJustification string         `json:"request_justification,omitempty"`
	GrantType            string         `json:"grant_type,omitempty"`
	GrantStart           string         `json:"grant_start,omitempty"`
	GrantEnd             string         `json:"grant_end,omitempty"`
	Action               string         `json:"action,omitempty"`
	Status               string         `json:"status,omitempty"`
	Comment              string         `json:"comment,omitempty"`
	WorkflowID           string         `json:"workflow,omitempty"`
	FloatingLength       int            `json:"floating_length,omitempty"`
	TargetRoles          []WorkflowRole `json:"target_roles,omitempty"`
	Steps                []WorkflowStep `json:"steps,omitempty"`
	TargetUser           WorkflowUser   `json:"target_user,omitempty"`
	Requester            WorkflowUser   `json:"requester,omitempty"`
	RequestedRole        WorkflowRole   `json:"requested_role,omitempty"`
}
