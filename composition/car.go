package composition

type engine struct {
	hp int
}

func (e engine) HP() int {
	return e.hp
}

type wheel struct {
	wheelDimensions int
}

func (w wheel) WheelDimension() int {
	return w.wheelDimensions
}

type Car struct {
	CarName string
	engine
	wheel
}

func NewCar(carName string, hp, wd int) Car {
	return Car{
		CarName: carName,
		engine: engine{
			hp: hp,
		},
		wheel: wheel{
			wd,
		},
	}
}
