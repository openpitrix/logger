// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// TODO: parse file/line/msg/... from log

package logger

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readBuf(buf *bytes.Buffer) string {
	str := buf.String()
	buf.Reset()
	return str
}

var ctx = context.TODO()

func TestLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	SetOutput(buf)

	Debugf(ctx, "debug log, should ignore by default")
	assert.Empty(t, readBuf(buf))

	Infof(ctx, "info log, should visable")
	assert.Contains(t, readBuf(buf), "info log, should visable")

	Infof(ctx, "format [%d]", 111)
	log := readBuf(buf)
	assert.Contains(t, log, "format [111]")
	t.Log(log)

	Infof(ctxutil_SetMessageId(ctx, []string{"xxxxx", "yyyyy"}), "format [%d]", 111)
	log = readBuf(buf)
	assert.Contains(t, log, "format [111]")
	t.Log(log)

	SetLevelByString("debug")
	Debugf(ctx, "debug log, now it becomes visible")
	assert.Contains(t, readBuf(buf), "debug log, now it becomes visible")

	defaultLoggerHelper = New()
	defaultLoggerHelper.SetOutput(buf)

	defaultLoggerHelper.HideCallstack()
	defaultLoggerHelper.Warnf(nil, "log_content")
	log = readBuf(buf)
	assert.Regexp(t, " -WARNING- log_content", log)
	t.Log(log)
}

func TestReplacer(t *testing.T) {
	input := `
x
y
z
`
	output := `\nx\ny\nz\n`
	assert.Equal(t, output, replacer.Replace(input))
}
