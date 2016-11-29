SHELL := /bin/bash

DEV_VERSION := $(shell stepup version)

build:
	CGO_ENABLED=0 go build -a -ldflags '-s' -ldflags "-X main.version=${DEV_VERSION}" .

build-njen:

build-dist:
	bin/build-dist

build-docker:
	bin/docker-build

test:
	bin/test

update-dep:
	bin/up-dep

dist:
	bin/dist

release: dist build-docker

.PHONY: build
