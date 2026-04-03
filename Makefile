# Repository root (resolved via git, works from any directory)
REPO_ROOT := $(shell git rev-parse --show-toplevel)

# Setting SHELL to bash (cross-platform)
ifeq ($(OS),Windows_NT)
    SHELL := C:/Program Files/Git/bin/bash.exe
else
    SHELL := bash
endif
.SHELLFLAGS := -o pipefail -ec

.PHONY: build test lint fmt clean

build:
	go build ./cmd/aula-cli
	go build ./cmd/aula-mcp

test:
	go test ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -w .
	goimports -w .

clean:
	rm -f aula-cli aula-cli.exe aula-mcp aula-mcp.exe

run:
	go run ./cmd/aula-cli $(ARGS)
