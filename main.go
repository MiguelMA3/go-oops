package main

import (
	"fmt"

	"github.com/MiguelMA3/go-oops/composition"
)

func main() {
	c := composition.NewCar("a", "b", "c", 600, 123)
	fmt.Println(c.EngineName())
	fmt.Println(c.WheelName())
}
