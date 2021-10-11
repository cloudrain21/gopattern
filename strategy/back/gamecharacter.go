package strategy

type GameCharacter struct {
	name   string
	weapon Weapon
}

func NewGameCharacter(name string) *GameCharacter {
	return &GameCharacter{name, DefaultWeapon{}}
}

func (g *GameCharacter) SetWeapon(weapon Weapon) {
	g.weapon = weapon
}

// gamecharacter 의 Attack() 을 수행하면 현재 설정된 weapon 의 Attack() 을 수행한다.
func (g GameCharacter) Attack() {
	g.weapon.Attack()
}
