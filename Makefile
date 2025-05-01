BINPATH ?= build

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)
LOCAL_DP_RENDERER_IN_USE = $(shell grep -c "\"github.com/ONSdigital/dis-design-system-go\" =" go.mod)

.PHONY: audit
audit:
	go list -m all | nancy sleuth

.PHONY: build
build: generate-prod
	go build -tags 'production' -o $(BINPATH)/dp-frontend-homepage-controller -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"

.PHONY: lint
lint: ## Used in ci to run linters against Go code
	cp assets/assets.go assets/assets.go.bak
	echo 'func Asset(_ string) ([]byte, error) { return nil, nil }' >> assets/assets.go
	echo 'func AssetNames() []string { return []string{} }' >> assets/assets.go
	gofmt -w assets/assets.go
	golangci-lint run ./... || { echo "Linting failed, restoring original assets.go"; mv assets/assets.go.bak assets/assets.go; exit 1; }
	mv assets/assets.go.bak assets/assets.go

.PHONY: debug
debug: generate-debug
	go build -tags 'debug' -race -o $(BINPATH)/dp-frontend-homepage-controller -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"
	HUMAN_LOG=1 DEBUG=1 APP_RENDERER_VERSION=$(APP_RENDERER_VERSION) $(BINPATH)/dp-frontend-homepage-controller

.PHONY: debug-run
debug-run: generate-debug
	HUMAN_LOG=1 DEBUG=1 go run -tags 'debug' -race -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)" main.go

.PHONY: test
test: generate-prod
	go test -race -cover -tags 'production' ./...

.PHONY: convey
convey:
	goconvey ./...

.PHONY: generate-debug
generate-debug: fetch-dp-renderer
	# fetch the renderer library and build the dev version
	go install github.com/kevinburke/go-bindata/v4/...@latest
	cd assets; go-bindata -prefix $(CORE_ASSETS_PATH)/assets -debug -o data.go -pkg assets locales/... templates/... $(CORE_ASSETS_PATH)/assets/locales/... $(CORE_ASSETS_PATH)/assets/templates/...
	{ printf "//go:build debug\n"; cat assets/data.go; } > assets/debug.go.new
	mv assets/debug.go.new assets/data.go

.PHONY: generate-prod
generate-prod: fetch-dp-renderer
	# fetch the renderer library and build the prod version
	go install github.com/kevinburke/go-bindata/v4/...@latest
	cd assets; go-bindata -prefix $(CORE_ASSETS_PATH)/assets -o data.go -pkg assets locales/... templates/... $(CORE_ASSETS_PATH)/assets/locales/... $(CORE_ASSETS_PATH)/assets/templates/...
	{ printf "//go:build production\n"; cat assets/data.go; } > assets/data.go.new
	mv assets/data.go.new assets/data.go

.PHONY: fetch-dp-renderer
fetch-dp-renderer:
ifeq ($(LOCAL_DP_RENDERER_IN_USE), 1)
	$(eval CORE_ASSETS_PATH = $(shell grep -w "\"github.com/ONSdigital/dis-design-system-go\" =>" go.mod | awk -F '=> ' '{print $$2}' | tr -d '"'))
else
	$(eval APP_RENDERER_VERSION=$(shell grep "github.com/ONSdigital/dis-design-system-go" go.mod | cut -d ' ' -f2 ))
	$(eval CORE_ASSETS_PATH = $(shell go get github.com/ONSdigital/dis-design-system-go@$(APP_RENDERER_VERSION) && go list -f '{{.Dir}}' -m github.com/ONSdigital/dis-design-system-go))
endif

.PHONY: test-component
test-component:
	$(shell mkdir -p assets/dist)
	$(shell touch assets/dist/scripts.js)
	make generate-prod
	go test -cover -tags 'production' -coverpkg=github.com/ONSdigital/dp-frontend-homepage-controller/... -component
