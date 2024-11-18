//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package workflow

import (
	"net/url"

	"github.com/SSHcom/privx-sdk-go/api/filters"
	"github.com/SSHcom/privx-sdk-go/api/response"
	"github.com/SSHcom/privx-sdk-go/restapi"
)

// WorkflowEngine is a workflow client instance.
type WorkflowEngine struct {
	api restapi.Connector
}

// New workflow engine client constructor.
func New(api restapi.Connector) *WorkflowEngine {
	return &WorkflowEngine{api: api}
}

// MARK: Status
// Status get workflow engine microservice status.
func (c *WorkflowEngine) Status() (*response.ServiceStatus, error) {
	status := &response.ServiceStatus{}

	_, err := c.api.
		URL("/workflow-engine/api/v1/status").
		Get(status)

	return status, err
}

// MARK: Workflows
// GetWorkflows get workflows.
func (c *WorkflowEngine) GetWorkflows(opts ...filters.Option) (*response.ResultSet[Workflow], error) {
	workflows := &response.ResultSet[Workflow]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/workflow-engine/api/v1/workflows").
		Query(params).
		Get(&workflows)

	return workflows, err
}

// CreateWorkflow create workflow.
func (c *WorkflowEngine) CreateWorkflow(workflow *Workflow) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/workflow-engine/api/v1/workflows").
		Post(&workflow, &identifier)

	return identifier, err
}

// GetWorkflow get workflow by id.
func (c *WorkflowEngine) GetWorkflow(workflowID string) (*Workflow, error) {
	workflow := &Workflow{}

	_, err := c.api.
		URL("/workflow-engine/api/v1/workflows/%s", workflowID).
		Get(&workflow)

	return workflow, err
}

// UpdateWorkflow update workflow.
func (c *WorkflowEngine) UpdateWorkflow(workflowID string, workflow *Workflow) error {
	_, err := c.api.
		URL("/workflow-engine/api/v1/workflows/%s", workflowID).
		Put(&workflow)

	return err
}

// DeleteWorkflow delete a workflow.
func (c *WorkflowEngine) DeleteWorkflow(workflowID string) error {
	_, err := c.api.
		URL("/workflow-engine/api/v1/workflows/%s", workflowID).
		Delete()

	return err
}

// MARK: Requests
// GetRequests get the request queue for the user
func (c *WorkflowEngine) GetRequests(opts ...filters.Option) (*response.ResultSet[AccessRequest], error) {
	requests := &response.ResultSet[AccessRequest]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/workflow-engine/api/v1/requests").
		Query(params).
		Get(&requests)

	return requests, err
}

// CreateRequest create request.
func (c *WorkflowEngine) CreateRequest(request *AccessRequest) (response.Identifier, error) {
	identifier := response.Identifier{}

	_, err := c.api.
		URL("/workflow-engine/api/v1/requests").
		Post(&request, &identifier)

	return identifier, err
}

// GetRequest get request by id.
func (c *WorkflowEngine) GetRequest(requestID string) (*AccessRequest, error) {
	request := &AccessRequest{}

	_, err := c.api.
		URL("/workflow-engine/api/v1/requests/%s", requestID).
		Get(request)

	return request, err
}

// DeleteRequest delete request.
func (c *WorkflowEngine) DeleteRequest(requestID string) error {
	_, err := c.api.
		URL("/workflow-engine/api/v1/requests/%s", requestID).
		Delete()

	return err
}

// UpdateDecisionOnRequest update a request decision in queue.
func (c *WorkflowEngine) UpdateDecisionOnRequest(requestID string, request Decision) error {
	_, err := c.api.
		URL("/workflow-engine/api/v1/requests/%s/decision", requestID).
		Post(&request)

	return err
}

// RevokeTargetRole revoke target role in request from target user.
func (c *WorkflowEngine) RevokeTargetRole(requestID string) error {
	_, err := c.api.
		URL("/workflow-engine/api/v1/requests/%s/role/revoke", requestID).
		Post(nil)

	return err
}

// SearchRequests search access requests
func (c *WorkflowEngine) SearchRequests(search *AccessRequestSearch, opts ...filters.Option) (*response.ResultSet[AccessRequest], error) {
	requests := &response.ResultSet[AccessRequest]{}
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}

	_, err := c.api.
		URL("/workflow-engine/api/v1/requests/search").
		Query(params).
		Post(&search, &requests)

	return requests, err
}

// MARK: Settings
// GetSettings get settings for workflow engine.
func (c *WorkflowEngine) GetSettings() (*WorkflowSettings, error) {
	settings := &WorkflowSettings{}

	_, err := c.api.
		URL("/workflow-engine/api/v1/settings").
		Get(&settings)

	return settings, err
}

// UpdateSettings update settings for workflow engine.
func (c *WorkflowEngine) UpdateSettings(settings *WorkflowSettings) error {
	_, err := c.api.
		URL("/workflow-engine/api/v1/settings").
		Put(&settings)

	return err
}

// MARK: Test SMTP
// TestSMTP test SMTP settings.
func (c *WorkflowEngine) TestSMTP(settings *WorkflowSettings) (SMTPResponse, error) {
	testResponse := SMTPResponse{}

	_, err := c.api.
		URL("/workflow-engine/api/v1/testsmtp").
		Post(&settings, &testResponse)

	return testResponse, err
}
