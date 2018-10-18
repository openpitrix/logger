// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	messageIdKey = "x-message-id"
	requestIdKey = "x-request-id"
)

type getMetadataFromContext func(ctx context.Context) (md metadata.MD, ok bool)

var getMetadataFromContextFunc = []getMetadataFromContext{
	metadata.FromOutgoingContext,
	metadata.FromIncomingContext,
}

func ctxutil_GetValueFromContext(ctx context.Context, key string) []string {
	if ctx == nil {
		return []string{}
	}
	for _, f := range getMetadataFromContextFunc {
		md, ok := f(ctx)
		if !ok {
			continue
		}
		m, ok := md[key]
		if ok && len(m) > 0 {
			return m
		}
	}
	m, ok := ctx.Value(key).([]string)
	if ok && len(m) > 0 {
		return m
	}
	s, ok := ctx.Value(key).(string)
	if ok && len(s) > 0 {
		return []string{s}
	}
	return []string{}
}

func ctxutil_GetMessageId(ctx context.Context) []string {
	return ctxutil_GetValueFromContext(ctx, messageIdKey)
}

func ctxutil_SetMessageId(ctx context.Context, messageId []string) context.Context {
	ctx = context.WithValue(ctx, messageIdKey, messageId)
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	md[messageIdKey] = messageId
	return metadata.NewOutgoingContext(ctx, md)
}

func ctxutil_AddMessageId(ctx context.Context, messageId ...string) context.Context {
	m := ctxutil_GetMessageId(ctx)
	m = append(m, messageId...)
	return ctxutil_SetMessageId(ctx, m)
}

func ctxutil_ClearMessageId(ctx context.Context) context.Context {
	return ctxutil_SetMessageId(ctx, []string{})
}

func ctxutil_Copy(src, dst context.Context) context.Context {
	return ctxutil_SetMessageId(dst, ctxutil_GetMessageId(src))
}

func ctxutil_GetRequestId(ctx context.Context) string {
	rid := ctxutil_GetValueFromContext(ctx, requestIdKey)
	if len(rid) == 0 {
		return ""
	}
	return rid[0]
}

func ctxutil_SetRequestId(ctx context.Context, requestId string) context.Context {
	ctx = context.WithValue(ctx, requestIdKey, []string{requestId})
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	md[requestIdKey] = []string{requestId}
	return metadata.NewOutgoingContext(ctx, md)
}
