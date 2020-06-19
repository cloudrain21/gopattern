package strategy

// weapon 을 변경해가면서 공격을 수행할 수 있다.
// GameCharacter 와 Weapon 은 변경될 필요가 없고, 무기의 종류나 공격 내용은 쉽게 변경할 수 있게 된다.
func ExampleStrategy1() {
	character := NewGameCharacter("knight")
	character.Attack()

	character.SetWeapon(NewSword())
	character.weapon.Attack()  // 이렇게 사용하면 안됨. weapon 변수 노출

	character.SetWeapon(NewGun())
	character.Attack()

	// Output:
	// DefaultWeapon attack
	// Sword attack
	// Gun attack
}