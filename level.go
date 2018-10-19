// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"strings"
)

type Level uint32

const (
	CriticalLevel Level = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARNING"
	case ErrorLevel:
		return "ERROR"
	case CriticalLevel:
		return "CRITICAL"
	}

	return "UNKNOWN"
}

func StringToLevel(level string) Level {
	switch strings.ToUpper(level) {
	case "CRITICAL":
		return CriticalLevel
	case "ERROR":
		return ErrorLevel
	case "WARN", "WARNING":
		return WarnLevel
	case "DEBUG":
		return DebugLevel
	case "INFO":
		return InfoLevel
	}
	return InfoLevel
}
