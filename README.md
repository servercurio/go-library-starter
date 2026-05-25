<p align="center">
  <img src="docs/logo.svg" alt="Server Curio — Project Templates" width="600">
</p>

# go-library-starter

An opinionated Go starter template for **reusable libraries**. It ships a curated set of public packages under `pkg/`, a tiny exemplar that demonstrates the canonical package shape, and a hardened release pipeline that produces a signed git tag plus a signed CycloneDX SBOM — no binaries, no containers.

[![Go Reference](https://pkg.go.dev/badge/github.com/servercurio/go-library-starter.svg)](https://pkg.go.dev/github.com/servercurio/go-library-starter)

## Install

```
go get github.com/servercurio/go-library-starter@latest
```

## Quick start

```go
package main

import (
    "github.com/servercurio/go-library-starter/pkg/greeter"
    "github.com/servercurio/go-library-starter/pkg/logging"
    "github.com/servercurio/go-library-starter/pkg/version"
)

func main() {
    logging.Initialize(logging.NewConfigFromEnv("APP"))

    out, _ := greeter.New(greeter.Options{}).Greet("World")

    logging.Default.Info().
        Str("version", version.Number()).
        Str("greeting", out).
        Msg("quickstart ready")
}
```

The full runnable version lives at [`examples/quickstart/main.go`](examples/quickstart/main.go) and can be exercised with `go run ./examples/quickstart`.

## Packages

| Package | What it provides |
|---|---|
| [`pkg/errors`](pkg/errors) | `joomcode/errorx` namespaces (`FileSystemErrors`, `PoolErrors`); the canonical pattern for defining error categories. |
| [`pkg/version`](pkg/version) | Build-time version metadata (commit hash, semver) via `//go:embed`. Populated by `task generate`. |
| [`pkg/obfusicate`](pkg/obfusicate) | `ConcealPrefix(value, n)` — small string-redaction helper for log output. Spelling is intentional. |
| [`pkg/logging`](pkg/logging) | Opinionated `rs/zerolog` wrapper exposing the package-wide `Default` logger plus an `AsStdLogger` adapter. |
| [`pkg/env`](pkg/env) | Type-safe environment-variable parsers (string, bool, int, uint16, float, duration) and an `AddPrefix` helper. |
| [`pkg/config`](pkg/config) | YAML/JSON configuration-file loader with multi-path search. Consumers supply the schema. |
| [`pkg/health`](pkg/health) | In-process health-check `Registry` and `Report` model (Spring Boot Actuator-style). |
| [`pkg/pool`](pkg/pool) | `panjf2000/ants/v2` goroutine pool wrapper with cancellation-aware `SubmitWithContext`, stats, and graceful release. |
| [`pkg/greeter`](pkg/greeter) | Canonical exemplar — interface + `Options` + `New(opts)` + `errorx` namespace + table-driven tests + runnable `Example*` funcs. Model new packages after it. |

## Project layout

```
.
├── pkg/             # public library surface (see table above)
├── examples/
│   └── quickstart/  # tiny package main showing downstream consumption
├── docs/            # README assets
└── .github/         # workflows + governance
```

There is no `cmd/` directory and no `package main` outside `examples/`. This module exists to be imported.

## Working with the repo

```
task vendor     # go mod tidy + go mod vendor
task generate   # populates pkg/version/commit.txt (gitignored)
task lint       # go fmt + go vet + golangci-lint
task test       # go test -race -cover across every package
task sbom       # CycloneDX JSON + XML SBOMs of the source module, GPG-signed
```

`task` with no argument runs `vendor → lint → test`.

## Releases

Releases are driven by [semantic-release](https://github.com/semantic-release/semantic-release) using conventional commits. Each release publishes:

| Artefact | Path | Notes |
|---|---|---|
| Signed git tag | `vX.Y.Z` | Tag commit is GPG-signed by the release pipeline. |
| Release notes | GitHub release body | Generated from conventional-commit messages. |
| Source SBOM | `sbom.json`, `sbom.xml` | CycloneDX 1.5; covers the Go module dependency tree. |
| SBOM signatures | `sbom.json.asc`, `sbom.xml.asc` | Detached GPG signatures of each SBOM. |
| Sigstore attestations | GitHub attestation log | `actions/attest-build-provenance` + `actions/attest-sbom` against each SBOM. |

The signing model is: **the SBOM is a manifest, so it is signed directly.** There are no binaries to hash; if a downstream fork adds one, sign the hash file (`<binary>.sha256.asc`) and never the binary itself.

### Verifying a release

```
gh attestation verify sbom.json --owner servercurio
gpg --verify sbom.json.asc sbom.json
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Commits must be GPG-signed and carry a DCO `Signed-off-by:` trailer; the `.git/hooks/prepare-commit-msg` documented in [.claude/git-hooks.md](.claude/git-hooks.md) appends the trailer automatically once installed in each clone.

## License

[Apache-2.0](LICENSE).
