SHELL:=/usr/bin/env bash
.DEFAULT_GOAL:=all

MAKEFLAGS += --no-print-directory

DOCS_DEPLOY_USE_SSH ?= true
DOCS_DEPLOY_GIT_USER ?= git

VERSION := 0.0.0

PROJECT_ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: help # Print this help message.
help:
	@grep -E '^\.PHONY: [a-zA-Z_-]+ .*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = "(: |#)"}; {printf "%-30s %s\n", $$2, $$3}'

.PHONY: all # Generate API, frontend, and server assets.
all: api

.PHONY: api # Generate API assets.
api:
	@cd common && ../tools/buf.sh generate
	@cd server && ../tools/buf.sh generate

.PHONY: api-lint # Lint the generated API assets.
api-lint:
	@cd common && ../tools/buf.sh lint
	@cd server && ../tools/buf.sh lint

.PHONY: api-verify # Verify API changes.
api-verify:
	find common/api -mindepth 1 -maxdepth 1 -type d -exec rm -rf {} \;
	find server/config -mindepth 1 -maxdepth 1 -type d -exec rm -rf {} \;
	@$(MAKE) api
	tools/ensure-no-diff.sh common/api server/config

.PHONY: build # Build the standalone server.
build:
	@echo "build"
