package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func (self *PlayerTestDsl) Test_LooserLoseHisBet() {
	game := self.create.RollDiceGame().Please()
	looser := self.create.Looser().WithBet(10).Chips().Please()

	game.Play()

	self.assertThat(looser).Lost(10).Chips()
}

func (self *PlayerTestDsl) Test_WinnerWins6Bets() {
	game := self.create.RollDiceGame().Please()
	winner := self.create.Winner().WithBet(10).Chips().Please()

	game.Play()

	self.assertThat(winner).Won(10 * 6).Chips()
}

type PlayerTestDsl struct {
	suite.Suite
	state  *State
	create *Father
	player *Player
}

func (self *PlayerTestDsl) SetupTest() {
	self.create = new(Father)
}

func (self *PlayerTestDsl) assertThat(player *Player) *PlayerTestDsl {
	self.player = player
	return self
}

func (self *PlayerTestDsl) Won(chips int) *PlayerTestDsl {
	self.Equal(chips, self.player.AvailableChips()-self.state.InitialAvailableChips)
	return self
}

func (self *PlayerTestDsl) Lost(chips int) *PlayerTestDsl {
	self.Equal(chips, self.state.InitialAvailableChips-self.player.AvailableChips())
	return self
}

func (self *PlayerTestDsl) Chips() *PlayerTestDsl {
	return self
}

func Test_RollDiceGameDsl(t *testing.T) {
	suite.Run(t, new(PlayerTestDsl))
}

type State struct {
	InitialAvailableChips int
}
