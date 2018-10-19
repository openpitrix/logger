// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel(t *testing.T) {
	ts := []struct {
		Level  Level
		expect string
	}{
		{CriticalLevel, "CRITICAL"},
		{ErrorLevel, "ERROR"},
		{WarnLevel, "WARNING"},
		{InfoLevel, "INFO"},
		{DebugLevel, "DEBUG"},

		{Level(100), "UNKNOWN"},
		{Level(CriticalLevel + 100), "UNKNOWN"},
	}

	for i, v := range ts {
		assert.Equal(t, v.Level.String(), v.expect, "i = %d, v = %v", i, v)
	}

	// Warn's string is not warn!
	assert.NotEqual(t, WarnLevel.String(), "warn")
}

func TestStringToLevel(t *testing.T) {
	ts := []struct {
		Level Level
		name  string
	}{
		{CriticalLevel, "critical"},
		{ErrorLevel, "error"},
		{WarnLevel, "warning"},
		{WarnLevel, "warn"},
		{InfoLevel, "info"},
		{DebugLevel, "debug"},

		{CriticalLevel, "Critical"},
		{ErrorLevel, "Error"},
		{WarnLevel, "Warning"},
		{WarnLevel, "Warn"},
		{InfoLevel, "Info"},
		{DebugLevel, "Debug"},

		{InfoLevel, "unknown"},
		{InfoLevel, "info2"},
	}

	for i, v := range ts {
		assert.Equal(t, StringToLevel(v.name), v.Level, "i = %d, v = %v", i, v)
	}
}
