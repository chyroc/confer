# go-loader

[![codecov](https://codecov.io/gh/chyroc/go-loader/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/go-loader)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/go-loader "go report card")](https://goreportcard.com/report/github.com/chyroc/go-loader)
[![test status](https://github.com/chyroc/go-loader/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/go-loader/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/go-loader)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Fgo-loader.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Fgo-loader)

![](./header.png)

## Install

```shell
go get github.com/chyroc/go-loader
```

## Go-Loader

### Software Architecture

The go-loader software is divided into two parts: data extract layer: `Extractor` and data transform layer: `Transformer`.

Use the tag of the structure in go to define what `Extractor` and `Transformer` are used by the field, and what their parameters are.

### `Extractor` and `Transformer`

go-loader has several built-in data `Extractor` and `Transformer`.

- build-in `Extractor`
  - `env`
  - `file`
- build-in `Transformer`

Of course, users can also customize:

```go
WithExtractor(extractor1, extractor2)

WithTransformer(transformer1, transformer2)
```

### Tag Grammar Rule

The structure field uses `loader` as the tag name to accept parameters from the user. such as:

```go
type Conf struct {
	GitHubToken string `loader:"env,GITHUB_TOKEN"`
}
```

#### Extractor Grammar

In tag `loader`, read several strings separated by `,` in turn, the first of which is the name of the `Extractor`,

and the subsequent list of strings are the parameters of the `Extractor` function execution

In the above example, the `Extractor` is `env`, and `GITHUB_TOKEN` is passed as a parameter to the `env` `Extractor` for processing

You can define only `Extractor` without defining `Transformer`. The `loader:"env,GITHUB_TOKEN"` listed in the previous example only has `Extractor`

#### Transformer Grammar

If you still need to define `Transformer`, separate it with `;`

and then a list of strings separated by `,` where the first string is the name of `Transformer`, and the following string list is the parameters of `Transformer`.

```go
type Conf struct {
	GitHubToken string `loader:"env,JSON_TOKEN;json,.GITHUB"`
}
```

In this example, env `JSON_TOKEN` stores the json data of a token, then after the `Extractor` extract the data, 

it also needs to use the json `Transformer` to convert the final data from the `GITHUB` key

## Usage

- [load_env](./_examples/load_env/main.go): load data from env
