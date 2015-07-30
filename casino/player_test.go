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
	player := NewPlayer()

	suite.AssertFalse(player.IsInGame())
	suite.AssertEquals(0, player.AvailableChips())
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

func (suite *PlayerSuite) TestPlayer_BuyZeroChips_Fail() {
	err := suite.player.BuyChips(0)

	suite.AssertNotNil(err)
}

func (suite *PlayerSuite) TestPlayer_BuyNegativeAmountOfChips_Fail() {
	err := suite.player.BuyChips(-1)

	suite.AssertNotNil(err)
}

func (suite *PlayerSuite) TestPlayer_With10Chips_Buy10Chips_Success() {
	suite.player.BuyChips(10)

	suite.player.BuyChips(10)

	suite.Equal(10+10, suite.player.AvailableChips())
}

func (suite *PlayerSuite) TestPlayer_WithChips_MakingBet_Success() {
	suite.player.BuyChips(10)
	bet := suite.bet(10, 6)

	suite.player.Bet(bet)

	suite.Equal(0, suite.player.AvailableChips())
	suite.Equal(10, suite.player.GetBetOn(6))
}

func (suite *PlayerSuite) TestPlayer_WithoutChips_MakingBet_Fail() {
	bet := suite.bet(10, 6)

	err := suite.player.Bet(bet)

	suite.AssertNotNil(err)
}

func (suite *PlayerSuite) TestPlayer_WithChips_MakingBetBiggerThanAvailableChips_Fail() {
	suite.player.BuyChips(5)
	bet := suite.bet(10, 6)

	err := suite.player.Bet(bet)

	suite.AssertNotNil(err)
}

func (suite *PlayerSuite) TestPlayer_WithChips_MakingBetOnWrongScore_Fail() {
	suite.player.BuyChips(5)
	bet := suite.bet(5, 7)

	err := suite.player.Bet(bet)

	suite.AssertNotNil(err)
}