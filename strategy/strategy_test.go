package strategy

func ExampleGameCharacter() {
	char := NewGameCharacter()
	char.Attack()

	char.SetWeapon(NewSword())
	char.Attack()

	char.SetWeapon(NewGun())
	char.Attack()

	// Output:
	// Default Attack
	// Sword Attack
	// Gun Attack
}
