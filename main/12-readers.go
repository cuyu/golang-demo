package main

import (
	"golang.org/x/tour/reader"
	"io"
	"os"
	"strings"
)

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	copy(b, "A")
	return 1, nil
}

// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	length, err := r.r.Read(b)
	for i := range b {
		switch {
		case b[i] < 'n' && b[i] >= 'a':
			b[i] += 13
		case b[i] >= 'n' && b[i] <= 'z':
			b[i] -= 13
		case b[i] < 'N' && b[i] >= 'A':
			b[i] += 13
		case b[i] >= 'N' && b[i] <= 'Z':
			b[i] -= 13
		}
	}
	return length, err
}

func main() {
	reader.Validate(MyReader{})
	// Another example
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
