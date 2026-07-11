# Go build targets.
GO_VERSION := $(shell sed -n 's/^go //p' go.mod)
VERSION    := $(shell \
  git describe --tags --exact-match 2>/dev/null \
  || echo "dev")
LDFLAGS    := -s -w \
  -X github.com/mrbrandao/ais/cmd.version=$(VERSION)
GOFLAGS    ?= -trimpath

MIN_COVERAGE ?= 60

.PHONY: build test vet lint fmt tidy \
        coverage coverage-badge

build: ## - build bin/ais binary
	@mkdir -p bin
	go build $(GOFLAGS) \
		-ldflags "$(LDFLAGS)" -o bin/ais .

test: ## - run test suite
	go test -race ./...

vet: ## - run go vet static analysis
	go vet ./...

lint: ## - run golangci-lint
	golangci-lint run ./...

fmt: ## - format code with gofmt
	gofmt -l -w .

tidy: ## - tidy go modules
	GOTOOLCHAIN=go$(GO_VERSION) go mod tidy

coverage: ## - run tests with coverage report
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

coverage-badge: coverage ## - update README.md badge
	$(eval COV := $(shell go tool cover \
		-func=coverage.out | grep ^total | \
		awk '{print $$3}'))
	sed -i \
		"s|coverage-[0-9.]*%25|coverage-$(COV)|g" \
		README.md
	@echo "Coverage: $(COV)"
