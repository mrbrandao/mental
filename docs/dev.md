# Developer Guide

## Prerequisites

| Tool | Required | Install |
|---|---|---|
| Go 1.25+ | yes | https://go.dev/dl |
| pre-commit | yes | https://pre-commit.com |
| golangci-lint | yes | `make dev-deps` |
| snyk | for pre-commit | `npm install -g snyk` |
| podman | for containers | https://podman.io |
| goreleaser | for releases | https://goreleaser.com |

## Setup

```bash
git clone https://github.com/mrbrandao/ais.git
cd ais
make dev-deps   # installs golangci-lint
make hooks      # installs pre-commit hooks
```

## Daily workflow

```bash
make          # build bin/ais
make test     # run tests
make lint     # run linter
make fmt      # format code
make coverage # coverage report
```

## Pre-commit hooks

Hooks run automatically on `git commit`:

| Hook | Purpose |
|---|---|
| gitleaks | secret scanning |
| go-fmt | formatting |
| go-vet | static analysis |
| golangci-lint | linting |
| gosec | SAST security scan |
| snyk | vulnerability scan |

snyk requires a separate install:
```bash
npm install -g snyk
snyk auth      # authenticate once
```

## Containers

```bash
# Build container image
make container-build

# Extract binary without Go installed
make container-binary  # produces bin/ais

# Run ais via container
make container-run ARGS="search -a opencode -s topic"
```

## Releases

Releases are automated via goreleaser on tag push:

```bash
git tag v0.1.0
git push origin v0.1.0
```

GitHub Actions runs goreleaser and uploads binaries
to the GitHub Release page.

Dry-run locally:
```bash
make release-dry
```

## Adding a provider

See `AGENTS.md` — "How to add a new assistant backend".
