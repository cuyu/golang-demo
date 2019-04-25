package main

import (
	"fmt"
	"math"
)

// Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// In this example, the Abs method has a receiver of type Point named v.
type Point struct {
	X, Y float64
}

func (v Point) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Methods with pointer receivers can modify the value to which the receiver points. (i.e. Value receivers cannot modify the value)
func (v *Point) Scale(factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

// The same effect but do not bind to the type `Point`
func ScalePointer(v *Point, factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

// This function will not modify the `v` (The value is copied into the function when called)
func ScaleValue(v Point, factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

func main() {
	v := Point{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// As a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
	// i.e. there is implicit conversion for methods of type (value => pointer or vice versa)
	v.Scale(5)
	fmt.Println(v)

	ScalePointer(&v, 0.1)
	fmt.Println(v)

	ScaleValue(v, 2)
	fmt.Println(v)
	// There are two reasons to use a pointer receiver.
	//The first is so that the method can modify the value that its receiver points to.
	//The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
	//In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)
}
