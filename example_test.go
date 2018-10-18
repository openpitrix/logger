// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger_test

import (
	"openpitrix.io/logger"
)

func Example() {
	logger.Infof(nil, "hello openpitrix.io/logger")
}
