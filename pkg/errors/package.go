// Package errors defines the joomcode/errorx namespaces and typed error
// values used across the starter. Centralising them here lets callers branch
// on category (filesystem, pool, …) without string-matching on error
// messages, and gives downstream consumers a single import path for every
// well-known failure mode.
package errors
