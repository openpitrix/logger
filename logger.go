// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

var (
	defaultGlobalLogLevel = InfoLevel
	defaultLoggerHelper   = New().WithDepth(1)

	replacer = strings.NewReplacer("\r", "\\r", "\n", "\\n")
)

func SetLevelByString(level string) {
	defaultGlobalLogLevel = StringToLevel(level)
	defaultLoggerHelper.SetLevelByString(level)
}

func Infof(ctx context.Context, format string, a ...interface{}) {
	defaultLoggerHelper.Infof(ctx, format, a...)
}

func Debugf(ctx context.Context, format string, a ...interface{}) {
	defaultLoggerHelper.Debugf(ctx, format, a...)
}

func Warnf(ctx context.Context, format string, a ...interface{}) {
	defaultLoggerHelper.Warnf(ctx, format, a...)
}

func Errorf(ctx context.Context, format string, a ...interface{}) {
	defaultLoggerHelper.Errorf(ctx, format, a...)
}

func Criticalf(ctx context.Context, format string, a ...interface{}) {
	defaultLoggerHelper.Criticalf(ctx, format, a...)
}

func SetOutput(output io.Writer) {
	defaultLoggerHelper.SetOutput(output)
}

type Logger struct {
	Level         Level
	output        io.Writer
	hideCallstack bool
	depth         int
}

func New() *Logger {
	return &Logger{
		Level:  defaultGlobalLogLevel,
		output: os.Stdout,
		depth:  0,
	}
}

func (p *Logger) level() Level {
	return Level(atomic.LoadUint32((*uint32)(&p.Level)))
}

func (p *Logger) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&p.Level), uint32(level))
}

func (p *Logger) SetLevelByString(level string) {
	p.SetLevel(StringToLevel(level))
}

func (p *Logger) Debugf(ctx context.Context, format string, a ...interface{}) {
	output := replacer.Replace(fmt.Sprintf(format, a...))
	p.logOutput(ctx, DebugLevel, output, p.depth+1)
}

func (p *Logger) Infof(ctx context.Context, format string, a ...interface{}) {
	output := replacer.Replace(fmt.Sprintf(format, a...))
	p.logOutput(ctx, InfoLevel, output, p.depth+1)
}

func (p *Logger) Warnf(ctx context.Context, format string, a ...interface{}) {
	output := replacer.Replace(fmt.Sprintf(format, a...))
	p.logOutput(ctx, WarnLevel, output, p.depth+1)
}

func (p *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	output := replacer.Replace(fmt.Sprintf(format, args...))
	p.logOutput(ctx, ErrorLevel, output, p.depth+1)
}

func (p *Logger) Criticalf(ctx context.Context, format string, args ...interface{}) {
	output := replacer.Replace(fmt.Sprintf(format, args...))
	p.logOutput(ctx, CriticalLevel, output, p.depth+1)
}

func (p *Logger) HideCallstack() *Logger {
	p.hideCallstack = true
	return p
}
func (p *Logger) ShowCallstack() *Logger {
	p.hideCallstack = false
	return p
}

func (p *Logger) SetOutput(w io.Writer) *Logger {
	p.output = w
	return p

}

func (p *Logger) WithDepth(depth int) *Logger {
	p.depth = depth
	return p
}

func (p *Logger) logOutput(ctx context.Context, level Level, output string, callerDepth int) {
	if p.level() < level {
		return
	}

	var (
		now = time.Now().Format("2006-01-02 15:04:05.99999")

		messageId = ctxutil_GetMessageId(ctx)
		requestId = ctxutil_GetRequestId(ctx)

		suffix string
	)

	if len(requestId) > 0 {
		messageId = append(messageId, requestId)
	}
	if len(messageId) > 0 {
		suffix = fmt.Sprintf("(%s)", strings.Join(messageId, "|"))
	}

	if p.hideCallstack {
		output = fmt.Sprintf("%-25s -%s- %s%s",
			now, strings.ToUpper(level.String()),
			output,
			suffix,
		)
	} else {
		file, line, _ := callerInfo(callerDepth + 1)

		// 2018-03-27 02:08:44.93894 -INFO- Api service start http://openpitrix-api-gateway:9100 (main.go:44)
		output = fmt.Sprintf("%-25s -%s- %s (%s:%d)%s",
			now, strings.ToUpper(level.String()),
			output, file, line,
			suffix,
		)
	}

	fmt.Fprintln(p.output, output)
}
