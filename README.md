# Go TSPath

> A [Typescript paths](https://www.typescriptlang.org/docs/handbook/module-resolution.html) replacer written in [Go](https://golang.org/). ⚡

__Status:__ WIP

## Installation

```sh
go get -u github.com/joseluisq/go-tspath
```

## Usage

```sh
# after the tsc compiling just run:
go-tspath -config=./tsconfig.json
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
