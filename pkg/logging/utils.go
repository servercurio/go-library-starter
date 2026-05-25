package logging

import (
	"log"

	"github.com/rs/zerolog"
)

// AsStdLogger wraps a zerolog.Logger so it can be passed to APIs that expect
// the standard library *log.Logger. Each Write call from the std logger is
// emitted as a single zerolog event at the wrapped logger's current level.
func AsStdLogger(logger zerolog.Logger) *log.Logger {
	return log.New(logger, "", 0)
}
