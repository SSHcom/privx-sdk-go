//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

// Package pkce implements the RFC 7636: "Proof Key for Code Exchange
// by OAuth Public Clients". This implementation support only the S256
// method.
package pkce

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

const (
	MethodS256               = "S256"
	ParamCodeChallenge       = "code_challenge"
	ParamCodeChallengeMethod = "code_challenge_method"
	ParamCodeVerifier        = "code_verifier"
)

type CodeVerifier string

func NewCodeVerifier() (CodeVerifier, error) {
	var buf [32]byte

	_, err := rand.Read(buf[:])
	if err != nil {
		return "", err
	}

	return CodeVerifier(base64.RawURLEncoding.EncodeToString(buf[:])), nil
}

func (v CodeVerifier) String() string {
	return string(v)
}

func (v CodeVerifier) ChallengeS256() (string, string) {
	digest := sha256.Sum256([]byte(v))
	return base64.RawURLEncoding.EncodeToString(digest[:]), MethodS256
}

func (v CodeVerifier) Verify(challenge, method string) bool {
	if method != MethodS256 {
		return false
	}
	computed, _ := v.ChallengeS256()
	return computed == challenge
}
