# Repository root (resolved via git, works from any directory)
REPO_ROOT := $(shell git rev-parse --show-toplevel)

# Setting SHELL to bash (cross-platform)
ifeq ($(OS),Windows_NT)
    SHELL := C:/Program Files/Git/bin/bash.exe
else
    SHELL := bash
endif
.SHELLFLAGS := -o pipefail -ec

.PHONY: build test lint fmt clean release-build

build:
	@mkdir -p $(REPO_ROOT)/bin
	cd $(REPO_ROOT) && go build -o bin/aula-cli ./cmd/aula-cli
	cd $(REPO_ROOT) && go build -o bin/aula-mcp ./cmd/aula-mcp

test:
	cd $(REPO_ROOT) && go test ./...

lint:
	cd $(REPO_ROOT) && golangci-lint run ./...

fmt:
	cd $(REPO_ROOT) && gofmt -w .
	cd $(REPO_ROOT) && goimports -w .

clean:
	rm -rf $(REPO_ROOT)/bin/

run:
	cd $(REPO_ROOT) && go run ./cmd/aula-cli $(ARGS)

run-mcp:
	cd $(REPO_ROOT) && go run ./cmd/aula-mcp $(ARGS)

release-build:
	@ext=""; if [ "$(GOOS)" = "windows" ]; then ext=".exe"; fi; \
	for cmd in aula-cli aula-mcp; do \
		cd $(REPO_ROOT) && go build -o "$${cmd}-$(GOOS)-$(GOARCH)$${ext}" ./cmd/$${cmd}; \
	done
