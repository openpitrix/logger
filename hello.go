// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// +build ignore

package main

import (
	"openpitrix.io/logger"
)

func main() {
	logger.Infof(nil, "hello1 openpitrix.io/logger")

	logger.HideCallstack()
	logger.Infof(nil, "hello2 openpitrix.io/logger")

	logger.ShowCallstack()
	logger.Infof(nil, "hello3 openpitrix.io/logger")
}
