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
}
