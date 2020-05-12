//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package api

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type Config struct {
	Endpoint    string
	Certificate *Certificate
}

type Certificate struct {
	X509 *x509.Certificate
}

func (cert *Certificate) UnmarshalText(text []byte) error {
	block, _ := pem.Decode(text)
	if block == nil {
		return fmt.Errorf("could not decode certificate PEM data")
	}
	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}
	cert.X509 = c
	return nil
}
