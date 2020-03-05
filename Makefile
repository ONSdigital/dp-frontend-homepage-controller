BINPATH ?= build

build:
	go build -tags 'production' -o $(BINPATH)/dp-frontend-homepage-controller -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"

debug:
	go build -tags 'debug' -o $(BINPATH)/dp-frontend-homepage-controller -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"
	HUMAN_LOG=1 DEBUG=1 $(BINPATH)/dp-frontend-homepage-controller

test:
	go test -race -cover ./...
.PHONY: test

convey:
	goconvey ./...

.PHONY: build debug