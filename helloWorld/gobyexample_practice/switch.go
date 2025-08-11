package main

//Switch statements express conditionals across many branches.

import (
	"fmt"
	"time"
)

func myswitch() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before Noon")
	default:
		fmt.Println("It's after Noon")
	}

	// A type switch compares types instead of values.
	// You can use this to discover the type of an interface value.
	// In this example, the variable t will have the type corresponding to its clause.

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Println("Don't know the type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")
}
