package greeter_test

import (
	"testing"

	ex "github.com/joomcode/errorx"
	"github.com/servercurio/go-library-starter/pkg/greeter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGreet(t *testing.T) {
	cases := []struct {
		name string
		opts greeter.Options
		in   string
		want string
	}{
		{
			name: "default options",
			opts: greeter.Options{},
			in:   "World",
			want: "Hello, World!",
		},
		{
			name: "custom salutation",
			opts: greeter.Options{Salutation: "Bonjour"},
			in:   "Marie",
			want: "Bonjour, Marie!",
		},
		{
			name: "custom punctuation",
			opts: greeter.Options{Punctuation: "."},
			in:   "World",
			want: "Hello, World.",
		},
		{
			name: "name is trimmed",
			opts: greeter.Options{},
			in:   "  World  ",
			want: "Hello, World!",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := greeter.New(tc.opts).Greet(tc.in)
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGreet_EmptyName(t *testing.T) {
	cases := []string{"", "   ", "\t\n"}

	for _, in := range cases {
		t.Run("input="+in, func(t *testing.T) {
			_, err := greeter.New(greeter.Options{}).Greet(in)
			require.Error(t, err)
			assert.True(t, ex.IsOfType(err, greeter.ErrEmptyName))
		})
	}
}
