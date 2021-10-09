# go-conf

[![codecov](https://codecov.io/gh/chyroc/go-conf/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/go-conf)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/go-conf "go report card")](https://goreportcard.com/report/github.com/chyroc/go-conf)
[![test status](https://github.com/chyroc/go-conf/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/go-conf/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/go-conf)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Fgo-conf.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Fgo-conf)

![](./header.png)

## Install

```shell
go get github.com/chyroc/go-conf
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/chyroc/go-conf"
)

func main() {
	res := go_project_template.Incr(1)
	fmt.Println(res) // output: 2
}
```
