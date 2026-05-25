# Contributing

Thanks for your interest in `go-library-starter`. The repository is a starter template, so contributions are oriented toward keeping the scaffold clean, generic, and reusable rather than expanding feature scope.

## Before you start

- **Open an issue first** for anything beyond a typo or one-line bug fix. Discussing scope and approach saves rework.
- **Don't add downstream concerns.** Persistence drivers, auth schemes, validation libraries, and domain models are intentionally out of scope. The template stays generic so consumers can layer their own choices.

## Local setup

```sh
git clone https://github.com/servercurio/go-library-starter.git
cd go-library-starter
task vendor
task
```

`task` is the canonical build runner — `task` with no argument runs `vendor → lint → test`. See `.claude/build-commands.md` for the full target list. Don't invoke `go build` / `go test` directly except for one-off debugging — go through the Taskfile so CI and local builds stay aligned.

## Required local git hook

Every clone needs a per-clone `prepare-commit-msg` hook that auto-appends a DCO `Signed-off-by:` line. The hook is **not** version-controlled; install it manually after cloning. Contents and rationale live in [`.claude/git-hooks.md`](.claude/git-hooks.md).

```sh
ls -l .git/hooks/prepare-commit-msg
```

If the hook is missing, copy the script from `.claude/git-hooks.md` and `chmod +x` it.

## Required commit signing

Every contributor commit must be **GPG-signed** in addition to carrying the DCO `Signed-off-by:` trailer. The release workflow already produces signed `chore(release)` commits via an imported key (`.github/workflows/800-call-semantic-release.yaml`); contributor commits must match so the entire `main` history verifies.

Configure once per clone (or globally):

```sh
git config commit.gpgsign true
git config user.signingkey <YOUR-GPG-KEY-ID>
git config tag.gpgsign true
```

Verify a commit you just authored:

```sh
git log -1 --show-signature
```

Don't bypass either requirement with `--no-gpg-sign` or `--no-verify`. If `gpg` prompts you for a passphrase on every commit, configure `gpg-agent` rather than disabling signing.

## Commit-message conventions

This repo uses [Conventional Commits](https://www.conventionalcommits.org). The release flow (semantic-release) reads commit types to decide version bumps:

| Type                | Bump  | Example                                 |
|---------------------|-------|-----------------------------------------|
| `feat`              | minor | `feat(greeter): support custom locale`  |
| `fix`               | patch | `fix(pool): drain on cancelled context` |
| `refactor`, `build` | patch | `refactor(env): consolidate parsers`    |
| breaking change     | minor | `feat!: rename APP_LOG_LEVEL` (or `BREAKING CHANGE:` footer) |
| `chore`, `ci`, `docs`, `style`, `test` | none | `docs: clarify pool sizing default` |

Subject line: imperative mood, no trailing period, ≤72 chars. The `prepare-commit-msg` hook will append your DCO line automatically.

## What every PR must include

For any change that adds or modifies exported library surface:

- a godoc comment on every new exported identifier (function, type, variable, constant);
- at least one runnable `Example*` function (in `example_test.go`, package suffix `_test`) with an `// Output:` block;
- a table-driven test (`t.Run` per case) covering the new behaviour;
- a corresponding doc-block update in the package's `package.go` if the package's contract is changing.

The exemplar in [`pkg/greeter`](pkg/greeter) is the canonical reference — mirror its shape.

## Pull request workflow

1. Branch from `main` (or a `release/X.Y` branch for backports).
2. Run `task` locally before pushing — vendor → lint → test.
3. Run `go build ./...` to confirm `examples/quickstart` still compiles.
4. Open the PR. The "PR Checks" workflow runs `task lint`, `task test`, and `govulncheck`; the "PR Formatting" workflow validates the title against Conventional Commits.
5. Address review feedback in additional commits — don't force-push during review unless asked.
6. Squash on merge is fine; the squash subject must still be a valid Conventional Commit.

## Code conventions

See [`.claude/conventions.md`](.claude/conventions.md). Highlights:

- **Library-first design.** No `main`, no Cobra, no daemon. The only `package main` lives at `examples/quickstart/main.go` and exists solely to demonstrate consumption.
- **Public packages live under `pkg/`.** Mirror the shape of `pkg/greeter`.
- **Options-struct constructors.** `New(opts Options)` with documented defaults inside `New`.
- **Errors via `joomcode/errorx` namespaces** (see `pkg/errors/`, `pkg/greeter.GreeterErrors`).
- **ants/v2 only** for goroutine pools — use `pool.New(...)`; don't reach for `ants` directly.

## Reporting security issues

Please **don't** file public issues for vulnerabilities. See [`SECURITY.md`](SECURITY.md) for the private reporting flow.
