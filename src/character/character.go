package character

type Character struct {
	PlayerID              string
	Name                  string
	Classes               []Class
	Abilities             Abilities
	Hitpoints, Armorclass int
	Inventory             []string
}
