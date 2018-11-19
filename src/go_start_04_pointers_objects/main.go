package main

import (
	"fmt"
	"go_start_04_pointers_objects/insect"
)

func main() {

	var buba insect.Insect
	var roach insect.Insect

	buba = insect.Insect{"Scolopendra", 20.3, 11.8, "hungredLegs beast"}
	buba.Print()

	roach = insect.Insect{"Roach", 4.21, 3.22, "bugDigger"}
	roach.Print()

	fly := &insect.Insect{Name: "BaksterStockamn", Spice: "fly"}
	fly.Print()
	fmt.Printf("%p\n", &fly)

	fly1 := insect.Insect{Name: "BaksterStockamn", Spice: "fly"}
	fly1.Print()
	fmt.Printf("%p\n", &fly1)

	fly1.SetName("Horse-fly")
	fly1.Print()

	fmt.Printf(" \n\n\n Hejjj druze \n")

}
