package casino_new

import (
	"math/rand"
	"time"
)

type IDice interface {
	Roll() int
}

type RollDiceGame struct {
	players map[*Player]struct{}
	winningScore int
	forceScore int
}

func NewRollDiceGame() *RollDiceGame {
	return &RollDiceGame{
		players: make(map[*Player]struct{}),
	}
}

func (self *RollDiceGame) Play() {
	rand.Seed(time.Now().UTC().UnixNano())
	self.winningScore = rand.Int()%6 + 1

	for player, _ := range self.players {
		player.Win(player.GetBetOn(self.GetWiningScore()) * 6)
		player.Lose()
	}
}

func (self *RollDiceGame) GetWiningScore() int {
	if self.forceScore != 0 {
		return self.forceScore
	}
	return self.winningScore
}


func (self *RollDiceGame) SetForceScore(score int) {
	self.forceScore = score
}

func (self *RollDiceGame) Add(player *Player) {
	self.players[player] = struct{}{}
}

func (self *RollDiceGame) Remove(player *Player) {
	delete(self.players, player)
}
