package rolls

import (
	"../dice"
)

func AttackRoll() int {
	return (dice.D(20))
}
