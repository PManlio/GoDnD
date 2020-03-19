package character

type Character struct {
	PlayerID   string    `json:"playerid"`
	PlayerName string    `json:"playername"`
	CharName   string    `json:"charname"`
	Classes    []Class   `json:"classes"`
	Ability    Abilities `json:"ability"`
	Competence int       `json:"competence"`
	Hitpoints  int       `json:"hitpoints"`
	ArmorClass int       `json:"armorclass"`
	Inventory  []string  `json:"inventory"`
}
