SHELL                 = /bin/bash

APP_NAME              = kibitalk
VERSION               = $(shell git describe --always --tags)
GIT_COMMIT            = $(shell git rev-parse HEAD)
GIT_DIRTY             = $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE            = $(shell date '+%Y-%m-%d-%H:%M:%S')
SQUAD                 = backend
BUSINESS              = platform

.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make cross-build-integration-test                           Compile the project.'
	@echo

.PHONY: cross-build-integration-test
cross-build-integration-test:
	@echo "GOOS=linux GOARCH=amd64"
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -X github.com/kitabisa/${APP_NAME}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/kitabisa/${APP_NAME}/version.Version=${VERSION} -X github.com/kitabisa/${APP_NAME}/version.Environment=${ENV} -X github.com/kitabisa/${APP_NAME}/version.BuildDate=${BUILD_DATE}" -o integration_test/bin/${APP_NAME}

