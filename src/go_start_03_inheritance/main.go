package main

import (
	"fmt"
	"go_start_03_inheritance/models"
)

// /* This is structure of human */

func main() {

	var vizard models.Human
	var attack models.Energy
	var magicK models.Magic

	vizard = models.Human{93, 189, "hermafrodit", "Leyla"}

	attack = models.Energy{Intensity: 101, Manifest: "Electricion"}

	magicK = models.Magic{models.Energy{12, "Fire_burst"}, 101}

	fmt.Printf("Hejj %s \n", vizard.Name)
	fmt.Printf("Magic wich he use : %s \n", attack.Manifest)
	fmt.Printf("Print of methode of structure: %s \n", attack.Explode())
	fmt.Printf("Magic operand: %s \n", magicK.Explode())

}
