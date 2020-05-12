//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package pkce

import (
	"testing"
)

func TestPKCE(t *testing.T) {
	verifier, err := NewCodeVerifier()
	if err != nil {
		t.Fatal(err)
	}

	challenge, method := verifier.ChallengeS256()

	if !verifier.Verify(challenge, method) {
		t.Fatalf("Challenge verification failed")
	}
}
