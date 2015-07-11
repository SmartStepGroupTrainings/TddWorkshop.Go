package casino_new

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RollDiceGameTestSuite struct {
	suite.Suite
	player *Player
	game *RollDiceGame
}

func (s *RollDiceGameTestSuite) SetupTest() {
	s.player = NewPlayer()
	s.player.BuyChips(1000)
	s.game = NewRollDiceGame()
	s.player.Join(s.game)
}

func (s *RollDiceGameTestSuite) Ignore_Test_Player_CanLose() {
	s.player.Bet(Bet{ Amount: 10, Score: 1})

	s.game.Play()

	assert.Equal(s.T(), 0, s.player.GetBetOn(1))
	assert.Equal(s.T(), 1000 - 10, s.player.AvailableChips())
}


func (s *RollDiceGameTestSuite) Ignore_Test_Player_CanWin() {
	s.player.Bet(Bet{ Amount: 10, Score: 6})

	s.game.Play()

	assert.Equal(s.T(), 0, s.player.GetBetOn(1))
	assert.Equal(s.T(), 1000 - 10 + 10*6, s.player.AvailableChips())
}

func TestRollDiceGameTestSuite (t *testing.T) {
	suite.Run(t, new(RollDiceGameTestSuite))
}
