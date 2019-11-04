#!/usr/bin/make -f

export CGO_ENABLED=0

PROJECT=github.com/previousnext/go-deploy-status
VERSION=$(shell git describe --tags --always)
COMMIT=$(shell git rev-list -1 HEAD)

# Builds the project.
define go_build
	GOOS=${1} GOARCH=${2} go build -o bin/cognito_auth_${1}_${2} -ldflags='-extldflags "-static"' github.com/previousnext/go-deploy-status
endef

# Builds the project.
build:
	$(call go_build,linux,amd64)
	$(call go_build,darwin,amd64)

# Run all lint checking with exit codes for CI.
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

# Run tests with coverage reporting.
test:
	go test -cover ./...

IMAGE=previousnext/go-deploy-status

release-github: build
	ghr -u previousnext "${VERSION}" ./bin/

release: release-github

.PHONY: *
