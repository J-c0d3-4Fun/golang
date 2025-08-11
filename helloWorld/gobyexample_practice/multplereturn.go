package main

// Go has built-in support for multiple return values.
// This feature is used often in idiomatic Go,
// for example to return both result and error values from a function.
import "fmt"

// The (int, int) in this function signature shows
// that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

func multiple() {

	// Here we use the 2 different
	// return values from the call with multiple assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
