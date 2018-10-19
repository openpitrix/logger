// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLogs(t *testing.T) {
	var logs = readLogs(`
		2018-10-19 15:55:30.50504 -INFO- hello1 openpitrix (hello.go:20)
		2018-10-19 15:55:30.50521 -INFO- hello2 openpitrix
		2018-10-19 15:55:30.50522 -INFO- hello3 openpitrix (hello.go:26)
		2018-10-19 15:55:30.50524 -INFO- hello4 openpitrix (hello.go:32)
		2018-10-19 15:55:30.50526 -INFO- hello5 openpitrix\nlogger (hello.go:33)

		2018-10-19 15:55:30.50527 -INFO-  (hello.go:34)
	`)

	assert.Equal(t, len(logs), 6)

	assert.Equal(t, logs[0].Level, "INFO")
	assert.Equal(t, logs[0].Text, "hello1 openpitrix")
	assert.Equal(t, logs[0].File, "hello.go")
	assert.Equal(t, logs[0].Line, 20)

	assert.Equal(t, logs[1].Level, "INFO")
	assert.Equal(t, logs[1].Text, "hello2 openpitrix")
	assert.Equal(t, logs[1].File, "")
	assert.Equal(t, logs[1].Line, 0)

	assert.Equal(t, logs[5].Level, "INFO")
	assert.Equal(t, logs[5].Text, "")
	assert.Equal(t, logs[5].File, "hello.go")
	assert.Equal(t, logs[5].Line, 34)
}
