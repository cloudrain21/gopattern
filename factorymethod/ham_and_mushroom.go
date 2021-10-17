package factorymethod

import "fmt"

type HamAndMushroom struct {
}

func NewHamAndMushroom() *HamAndMushroom {
	return &HamAndMushroom{}
}

func (m *HamAndMushroom) GetPrice() {
	fmt.Println("ham and mushroom : 100")
}
