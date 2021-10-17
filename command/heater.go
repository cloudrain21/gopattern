package command

import "fmt"

type Heater struct {
}

func NewHeater() *Heater {
	return &Heater{}
}

func (h *Heater) powerOn() {
	fmt.Println("Heater on")
}
