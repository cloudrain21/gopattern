package strategy

import "fmt"

type DefaultWeapon struct {}

func (d *DefaultWeapon)Attack() {
	fmt.Println("Default Attack")
}

func NewDefaultWeapon() Weapon {
	return &DefaultWeapon{}
}
