package strategy

type GameCharacter struct {
	weapon Weapon
}

func NewGameCharacter() *GameCharacter {
	return &GameCharacter{NewDefaultWeapon()}
}

func (g *GameCharacter)SetWeapon(weapon Weapon) {
	g.weapon = weapon
}

func (g GameCharacter)Attack() {
	g.weapon.Attack()
}