# List all back service on the project
_SERVICES=link

# Name of binary
_APP_NAME=links123

# Versioning
_VERSION = $(shell git tag --sort="v:refname" | tail -n1)
_GITHASH = $(shell git rev-parse --short HEAD)

_LDFLAGS += -X github.com/links-123/links123/shared/version.Version=$(_VERSION)
_LDFLAGS += -X github.com/links-123/links123/shared/version.CommitHash=$(_GITHASH)
_LDFLAGS += -X github.com/links-123/links123/shared/version.BuiltAt=$(shell date +"%D-%T")

# Building

build: clean
	@go build -o $(_APP_NAME) -ldflags "$(_LDFLAGS)" entry/entry.go
	@./$(_APP_NAME) --version

clean:
	rm -f $(_APP_NAME)

generate-proto:
	for service in $(_SERVICES) ; do \
		make -C app/services/$$service gen-proto; \
	done

# Linting

install-linter:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.24.0

run-linter:
	golangci-lint run -v

# 3-rd parties tools installation

install-proto-compiler:
	@echo "Go to: https://github.com/protocolbuffers/protobuf" && exit 1

install-proto-go-plugin:
	go get -u github.com/golang/protobuf/protoc-gen-go

install-proto-micro-plugin:
	go get github.com/micro/protoc-gen-micro/v2

.SILENT: clean
