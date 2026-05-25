package greeter_test

import (
	"fmt"

	"github.com/servercurio/go-library-starter/pkg/greeter"
)

func ExampleNew() {
	g := greeter.New(greeter.Options{})
	out, _ := g.Greet("World")
	fmt.Println(out)
	// Output: Hello, World!
}

func ExampleNew_customSalutation() {
	g := greeter.New(greeter.Options{Salutation: "Bonjour"})
	out, _ := g.Greet("Marie")
	fmt.Println(out)
	// Output: Bonjour, Marie!
}

func ExampleGreeter_Greet() {
	g := greeter.New(greeter.Options{Punctuation: "."})
	out, _ := g.Greet("World")
	fmt.Println(out)
	// Output: Hello, World.
}
