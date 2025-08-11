package main

import "fmt"

type FlyRocketPowered struct{}

func (r FlyRocketPowered) Fly() {
	fmt.Println("I'm flying witha. rocket!")
}
