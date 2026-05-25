# Key Build Commands

- `task` — Default pipeline: vendor → lint → test.
- `task lint` — Runs `go fmt`, a strict `go vet` (atomic, defers, assign, bools, buildtag, framepointer, lostcancel, loopclosure, nilfunc, shift, stdmethods, stringintconv, structtag), and `golangci-lint run` against `.golangci.yml`. The `golangci-lint` binary is installed via `go install` on the fly (pinned version inside the task), so no system-level prerequisite.
- `task test` — Runs `go test -parallel 4 -cover -coverprofile cover.out -race -v ./...` (with `CGO_ENABLED=1` for the race detector). The `cover.out` profile is written into the repo root and removed by `task clean`.
- `task vendor` — Runs `go mod tidy` followed by `go mod vendor`. Use after any `go.mod` change.
- `task generate` — Runs `go generate ./...`. Required after a fresh clone (and before `task test`/`task sbom`) because `pkg/version/commit.txt` is gitignored but is `//go:embed`-ed.
- `task build` — Runs `go build ./...` (depends on `task generate`). No binary output — this is a library — but it validates that every package in the module compiles end-to-end, including `examples/quickstart`. Used by the CI "Code Compiles" check.
- `task sbom` — Generates CycloneDX 1.5 source-module SBOMs at `sbom.json` and `sbom.xml` via `cyclonedx-gomod mod`, then GPG-signs each output (`sbom.{json,xml}.asc`). Listed in `.releaserc.json`'s GitHub assets and run by `publishCmd`, so each release ships both flavours plus their signatures.
- `task clean` — Removes `cover.out`, the SBOM artefacts, and `pkg/version/commit.txt`.
- `task clean:cache` — Calls `go clean -cache -testcache`.

There are no binary, container, or run-the-app tasks — this is a library starter and the release pipeline does not produce executables.

## Running the quickstart consumer

`examples/quickstart/main.go` is a tiny `package main` that imports the starter's packages so you can confirm an end-to-end build works:

```
go run ./examples/quickstart
APP_LOG_LEVEL=debug go run ./examples/quickstart
```

It is built by `go build ./...` (and therefore by the CI `Compile Code` step) but not invoked by any task target.
