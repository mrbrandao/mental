SHELL := /bin/bash
.DEFAULT_GOAL := build

.PHONY: help
help: ## - print help and usage
	@printf "mental — cross-session memory and AI session manager\n\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' \
		$(MAKEFILE_LIST) | \
		sed 's/^[^:]*://' | \
		awk 'BEGIN {FS = ":.*?## "}; \
		{printf "\033[36m%-20s\033[0m %s\n", \
		$$1, $$2}'

include mk/go.mk
include mk/install.mk
include mk/release.mk
include mk/container.mk
include mk/dev.mk
