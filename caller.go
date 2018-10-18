// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"path/filepath"
	"runtime"
)

func callerInfo(skip int) (file string, line int, ok bool) {
	_, file, line, ok = runtime.Caller(skip + 1)
	if !ok {
		file = "???"
		line = 0
	}

	// short file name
	file = filepath.Base(file)
	return
}
