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
	@mkdir -p bin
	go build -o bin/aula-cli ./cmd/aula-cli
	go build -o bin/aula-mcp ./cmd/aula-mcp

test:
	go test ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -w .
	goimports -w .

clean:
	rm -rf bin/

run:
	go run ./cmd/aula-cli $(ARGS)

run-mcp:
	go run ./cmd/aula-mcp $(ARGS)
