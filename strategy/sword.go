package strategy

import "fmt"

type Sword struct {}

func (d *Sword)Attack() {
	fmt.Println("Sword Attack")
}

func NewSword() Weapon {
	return &Sword{}
}
