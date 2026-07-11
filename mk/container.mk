IMAGE ?= ghcr.io/mrbrandao/ais
TAG   ?= latest

.PHONY: container-build container-binary \
        container-run container-push

container-build: ## - build container image
	podman build -t $(IMAGE):$(TAG) .

container-binary: ## - extract binary (no Go needed)
	podman build --target builder \
		-t ais-builder .
	podman run --rm \
		-v $(PWD)/bin:/out:Z \
		ais-builder cp /ais /out/ais

container-run: ## - run ais via container
	podman run --rm \
		-v $(HOME)/.local/share:/data:ro:Z \
		$(IMAGE):$(TAG) $(ARGS)

container-push: ## - push image to registry
	podman push $(IMAGE):$(TAG)
