package main

func MiniDuckSimulator() {
	duckMallard := Duck_class()
	duckMallard.performFly()
	duckMallard.performQuack()

	duckModel := ModelDuck()
	duckModel.performFly()
	duckModel.setFlyBehavior(FlyRocketPowered{})
	duckModel.performFly()

}
