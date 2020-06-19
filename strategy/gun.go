package strategy

import "fmt"

type Gun struct {}

func (d *Gun)Attack() {
	fmt.Println("Gun Attack")
}

func NewGun() Weapon {
	return &Gun{}
}
