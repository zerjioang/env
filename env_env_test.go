//
// Created by zerjioang
// https://github.com/zerjioang/env
// Copyright (c) 2021. All rights reserved.
//
// SPDX-License-Identifier: GPL-3.0
//

package env // github.com/zerjioang/env

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	t.Run("noop", func(t *testing.T) {
		t.Log("boilerplate test to avoid FAIL in travis CI due to missing _test file")
	})
	t.Run("load-env", func(t *testing.T) {
		envData := loadEnv()
		assert.NotNil(t, envData)
	})
	t.Run("GetEnvironmentConfig", func(t *testing.T) {
		envData := Variables()
		assert.NotNil(t, envData)
	})
}

func ExampleVariables() {
	env := Variables()
	val := env.GetInt64("MY_INT64_VALUE", 20)
	fmt.Println("val:", val)
	// Output: 20
}