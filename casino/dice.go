package casino_new

import "math/rand"

type iDice interface {
	Roll() int
}

type RandomDice struct{}

func (d *RandomDice) Roll() int {
	return rand.Intn(5) + 1
}
