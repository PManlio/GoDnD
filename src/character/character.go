package character

type Character struct {
	PlayerID              string
	PlayerName            string
	CharName              string
	Classes               []Class
	Ability               Abilities
	Competence            int
	Hitpoints, ArmorClass int
	Inventory             []string
}
