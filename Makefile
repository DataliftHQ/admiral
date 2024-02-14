SHELL:=/usr/bin/env bash
.DEFAULT_GOAL:=all

MAKEFLAGS += --no-print-directory

VERSION ?= 0.0.0-dev

PROJECT_ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: help # Print this help message.
help:
	@grep -E '^\.PHONY: [a-zA-Z_-]+ .*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = "(: |#)"}; {printf "%-30s %s\n", $$2, $$3}'

.PHONY: all # Generate proto, cli, ui, and server assets.
all: proto server-with-assets cli web

.PHONY: lint # Lint all of the code.
lint: proto-lint server-lint cli-lint web-lint

.PHONY: lint-fix # Lint and fix all of the code.
lint-fix: server-lint-fix cli-lint-fix web-lint-fix

.PHONY: verify # Verify all of the code.
verify: proto-verify server-verify cli-verify web-verify

.PHONY: proto # Generate proto assets.
proto:
	@cd common && rm -rf api && ../tools/buf.sh generate
	@cd server && rm -rf config && ../tools/buf.sh generate

.PHONY: proto-lint # Lint the generated proto assets.
proto-lint:
	@cd common && ../tools/buf.sh lint
	@cd server && ../tools/buf.sh lint

.PHONY: proto-verify # Verify proto changes.
proto-verify:
	find common/api -mindepth 1 -maxdepth 1 -type d -exec rm -rf {} \;
	find server/config -mindepth 1 -maxdepth 1 -type d -exec rm -rf {} \;
	@$(MAKE) proto
	tools/ensure-no-diff.sh common/api server/config

.PHONY: server # Build the server.
server:
	cd server && CGO_ENABLED=0 go build -o ../build/admiral-server -ldflags "-s -w -X main.version=$(VERSION)"

.PHONY: server-with-assets # Build the server with ui assets.
server-with-assets: web
	cd server && go run cmd/assets/generate.go ../web/build && CGO_ENABLED=0 go build -tags withAssets -o ../build/admiral-server -ldflags="-X main.version=$(VERSION)"

.PHONY: server-dev # Start the server in development mode.
server-dev:
	tools/air.sh

.PHONY: server-lint # Lint the server code.
server-lint:
	tools/golangci-lint.sh run

.PHONY: server-lint-fix # Lint and fix the server code.
server-lint-fix:
	tools/golangci-lint.sh run --fix
	cd server && go mod tidy

.PHONY: server-verify # Verify go modules' requirements files are clean.
server-verify:
	cd server && go mod tidy
	tools/ensure-no-diff.sh server

.PHONY: cli # Build the CLI.
cli:
	CGO_ENABLED=0 go build -C cli -o ../build/admiral -ldflags "-s -w -X main.version=$(VERSION)"

.PHONY: cli-lint # Lint the cli code.
cli-lint:
	cd cli && tools/golangci-lint.sh run

.PHONY: cli-lint-fix # Lint and fix the cli code.
cli-lint-fix:
	cd cli && tools/golangci-lint.sh run --fix && go mod tidy

.PHONY: cli-verify # Verify go modules' requirements files are clean.
cli-verify:
	cd cli && go mod tidy
	tools/ensure-no-diff.sh cli

.PHONY: npm-install # Install web dependencies.
npm-install:
	npm --prefix web ci

.PHONY: web # Build the web code.
web: npm-install
	npm --prefix web run build

.PHONY: web-lint # Lint the web code.
web-lint:
	npm --prefix web run lint

.PHONY: web-lint-fix # Lint and fix the web code.
web-lint-fix:
	npm --prefix web run lint:fix

.PHONY: web-verify # Verify web packages are sorted.
web-verify:
	npm --prefix web run lint:packages

.PHONY: dev # Start the start in development mode.
dev:
	$(MAKE) -j2 server-dev web-dev

.PHONY: cli-completions # Generate cli shell completion scripts.
cli-completions:
	cd cli && ./scripts/completions.sh

.PHONY: cli-manpages # Generate cli manpages.
cli-manpages:
	cd cli && ./scripts/manpages.sh

.PHONY: clean # Clean all build artifacts.
clean:
	@rm -rf build
	@rm -rf dist
	@rm -rf cli/completions
	@rm -rf cli/manpages
	@rm -rf web/build
	@rm -rf web/node_modules
