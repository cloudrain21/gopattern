package strategy

import "fmt"

type Sword struct {}

func NewSword() Weapon {
	return &Sword{}
}

func (s Sword)Attack() {
	fmt.Println("Sword attack")
}

