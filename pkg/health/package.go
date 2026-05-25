// Package health implements an in-process health-check registry and the
// report model snapshotted by callers (a server's readiness loop, a CLI
// "health" command, or any other consumer).
//
// The model intentionally mirrors what Spring Boot Actuator, Quarkus
// SmallRye Health, and Micronaut return: an overall status plus a map of
// per-component statuses with optional details. Components register
// themselves with a Registry; consumers ask the registry for a Snapshot.
//
// The Registry is an explicit dependency — construct one per consumer and
// pass it via DI rather than relying on a package-level singleton.
package health
