package pryl

import (
	"fmt"
	"strings"
)

func (p *Prompt) Input(label string, placeholder string) string {
	fmt.Printf("%s %s ", p.style.Question, label)
	if placeholder != "" {
		fmt.Printf(p.style.Placeholder+"\n", placeholder)
	}
	fmt.Print(p.style.Cursor + " ")

	input, _ := p.reader.ReadLine()
	return strings.TrimSpace(input)
}
