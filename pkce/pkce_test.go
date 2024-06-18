//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package pkce

import (
	"reflect"
	"testing"
)

func TestPKCE(t *testing.T) {
	verifier, err := NewCodeVerifier()
	if err != nil {
		t.Fatal(err)
	}

	challenge, method := verifier.ChallengeS256()
	if !verifier.Verify(challenge, method) {
		t.Error("Expected verifier.Verify to return true")
	}
}

func TestPKCEVerify(t *testing.T) {
	verifier, err := NewCodeVerifier()
	if err != nil {
		t.Fatal(err)
	}

	challenge, _ := verifier.ChallengeS256()
	if verifier.Verify(challenge, "S512") {
		t.Error("Expected verifier.Verify to return false")
	}
}

func TestString(t *testing.T) {
	verifier, err := NewCodeVerifier()
	if err != nil {
		t.Fatal(err)
	}

	result := reflect.TypeOf(verifier.String()).Kind()

	if result != reflect.String {
		t.Errorf("Expected type string, but got %v", result)
	}
}
