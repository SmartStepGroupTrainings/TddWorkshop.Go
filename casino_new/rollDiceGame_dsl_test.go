package casino_new

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Test struct {
	suite.Suite
	create Father
}

func (self *Test) Test_Player_CanLose_Dsl() {
	game := self.create.RollDiceGame().
					WithWinningScore(6).
					Please()
	player := self.create.Player().
						Rich().
						WithBet(10).On(1).
						Joined(game).
						Please()

	game.Play()

	assert.Equal(self.T(), 0, player.GetBetOn(1))
	assert.Equal(self.T(), 1000 - 10, player.AvailableChips())
}

func Test_RollDiceGame(t *testing.T) {
	suite.Run(t, new(Test))
}