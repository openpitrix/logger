// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Package ctxutil provides utilities for grpc context in OpenPitrix.
package ctxutil

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	MessageIdKey = "x-message-id"
	RequestIdKey = "x-request-id"
)

type getMetadataFromContext func(ctx context.Context) (md metadata.MD, ok bool)

var getMetadataFromContextFunc = []getMetadataFromContext{
	metadata.FromOutgoingContext,
	metadata.FromIncomingContext,
}

func GetRequestId(ctx context.Context) string {
	rid := GetValue(ctx, RequestIdKey)
	if len(rid) == 0 {
		return ""
	}
	return rid[0]
}
func SetRequestId(ctx context.Context, requestId string) context.Context {
	ctx = context.WithValue(ctx, RequestIdKey, []string{requestId})
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	md[RequestIdKey] = []string{requestId}
	return metadata.NewOutgoingContext(ctx, md)
}

func GetMessageId(ctx context.Context) []string {
	return GetValue(ctx, MessageIdKey)
}
func SetMessageId(ctx context.Context, messageId ...string) context.Context {
	ctx = context.WithValue(ctx, MessageIdKey, messageId)
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	md[MessageIdKey] = messageId
	return metadata.NewOutgoingContext(ctx, md)
}

func AppendMessageId(ctx context.Context, messageId ...string) context.Context {
	m := GetMessageId(ctx)
	m = append(m, messageId...)
	return SetMessageId(ctx, m...)
}
func ClearMessageId(ctx context.Context) context.Context {
	return SetMessageId(ctx)
}

func GetValue(ctx context.Context, key string) []string {
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
