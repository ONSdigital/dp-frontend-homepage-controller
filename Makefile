BINPATH ?= build

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)

all: audit test build

audit:
	go list -m all | nancy sleuth
build:
	go build -tags 'production' -o $(BINPATH)/dp-frontend-homepage-controller -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"

debug:
	go build -tags 'debug' -race -o $(BINPATH)/dp-frontend-homepage-controller -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"
	HUMAN_LOG=1 DEBUG=1 $(BINPATH)/dp-frontend-homepage-controller

test:
	go test -race -cover ./...
.PHONY: test

convey:
	goconvey ./...

.PHONY: all build debug audit