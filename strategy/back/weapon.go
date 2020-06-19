package strategy

import "fmt"

type Weapon interface {
	Attack()
}

type DefaultWeapon struct {}

func (d DefaultWeapon)Attack() {
	fmt.Println("DefaultWeapon attack")
}
