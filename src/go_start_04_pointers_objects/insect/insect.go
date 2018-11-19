package insect

import "fmt"

type Insect struct {
	Name   string
	Width  float32
	Height float32
	Spice  string
}

func (i Insect) Print() {

	fmt.Printf("************* \n"+
		"Name: %s \nWidth: %f \nHeight: %f \nSpice: %s \n", i.Name, i.Width, i.Height, i.Spice)
}

func (i *Insect) SetName(name string) {

	i.Name = name
}
