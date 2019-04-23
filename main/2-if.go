package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// if statement can start with a short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if.
	// Here is the use case which `x:=y` cannot be replaced by `var x=y`
	a := 4
	if v := a * a; v < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println(v)
	}
	fmt.Println(v)  // raise exception as v is not defined outside if statement

	// For switch statement, Go only runs the selected case, not all the cases that follow. (i.e. no need to add `break` under each case)
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	// Switch without a condition is the same as `switch true`.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
