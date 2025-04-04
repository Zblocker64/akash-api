UNAME_OS              := $(shell uname -s)
UNAME_ARCH            := $(shell uname -m)
PROTO_LEGACY          ?= true

ifeq (0, $(shell id -u))
$(warning "make was started with superuser privileges. it may cause issues with direnv")
endif

ifeq (, $(shell which direnv))
$(error "No direnv in $(PATH), consider installing. https://direnv.net")
endif

ifneq (1, $(AKASH_DIRENV_SET))
$(error "no envrc detected. might need to run \"direnv allow\"")
endif

# AKASH_ROOT may not be set if environment does not support/use direnv
# in this case define it manually as well as all required env variables
ifndef AKASH_ROOT
$(error "AKASH_ROOT is not set. might need to run \"direnv allow\"")
endif

ifeq (, $(GOTOOLCHAIN))
$(error "GOTOOLCHAIN is not set")
endif

GO                           := GO111MODULE=$(GO111MODULE) go
GO_MOD_NAME                  := $(shell go list -m 2>/dev/null)

BUF_VERSION                     ?= 1.28.1
PROTOC_VERSION                  ?= 21.12
GOGOPROTO_VERSION               ?= $(shell $(GO) list -mod=readonly -m -f '{{ .Version }}' github.com/cosmos/gogoproto)
# TODO https://github.com/akash-network/support/issues/77
PROTOC_GEN_GOCOSMOS_VERSION     ?= $(shell $(GO) list -mod=readonly -m -f '{{ .Version }}' github.com/regen-network/cosmos-proto)
PROTOC_GEN_GO_PULSAR_VERSION    ?= $(shell $(GO) list -mod=readonly -m -f '{{ .Version }}' github.com/cosmos/cosmos-proto)
PROTOC_GEN_GO_VERSION           ?= $(shell $(GO) list -mod=readonly -m -f '{{ .Version }}' google.golang.org/protobuf)
PROTOC_GEN_GRPC_GATEWAY_VERSION := $(shell $(GO) list -mod=readonly -m -f '{{ .Version }}' github.com/grpc-ecosystem/grpc-gateway)
PROTOC_GEN_DOC_VERSION          := $(shell $(GO) list -mod=readonly -m -f '{{ .Version }}' github.com/pseudomuto/protoc-gen-doc)

PROTOC_GEN_SWAGGER_VERSION      := $(PROTOC_GEN_GRPC_GATEWAY_VERSION)
MODVENDOR_VERSION               ?= v0.5.0
MOCKERY_VERSION                 ?= 2.42.0
GOLANGCI_LINT_VERSION           ?= v1.63.4

BUF_VERSION_FILE                     := $(AKASH_DEVCACHE_VERSIONS)/buf/$(BUF_VERSION)
PROTOC_VERSION_FILE                  := $(AKASH_DEVCACHE_VERSIONS)/protoc/$(PROTOC_VERSION)
GOGOPROTO_VERSION_FILE               := $(AKASH_DEVCACHE_VERSIONS)/gogoproto/$(GOGOPROTO_VERSION)
PROTOC_GEN_GOCOSMOS_VERSION_FILE     := $(AKASH_DEVCACHE_VERSIONS)/protoc-gen-gocosmos/$(PROTOC_GEN_GOCOSMOS_VERSION)
PROTOC_GEN_GO_PULSAR_VERSION_FILE    := $(AKASH_DEVCACHE_VERSIONS)/protoc-gen-go-pulsar/$(PROTOC_GEN_GO_PULSAR_VERSION)
PROTOC_GEN_GO_VERSION_FILE           := $(AKASH_DEVCACHE_VERSIONS)/protoc-gen-go/$(PROTOC_GEN_GO_VERSION)
PROTOC_GEN_GRPC_GATEWAY_VERSION_FILE := $(AKASH_DEVCACHE_VERSIONS)/protoc-gen-grpc-gateway/$(PROTOC_GEN_GRPC_GATEWAY_VERSION)
PROTOC_GEN_SWAGGER_VERSION_FILE      := $(AKASH_DEVCACHE_VERSIONS)/protoc-gen-swagger/$(PROTOC_GEN_SWAGGER_VERSION)
PROTOC_GEN_DOC_VERSION_FILE          := $(AKASH_DEVCACHE_VERSIONS)/protoc-gen-doc/$(PROTOC_GEN_DOC_VERSION)
MODVENDOR_VERSION_FILE               := $(AKASH_DEVCACHE_VERSIONS)/modvendor/$(MODVENDOR_VERSION)
GIT_CHGLOG_VERSION_FILE              := $(AKASH_DEVCACHE_VERSIONS)/git-chglog/$(GIT_CHGLOG_VERSION)
MOCKERY_VERSION_FILE                 := $(AKASH_DEVCACHE_VERSIONS)/mockery/v$(MOCKERY_VERSION)
GOLANGCI_LINT_VERSION_FILE           := $(AKASH_DEVCACHE_VERSIONS)/golangci-lint/$(GOLANGCI_LINT_VERSION)

BUF                              := $(AKASH_DEVCACHE_BIN)/buf
PROTOC                           := $(AKASH_DEVCACHE_BIN)/protoc
# TODO https://github.com/akash-network/support/issues/77
PROTOC_GEN_GOCOSMOS              := $(AKASH_DEVCACHE_BIN)/legacy/protoc-gen-gocosmos
PROTOC_GEN_GO_PULSAR             := $(AKASH_DEVCACHE_BIN)/protoc-gen-go-pulsar
PROTOC_GEN_GO                    := $(AKASH_DEVCACHE_BIN)/protoc-gen-go
PROTOC_GEN_GRPC_GATEWAY          := $(AKASH_DEVCACHE_BIN)/protoc-gen-grpc-gateway
PROTOC_GEN_SWAGGER               := $(AKASH_DEVCACHE_BIN)/protoc-gen-swagger
PROTOC_GEN_DOC                   := $(AKASH_DEVCACHE_BIN)/protoc-gen-doc
MODVENDOR                        := $(AKASH_DEVCACHE_BIN)/modvendor
GIT_CHGLOG                       := $(AKASH_DEVCACHE_BIN)/git-chglog
SWAGGER_COMBINE                  := $(AKASH_DEVCACHE_NODE_BIN)/swagger-combine
MOCKERY                          := $(AKASH_DEVCACHE_BIN)/mockery
GOLANGCI_LINT                    := $(AKASH_DEVCACHE_BIN)/golangci-lint

GOLANGCI_LINT_RUN                := $(GOLANGCI_LINT) run
GOLINT                           := $(GOLANGCI_LINT_RUN) ./... --disable-all --timeout=5m --enable

DOCKER_RUN            := docker run --rm -v $(shell pwd):/workspace -w /workspace
DOCKER_BUF            := $(DOCKER_RUN) bufbuild/buf:$(BUF_VERSION)

include $(AKASH_ROOT)/make/setup-cache.mk
include $(AKASH_ROOT)/make/mod.mk
include $(AKASH_ROOT)/make/test.mk
include $(AKASH_ROOT)/make/codegen.mk
include $(AKASH_ROOT)/make/lint.mk
include $(AKASH_ROOT)/make/release-ts.mk
include $(AKASH_ROOT)/make/code-style.mk

.PHONY: clean
clean:
	rm -rf $(AKASH_DEVCACHE)
	rm -rf $(AKASH_TS_ROOT)/node_modules
	rm -rf $(AKASH_TS_ROOT)/dist
