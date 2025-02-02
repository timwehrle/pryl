package pryl

import (
	"fmt"
	"github.com/timwehrle/pryl/internal/reader"
	"github.com/timwehrle/pryl/internal/style"
)

type Prompt struct {
	style    *style.Style
	reader   Reader
	validate func(string) error
}

type Reader interface {
	ReadLine() (string, error)
}

func New() *Prompt {
	return &Prompt{
		style:  style.Default(),
		reader: reader.NewDefaultReader(),
	}
}

func (p *Prompt) WithStyle(s *style.Style) *Prompt {
	p.style = s
	return p
}

func (p *Prompt) WithValidation(fn func(string) error) *Prompt {
	p.validate = fn
	return p
}

func (p *Prompt) clearLine() {
	fmt.Print("\r\033[K")
}

func (p *Prompt) moveUpOneLine() {
	fmt.Print("\033[A")
}
