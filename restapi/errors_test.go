//
// Copyright (c) 2021 SSH Communications Security Inc.
//
// All rights reserved.
//

package restapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmptyResponseBody(t *testing.T) {
	var emptyRespBody []byte

	resp, _ := mockResponse()
	result := ErrorFromResponse(resp, emptyRespBody)

	expectedError := fmt.Errorf("HTTP error: %s", resp.Status)
	if result.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, result)
	}
}

func TestUnexpectedResponseBody(t *testing.T) {
	resp, body := mockResponse()
	result := ErrorFromResponse(resp, body)

	expectedError := fmt.Errorf("HTTP error: 200 OK (unexpected response body: invalid character '<' looking for beginning of value)")
	if result.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, result)
	}
}

func TestDetailsErrorMessage(t *testing.T) {
	errorDetail := []ErrorDetail{{
		ErrorCode:    "42",
		ErrorMessage: "DtlTest",
		Property:     "Detail",
	}}

	body, _ := json.Marshal(ErrorResponse{
		ErrorCode:    "42",
		ErrorMessage: "ErrRspTest",
		Property:     "ErrRsp",
		Details:      errorDetail,
	})

	resp, _ := mockResponse()
	result := ErrorFromResponse(resp, body)

	expectedError := fmt.Errorf("error: 42, message: ErrRspTest, property: ErrRsp, {error: 42, message: DtlTest, property: Detail}")
	if result.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, result)
	}
}

func mockResponse() (*http.Response, []byte) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Test Body!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	return resp, body
}
