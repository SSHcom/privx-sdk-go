//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package pkce

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPKCE(t *testing.T) {
	verifier, err := NewCodeVerifier()
	if err != nil {
		t.Fatal(err)
	}

	challenge, method := verifier.ChallengeS256()
	assert.True(t, verifier.Verify(challenge, method), "Should return true")
}

func TestPKCEVerify(t *testing.T) {
	verifier, err := NewCodeVerifier()
	if err != nil {
		t.Fatal(err)
	}

	challenge, _ := verifier.ChallengeS256()
	assert.False(t, verifier.Verify(challenge, "S512"), "Should return false")
}

func TestString(t *testing.T) {
	verifier, err := NewCodeVerifier()
	if err != nil {
		t.Fatal(err)
	}

	result := reflect.TypeOf(verifier.String()).Kind()

	assert.EqualValues(t, reflect.String, result, "Expected type string")
}
