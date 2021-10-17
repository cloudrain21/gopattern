package factorymethod

import "fmt"

type Deluxe struct {
}

func NewDeluxe() *Deluxe {
	return &Deluxe{}
}

func (m *Deluxe) GetPrice() {
	fmt.Println("deluxe : 200")
}
