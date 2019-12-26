package file

import (
	"io"
	"io/ioutil"
	"os"
)

type (
	File interface {
		Read(path string) ([]byte, error)
	}

	file struct{}
)

func New() *file {
	return &file{}
}

func (f *file) Read(path string) ([]byte, error) {
	var err error
	input := io.ReadCloser(os.Stdin)

	if input, err = os.Open(path); err != nil {
		return nil, err
	}

	defer input.Close()

	return ioutil.ReadAll(input)
}
