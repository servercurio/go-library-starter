package greeter

import (
	"fmt"
	"strings"

	ex "github.com/joomcode/errorx"
)

// GreeterErrors is the errorx namespace for failures raised by this
// package. Downstream consumers can branch on it with errorx.IsOfType.
var GreeterErrors = ex.NewNamespace("greeter")

// ErrEmptyName is returned by Greeter.Greet when the supplied name is
// blank or contains only whitespace.
var ErrEmptyName = GreeterErrors.NewType("empty_name")

// Greeter renders a greeting for a given name.
type Greeter interface {
	Greet(name string) (string, error)
}

// Options configures a Greeter constructed by New. The zero value is
// valid and produces "Hello, <name>!".
type Options struct {
	// Salutation is the leading word of the greeting. Defaults to "Hello".
	Salutation string

	// Punctuation is the trailing character. Defaults to "!".
	Punctuation string
}

// New returns a Greeter configured by opts. Empty Options fields fall
// back to the documented defaults.
func New(opts Options) Greeter {
	if strings.TrimSpace(opts.Salutation) == "" {
		opts.Salutation = "Hello"
	}
	if opts.Punctuation == "" {
		opts.Punctuation = "!"
	}
	return &defaultGreeter{opts: opts}
}

type defaultGreeter struct {
	opts Options
}

// Greet returns the formatted greeting for name, or ErrEmptyName when
// name is blank.
func (g *defaultGreeter) Greet(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", ErrEmptyName.New("name must not be empty")
	}
	return fmt.Sprintf("%s, %s%s", g.opts.Salutation, name, g.opts.Punctuation), nil
}
