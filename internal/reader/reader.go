package reader

import (
	"bufio"
	"os"
)

type DefaultReader struct {
	*bufio.Reader
}

func NewDefaultReader() *DefaultReader {
	return &DefaultReader{
		Reader: bufio.NewReader(os.Stdin),
	}
}

func (r *DefaultReader) ReadLine() (string, error) {
	return r.ReadString('\n')
}
