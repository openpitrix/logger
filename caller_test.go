// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallerInfo(t *testing.T) {
	file, line, ok := callerInfo(0)
	assert.Equal(t, file, "caller_test.go")
	assert.Equal(t, line, 14)
	assert.True(t, ok)
}
