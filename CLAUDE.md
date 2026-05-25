# CLAUDE.md

Guidance for Claude Code working in this repository. This file covers project intent, conventions, and procedural guidance that isn't captured in the agent reference docs:

- [`.claude/instructions.md`](.claude/instructions.md) — tech stack, personality, requirements
- [`.claude/build-commands.md`](.claude/build-commands.md) — `task` targets and what they run
- [`.claude/module-structure.md`](.claude/module-structure.md) — directory-by-directory roles
- [`.claude/conventions.md`](.claude/conventions.md) — coding conventions for this repo
- [`.claude/git-hooks.md`](.claude/git-hooks.md) — required local git hooks (must be installed per clone)

@.claude/instructions.md
@.claude/build-commands.md
@.claude/module-structure.md
@.claude/conventions.md
@.claude/git-hooks.md

## What this project is

A **starter template** for reusable Go libraries — not an application, not a CLI, not a daemon. It provides:

- a curated set of eight public packages under `pkg/` (errors, version, obfusicate, logging, env, config, health, pool),
- a single exemplar package (`pkg/greeter`) that shows the canonical shape every new package should follow,
- an `examples/quickstart/main.go` that consumes the library exactly as a downstream importer would,
- a release pipeline that ships a signed git tag, release notes, and a signed CycloneDX SBOM (no binaries, no containers).

When asked to add functionality, prefer keeping the codebase generic and composable. Don't introduce app-specific concerns (a particular database driver, a specific business model, an auth scheme) unless explicitly requested.

## When adding code

- **New public package**: add it under `pkg/<name>/` and mirror `pkg/greeter`. Required files: `package.go` (or `doc.go`) opening with `// Package <name> ...`; one or more implementation files; `*_test.go` with table-driven tests; `example_test.go` (package suffix `_test`) with at least one runnable `Example*` function whose `// Output:` block doubles as a smoke test and as the godoc snippet.
- **New exported function or type**: requires (1) a godoc comment beginning with the identifier name, (2) at least one runnable `Example*` function, (3) test coverage via a table-driven test.
- **New error category**: define a fresh `errorx.Namespace` either inside the owning package (see `pkg/greeter.GreeterErrors`) or in `pkg/errors/` for cross-package error families. Don't introduce ad-hoc `errors.New` / `fmt.Errorf` for new categories.
- **New config field**: add to the appropriate struct's `Config`, give it a sensible default in the package's `DefaultXxxConfig`, wire env-var loading using the helpers in `pkg/env`, and extend `Validate() error` to reject obviously-bad values. The consumer's `Validate` calls bubble back to them — there is no `Application.Configure` to short-circuit boot.
- **New CI workflow**: file under `.github/workflows/` following the naming convention in `.github/workflows/docs/naming-standards.md` (`ddd-xxxx-name.yaml` file, matching `ddd: [XXXX] Name` workflow `name:`). PR-triggered workflows use the **200** prefix; main-branch push-triggered workflows use **300**; reusable workflows use **800**. New `go:generate` directives don't need a workflow change — `task generate` is invoked by the PR reusable workflow and by the release flow, and picks them up via `./...`.
- **New release behaviour**: edit `.releaserc.json` for semantic-release plugin config (commit-analyzer rules, release notes preset, exec hooks, branch channels, GitHub asset list). The asset list ships only `sbom.{json,xml}{,.asc}`; don't add binary asset entries.

## Things to leave alone unless asked

- This repo uses `pkg/env` + `pkg/config`, **not Viper**. Don't pull Viper in.
- **ants/v2** is the only goroutine-pool library. Don't add a second.
- `pkg/obfusicate/` exposes one tiny helper (`ConcealPrefix`) used for log redaction — leave it alone unless the task is explicitly about redaction; the misspelled package name is intentional and matches the import path.
- The `.git/hooks/prepare-commit-msg` DCO sign-off hook is **per-clone** and **not version-controlled**. Document it in `.claude/git-hooks.md`; don't try to commit the hook itself or switch to `core.hooksPath`.
- The release flow's signing model is "sign manifests (the SBOM) directly; sign hashes for any future binaries". Don't add bare `<binary>.asc` entries to the release asset list.
