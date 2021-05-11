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

// Engine is a workflow client instance.
type Engine struct {
	api restapi.Connector
}

type workflowsResult struct {
	Count int        `json:"count"`
	Items []Workflow `json:"items"`
}

type requestsResult struct {
	Count int       `json:"count"`
	Items []Request `json:"items"`
}

// New creates a new workflow client instance, using the
// argument SDK API client.
func New(api restapi.Connector) *Engine {
	return &Engine{api: api}
}

// Workflows get all workflows
func (store *Engine) Workflows(offset, limit int) ([]Workflow, error) {
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

// CreateWorkflow create a new workflow
func (store *Engine) CreateWorkflow(workflow *Workflow) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows").
		Post(&workflow, &object)

	return object.ID, err
}

// Workflow return workflow object by ID
func (store *Engine) Workflow(workflowID string) (*Workflow, error) {
	workflow := &Workflow{}

	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows/%s", url.PathEscape(workflowID)).
		Get(workflow)

	return workflow, err
}

// DeleteWorkflow delete a workflow by ID
func (store *Engine) DeleteWorkflow(workflowID string) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows/%s", workflowID).
		Delete()

	return err
}

// UpdateWorkflow update  a workflow
func (store *Engine) UpdateWorkflow(workflowID string, workflow *Workflow) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows/%s", url.PathEscape(workflowID)).
		Put(workflow)

	return err
}

// Requests get the request queue for the user
func (store *Engine) Requests(offset, limit int, filter string) ([]Request, error) {
	result := requestsResult{}
	filters := Params{
		Offset: offset,
		Limit:  limit,
		Filter: filter,
	}

	_, err := store.api.
		URL("/workflow-engine/api/v1/requests").
		Query(&filters).
		Get(&result)

	return result.Items, err
}

// CreateRequest add a workflow to the request queue.
func (store *Engine) CreateRequest(request *Request) (string, error) {
	var object struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/workflow-engine/api/v1/requests").
		Post(&request, &object)

	return object.ID, err
}

// Request return a request object by ID.
func (store *Engine) Request(requestID string) (*Request, error) {
	request := &Request{}

	_, err := store.api.
		URL("/workflow-engine/api/v1/requests/%s", url.PathEscape(requestID)).
		Get(request)

	return request, err
}

// DeleteRequest delete request item by ID.
func (store *Engine) DeleteRequest(requestID string) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/requests/%s", requestID).
		Delete()

	return err
}

// MakeDecisionOnRequest update a request in queue
func (store *Engine) MakeDecisionOnRequest(requestID string, request Decision) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/requests/%s/decision", url.PathEscape(requestID)).
		Post(&request)

	return err
}

// SearchRequests search access requests
func (store *Engine) SearchRequests(
	offset, limit int, sortdir, sortkey, filter string, searchObject *Search) ([]Request, error) {
	result := requestsResult{}
	filters := Params{
		Offset:  offset,
		Limit:   limit,
		Sortkey: sortkey,
		Sortdir: sortdir,
		Filter:  filter,
	}

	_, err := store.api.
		URL("/workflow-engine/api/v1/requests/search").
		Query(&filters).
		Post(&searchObject, &result)

	return result.Items, err
}

// Settings get settings for the microservice
func (store *Engine) Settings() (*Settings, error) {
	settings := &Settings{}

	_, err := store.api.
		URL("/workflow-engine/api/v1/settings").
		Get(&settings)

	return settings, err
}

// UpdateSettings store microservice settings
func (store *Engine) UpdateSettings(settings *Settings) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/settings").
		Put(settings)

	return err
}

// TestEmailNotification test the email settings
func (store *Engine) TestEmailNotification(settings *Settings) (SMTPResponse, error) {
	var result SMTPResponse

	_, err := store.api.
		URL("/workflow-engine/api/v1/testsmtp").
		Post(&settings, &result)

	return result, err
}
