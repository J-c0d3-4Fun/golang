package main

// Remember import is a function
import (
	"fmt"
	"math"
)

const s string = "constant"

func constant() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call. For example, here math.Sin expects a float64.

}
