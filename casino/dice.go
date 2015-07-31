package casino_new

import (
	"math/rand"
	"time"
)

type IDice interface {
	Roll() int
}

type Dice struct {
}

func NewDice() *Dice {
	return new(Dice)
}

func (self *Dice) Roll() int {
	rand.Seed(time.Now().UTC().UnixNano())
	winningScore := rand.Int()%6 + 1
	return winningScore
}
