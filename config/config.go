//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/SSHcom/privx-sdk-go/api"
	"github.com/SSHcom/privx-sdk-go/oauth"
)

const (
	etcDir         = "/opt/etc/privx"
	configFileName = "privx-sdk.toml"
)

type Config struct {
	API  api.Config
	Auth oauth.Config
}

func Default() string {
	var defaultConfig string

	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("failed to get user's home directory: %s", err)
		defaultConfig = path.Join(etcDir, configFileName)
		log.Printf("fallback to '%s'", defaultConfig)
	} else {
		defaultConfig = path.Join(home, fmt.Sprintf(".%s", configFileName))
	}

	return defaultConfig
}

func Read(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// First, read configuration file.
	config := new(Config)
	err = toml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	// Next, apply environment variables.
	val, ok := os.LookupEnv("OAUTH_CLIENT_ID")
	if ok {
		config.Auth.ClientID = val
	}
	val, ok = os.LookupEnv("OAUTH_CLIENT_SECRET")
	if ok {
		config.Auth.ClientSecret = val
	}
	val, ok = os.LookupEnv("API_CLIENT_ID")
	if ok {
		config.Auth.APIClientID = val
	}
	val, ok = os.LookupEnv("API_CLIENT_SECRET")
	if ok {
		config.Auth.APIClientSecret = val
	}

	return config, err
}
