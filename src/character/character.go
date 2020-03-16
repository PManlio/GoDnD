package character

type Character struct {
	PlayerID              string
	PlayerName            string
	CharName              string
	Classes               []Class
	Abilities             Abilities
	Competence            int
	Hitpoints, Armorclass int
	Inventory             []string
}
