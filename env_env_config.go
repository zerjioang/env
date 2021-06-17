//
// Created by zerjioang
// https://github.com/zerjioang/env
// Copyright (c) 2021. All rights reserved.
//
// SPDX-License-Identifier: GPL-3.0
//

package env // github.com/zerjioang/env

import (
	"strconv"
)

// Create our environment based configuration
type Config map[string]interface{}

// DefaultInt sets default INT configuration value
func (c *Config) DefaultInt(key string, defaultValue int) {
	_, ok := (*c)[key]
	if !ok {
		(*c)[key] = defaultValue
	}
}

// DefaultString sets default STRING configuration value
func (c *Config) DefaultString(key string, defaultValue string) {
	_, ok := (*c)[key]
	if !ok {
		(*c)[key] = defaultValue
	}
}

// Default sets default interface{} configuration value
func (c *Config) Default(key string, defaultValue interface{}) {
	_, ok := (*c)[key]
	if !ok {
		(*c)[key] = defaultValue
	}
}

// GetString Returns a environment variable as string by its key
func (c *Config) GetString(key string, fallback string) string {
	//log.Log.Info("reading env key: ", key)
	v, ok := (*c)[key]
	if ok {
		str, ok1 := v.(string)
		if ok1 {
			return str
		}
	}
	return fallback
}

// Get Returns a environment variable as interface{} by its key
func (c *Config) Get(key string) interface{} {
	v, ok := (*c)[key]
	if ok {
		return v
	}
	return nil
}

// GetInt Returns a environment variable as int by its key
func (c *Config) GetInt(key string, fallback int) int {
	return int(c.GetInt64(key, int64(fallback)))
}

// GetInt64 Returns a environment variable as int64 by its key
func (c *Config) GetInt64(key string, fallback int64) int64 {
	v := c.Get(key)
	if v != nil {
		switch v.(type) {
		case int:
			return int64(v.(int))
		case int64:
			return v.(int64)
		case string:
			val, err := strconv.Atoi(v.(string))
			if err == nil {
				return int64(val)
			}
		}
	}
	return fallback
}
