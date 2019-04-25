package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
type Abser interface {
	MyAbs() float64
}

type InterFloat float64

// This method means type `InterFloat` implements the interface `Abser`, but we don't need to explicitly declare that it does so.
// Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
func (f InterFloat) MyAbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type InterPoint struct {
	X float64
	Y float64
}

func (v *InterPoint) MyAbs() float64 {
	if v == nil {
		fmt.Println("Got <nil>")
		return 0
	} else {
		return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}
}

func (v *InterPoint) String() string {
	return fmt.Sprintf("<X=%v,Y=%v>", v.X, v.Y)
}

func main() {
	var a Abser
	f := InterFloat(-math.Sqrt2)
	v := InterPoint{3, 4}

	a = f  // a InterFloat implements Abser
	a = &v // a *InterPoint implements Abser

	// In the following line, v is a Point (not *Point) and does NOT implement Abser, will raise exception.
	//a = v
	fmt.Println(a.MyAbs())

	// Under the hood, interface values can be thought of as a tuple of a value and a concrete type: (value, type)
	// An interface value holds a value of a specific underlying concrete type.
	// Calling a method on an interface value executes the method of the same name on its underlying type.
	describe := func (i Abser) {
		fmt.Printf("(%v, %T)\n", i, i)
	}
	describe(f)
	describe(&v)

	// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
	// In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver.
	// Note that an interface value that holds a nil concrete value is itself non-nil.
	var b Abser
	var f2 *InterPoint
	b = f2
	describe(b)
	b.MyAbs()

	// A nil interface value holds neither value nor concrete type.
	// Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	var c Abser
	describe(c)
	//c.MyAbs()

	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
	var any interface{}
	any = 42
	fmt.Printf("(%v, %T)\n", any, any)
	any = "hello"
	fmt.Printf("(%v, %T)\n", any, any)

	// A type assertion provides access to an interface value's underlying concrete value. Format is:
	// t := interface.(Type)
	d, ok := b.(*InterPoint)
	fmt.Println(d, ok)
	e := b.(*InterPoint)
	fmt.Println(e)

	// A type switch is a construct that permits several type assertions in series.
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case *InterPoint:
			fmt.Printf("Type is *InterPoint\n")
		case InterFloat:
			fmt.Printf("Type is InterFloat\n")
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}
	checkType(b)
	checkType(22)

	// One of the most ubiquitous interfaces is Stringer defined by the fmt package.
	// type Stringer interface {
	//     String() string
	// }
	// So, when we defined a `String` method in our type, the fmt.Println will call our `String` method when print (kind like magic method in Python)
	y := InterPoint{11, 42}
	z := InterPoint{21, 9001}
	fmt.Println(y, z)
	fmt.Println(&y, &z)
}