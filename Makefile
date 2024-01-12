NAME := $(shell basename $(CURDIR))

GO ?= go
GOOS := $(shell $(GO) env GOOS)
GOARCH := $(shell $(GO) env GOARCH)

DOCKER ?= docker

ROOT_PATH := $(abspath .)
BIN_PATH := $(abspath ./bin/$(GOOS)_$(GOARCH))
SERVER_PATH := $(abspath ./server)

GO_ENV ?= CGO_ENABLED=0 GOBIN=$(BIN_PATH)


$(shell mkdir -p $(BIN_PATH))

OAPI_CODEGEN_VERSION := v2.0.0
$(BIN_PATH)/oapi-codegen-$(OAPI_CODEGEN_VERSION):
	unlink $(BIN_PATH)/oapi-codegen || true
	$(GO_ENV) $(GO) install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION)
	mv $(BIN_PATH)/oapi-codegen $(BIN_PATH)/oapi-codegen-$(OAPI_CODEGEN_VERSION)
	ln -s $(BIN_PATH)/oapi-codegen-$(OAPI_CODEGEN_VERSION) $(BIN_PATH)/oapi-codegen

.PHONY: generate-openapi-server
generate-openapi-server: $(BIN_PATH)/oapi-codegen-$(OAPI_CODEGEN_VERSION)
	mkdir -p $(SERVER_PATH)/openapi
	rm -f $(SERVER_PATH)/openapi/*.gen.go
	$(BIN_PATH)/oapi-codegen -package openapi \
	-generate types \
	schema/rest/openapi.yaml > server/openapi/types.gen.go
	$(BIN_PATH)/oapi-codegen -package openapi \
	-generate chi-server \
	schema/rest/openapi.yaml > server/openapi/chi_server.gen.go

.PHONY: run-server
run-server:
	cd $(SERVER_PATH) && $(GO_ENV) $(GO) run cmd/main.go

.PHONY: build-server
build-server:
	cd $(SERVER_PATH) && $(GO_ENV) $(GO) build -o $(SERVER_PATH)/server cmd/main.go

.PHONY: go-tidy
go-tidy:
	cd $(SERVER_PATH) && $(GO_ENV) $(GO) mod tidy

.PHONY: go-fmt
go-fmt:
	cd $(SERVER_PATH) && $(GO_ENV) $(GO) fmt ./...

.PHONY: go-vet
go-vet:
	cd $(SERVER_PATH) && $(GO_ENV) $(GO) vet ./...

.PHONY: fmt
fmt: go-tidy go-fmt go-vet
