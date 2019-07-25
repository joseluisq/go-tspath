GO_BINARY=./bin/go-tspath
GO_OS := $(shell uname -s | tr A-Z a-z)

install:
	@go get -v golang.org/x/lint/golint
	@go get github.com/markbates/refresh
.PHONY: install

test:
	@go test -v -timeout 30s -race -coverprofile=coverage.txt -covermode=atomic ./...
.PHONY: test

coverage:
	@bash -c "bash <(curl -s https://codecov.io/bash)"
.PHONY: coverage

watch:
	@refresh run
.PHONY: watch

tidy:
	@go mod tidy
.PHONY: tidy

release:
	@goreleaser release --rm-dist
.PHONY: release

release-test:
	@goreleaser release --skip-publish --rm-dist --snapshot
.PHONY: release-test

build:
	env \
		CGO_ENABLED=0 \
		GOOS=$(GO_OS) \
		go build \
			-ldflags="-s -w" \
			-a -installsuffix cgo \
			-o $(GO_BINARY)
.PHONY: build
