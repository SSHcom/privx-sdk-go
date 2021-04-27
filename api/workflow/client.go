//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package workflow

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

// Workflow is a workflow client instance.
type WorkflowEngine struct {
	api restapi.Connector
}

type workflowsResult struct {
	Count int        `json:"count"`
	Items []Workflow `json:"items"`
}

// New creates a new workflow client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *WorkflowEngine {
	return &WorkflowEngine{api: api}
}

// UpdateWorkflow update  a workflow
func (store *WorkflowEngine) UpdateWorkflow(workflowID string, workflow *Workflow) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows/%s", url.PathEscape(workflowID)).
		Put(workflow)

	return err
}

// DeleteWorkflow delete a workflow by ID
func (store *WorkflowEngine) DeleteWorkflow(workflowID string) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows/%s", workflowID).
		Delete()

	return err
}

// Workflow return workflow object by ID
func (store *WorkflowEngine) Workflow(workflowID string) (workflow *Workflow, err error) {
	workflow = new(Workflow)

	_, err = store.api.
		URL("/workflow-engine/api/v1/workflows/%s", url.PathEscape(workflowID)).
		Get(workflow)

	return workflow, err
}

// CreateWorkflow create a new workflow
func (store *WorkflowEngine) CreateWorkflow(workflow Workflow) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows").
		Post(workflow, &id)

	return id.ID, err
}

// Workflows get all workflows
func (store *WorkflowEngine) Workflows(offset, limit string) ([]Workflow, error) {
	result := workflowsResult{}
	filters := Params{
		Offset: offset,
		Limit:  limit,
	}

	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows").
		Query(&filters).
		Get(&result)

	return result.Items, err
}