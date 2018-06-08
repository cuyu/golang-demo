package main

import (
	"fmt"
	"runtime"
)

func main() {
	// if statement can start with a short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if.
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
}
