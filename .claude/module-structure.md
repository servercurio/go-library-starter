# Module Structure

Public surface lives under `pkg/`. Each package is exported (no `internal/` segment) so downstream consumers can `go get` and import directly.

- `pkg/errors/` — `joomcode/errorx` namespaces (`FileSystemErrors`, `PoolErrors`); the canonical pattern for defining error categories. Add a new namespace per concern; don't introduce ad-hoc `errors.New` / `fmt.Errorf` for new categories.
- `pkg/version/` — Build-time version metadata (commit, semver, tag) using `Masterminds/semver`. `commit.txt` is gitignored and populated by `go generate`; `version.txt` is content-tracked and bumped by semantic-release's `prepareCmd`.
- `pkg/obfusicate/` — Small string-redaction helper exposing `ConcealPrefix` for masking values in log output. Spelling is intentional — keep it as-is.
- `pkg/logging/` — One named zerolog logger (`Default`) initialized from a `Config` plus a `*log.Logger` adapter (`AsStdLogger`). The consumer is responsible for invoking `logging.Initialize(cfg)` before the first `Default.*` call.
- `pkg/env/` — Type-safe environment-variable parsers (string, bool, int, uint16, float, duration) plus the `AddPrefix` helper. Consumers thread their own application-specific prefix; the `APP` literal seen in tests is illustrative.
- `pkg/config/` — YAML/JSON config-file loader with multi-path search. Consumers supply a list of search paths and a base name; `FromPaths` walks them and unmarshals the first match into the caller's struct.
- `pkg/health/` — Per-component health-check registry and the `Report` model (Spring Boot Actuator-style). The `Registry` is an explicit dependency; consumers construct one and pass it via DI.
- `pkg/pool/` — Goroutine-pool subsystem wrapping `github.com/panjf2000/ants/v2`. Exposes `New`, `Submit`, `SubmitWithContext` (cancellation-aware), `Stats`, `Release`/`ReleaseTimeout`. Submit errors come pre-wrapped in the `PoolErrors` errorx namespace.
- `pkg/greeter/` — Tiny exemplar package showing the canonical library shape: `package.go` doc, exported interface (`Greeter`), unexported default implementation, value-typed `Options`, `New(opts Options) Greeter` constructor, `GreeterErrors` `errorx` namespace, table-driven tests, runnable `Example*` funcs.

## Sibling directories

- `examples/quickstart/` — a `package main` that imports `pkg/greeter`, `pkg/logging`, and `pkg/version` to show what downstream consumption looks like. Builds with `go build ./...` so CI's `Compile Code` step exercises it.
- `docs/` — Non-Go documentation assets. `logo.svg` is embedded at the top of the README and is not consumed by any Go code.
- `.github/workflows/` — CI workflows, numbered per `.github/workflows/docs/naming-standards.md`. The 200/300-series flows run on PRs and `main` respectively; both delegate to 800-series reusables: `800-call-code-compiles.yaml`, `800-call-unit-test.yaml`, `800-call-vulncheck.yaml`. Releases run via `800-call-semantic-release.yaml` (tag + notes + signed source SBOM). Semantic-release plugin config lives in `.releaserc.json` at the repo root.

There is no `cmd/` directory: this module has no `main` package shipped under `cmd/`. The only `package main` lives at `examples/quickstart/main.go` and exists solely to demonstrate consumption.
