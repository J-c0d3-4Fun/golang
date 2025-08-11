package main

import "fmt"

type MuteQuack struct{}

func (m MuteQuack) Quack() {
	fmt.Println("<< SILENCE >>")
}
