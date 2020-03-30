package dice

import (
	"math/rand"
)

/*
rand.Intn(number) returns a value between [0, number]
with extremes included
*/

func D(n int) int {
	return rand.Intn(n) + 1
}
