# List all back service on the project
SERVICES=link

# Name of binary
APP_NAME=links123

# Versioning
VERSION = $(shell git tag --sort="v:refname" | tail -n1)
GITHASH = $(shell git rev-parse --short HEAD)

LDFLAGS += -X github.com/links-123/links123/shared/version.Version=$(VERSION)
LDFLAGS += -X github.com/links-123/links123/shared/version.CommitHash=$(GITHASH)
LDFLAGS += -X github.com/links-123/links123/shared/version.BuiltAt=$(shell date +"%D-%T")

install-dependency-manager:
	go get -u github.com/golang/dep/cmd/dep

run-dependency-manager:
	dep ensure -v

build:
	@go build -o $(APP_NAME) -ldflags "${LDFLAGS}" entry/entry.go
	@./$(APP_NAME) --version

generate-proto:
	for service in $(SERVICES) ; do \
		make -C app/services/$$service gen-proto; \
	done

install-linter:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.12.3

run-linter:
	golangci-lint run -v

run-tests:
	cd test-pack && $(MAKE) run-all-tests
