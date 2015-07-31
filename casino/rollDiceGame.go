package casino_new

import (
	"math/rand"
	"time"
)

//этопростокубик
type Dice struct {
}

func (self *Dice) Roll() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Int()%6 + 1
}

type IDicer interface {
	Roll() int
}

type RollDiceGame struct {
	players map[*Player]struct{}
	roller  IDicer
}

func NewRollDiceGame(roller IDicer) *RollDiceGame {
	return &RollDiceGame{
		players: make(map[*Player]struct{}),
		roller:  roller,
	}
}

func (self *RollDiceGame) Play() {
	winningScore := self.roller.Roll()

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
