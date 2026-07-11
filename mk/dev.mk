.PHONY: dev-deps hooks pre-commit clean

dev-deps: ## - install local dev tools
	go install \
		github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Install snyk: npm install -g snyk"
	@echo "  or: brew install snyk"

hooks: ## - install pre-commit hooks
	pre-commit install

pre-commit: ## - run pre-commit on all files
	pre-commit run --all-files

clean: ## - remove build artifacts
	rm -rf bin/ coverage.out
