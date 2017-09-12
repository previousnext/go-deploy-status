#!/usr/bin/make -f

VERSION=$(shell git describe --tags --always)
IMAGE=previousnext/CHANGE_ME

release: build push

build:
	docker build -t ${IMAGE}:${VERSION} .

push:
	docker push ${IMAGE}:${VERSION}
