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
	@echo "generate API assets"

.PHONY: build # Build the standalone server.
build:
	@echo "build"
