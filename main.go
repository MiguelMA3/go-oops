package main

import (
	"fmt"

	"github.com/MiguelMA3/go-oops/composition"
)

func main() {
	c := composition.NewCar("a", 600, 123)
	fmt.Println(c.HP())
	fmt.Println(c.WheelDimension())
}
