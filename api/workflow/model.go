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

// StepApprover workflow step approver defintion
type StepApprover struct {
	ID   string `json:"id,omitempty"`
	Role Role   `json:"role,omitempty"`
}

// Step workflow step definition
type Step struct {
	ID        string         `json:"id,omitempty"`
	Name      string         `json:"name,omitempty"`
	Match     string         `json:"match,omitempty"`
	Approvers []StepApprover `json:"approvers,omitempty"`
}

// Role workflow frole definition
type Role struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

// User workflow user definition
type User struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

// Workflow workflow response definition
type Workflow struct {
	ID                   string `json:"id,omitempty"`
	Author               string `json:"author,omitempty"`
	Created              string `json:"created,omitempty"`
	Updated              string `json:"updated,omitempty"`
	UpdatedBy            string `json:"updated_by,omitempty"`
	Name                 string `json:"name,omitempty"`
	RequestJustification string `json:"request_justification,omitempty"`
	GrantType            string `json:"grant_type,omitempty"`
	GrantStart           string `json:"grant_start,omitempty"`
	GrantEnd             string `json:"grant_end,omitempty"`
	Action               string `json:"action,omitempty"`
	Status               string `json:"status,omitempty"`
	Comment              string `json:"comment,omitempty"`
	WorkflowID           string `json:"workflow,omitempty"`
	FloatingLength       int    `json:"floating_length,omitempty"`
	TargetRoles          []Role `json:"target_roles,omitempty"`
	Steps                []Step `json:"steps,omitempty"`
	TargetUser           User   `json:"target_user,omitempty"`
	Requester            User   `json:"requester,omitempty"`
	RequestedRole        Role   `json:"requested_role,omitempty"`
}

// Search request search definition
type Search struct {
	Keywords  string `json:"keywords,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

// Decision request decision definition
type Decision struct {
	Step     int    `json:"step"`
	Decision string `json:"decision"`
	Comment  string `json:"comment,omitempty"`
}

// RequestStepApprover request step approver definition
type RequestStepApprover struct {
	ID           string `json:"id,omitempty"`
	Decision     string `json:"decision,omitempty"`
	DecisionTime string `json:"decision_time,omitempty"`
	Comment      string `json:"comment,omitempty"`
	User         User   `json:"user,omitempty"`
	Role         Role   `json:"role,omitempty"`
}

// RequestStep request step definition
type RequestStep struct {
	ID        string                `json:"id,omitempty"`
	Name      string                `json:"name,omitempty"`
	Match     string                `json:"match,omitempty"`
	Approvers []RequestStepApprover `json:"approvers,omitempty"`
}

// Request request response definition
type Request struct {
	ID                   string        `json:"id,omitempty"`
	Author               string        `json:"author,omitempty"`
	Created              string        `json:"created,omitempty"`
	Updated              string        `json:"updated,omitempty"`
	UpdatedBy            string        `json:"updated_by,omitempty"`
	Name                 string        `json:"name,omitempty"`
	RequestJustification string        `json:"request_justification,omitempty"`
	GrantType            string        `json:"grant_type,omitempty"`
	GrantStart           string        `json:"grant_start,omitempty"`
	GrantEnd             string        `json:"grant_end,omitempty"`
	Action               string        `json:"action,omitempty"`
	Status               string        `json:"status,omitempty"`
	Comment              string        `json:"comment,omitempty"`
	WorkflowID           string        `json:"workflow,omitempty"`
	FloatingLength       int           `json:"floating_length,omitempty"`
	TargetRoles          []string      `json:"target_roles,omitempty"`
	Steps                []RequestStep `json:"steps,omitempty"`
	TargetUser           User          `json:"target_user,omitempty"`
	Requester            User          `json:"requester,omitempty"`
	RequestedRole        Role          `json:"requested_role,omitempty"`
}
