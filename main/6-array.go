package main

import (
	"fmt"
	"strings"
)

func main() {
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
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	s[0] = -1
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

	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	// A nil slice has a length and capacity of 0 and has no underlying array.
	var s1 []int
	printSlice(s1)
	if s1 == nil {
		fmt.Println("nil!")
	}

	// Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.
	// The `make` function allocates a zeroed array and returns a slice that refers to that array
	s2 := make([]int, 5)
	printSlice(s2)

	s3 := make([]int, 0, 5)
	printSlice(s3)

	s4 := s3[:2]
	printSlice(s4)

	s5 := s4[2:5]
	printSlice(s5)

	// Slices can contain any type, including other slices.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// append works on nil slices.
	var s6 []int
	s6 = append(s6, 0)
	printSlice(s6)

	// The slice grows as needed.
	s6 = append(s6, 1)
	printSlice(s6)

	// We can add more than one element at a time.
	s6 = append(s6, 2, 3, 4)
	printSlice(s6)

	// When ranging over a slice, two values are returned for each iteration.
	// The first is the index, and the second is a copy of the element at that index.
	for i, v := range primes {
		fmt.Println(i, v)
	}

	// If you only want the index, you can omit the second variable.
	for i := range primes {
		fmt.Println(i)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}