PKG_TARGET=linux
PKG_BIN=./bin/go-tspath
PKG_TAG=$(shell git tag -l --contains HEAD)


#######################################
############# Development #############
#######################################

install:
	@go version
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
	@go version
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
	@go version
	@env \
		CGO_ENABLED=0 \
		GOOS=$(PKG_TARGET) \
		go build \
			-ldflags="-s -w" \
			-a -installsuffix cgo \
			-o $(PKG_BIN)
	@du -sh 
.ONESHELL: prod.release.build

prod.release.ci:
	set -e
	set -u

	@go version
	@git tag $(DRONE_TAG)
	@curl -sL https://git.io/goreleaser | bash
.ONESHELL: prod.release.ci
