package main

import (
	"firstPackage/dateModule"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//setting random seed
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Gde si brej baraboZoica: %d  \n", rand.Intn(10))
	fmt.Printf("Akko si tako !!!  reverse: [ %s  ]", string.Reverse("Akko si tako!!! \n"))

	bytes := rand.Intn(6) + 1
	fmt.Printf("Hejj druze ovo je tvoj srecni random broj: %d \n", bytes)

}
