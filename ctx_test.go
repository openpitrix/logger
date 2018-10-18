// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddMessageId(t *testing.T) {
	ctx := context.TODO()
	ctx = ctxutil_SetMessageId(ctx, []string{"1", "2", "3"})

	messageId := ctxutil_GetMessageId(ctx)
	require.Equal(t, messageId, []string{"1", "2", "3"})

	ctx = ctxutil_AddMessageId(ctx, "4")

	messageId = ctxutil_GetMessageId(ctx)
	require.Equal(t, messageId, []string{"1", "2", "3", "4"})

	ctx = ctxutil_ClearMessageId(ctx)

	messageId = ctxutil_GetMessageId(ctx)
	require.Equal(t, messageId, []string{})
}

func TestGetRequestId(t *testing.T) {
	ctx := context.TODO()
	requestId := "abcdef"
	ctx = ctxutil_SetRequestId(ctx, requestId)

	require.Equal(t, requestId, ctxutil_GetRequestId(ctx))

	ctx = context.TODO()
	requestId = "12345"
	ctx = ctxutil_SetRequestId(ctx, requestId)

	require.Equal(t, requestId, ctxutil_GetRequestId(ctx))

	ctx = context.TODO()
	requestId = "qwert"
	ctx = ctxutil_SetRequestId(ctx, requestId)

	require.Equal(t, requestId, ctxutil_GetRequestId(ctx))
}
