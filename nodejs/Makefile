SHELL := /bin/bash

# Configuration variables
REGISTRY           ?= alemos
IMAGE_NAME         ?= golden_images
LANGUAGE           ?= node
VERSION            ?= 24
LTS                ?= false
NODE_ENV           ?= production
DISTROLESS_VARIANT ?= nonroot
DEBUG_VARIANT      ?= debug-nonroot
DOCKERFILE         ?= dockerfile.nodejs24
PLATFORMS          ?= linux/amd64,linux/arm64

# Version descriptor: version_number or version-lts
VERSION_LABEL      := $(if $(filter true 1 yes,$(LTS)),$(VERSION)-lts,$(VERSION))

# Pattern: (node|go)(version_number|version-LTS)-(debug|regular)
TAG_REGULAR        := $(LANGUAGE)$(VERSION_LABEL)-regular
TAG_DEBUG          := $(LANGUAGE)$(VERSION_LABEL)-debug

# Image references
IMAGE_BASE         := $(if $(REGISTRY),$(REGISTRY)/$(IMAGE_NAME),$(IMAGE_NAME))

IMAGE_REGULAR_TAG  := $(IMAGE_BASE):$(TAG_REGULAR)
IMAGE_DEBUG_TAG    := $(IMAGE_BASE):$(TAG_DEBUG)

# Additional alias tags for convenience and backwards compatibility
IMAGE_REGULAR_ALIAS:= $(IMAGE_BASE):$(LANGUAGE)$(VERSION_LABEL)
IMAGE_DEBUG_ALIAS  := $(IMAGE_BASE):$(LANGUAGE)$(VERSION_LABEL)-debug

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help message
	@echo "Usage: make [target] [VARIABLE=value]"
	@echo ""
	@echo "Naming standard: (node|go)(version_number|version-LTS)-(debug|regular)"
	@echo "  Regular image: $(IMAGE_REGULAR_TAG)"
	@echo "  Debug image:   $(IMAGE_DEBUG_TAG)"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  %-18s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "Variables:"
	@echo "  REGISTRY           Docker registry/username (default: $(REGISTRY))"
	@echo "  IMAGE_NAME         Base repository name (default: $(IMAGE_NAME))"
	@echo "  LANGUAGE           Language prefix: node or go (default: $(LANGUAGE))"
	@echo "  VERSION            Version number (default: $(VERSION))"
	@echo "  LTS                Append -lts to version: true/false (default: $(LTS))"
	@echo "  NODE_ENV           Node environment (default: $(NODE_ENV))"
	@echo "  DISTROLESS_VARIANT Standard distroless variant (default: $(DISTROLESS_VARIANT))"
	@echo "  DEBUG_VARIANT      Debug distroless variant (default: $(DEBUG_VARIANT))"
	@echo "  PLATFORMS          Target platforms for buildx (default: $(PLATFORMS))"

# ==============================================================================
# Standard Image Targets (regular)
# ==============================================================================

.PHONY: build
build: ## Build regular docker image: $(IMAGE_REGULAR_TAG)
	docker build \
		-f $(DOCKERFILE) \
		--build-arg DISTROLESS_VARIANT=$(DISTROLESS_VARIANT) \
		--build-arg NODE_ENV=$(NODE_ENV) \
		-t $(IMAGE_REGULAR_TAG) \
		-t $(IMAGE_REGULAR_ALIAS) \
		.

.PHONY: build-nc
build-nc: ## Build regular docker image without cache
	docker build \
		--no-cache \
		-f $(DOCKERFILE) \
		--build-arg DISTROLESS_VARIANT=$(DISTROLESS_VARIANT) \
		--build-arg NODE_ENV=$(NODE_ENV) \
		-t $(IMAGE_REGULAR_TAG) \
		-t $(IMAGE_REGULAR_ALIAS) \
		.

.PHONY: test
test: ## Smoke test regular image (checks version and NODE_ENV)
	docker run --rm $(IMAGE_REGULAR_TAG) -e "console.log('Node: ' + process.version + ' | NODE_ENV: ' + process.env.NODE_ENV)"

.PHONY: run
run: ## Run regular image in an interactive container
	docker run --rm -it $(IMAGE_REGULAR_TAG)

