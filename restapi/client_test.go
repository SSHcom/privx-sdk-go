//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package restapi_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SSHcom/privx-sdk-go/restapi"
)

func TestGet(t *testing.T) {
	ts := mockStatus()
	defer ts.Close()

	_, err := restapi.New(restapi.BaseURL(ts.URL)).
		URL("/users/%v", 1).Status()

	if err != nil {
		t.Errorf("client get fails: %v", err)
	}
}

func TestGetFails(t *testing.T) {
	ts := mockStatus()
	defer ts.Close()

	_, err := restapi.New(restapi.BaseURL(ts.URL)).
		URL("/users/%v", 2).Status()

	if err == nil {
		t.Errorf("client get is not failing.")
	} else if err.Error() !=
		"error: error42, message: broken request, property: mock" {
		t.Errorf("unexpected error: %s", err)
	}
}

type T struct {
	ID string `json:"id"`
}

func TestPut(t *testing.T) {
	ts := mockStatus()
	defer ts.Close()

	eg := T{ID: "id"}
	in := T{}

	_, err := restapi.New(restapi.BaseURL(ts.URL)).
		URL("/echo").Put(eg, &in)

	if err != nil {
		t.Errorf("client fails: %v", err)
	}

	if eg.ID != in.ID {
		t.Errorf("unexpected response: %v", in)
	}
}

func TestPost(t *testing.T) {
	ts := mockStatus()
	defer ts.Close()

	eg := T{ID: "id"}
	in := T{}

	_, err := restapi.New(restapi.BaseURL(ts.URL)).
		URL("/echo").Post(eg, &in)

	if err != nil {
		t.Errorf("client fails: %v", err)
	}

	if eg.ID != in.ID {
		t.Errorf("unexpected response: %v", in)
	}
}

func TestRecvNoIdP(t *testing.T) {
	ts := mock()
	defer ts.Close()

	var data struct {
		ID string `json:"id"`
	}

	_, err := restapi.New(restapi.BaseURL(ts.URL)).
		URL("/").Get(&data)

	if err != nil {
		t.Errorf("client fails: %v", err)
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

//
func mockStatus() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch {
			case r.URL.Path == "/users/1":
				w.WriteHeader(http.StatusOK)
			case strings.HasPrefix(r.URL.Path, "/users/"):
				w.WriteHeader(http.StatusBadRequest)
				body, _ := json.Marshal(map[string]string{
					"error_code":    "error42",
					"error_message": "broken request",
					"property":      "mock",
				})
				w.Write(body)

			case r.URL.Path == "/echo":
				b, _ := io.ReadAll(r.Body)
				w.Write(b)
			}
		}),
	)
}
