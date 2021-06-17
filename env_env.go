//
// Created by zerjioang
// https://github.com/zerjioang/env
// Copyright (c) 2021. All rights reserved.
//
// SPDX-License-Identifier: GPL-3.0
//

package env // github.com/zerjioang/env

import (
	"os"
	"strings"
)

var cfg *Config

func init() {
	cfg = loadEnv()
}

// loadEnv Loads server configuration from environment keys
func loadEnv() *Config {
	var config = Config{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			config[pair[0]] = pair[1]
		}
	}
	return &config
}

// Variables returns environment configuration variables
// as key-value map
func Variables() *Config {
	return cfg
}
