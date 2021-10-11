package strategy

import "fmt"

type Gun struct {
	bullets int
}

func NewGun() Weapon {
	return &Gun{10}
}

func (g Gun) Attack() {
	fmt.Println("Gun attack")
}
