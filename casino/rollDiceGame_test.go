package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DiceTestSuite struct {
	suite.Suite
	game   *RollDiceGame
	player *Player
}

func (suite *DiceTestSuite) SetupTest() {
	suite.game = NewRollDiceGame(NewRandomizerMock(1))
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
	game := NewRollDiceGame(NewRandomizerMock(3))
	game.Add(suite.player)
	suite.player.BuyChips(2)
	suite.player.Bet(Bet{Amount: 2, Score: 3})
	suite.player.Join(game)

	game.Play()

	suite.Equal(2*6, suite.player.AvailableChips())
}
