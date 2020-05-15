//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SSHcom/privx-sdk-go/api"
)

//
type MockIdP struct {
	string
}

func (idp MockIdP) Token() (string, error) {
	return idp.string, nil
}

var idp = api.IdP(MockIdP{"trusted"})

func TestRecv(t *testing.T) {
	ts := mock()
	defer ts.Close()

	var data struct {
		ID string `json:"id"`
	}

	err := api.NewClient(idp, api.Endpoint(ts.URL)).
		Get("/").
		Recv(&data)

	if err != nil {
		t.Errorf("client fails: %w", err)
	}

	if data.ID != "trusted" {
		t.Errorf("unexpected response: %v", data)
	}
}

func TestRecvNoIdP(t *testing.T) {
	ts := mock()
	defer ts.Close()

	var data struct {
		ID string `json:"id"`
	}

	err := api.NewClient(api.Endpoint(ts.URL)).
		Get("/").
		Recv(&data)

	if err != nil {
		t.Errorf("client fails: %w", err)
	}

	if data.ID != "untrusted" {
		t.Errorf("unexpected response: %v", data)
	}
}

//
func mock() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch {
			case r.Header.Get("Authorization") == "Bearer trusted":
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`{"id": "trusted"}`))
			default:
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`{"id": "untrusted"}`))
			}
		}),
	)
}
