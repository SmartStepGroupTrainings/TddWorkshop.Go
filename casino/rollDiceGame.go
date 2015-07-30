package casino_new

import (
	"math/rand"
	"time"
)

type IDice interface {
	Roll() int
}

type Dice struct{}

func (self *Dice) Roll() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Int()%6 + 1
}

type RollDiceGame struct {
	dice    IDice
	players map[*Player]struct{}
}

func NewRollDiceGame(dice IDice) *RollDiceGame {
	return &RollDiceGame{
		dice:    dice,
		players: make(map[*Player]struct{}),
	}
}

func (self *RollDiceGame) Play() {
	var winningScore = self.dice.Roll()

	for player, _ := range self.players {
		player.Win(player.GetBetOn(winningScore) * 6)
		player.Lose()
	}
}

func (self *RollDiceGame) Add(player *Player) {
	self.players[player] = struct{}{}
}

func (self *RollDiceGame) Remove(player *Player) {
	delete(self.players, player)
}
