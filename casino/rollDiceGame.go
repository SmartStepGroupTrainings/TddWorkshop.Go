package casino_new

import (
	"math/rand"
	"time"
)

type IRandomizer interface {
	GetValue() int
}

type DefaultRandomizer struct {
}

func (self *DefaultRandomizer) GetValue() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Int()%6 + 1
}

type IDice interface {
	Roll() int
}

type RollDiceGame struct {
	players    map[*Player]struct{}
	randomizer IRandomizer
}

func NewRollDiceGame(randomizer IRandomizer) *RollDiceGame {
	return &RollDiceGame{
		players:    make(map[*Player]struct{}),
		randomizer: randomizer,
	}
}

func (self *RollDiceGame) Play() {
	winningScore := self.randomizer.GetValue()

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

func (self *RollDiceGame) GetPlayers() map[*Player]struct{} {
	return self.players
}
