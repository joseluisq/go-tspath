PKG_TARGET=linux
PKG_BIN_PATH=./bin

PKG_NAME=go-tspath
PKG_TAG=$(shell git tag -l --contains HEAD)


#######################################
############# Development #############
#######################################

install:
	@go get -v golang.org/x/lint/golint
	@go mod download
.PHONY: install

watch:
	@refresh run
.PHONY: watch

tidy:
	@go mod tidy
.PHONY: tidy

dev.release:
	set -e
	set -u

	@goreleaser release --skip-publish --rm-dist --snapshot
.ONESHELL: dev.release


#######################################
########### Utility tasks #############
#######################################

test:
	@golint -set_exit_status ./...
	@go test -v -timeout 30s -race -coverprofile=coverage.txt -covermode=atomic ./...
.PHONY: test

coverage:
	@bash -c "bash <(curl -s https://codecov.io/bash)"
.PHONY: coverage


#######################################
########## Production tasks ###########
#######################################

prod.release.build:
	@env \
		CGO_ENABLED=0 \
		GOOS=$(GO_OS) \
		go build \
			-ldflags="-s -w" \
			-a -installsuffix cgo \
			-o $(GO_BINARY)
.ONESHELL: prod.release.build

prod.release:
	set -e
	set -u

	@goreleaser release --rm-dist
.ONESHELL: prod.release
