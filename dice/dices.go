package dice

import (
	"math/rand"
)

func D4() int {
	return rand.Intn(3) + 1
}

func D6() int {
	return rand.Intn(5) + 1
}

func D8() int {
	return rand.Intn(7) + 1
}

func D10() int {
	return rand.Intn(9) + 1
}

func D12() int {
	return rand.Intn(11) + 1
}

func D20() int {
	return rand.Intn(19) + 1
}

func D100() int {
	return rand.Intn(99) + 1
}
