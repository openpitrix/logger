// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	logger := tNewLogger(t)

	logger.Infof(nil, "hello")
	logger.Warnf(nil, "error message")

	logs := logger.ReadAllMessage()

	assert.Equal(t, len(logs), 2)

	assert.Equal(t, logs[0].Level, "INFO")
	assert.Equal(t, logs[0].Text, "hello")
	assert.Equal(t, logs[0].File, "logger_test.go")
	assert.Equal(t, logs[0].Line, 16)

	assert.Equal(t, logs[1].Level, "WARNING")
	assert.Equal(t, logs[1].Text, "error message")
	assert.Equal(t, logs[1].File, "logger_test.go")
	assert.Equal(t, logs[1].Line, 17)
}

func TestEscapeNewline(t *testing.T) {
	input := `
x
y
z
`
	output := `\nx\ny\nz\n`

	assert.Equal(t, output, escapeNewline(input))
}
