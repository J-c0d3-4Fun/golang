package main

import (
	"fmt"
)

func arrays() {
	var a [5]int
	fmt.Println(a)

	// We can set a value at an index using the array[index] = value syntax,
	// and get a value with array[index].

	a[4] = 100 // adding to the index, by setting the value of the 4th position of a
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])   // print the 4th position fo the index
	fmt.Println("len:", len(a)) // Grab length

	// Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// You can also have the compiler count the number of elements for you with "..."
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// If you specify the index with :, the elements in between will be zeroed.
	b = [...]int{1, 2, 3: 400, 500}
	fmt.Println("idx:", b)

	// array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// You can create and initialize multi-dimensional arrays at once too.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

}
