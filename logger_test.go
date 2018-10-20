// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"openpitrix.io/logger/ctxutil"
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
	assert.Equal(t, logs[0].Line, 18)

	assert.Equal(t, logs[1].Level, "WARNING")
	assert.Equal(t, logs[1].Text, "error message")
	assert.Equal(t, logs[1].File, "logger_test.go")
	assert.Equal(t, logs[1].Line, 19)
}

func TestLogger_withContext(t *testing.T) {
	ctx := context.Background()
	ctx = ctxutil.SetRequestId(ctx, "req-id-001")

	logger := tNewLogger(t)

	logger.Infof(ctx, "hello context1")

	ctx = ctxutil.SetMessageId(ctx, "msg-001", "msg-002")
	logger.Infof(ctx, "hello context2")

	ctx = ctxutil.SetRequestId(ctx, "")
	logger.Infof(ctx, "hello context3")

	logs := logger.ReadAllMessage()
	assert.Equal(t, len(logs), 3)

	assert.Equal(t, logs[0].RequestId, "req-id-001")
	assert.Equal(t, logs[1].RequestId, "req-id-001")
	assert.Equal(t, logs[2].RequestId, "")

	assert.Equal(t, logs[0].MessageId, []string(nil))
	assert.Equal(t, logs[1].MessageId, []string{"msg-001", "msg-002"})
	assert.Equal(t, logs[2].MessageId, []string{"msg-001", "msg-002"})

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
