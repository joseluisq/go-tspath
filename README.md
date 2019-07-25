# Go TSPath [![Build Status](https://travis-ci.org/joseluisq/go-tspath.svg?branch=master)](https://travis-ci.org/joseluisq/go-tspath) [![codecov](https://codecov.io/gh/joseluisq/go-tspath/branch/master/graph/badge.svg)](https://codecov.io/gh/joseluisq/go-tspath) [![Go Report Card](https://goreportcard.com/badge/github.com/joseluisq/go-tspath)](https://goreportcard.com/report/github.com/joseluisq/go-tspath) [![GoDoc](https://godoc.org/github.com/joseluisq/go-tspath?status.svg)](https://godoc.org/github.com/joseluisq/go-tspath)

> A fast [Typescript paths](https://www.typescriptlang.org/docs/handbook/module-resolution.html) replacer written in [Go](https://golang.org/). ⚡

**go-tspath** replaces directly [Typescript paths aliases](https://www.typescriptlang.org/docs/handbook/module-resolution.html) into JS files with real paths based on `tsconfig.json`, no more runtime replacers.

__Status:__ Beta

_🚀 View current beta releases at [go-tspath/releases](https://github.com/joseluisq/go-tspath/releases)._

## Installation

```sh
go get -u github.com/joseluisq/go-tspath
```

## Usage

```sh
# 1. Build TS files via tsc
# 2. Replace TS paths
go-tspath -config=./tsconfig.json
# 3. Just run your app
# node main.js
```

## API

```sh
~> go-tspath --help
Usage of go-tspath:
  -config string
    	Specifies the Typescript configuration file. (default "./tsconfig.json")
```

## Contributions

Feel free to send some [Pull request](https://github.com/joseluisq/go-tspath/pulls) or [issue](https://github.com/joseluisq/go-tspath/issues).

## License
MIT license

© 2019 [Jose Quintana](https://git.io/joseluisq)
