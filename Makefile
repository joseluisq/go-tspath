# Go

ifndef ENV_OS
override ENV_OS = linux
endif

GO_OS=$(ENV_OS)
GO_DEST=./bin
GO_BINARY=$(GO_DEST)/go-tspath

install:
	-go get github.com/oxequa/realize
.PHONY: install

watch:
	-realize start
.PHONY: watch

release:
	-goreleaser release --skip-publish --rm-dist
.PHONY: release

build:
	-env \
		CGO_ENABLED=0 \
		GOOS=$(GO_OS) \
		go build \
			-ldflags="-s -w" \
			-a -installsuffix cgo \
			-o $(GO_BINARY)
.PHONY: build
