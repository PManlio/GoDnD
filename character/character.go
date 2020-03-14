package character

type Character struct {
	name                  string
	classes               []Class
	abilities             Abilities
	hitpoints, armorclass int
	inventory             []string
}
