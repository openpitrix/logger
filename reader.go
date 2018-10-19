// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type logMessage struct {
	Time  time.Time
	Level string
	Text  string
	File  string
	Line  int
}

func readLogs(logs string) []logMessage {
	const (
		timeLayout = "2006-01-02 15:04:05.99999"
		minLogLen  = len(timeLayout + " " + "-INFO-")
	)
	var (
		results []logMessage
	)

	for _, s := range strings.Split(logs, "\n") {
		s = strings.TrimSpace(s)
		if len(s) < minLogLen {
			continue
		}

		// 1. parse time
		when, err := time.Parse(timeLayout, strings.TrimSpace(s[:len(timeLayout)]))
		if err != nil {
			continue
		}

		// skip time
		s = s[len(timeLayout):]

		// 2. parse level
		var level string
		fmt.Sscanf(s, "%s", &level)
		if level == "" {
			continue
		}

		// skip level
		if len(s) > len(level) {
			s = s[len(level)+1:]
		} else {
			s = s[len(level):]
		}

		// fix level
		level = strings.Trim(level, "-")

		// 3. parse file:line
		filenameStartPos := strings.LastIndex(s, "(")
		lineStartPos := strings.LastIndex(s, ":")
		lineEndPos := strings.LastIndex(s, ")")

		var filename string
		var fileline int
		if filenameStartPos >= 0 && lineStartPos > (filenameStartPos+1) {
			filename = s[filenameStartPos+1 : lineStartPos]
		}
		if lineStartPos >= 0 && lineEndPos > (lineStartPos+1) {
			fileline, _ = strconv.Atoi(s[lineStartPos+1 : lineEndPos])
		}

		// 4. parse text
		var text string
		if filename != "" {
			text = strings.TrimSpace(s[:filenameStartPos])
		} else {
			text = strings.TrimSpace(s)
		}

		// OK
		results = append(results, logMessage{
			Time:  when,
			Level: strings.Trim(level, "-"),
			Text:  text,
			File:  filename,
			Line:  fileline,
		})
	}

	return results
}
