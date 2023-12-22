NAME := $(shell basename $(CURDIR))

GO ?= go
GOOS := $(shell $(GO) env GOOS)
GOARCH := $(shell $(GO) env GOARCH)

BIN := $(abspath ./bin/$(GOOS)_$(GOARCH))
PATH := $(BIN):$(PATH)

GO_ENV ?= CGO_ENABLED=0 GOBIN=$(BIN)


$(shell mkdir -p $(BIN))

BUF_VERSION := v1.28.1
$(BIN)/buf-$(BUF_VERSION):
	unlink $(BIN)/buf || true
	$(GO_ENV) $(GO) install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	mv $(BIN)/buf $(BIN)/buf-$(BUF_VERSION)
	ln -s $(BIN)/buf-$(BUF_VERSION) $(BIN)/buf


.PHONY: init-proto
init-proto: $(BIN)/buf-$(BUF_VERSION)
	$(BIN)/buf mod init -o proto/

.PHONY: build-proto
build-proto: $(BIN)/buf-$(BUF_VERSION)
	$(BIN)/buf build --path proto/

.PHONY: generate-proto
generate-proto: $(BIN)/buf-$(BUF_VERSION) format-proto lint-proto
	cd proto && $(BIN)/buf generate

.PHONY: lint-proto
lint-proto: $(BIN)/buf-$(BUF_VERSION)
	$(BIN)/buf lint proto/

.PHONY: format-proto
format-proto: $(BIN)/buf-$(BUF_VERSION)
	$(BIN)/buf format -w proto/

.PHONY: grpc-curl-local
grpc-curl-local: DATA := {"message":"echo"}
grpc-curl-local: SERVICE := health.v1.HealthCheckService/Call
grpc-curl-local:
	$(BIN)/buf curl --protocol grpc --http2-prior-knowledge  --data '$(DATA)' http://localhost:8080/$(SERVICE)

.PHONY: run-server
run-server:
	cd server && $(GO_ENV) $(GO) run cmd/main.go

.PHONY: go-fmt
go-fmt:
	cd server && $(GO_ENV) $(GO) fmt ./...

.PHONY: go-vet
go-vet:
	cd server && $(GO_ENV) $(GO) vet ./...

fmt: format-proto format-proto go-fmt go-vet

