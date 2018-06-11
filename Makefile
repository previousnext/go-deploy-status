#!/usr/bin/make -f

export CGO_ENABLED=0

PROJECT=github.com/previousnext/deploy-status

# Builds the project
build:
	gox -os='linux darwin' -arch='amd64' -output='bin/deploy-status_{{.OS}}_{{.Arch}}' $(PROJECT)

# Run all lint checking with exit codes for CI
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

# Run tests with coverage reporting
test:
	go test -cover ./...

IMAGE=previousnext/deploy-status
VERSION=$(shell git describe --tags --always)

# Releases the project Docker Hub
release:
	docker build -t ${IMAGE}:${VERSION} -t ${IMAGE}:latest .
	docker push ${IMAGE}:${VERSION}
	docker push ${IMAGE}:latest
