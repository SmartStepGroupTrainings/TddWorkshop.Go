package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PlayerSuite struct {
	player *Player
	CasinoBasicSuite
}

func TestPlayerSuite(t *testing.T) {
	suite.Run(t, new(PlayerSuite))
}

func (suite *PlayerSuite) SetupTest() {
	suite.player = NewPlayer()
}

func (suite *PlayerSuite) TestPlayer_CreateNew_Success() {
	suite.AssertFalse(suite.player.IsInGame())
	suite.AssertEquals(0, suite.player.AvailableChips())
}

func (suite *PlayerSuite) TestPlayer_JoinGame_Success() {
	game := NewRollDiceGame()

	suite.player.Join(game)

	suite.AssertTrue(suite.player.IsInGame())
}

func (suite *PlayerSuite) TestPlayer_JoinSimultaneouslySecondGame_Fail() {
	game_one := NewRollDiceGame()
	game_two := NewRollDiceGame()
	suite.player.Join(game_one)

	err := suite.player.Join(game_two)

	suite.AssertNotNil(err)
}

func (suite *PlayerSuite) TestPlayer_LeaveGame_Success() {
	game := NewRollDiceGame()
	suite.player.Join(game)

	suite.player.Leave()

	suite.AssertFalse(suite.player.IsInGame())
}

func (suite *PlayerSuite) TestPlayer_LeaveGameBeforeJoin_Fail() {
	err := suite.player.Leave()

	suite.AssertNotNil(err)
}
