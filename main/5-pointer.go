package main

import "fmt"

// Go do not have `class` keyword, instead, it uses `type` to define classes/structures
type Vertex struct {
	X int
	Y int
}

func main() {
	// The type *T is a pointer to a T value. Its zero value is nil.
	var p0 *int
	fmt.Println(p0)

	// Unlike C, Go has no pointer arithmetic. (e.g do not support pointer++)
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// `struct` is like classes in other languages, it has properties
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Struct fields can be accessed through a struct pointer.
	p2 := &v
	p2.X = 1e9
	fmt.Println(v)

	var (
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p3 = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(p3, v2, v3)

	// An array's length is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// A slice is formed by specifying two indices, which includes the first element, but excludes the last one
	var s []int = primes[1:4]
	fmt.Println(s)

	// Slices are like references to arrays, they do not store any data
	s[0] = 1
	fmt.Println(primes)

	// This creates an array, then builds a slice that references it:
	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)

	// These slice expressions are equivalent
	fmt.Println(primes[0:6])
	fmt.Println(primes[0:])
	fmt.Println(primes[:6])
	fmt.Println(primes[:])
}
