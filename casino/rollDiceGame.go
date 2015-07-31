package casino_new

import (
	"math/rand"
	"time"
)

type IDice interface {
	Roll() int
}

type RandomDice struct{}

type RollDiceGame struct {
	players map[*Player]struct{}
	dice    IDice
}

func NewRollDiceGame() *RollDiceGame {
	return &RollDiceGame{
		players: make(map[*Player]struct{}),
		dice:    RandomDice{},
	}
}

func (self *RollDiceGame) Play() {
	winningScore := self.dice.Roll()

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

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (dice RandomDice) Roll() int {
	return rand.Int()%6 + 1
}
