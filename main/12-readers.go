package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	copy(b, "A")
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
