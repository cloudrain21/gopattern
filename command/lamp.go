package command

import "fmt"

type Lamp struct {
}

func NewLamp() *Lamp {
	return &Lamp{}
}

func (l *Lamp) turnOn() {
	fmt.Println("Lamp on")
}