.PHONY: push
push: ## Push regular image to registry
	docker push $(IMAGE_REGULAR_TAG)
	docker push $(IMAGE_REGULAR_ALIAS)

.PHONY: buildx
buildx: ## Build regular multi-platform image locally
	docker buildx build \
		--platform $(PLATFORMS) \
		-f $(DOCKERFILE) \
		--build-arg DISTROLESS_VARIANT=$(DISTROLESS_VARIANT) \
		--build-arg NODE_ENV=$(NODE_ENV) \
		-t $(IMAGE_REGULAR_TAG) \
		-t $(IMAGE_REGULAR_ALIAS) \
		.

.PHONY: buildx-push
buildx-push: ## Build and push regular multi-platform image to registry
	docker buildx build \
		--platform $(PLATFORMS) \
		-f $(DOCKERFILE) \
		--build-arg DISTROLESS_VARIANT=$(DISTROLESS_VARIANT) \
		--build-arg NODE_ENV=$(NODE_ENV) \
		-t $(IMAGE_REGULAR_TAG) \
		-t $(IMAGE_REGULAR_ALIAS) \
		--push \
		.

# ==============================================================================
# Debug Image Targets (debug)
# ==============================================================================

.PHONY: build-debug
build-debug: ## Build debug docker image: $(IMAGE_DEBUG_TAG)
	docker build \
		-f $(DOCKERFILE) \
		--build-arg DISTROLESS_VARIANT=$(DEBUG_VARIANT) \
		--build-arg NODE_ENV=$(NODE_ENV) \
		-t $(IMAGE_DEBUG_TAG) \
		-t $(IMAGE_DEBUG_ALIAS) \
		.

.PHONY: build-debug-nc
build-debug-nc: ## Build debug docker image without cache
	docker build \
		--no-cache \
		-f $(DOCKERFILE) \
		--build-arg DISTROLESS_VARIANT=$(DEBUG_VARIANT) \
		--build-arg NODE_ENV=$(NODE_ENV) \
		-t $(IMAGE_DEBUG_TAG) \
		-t $(IMAGE_DEBUG_ALIAS) \
		.

.PHONY: test-debug
test-debug: ## Smoke test debug image
	docker run --rm $(IMAGE_DEBUG_TAG) -e "console.log('Node: ' + process.version + ' | NODE_ENV: ' + process.env.NODE_ENV)"

.PHONY: run-debug
run-debug: ## Run debug image with interactive busybox shell
	docker run --rm -it --entrypoint /busybox/sh $(IMAGE_DEBUG_TAG)

.PHONY: push-debug
push-debug: ## Push debug image to registry
	docker push $(IMAGE_DEBUG_TAG)
	docker push $(IMAGE_DEBUG_ALIAS)

.PHONY: buildx-debug
buildx-debug: ## Build debug multi-platform image locally
	docker buildx build \
		--platform $(PLATFORMS) \
		-f $(DOCKERFILE) \
		--build-arg DISTROLESS_VARIANT=$(DEBUG_VARIANT) \
		--build-arg NODE_ENV=$(NODE_ENV) \
		-t $(IMAGE_DEBUG_TAG) \
		-t $(IMAGE_DEBUG_ALIAS) \
		.

.PHONY: buildx-debug-push
buildx-debug-push: ## Build and push debug multi-platform image to registry
	docker buildx build \
		--platform $(PLATFORMS) \
		-f $(DOCKERFILE) \
		--build-arg DISTROLESS_VARIANT=$(DEBUG_VARIANT) \
		--build-arg NODE_ENV=$(NODE_ENV) \
		-t $(IMAGE_DEBUG_TAG) \
		-t $(IMAGE_DEBUG_ALIAS) \
		--push \
		.

# ==============================================================================
# Combined Targets
# ==============================================================================

.PHONY: build-all
build-all: build build-debug ## Build both regular and debug images

.PHONY: push-all
push-all: push push-debug ## Push both regular and debug images

.PHONY: buildx-all-push
buildx-all-push: buildx-push buildx-debug-push ## Build and push multi-platform regular and debug images

.PHONY: clean
clean: ## Remove local regular and debug images
	-docker rmi $(IMAGE_REGULAR_TAG) $(IMAGE_REGULAR_ALIAS) $(IMAGE_DEBUG_TAG) $(IMAGE_DEBUG_ALIAS)
