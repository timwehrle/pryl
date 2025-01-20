package pryl

import (
	"fmt"
	"strings"
)

const (
	clearLine     = "\r\033[K" // Clear the current line
	moveUpOneLine = "\033[A"   // Move cursor up one line
)

func (p *Prompt) Input(label string, placeholder string) string {
	if placeholder != "" {
		fmt.Printf("%s %s\n%s\n",
			p.style.Question,
			label,
			fmt.Sprintf(p.style.Placeholder, placeholder),
		)

		p.moveUpOneLine()
		p.clearLine()
	} else {
		fmt.Printf("%s %s ", p.style.Question, label)
	}

	fmt.Print(p.style.Cursor + " ")
	input, _ := p.reader.ReadLine()
	result := strings.TrimSpace(input)

	fmt.Printf("%s %s %s%s%s\n",
		p.style.Question,
		label,
		p.style.SuccessColor,
		result,
		"\033[0m",
	)

	return result
}
