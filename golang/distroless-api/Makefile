SHELL := /bin/bash

# Configuration variables
DEBUG          ?= true
LTS            ?= false
LANGUAGE       ?= go
VERSION        ?= 1.24
IMAGE_NAME     ?= alemos/distroless-api
PORT           ?= 8080
APP_ENV        ?= production
CONTAINER_NAME ?= distroless-api-go

# Version descriptor: version_number or version-lts
VERSION_LABEL  := $(if $(filter true 1 yes,$(LTS)),$(VERSION)-lts,$(VERSION))

# Variant: regular or debug
IMAGE_VARIANT  := $(if $(filter false 0 no,$(DEBUG)),regular,debug)

# Pattern: (node|go)(version_number|version-LTS)-(debug|regular)
TAG            := $(LANGUAGE)$(VERSION_LABEL)-$(IMAGE_VARIANT)

# Base golden image following: (node|go)(version_number|version-LTS)-(debug|regular)
DEFAULT_BASE_IMAGE := alemos/golden_images:$(TAG)
BASE_IMAGE     ?= $(DEFAULT_BASE_IMAGE)

# Full API image reference
FULL_IMAGE     := $(IMAGE_NAME):$(TAG)

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help message
	@echo "Usage: make [target] [VARIABLE=value]"
	@echo ""
	@echo "Naming standard: (node|go)(version_number|version-LTS)-(debug|regular)"
	@echo "  Current target image: $(FULL_IMAGE)"
	@echo "  Current base image:   $(BASE_IMAGE)"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  %-18s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "Variables:"
	@echo "  DEBUG              Image variant: true -> debug, false -> regular (default: $(DEBUG))"
	@echo "  LTS                Version suffix: true -> $(VERSION)-lts (default: $(LTS))"
	@echo "  LANGUAGE           Language prefix: node or go (default: $(LANGUAGE))"
	@echo "  VERSION            Language version (default: $(VERSION))"
	@echo "  IMAGE_NAME         API image name (default: $(IMAGE_NAME))"
	@echo "  TAG                Full tag (default: $(TAG))"
	@echo "  BASE_IMAGE         Golden Image base (default: $(BASE_IMAGE))"
	@echo "  PORT               Host port to expose (default: $(PORT))"
	@echo "  APP_ENV            Application environment (default: $(APP_ENV))"
	@echo "  CONTAINER_NAME     Container name for run/start (default: $(CONTAINER_NAME))"

# ==============================================================================
# Base Image Targets (dockerfile.golang124)
# ==============================================================================

.PHONY: build-base
build-base: ## Build base image via dockerfile.golang124 (respects DEBUG and LTS)
ifeq ($(filter false 0 no,$(DEBUG)),)
	$(MAKE) -C .. build-debug APP_ENV=$(APP_ENV) VERSION=$(VERSION) LTS=$(LTS)
else
	$(MAKE) -C .. build APP_ENV=$(APP_ENV) VERSION=$(VERSION) LTS=$(LTS)
endif

.PHONY: build-base-prod
build-base-prod: ## Build regular base image (e.g. go1.24-regular)
	$(MAKE) -C .. build APP_ENV=$(APP_ENV) VERSION=$(VERSION) LTS=$(LTS)

.PHONY: build-base-debug
build-base-debug: ## Build debug base image (e.g. go1.24-debug)
	$(MAKE) -C .. build-debug APP_ENV=$(APP_ENV) VERSION=$(VERSION) LTS=$(LTS)

# ==============================================================================
# API Application Targets
# ==============================================================================

.PHONY: build
build: build-base ## Build the Go API docker image using local base image
	docker build \
		--build-arg BASE_IMAGE=$(BASE_IMAGE) \
		-t $(FULL_IMAGE) \
		.

.PHONY: build-nc
build-nc: build-base ## Build the Go API docker image without cache
	docker build \
		--no-cache \
		--build-arg BASE_IMAGE=$(BASE_IMAGE) \
		-t $(FULL_IMAGE) \
		.

.PHONY: start
start: build stop ## Build (if needed) and run container in background
	docker run -d \
		--pull=never \
		--name $(CONTAINER_NAME) \
		-p $(PORT):8080 \
		-e APP_ENV=$(APP_ENV) \
		-e PORT=8080 \
		$(FULL_IMAGE)
	@echo "Go API started on http://localhost:$(PORT) [$(FULL_IMAGE)]"

.PHONY: run
run: build ## Build (if needed) and run container in foreground
	docker run --rm \
		--pull=never \
		--name $(CONTAINER_NAME) \
		-p $(PORT):8080 \
		-e APP_ENV=$(APP_ENV) \
		-e PORT=8080 \
		$(FULL_IMAGE)

.PHONY: stop
stop: ## Stop and remove background container
	-docker stop $(CONTAINER_NAME)
	-docker rm $(CONTAINER_NAME)

.PHONY: logs
logs: ## Follow logs from running container
	docker logs -f $(CONTAINER_NAME)

.PHONY: test
test: ## Test the API endpoints (/health and /)
	@echo "Testing /health endpoint:"
	@curl -s -i http://localhost:$(PORT)/health || echo "Failed to connect to /health"
	@echo -e "\n\nTesting root endpoint /:"
	@curl -s -i http://localhost:$(PORT)/ || echo "Failed to connect to /"
	@echo ""

.PHONY: shell
shell: ## Open busybox shell in running container (available when variant is debug)
	docker exec -it $(CONTAINER_NAME) /busybox/sh

.PHONY: clean
clean: stop ## Stop container and remove the API image
	-docker rmi $(FULL_IMAGE)
