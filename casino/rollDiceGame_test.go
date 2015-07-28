package casino_new

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
  "github.com/stretchr/testify/mock"
	"testing"
)

type DiceStub struct {
	mock.Mock
}

func (self DiceStub) Roll() int {
	args := self.Called()
	return args.Int(0)
}

type RollDiceGameTestSuite struct {
	suite.Suite

	dice *DiceStub
	player *Player
	game *RollDiceGame
}

func (s *RollDiceGameTestSuite) SetupTest() {
	s.dice = new(DiceStub)
	s.player = NewPlayer()
	s.player.BuyChips(1000)
	s.game = NewRollDiceGame(s.dice)
	s.player.Join(s.game)
}

func (s *RollDiceGameTestSuite) Test_Player_CanLose() {
	s.dice.On("Roll").Return(6)
	s.player.Bet(Bet{ Amount: 10, Score: 1})

	s.game.Play()

	assert.Equal(s.T(), 0, s.player.GetBetOn(1))
	assert.Equal(s.T(), 1000 - 10, s.player.AvailableChips())
}

func (s *RollDiceGameTestSuite) Test_Player_CanWin() {
	s.dice.On("Roll").Return(6)
	s.player.Bet(Bet{ Amount: 10, Score: 6})

	s.game.Play()

	assert.Equal(s.T(), 0, s.player.GetBetOn(1))
	assert.Equal(s.T(), 1000 - 10 + 10*6, s.player.AvailableChips())
}

func (s *RollDiceGameTestSuite) Test_PlayerNotInGame_CanNotLose() {
	s.dice.On("Roll").Return(6)
	s.player.Bet(Bet{ Amount: 10, Score: 1})
	s.player.Leave()

	s.game.Play()

	assert.Equal(s.T(), 0, s.player.GetBetOn(1))
	assert.Equal(s.T(), 1000, s.player.AvailableChips())
}

func TestRollDiceGameTestSuite (t *testing.T) {
	suite.Run(t, new(RollDiceGameTestSuite))
}
