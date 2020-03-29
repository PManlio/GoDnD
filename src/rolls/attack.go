package rolls

import (
	"math/rand"
)

func AttackRoll() int {
	return (rand.Intn(20) + 1)
}
