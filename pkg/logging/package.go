// Package logging exposes an opinionated structured logger on top of
// zerolog. Initialize swaps in the package-wide Default logger from a Config;
// AsStdLogger adapts a zerolog.Logger so it can be passed to APIs that want
// the standard library *log.Logger.
package logging
