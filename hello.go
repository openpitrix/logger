// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"openpitrix.io/logger"
)

func main() {
	logger.Infof(nil, "hello1 openpitrix")

	logger.HideCallstack()
	logger.Infof(nil, "hello2 openpitrix")

	logger.ShowCallstack()
	logger.Infof(nil, "hello3 openpitrix")

	{
		var buf bytes.Buffer
		logger := logger.New().SetOutput(&buf)

		logger.Infof(nil, "hello4 openpitrix")
		logger.Infof(nil, "hello5 openpitrix\nlogger")
		logger.Infof(nil, "")
		fmt.Println(buf.String())
	}

	t, err := time.Parse("2006-01-02 15:04:05.99999", "2018-10-19 15:12:48.53323 aaaa")

	fmt.Println(t, err)

	var xs string
	var line int
	fmt.Sscanf("(aa.go:12)", "%s:%d", &xs, &line)
	fmt.Println(xs, line)
}

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
		level = strings.Trim(level, "-")

		// skip level
		s = s[len(level):]

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

/*

	"2018-10-18 23:06:42.34363 -INFO- hello1 openpitrix (hello.go:17)",
	"2018-10-18 23:06:42.34379 -INFO- hello2 openpitrix",
	"2018-10-18 23:06:42.3438  -INFO- hello3 openpitrix (hello.go:23)",
	"2018-10-18 23:06:42.34381 -INFO- hello4 openpitrix (hello.go:29)",
*/
