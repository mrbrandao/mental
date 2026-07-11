# AGENTS.md

Guidance for coding agents and contributors
working in this repository.

## What is ais

`ais` is an AI session manager CLI that searches
OpenCode sessions from the terminal. The provider
architecture is designed to add more assistants
in future releases.

Binary: `ais` | Module: `github.com/mrbrandao/ais`

## Architecture

Commands call a `provider.Provider` interface.
Each assistant backend is an isolated package.
Output formatting is independent of both.

```
cmd/           Cobra commands — thin, no logic
internal/
  model/       Pure data types (Session, Query)
  provider/    Provider interface + per-assistant pkg
    opencode/  OpenCode SQLite backend
  output/      Formatters: table (pterm), json, plain
```

## How to add a new assistant backend

1. Create `internal/provider/<name>/` package
2. Implement the `Provider` interface:
   ```go
   func (p *Provider) Name() string
   func (p *Provider) Search(
       ctx context.Context,
       q model.Query,
   ) ([]model.Session, error)
   ```
3. Register in `cmd/search.go` `resolveProvider()`
4. Add table-driven tests in `<name>_test.go`

## How to add a new command

1. Create `cmd/<command>.go`
2. Define a `cobra.Command`, register with
   `rootCmd.AddCommand`
3. All logic goes in `internal/` — commands only
   parse flags and call domain functions

## How to add a new output format

1. Add a struct implementing `output.Formatter`
2. Add a case in `output.New(format string)`

Both changes live in `internal/output/output.go`.

## Build and test

```bash
make            # build bin/ais
make test       # go test -race ./...
make lint       # golangci-lint run ./...
make coverage   # coverage report
make install    # install to /usr/local/bin
sudo make install        # system-wide
PREFIX=~/.local make install  # user-local
```

See `docs/dev.md` for full developer setup.

## Code standards

- Go 1.25, follow https://go.dev/doc/effective_go
- 80 characters per line — hard wrap
- Errors last return value; wrap with fmt.Errorf
- context.Context always first param
- defer for all cleanup (db.Close, rows.Close)
- No else after return
- Table-driven tests in `_test.go` files
- No CGO — use modernc.org/sqlite for SQLite

## Commit rules (tpope)

- Conventional prefix: feat/fix/docs/test/chore/ci
- Subject: imperative mood, ≤50 chars, no period
- Body: wrapped at 72 chars, explain what and why
- Commit after ≤150 lines changed
- Never `git add .` — always explicit file paths

## Security — NEVER include in any file or commit

- Secrets, tokens, API keys, passwords
- Local filesystem paths revealing real environments
  (use /path/to/... or ~/.config/... as examples)
- Internal hostnames, IPs, org-internal URLs
- Real session content containing private data
- Personal information beyond public git metadata
- Any data identifying a real private environment

All examples must use generic, public-safe values.
Only commit content safe to publish publicly.
