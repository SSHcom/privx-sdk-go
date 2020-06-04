//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package restapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ErrorResponse contains REST endpoint error response information.
type ErrorResponse struct {
	ErrorCode    string        `json:"error_code"`
	ErrorMessage string        `json:"error_message,omitempty"`
	Property     string        `json:"property,omitempty"`
	Details      []ErrorDetail `json:"details,omitempty"`
}

// ErrorDetail contains detailed error information, linked with the
// error response.
type ErrorDetail struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message,omitempty"`
	Property     string `json:"property,omitempty"`
}

// ErrorFromResponse creates an error value from the REST API error
// response.
func ErrorFromResponse(r *http.Response, responseBody []byte) error {
	if len(responseBody) == 0 {
		return fmt.Errorf("HTTP error: %s", r.Status)
	}

	response := new(ErrorResponse)
	err := json.Unmarshal(responseBody, response)
	if err != nil {
		return fmt.Errorf("HTTP error: %s (unexpected response body: %s)",
			r.Status, err)
	}

	msg := fmt.Sprintf("error: %s", response.ErrorCode)
	if len(response.ErrorMessage) > 0 {
		msg += fmt.Sprintf(", message: %s", response.ErrorMessage)
	}
	if len(response.Property) > 0 {
		msg += fmt.Sprintf(", property: %s", response.Property)
	}
	if len(response.Details) > 0 {
		for _, detail := range response.Details {
			msg += fmt.Sprintf(", {error: %s", detail.ErrorCode)
			if len(detail.ErrorMessage) > 0 {
				msg += fmt.Sprintf(", message: %s", detail.ErrorMessage)
			}
			if len(detail.Property) > 0 {
				msg += fmt.Sprintf(", property: %s", detail.Property)
			}
			msg += "}"
		}
	}

	return errors.New(msg)
}
