package main

import (
	"fmt"
	"math"
)

const smallValue = 0.00001

func Sqrt(x float64) float64 {
	// The init and post statement are optional. (Note the first `;` is required)
	// In this scenario, for loop is the same as while loop in other languages
	result := 1.0
	for ; math.Abs(result*result-x) > smallValue; {
		result -= (result*result - x) / (2 * x)
	}
	return result
}

func main() {
	// Go has only one looping construct, the for loop.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// This is the forever loop
	for {
		break
	}

	fmt.Println(Sqrt(4))
}
