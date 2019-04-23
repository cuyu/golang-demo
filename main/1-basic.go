package main

import (
	"fmt"
	"math"
)

func add(a int, b int) int {
	return a + b
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add2(a, b float32) float32 {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

// A return statement without arguments returns the named return values. This is known as a "naked" return.
func swap2(a, b int) (x, y int) {
	x, y = b, a
	return
}


func main() {
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	name, location := "Prince Oberyn", "Dorne"
	age := 32
	fmt.Printf("%s (%d) of %s", name, age, location)

	fmt.Println(math.Pi)

	fmt.Println(add(1, 2))

	fmt.Println(add2(1.1, 2.2))

	fmt.Println(swap("aaa", "bbb"))

	fmt.Println(swap2(3, 4))

	// The var statement declares a list of variables; as in function argument lists, the type is last.
	var var1, var2 bool
	fmt.Println(var1, var2)
	var1 = true
	fmt.Println(var1, var2)

	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var python, java = false, "no!"
	python = 22  // will raise exception! (Comparing with Python, Golang is more static~)

	// An untyped constant takes the type needed by its context.
	const two = 2
	var floatTwo float32 = two + 0.1
	fmt.Println(floatTwo)
	const intTwo int = 2
	var floatTwo2 float32 = intTwo + 0.1 // raise exception!
}
