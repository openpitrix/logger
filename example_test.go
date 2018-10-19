// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger_test

import (
	"context"

	"openpitrix.io/logger"
	"openpitrix.io/logger/ctxutil"
)

func Example() {
	logger.Infof(nil, "hello openpitrix")
}

func Example_withContext() {
	ctx := context.Background()
	ctx = ctxutil.SetRequestId(ctx, "req-id-001")
	ctx = ctxutil.SetMessageId(ctx, "msg-001", "msg-002")

	logger.Infof(ctx, "hello openpitrix")
}
