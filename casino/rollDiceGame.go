package casino_new

import (
	"math/rand"
	"time"
)

type IDice interface {
	Roll() int
}

type randomDice struct {
}

func (d randomDice) Roll() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Int()%6 + 1
}

type RollDiceGame struct {
	players map[*Player]struct{}
	dice IDice
}

func NewRollDiceGame() *RollDiceGame {
	return &RollDiceGame{
		players: make(map[*Player]struct{}),
		dice: &randomDice{},
	}
}

func (self *RollDiceGame) Play() {
	for player, _ := range self.players {
		player.Win(player.GetBetOn(self.dice.Roll()) * 6)
		player.Lose()
	}
}

func (self *RollDiceGame) Add(player *Player) {
	self.players[player] = struct{}{}
}

func (self *RollDiceGame) Remove(player *Player) {
	delete(self.players, player)
}
