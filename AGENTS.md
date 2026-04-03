# AGENTS.md

## Project

go-aula is a Go client for Denmark's Aula school communication platform. Ported from the Rust [aulalibre](https://github.com/eisbaw/aulalibre) project. The original Rust source is at `C:\src\aulalibre` for cross-referencing.

## Structure

- `cmd/aula-cli/` — Cobra CLI binary
- `internal/aulaapi/` — API client library (auth, client, session, tokenstore, error, response)
- `internal/aulaapi/enums/` — String-based enums (`type Foo string` with const blocks)
- `internal/aulaapi/models/` — Data model structs
- `internal/aulaapi/services/` — API service functions (one per domain)
- `internal/cli/` — CLI support (config, output)
- `internal/cli/commands/` — Cobra command implementations

## Key design decisions

- Enums use `type Foo string` with const blocks (not iota) — marshals to JSON for free
- HTTP: stdlib `net/http` + cookiejar (no external HTTP client)
- Generics: `SessionGet[T]()` etc. for typed response envelope parsing
- Access token goes as `?access_token=<jwt>` query param (not Bearer header)
- CSRF: reads `Csrfp-Token` cookie, injects as `csrfp-token` header

## Code style

- Use `*string` / `*int64` for optional JSON fields (matching the existing model pattern)
- Use `json:"camelCase,omitempty"` tags on optional fields
- Services: one file per API domain in `internal/aulaapi/services/`
- CLI commands: one file per command group in `internal/cli/commands/`
- Error handling: return errors, don't panic. Use `fmt.Fprintf(os.Stderr, ...)` + `os.Exit(1)` in CLI commands
- Table output: use `cli.NewTable` / `cli.PrintRow` from `internal/cli/output.go`
- HTML content: use `cli.StripHTMLTags()` for terminal display

## Build & test

```bash
make build      # go build ./cmd/aula-cli
make test       # go test ./...
make lint       # golangci-lint run ./...
make fmt        # gofmt + goimports
```

## Git rules

- Do not commit or push unless explicitly told to
- Do not include `Co-Authored-By: Claude` in commit messages
