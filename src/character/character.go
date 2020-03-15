package character

type Character struct {
	player                string
	name                  string
	classes               []Class
	abilities             Abilities
	hitpoints, armorclass int
	inventory             []string
}
