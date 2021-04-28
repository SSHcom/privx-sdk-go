//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package workflow

import (
	"encoding/json"
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

// SetSettings store microservice settings
func (store *Engine) SetSettings(settings *json.RawMessage) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/settings").
		Put(settings)

	return err
}

// Settings get settings for the microservice
func (store *Engine) Settings() (settings *json.RawMessage, err error) {
	_, err = store.api.
		URL("/workflow-engine/api/v1/settings").
		Get(&settings)

	return
}

// SearchRequests search access requests
func (store *Engine) SearchRequests(
	offset, limit, sortdir, sortkey, filter string, searchObject *Search) ([]Request, error) {
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

// UserRequestDecision update a request in queue
func (store *Engine) UserRequestDecision(requestID string, request Decision) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/requests/%s/decision", url.PathEscape(requestID)).
		Post(&request)

	return err
}

// DeleteRequest delete request item by ID.
func (store *Engine) DeleteRequest(requestID string) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/requests/%s", requestID).
		Delete()

	return err
}

// Request return a request object by ID.
func (store *Engine) Request(requestID string) (request *Request, err error) {
	request = new(Request)

	_, err = store.api.
		URL("/workflow-engine/api/v1/requests/%s", url.PathEscape(requestID)).
		Get(request)

	return request, err
}

// CreateRequest add a workflow to the request queue.
func (store *Engine) CreateRequest(request *Request) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/workflow-engine/api/v1/requests").
		Post(&request, &id)

	return id.ID, err
}

// Requests get the request queue for the user
func (store *Engine) Requests(offset, limit, filter string) ([]Request, error) {
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

// UpdateWorkflow update  a workflow
func (store *Engine) UpdateWorkflow(workflowID string, workflow *Workflow) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows/%s", url.PathEscape(workflowID)).
		Put(workflow)

	return err
}

// DeleteWorkflow delete a workflow by ID
func (store *Engine) DeleteWorkflow(workflowID string) error {
	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows/%s", workflowID).
		Delete()

	return err
}

// Workflow return workflow object by ID
func (store *Engine) Workflow(workflowID string) (workflow *Workflow, err error) {
	workflow = new(Workflow)

	_, err = store.api.
		URL("/workflow-engine/api/v1/workflows/%s", url.PathEscape(workflowID)).
		Get(workflow)

	return workflow, err
}

// CreateWorkflow create a new workflow
func (store *Engine) CreateWorkflow(workflow *Workflow) (string, error) {
	var id struct {
		ID string `json:"id"`
	}

	_, err := store.api.
		URL("/workflow-engine/api/v1/workflows").
		Post(&workflow, &id)

	return id.ID, err
}

// Workflows get all workflows
func (store *Engine) Workflows(offset, limit string) ([]Workflow, error) {
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
