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

func TestIsTrue(t *testing.T) {
	t.Run("case-yes", func(t *testing.T) {
		assert.True(t, IsTrue("YES"))
		assert.True(t, IsTrue("Yes"))
		assert.True(t, IsTrue("yes"))
		assert.True(t, IsTrue("on"))
		assert.True(t, IsTrue("ACTIVE"))
		assert.True(t, IsTrue("Active"))
		assert.True(t, IsTrue("active"))
		assert.True(t, IsTrue("TRUE"))
		assert.True(t, IsTrue("True"))
		assert.True(t, IsTrue("true"))
		assert.True(t, IsTrue("1"))
		assert.True(t, IsTrue("ENABLED"))
		assert.True(t, IsTrue("Enabled"))
		assert.True(t, IsTrue("enabled"))
	})
}

func ExampleVariables() {
	env := Variables()
	val := env.GetInt64("MY_INT64_VALUE", 20)
	fmt.Println("val:", val)
	// Output: val: 20
}
