package main

import (
	"fmt"
	"strconv"
	"time"
)

// The error type is a built-in interface similar to fmt.Stringer:
// type error interface {
//     Error() string
// }
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Another example
	i, err := strconv.Atoi("4Q")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}
