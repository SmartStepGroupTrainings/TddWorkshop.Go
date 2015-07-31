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

func (dice *Dice) Roll() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Int()%6 + 1
}

type RollDiceGame struct {
	players map[*Player]struct{}
	dice    IDice
}

func NewRollDiceGame() *RollDiceGame {
	return &RollDiceGame{
		players: make(map[*Player]struct{}),
		dice:    &Dice{},
	}
}

func (self *RollDiceGame) setDice(dice IDice) {
	self.dice = dice
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
	player.currentGame = self
}

func (self *RollDiceGame) Remove(player *Player) {
	player.currentGame = nil
	delete(self.players, player)
}
