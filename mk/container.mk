IMAGE ?= ghcr.io/mrbrandao/mental
TAG   ?= latest

.PHONY: container-build container-binary \
        container-run container-push

container-build: ## - build container image
	podman build -t $(IMAGE):$(TAG) .

container-binary: ## - extract binary (no Go needed)
	podman build --target builder \
		-t mental-builder .
	podman run --rm \
		-v $(PWD)/bin:/out:Z \
		mental-builder cp /mental /out/mental

container-run: ## - run mental via container
	podman run --rm \
		-v $(HOME)/.local/share:/data:ro:Z \
		$(IMAGE):$(TAG) $(ARGS)

container-push: ## - push image to registry
	podman push $(IMAGE):$(TAG)
