package main

import (
	"fmt"
	"github.com/timwehrle/pryl"
	"github.com/timwehrle/pryl/internal/style"
)

func main() {
	p := pryl.New()

	name := p.Input("What is your name?", "John Doe")

	customStyle := style.NewStyle(
		style.WithQuestion("!"),
	)

	age := p.WithStyle(customStyle).Input("Whats your age?", "25")

	fmt.Printf("Hello %s, you are %s years old\n", name, age)
}
