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

func (suite *PlayerSuite) TestPlayer_NotInGame_JoinGame_Success() {
	game := NewRollDiceGame()

	suite.player.Join(game)

	suite.AssertTrue(suite.player.IsInGame())
}

func (suite *PlayerSuite) TestPlayer_NotInGame_LeaveGame_Fail() {
	err := suite.player.Leave()

	suite.AssertNotNil(err)
}

func (suite *PlayerSuite) TestPlayer_InGame_LeaveGame_Fail() {
	game_one := NewRollDiceGame()
	game_two := NewRollDiceGame()
	suite.player.Join(game_one)

	err := suite.player.Join(game_two)

	suite.AssertNotNil(err)
}

func (suite *PlayerSuite) TestPlayer_InGame_LeaveGame_Success() {
	game := NewRollDiceGame()
	suite.player.Join(game)

	suite.player.Leave()

	suite.AssertFalse(suite.player.IsInGame())
}

func (suite *PlayerSuite) TestPlayer_WithZeroChips_Buy10Chips_Success() {
	suite.player.BuyChips(10)

	suite.Equal(10, suite.player.AvailableChips())
}

func (suite *PlayerSuite) TestPlayer_With10Chips_Buy10Chips_Success() {
	suite.player.BuyChips(10)

	suite.player.BuyChips(10)

	suite.Equal(10+10, suite.player.AvailableChips())
}

func (suite *PlayerSuite) TestPlayer_With10Chips_Making10ChipsBetOn6_Success() {
	suite.player.BuyChips(10)
	bet := Bet{}
	bet.Amount = 10
	bet.Score = 6

	suite.player.Bet(bet)

	suite.Equal(0, suite.player.AvailableChips())
	suite.Equal(10, suite.player.GetBetOn(6))
}