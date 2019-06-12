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
go-tspath -source=./src/**/*.js -config=./tsconfig.json
```

## API

```sh
~> go-tspath --help
Usage of go-tspath:
  -config string
    	Specifies the Typescript configuration file. (default "./tsconfig.json")
  -source string
    	Specifies path of Javascript files emitted by tsc. (default "./dist/**/*.js")
```

## Contributions

Feel free to send some [Pull request](https://github.com/joseluisq/go-tspath/pulls) or [issue](https://github.com/joseluisq/go-tspath/issues).

## License
MIT license

© 2019 [Jose Quintana](https://git.io/joseluisq)
