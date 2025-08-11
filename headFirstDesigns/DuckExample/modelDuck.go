package main

import "fmt"

func ModelDuck() Duck {
	duckmodel := Duck_class()
	// we call the setFlyBehavior method which referneces the interface Flybehavior allowing us to grab FlyNoWay
	duckmodel.setFlyBehavior(FlyNoWay{})
	duckmodel.setQuackBehavior(Quack{})

	fmt.Println("I'm a model duck")

	return duckmodel

}
