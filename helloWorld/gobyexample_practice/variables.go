package main

import "fmt"

// var name type =
func myvars() {
	var a = "test"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "testing"
	fmt.Println(f)

	// The := syntax is shorthand for declaring and initializing a variable, e.g.
	// for var f string = "testing" in this case. This syntax is only available inside functions.

}
