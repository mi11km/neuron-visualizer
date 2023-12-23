NAME := $(shell basename $(CURDIR))

GO ?= go
GOOS := $(shell $(GO) env GOOS)
GOARCH := $(shell $(GO) env GOARCH)

BIN_PATH := $(abspath ./bin/$(GOOS)_$(GOARCH))
PROTO_PATH := $(abspath ./proto)
SERVER_PATH := $(abspath ./server)

GO_ENV ?= CGO_ENABLED=0 GOBIN=$(BIN_PATH)


$(shell mkdir -p $(BIN_PATH))


BUF_VERSION := v1.28.1
$(BIN_PATH)/buf-$(BUF_VERSION):
	unlink $(BIN_PATH)/buf || true
	$(GO_ENV) $(GO) install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	mv $(BIN_PATH)/buf $(BIN_PATH)/buf-$(BUF_VERSION)
	ln -s $(BIN_PATH)/buf-$(BUF_VERSION) $(BIN_PATH)/buf


.PHONY: init-proto
init-proto: $(BIN_PATH)/buf-$(BUF_VERSION)
	$(BIN_PATH)/buf mod init -o $(PROTO_PATH)

.PHONY: build-proto
build-proto: $(BIN_PATH)/buf-$(BUF_VERSION)
	$(BIN_PATH)/buf build --path $(PROTO_PATH)

.PHONY: generate-proto
generate-proto: $(BIN_PATH)/buf-$(BUF_VERSION) format-proto lint-proto
	cd $(PROTO_PATH) && $(BIN_PATH)/buf generate

.PHONY: lint-proto
lint-proto: $(BIN_PATH)/buf-$(BUF_VERSION)
	$(BIN_PATH)/buf lint $(PROTO_PATH)

.PHONY: format-proto
format-proto: $(BIN_PATH)/buf-$(BUF_VERSION)
	$(BIN_PATH)/buf format -w $(PROTO_PATH)

.PHONY: grpc-curl-local
grpc-curl-local: DATA := {"message":"from buf curl"}
grpc-curl-local: SERVICE := health.v1.HealthCheckService/Check
grpc-curl-local:
	$(BIN_PATH)/buf curl --protocol grpc --http2-prior-knowledge  --data '$(DATA)' http://localhost:8080/$(SERVICE)

.PHONY: run-server
run-server:
	cd $(SERVER_PATH) && $(GO_ENV) $(GO) run cmd/main.go

.PHONY: go-fmt
go-fmt:
	cd $(SERVER_PATH) && $(GO_ENV) $(GO) fmt ./...

.PHONY: go-vet
go-vet:
	cd $(SERVER_PATH) && $(GO_ENV) $(GO) vet ./...

fmt: format-proto lint-proto go-fmt go-vet

