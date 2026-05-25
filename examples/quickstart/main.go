// Quickstart demonstrates how a downstream consumer wires the starter's
// packages together: hydrate a logging.Config from env, initialize the
// default logger, and emit one structured event built from pkg/greeter.
//
// Run with:
//
//	go run ./examples/quickstart
//	APP_LOG_LEVEL=debug go run ./examples/quickstart
package main

import (
	"os"

	"github.com/servercurio/go-library-starter/pkg/greeter"
	"github.com/servercurio/go-library-starter/pkg/logging"
	"github.com/servercurio/go-library-starter/pkg/version"
)

func main() {
	logging.Initialize(logging.NewConfigFromEnv("APP"))

	out, err := greeter.New(greeter.Options{}).Greet("World")
	if err != nil {
		logging.Default.Error().Err(err).Msg("greet failed")
		os.Exit(1)
	}

	logging.Default.Info().
		Str("version", version.Number()).
		Str("commit", version.Commit()).
		Str("greeting", out).
		Msg("quickstart ready")
}
