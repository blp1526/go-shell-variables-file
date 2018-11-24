[![Build Status](https://travis-ci.org/blp1526/go-shell-variables-file.svg?branch=travis)](https://travis-ci.org/blp1526/go-shell-variables-file)
[![Go Report Card](https://goreportcard.com/badge/github.com/blp1526/go-shell-variables-file)](https://goreportcard.com/report/github.com/blp1526/go-shell-variables-file)
[![GoDoc](https://godoc.org/github.com/blp1526/go-shell-variables-file?status.svg)](https://godoc.org/github.com/blp1526/go-shell-variables-file)

# go-shell-variables-file

## Usage

### As A Package

```go
package main

import (
	"fmt"

	"github.com/blp1526/go-shell-variables-file/pkg/svf"
)

func main() {
	path := "/etc/os-release"
	s, _ := svf.ReadFile(path)
	key := "VERSION"
	value, _ := s.GetValue(key)
	fmt.Printf("path: %s, key: %s, value %s\n", path, key, value)
}
```

then, `go run main.go`

```
path: /etc/os-release, key: VERSION, value 18.04.1 LTS (Bionic Beaver)
```

### As A CLI

```
$ svf values /etc/os-release --key VERSION
18.04.1 LTS (Bionic Beaver)
```

## Build CLI

```
$ make
```
