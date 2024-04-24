package main

import (
	"fmt"

	"github.com/MiguelMA3/go-oops/structs"
)

func main() {
	p1 := structs.NewPerson("Obi-wan", "Kenobi", 33)
	//p2 := structs.NewPerson("Mace", "Windu", 46)
	p3 := structs.NewPerson("Obi-wan", "Kenobi", 33)

	if p1 == p3 {
		fmt.Println("Both the structs are equal")
	}
}
