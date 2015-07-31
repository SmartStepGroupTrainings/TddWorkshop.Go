package casino_new

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type DiceStub struct {
	mock.Mock
}

func (self DiceStub) GetValue() int {
	args := self.Called()
	return args.Int(0)
}

type DiceTestSuite struct {
	suite.Suite
	dice   *DiceStub
	game   *RollDiceGame
	player *Player
}

func (suite *DiceTestSuite) SetupTest() {
	suite.dice = new(DiceStub)
	suite.game = NewRollDiceGame(suite.dice)
	suite.player = NewPlayer()
}

func TestDiceTestSuite(t *testing.T) {
	suite.Run(t, new(DiceTestSuite))
}

func (suite *DiceTestSuite) TestDice_EmptyPlayers_NoPlayer_Success() {
	suite.Empty(suite.game.GetPlayers())
}

func (suite *DiceTestSuite) TestDice_EmptyPlayers_AddPlayer_Success() {
	suite.game.Add(suite.player)

	suite.NotNil(suite.game.GetPlayers())
}

func (suite *DiceTestSuite) TestDice_PlayerInGame_Play_NullAvailableChipsSuccess() {
	suite.dice.On("GetValue").Return(3)
	suite.game.Add(suite.player)
	suite.player.BuyChips(2)
	suite.player.Bet(Bet{Amount: 2, Score: 3})
	suite.player.Join(suite.game)

	suite.game.Play()

	suite.Equal(2*6, suite.player.AvailableChips())
}
