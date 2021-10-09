# confer

[![codecov](https://codecov.io/gh/chyroc/confer/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/confer)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/confer "go report card")](https://goreportcard.com/report/github.com/chyroc/confer)
[![test status](https://github.com/chyroc/confer/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/confer/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/confer)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Fconfer.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Fconfer)

![](./header.png)

## Install

```shell
go get github.com/chyroc/confer
```

## Config

### Software Architecture

The confer software is divided into two parts: data loading layer and data conversion.

Use the tag of the structure in go to define what loader and transfer are used by the field, and what their parameters are.

### Loader

Confer has several built-in data loading layers.

build-in loader:

- `env`

Of course, users can also customize loader:

`loader1` and `loader2` should impl interface: `Loader`

```go
WithLoader(loader1, loader2)
```

### Transfer

Confer has several built-in data data conversion layer.

build-in transfer:

- `TODO`

Of course, users can also customize transfer:

`transfer1` and `transfer2` should impl interface: `Transfer`

```go
WithTransfer(transfer1, transfer2)
```

### Conf Grammar

The structure field uses `conf` as the tag name to accept parameters from the user. such as:

```go
type Conf struct {
	GitHubToken string `conf:"env,GITHUB_TOKEN"`
}
```

#### Conf Loader Grammar

In tag conf, read several strings separated by `,` in turn, the first of which is the name of the loader layer,

and the subsequent list of strings are the parameters of the loader function execution

In the above example, the loader is `env`, and `GITHUB_TOKEN` is passed as a parameter to the `env` loader for processing

You can define only loader without defining transfer. The `conf:"env,GITHUB_TOKEN"` listed in the previous example only has loader

#### Conf Transfer Grammar

If you still need to define transfer, separate it with `;` and loader,

and then a list of strings separated by `,` where the first string is the name of transfer, and the following string list is the parameters of transfer.

Of course, loader The loaded string data will also be the parameters of transfer

```go
type Conf struct {
	GitHubToken string `conf:"env,JSON_TOKEN;json,.GITHUB"`
}
```

In this example, env `JSON_TOKEN` stores the json data of a token, then after the loader loads the data, 

it also needs to use the json transfer to convert the final data from the `GITHUB` key

## Usage

- [load_env](./_examples/load_env/main.go): load from env
