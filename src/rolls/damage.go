package rolls

import (
	"math/rand"
	"strings"

	"../dice"
)

func DamageRoll(n int, die string) int {
	die = strings.ToLower(die)
	var result int
	var dicetype int

	for i := 0; i < n; i++ {
		switch die {
		case "d4":
			dicetype = dice.D4()
		case "d6":
			dicetype = dice.D6()
		case "d8":
			dicetype = dice.D8()
		case "d10":
			dicetype = dice.D10()
		case "d12":
			dicetype = dice.D12()
		case "d20":
			dicetype = dice.D20()
		case "d100":
			dicetype = dice.D100()
		}
		result = result + dicetype
	}

	return (rand.Intn(20) + 1)
}
