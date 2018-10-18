# logger for OpenPitrix

[![Build Status](https://travis-ci.org/openpitrix/logger.svg)](https://travis-ci.org/openpitrix/logger)
[![Go Report Card](https://goreportcard.com/badge/openpitrix.io/logger)](https://goreportcard.com/report/openpitrix.io/logger)
[![GoDoc](https://godoc.org/openpitrix.io/logger?status.svg)](https://godoc.org/openpitrix.io/logger)
[![License](http://img.shields.io/badge/license-apache%20v2-blue.svg)](https://github.com/openpitrix/logger/blob/master/LICENSE)

## Install

1. Install Go1.11+
1. set `GO111MODULE=on`
1. `go get openpitrix.io/logger`
1. `go run hello.go`

## Example

```go
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
```

output:

```
2018-10-18 18:11:27.808   -INFO- hello1 openpitrix.io/logger (hello.go:14)
2018-10-18 18:11:27.80815 -INFO- hello2 openpitrix.io/logger
2018-10-18 18:11:27.80816 -INFO- hello3 openpitrix.io/logger (hello.go:20)
```
