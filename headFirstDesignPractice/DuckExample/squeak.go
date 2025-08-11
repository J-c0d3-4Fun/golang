package main

import "fmt"

type Squeak struct{}

func (S Squeak) Quack() {
	fmt.Println("Squeak")
}
