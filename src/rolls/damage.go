package rolls

import (
	"strings"

	"../dice"
)

func DamageRoll(n int, die string) int {
	die = strings.ToLower(die)
	var result int = 0
	var dicetype int

	switch die {
	case "d4":
		dicetype = 4
	case "d6":
		dicetype = 6
	case "d8":
		dicetype = 8
	case "d10":
		dicetype = 10
	case "d12":
		dicetype = 12
	case "d20":
		dicetype = 20
	case "d100":
		dicetype = 100
	}

	for i := 0; i < n; i++ {
		result = result + dice.D(dicetype)
	}

	return result
}
