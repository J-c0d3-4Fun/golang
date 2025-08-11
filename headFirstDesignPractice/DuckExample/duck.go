package main

import "fmt"

// these structs can hold fields so we refernce the interfaces we created in separate files
type Duck struct {
	Flybehavior   Flybehavior
	Quackbehavior Quackbehavior
}

// (d Duck) creates a method of the from the struct referencing the interface
func (d Duck) performFly() {
	d.Flybehavior.Fly()

}

func (d Duck) performQuack() {
	d.Quackbehavior.Quack()

}

func (d Duck) swim() {
	fmt.Println("All ducks float, even decoys!")
}

// Dynamically settign the Flybehavior we use * to get the pointer (data from Duck struct)
// d *Duck means the method receives a pointer to the Duck — so changes persist on the original object.
// fb FlyBehavior is the parameter (interface type) passed in.
// d.FlyBehavior = fb updates the duck’s behavior dynamically.
func (d *Duck) setFlyBehavior(fb Flybehavior) {
	d.Flybehavior = fb
}

func (d Duck) setQuackBehavior(qb Quackbehavior) {
	d.Quackbehavior = qb
}

// This is where we call the methods we created
func Duck_class() Duck {
	return Duck{
		Flybehavior:   FlyWithWings{},
		Quackbehavior: Quack{},
	}

}
