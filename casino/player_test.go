package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RandomizerMock struct {
	value int
}

func (self *RandomizerMock) GetValue() int {
	return self.value
}

func NewRandomizerMock(value int) *RandomizerMock {
	return &RandomizerMock{value: value}
}

type PlayerTestSuite struct {
	suite.Suite
	player *Player
	game   *RollDiceGame
}

func (suite *PlayerTestSuite) SetupTest() {
	suite.game = NewRollDiceGame(NewRandomizerMock(1))
	suite.player = NewPlayer()
}

func (suite *PlayerTestSuite) TestPlayer_NotInGame_FirstJoinGame_Success() {
	suite.player.Join(suite.game)

	suite.True(suite.player.IsInGame())
}

func (suite *PlayerTestSuite) TestPlayer_NotInGame_FirstJoinGame_NotReturnErrorForSuccessResult() {
	err := suite.player.Join(suite.game)

	suite.Nil(err, "Return value is null")
}

func (suite *PlayerTestSuite) TestPlayer_NotInGame_JoinAnotherGame_Fail() {
	suite.player.Join(suite.game)

	err := suite.player.Join(suite.game)

	suite.NotNil(err, "Return value is not null")
	suite.Equal("Unable to join another game", err.Error())
}

func (suite *PlayerTestSuite) TestPlayer_NotInGame_Leave_Fail() {
	err := suite.player.Leave()

	suite.NotNil(err, "Return value is not null")
	suite.Equal("Unable to leave the game before joining", err.Error())
}

func (suite *PlayerTestSuite) TestPlayer_NotInGame_Leave_Success() {
	suite.player.Join(suite.game)

	err := suite.player.Leave()

	suite.Nil(err, "Player error is not null")
}

func (suite *PlayerTestSuite) TestPlayer_NotInGameAndBets_BuyChips_Success() {
	suite.player.BuyChips(1)
	suite.player.Bet(Bet{Amount: 1, Score: 1})
	suite.player.Join(suite.game)

	suite.player.Leave()

	suite.Equal(1, suite.player.AvailableChips())
}

func (suite *PlayerTestSuite) TestPlayer_NotInGame_IsInGame_Fail() {
	suite.Equal(false, suite.player.IsInGame())
}

func (suite *PlayerTestSuite) TestPlayer_InGame_Join_Success() {
	suite.player.Join(suite.game)

	suite.Equal(true, suite.player.IsInGame())
}

func (suite *PlayerTestSuite) TestPlayer_InGame_BuyChipsWithInvalidValue_Fail() {
	suite.player.Join(suite.game)

	err := suite.player.BuyChips(-1)

	suite.NotNil(err, "Return value is not null")
	suite.Equal("Please buy positive amount", err.Error())
}

func (suite *PlayerTestSuite) TestPlayer_InGame_BuyChipsWithValidValue_Success() {
	suite.player.Join(suite.game)

	err := suite.player.BuyChips(1)

	suite.Nil(err, "Player error is not null")
	suite.Equal(1, suite.player.AvailableChips())
}

func (suite *PlayerTestSuite) TestPlayer_NotInGame_HasAvailableChipsIsZero() {
	suite.Equal(0, suite.player.AvailableChips())
}

func (suite *PlayerTestSuite) TestPlayer_InGame_BetAmountMoreThanAvailable_Fail() {
	suite.player.Join(suite.game)
	suite.player.BuyChips(1)

	err := suite.player.Bet(Bet{Amount: 2, Score: 1})

	suite.NotNil(err, "Return value is not null")
	suite.Equal("Unable to bet chips more than available", err.Error())

}

func (suite *PlayerTestSuite) TestPlayer_InGame_BetScoreNotValid_Fail() {
	suite.player.Join(suite.game)
	suite.player.BuyChips(1)

	err := suite.player.Bet(Bet{Amount: 1, Score: 7})

	suite.NotNil(err, "Return value is not null")
	suite.Equal("Bets on 1..6 only are allowed", err.Error())
}

func (suite *PlayerTestSuite) TestPlayer_InGame_BetScoreValid_Success() {
	suite.player.Join(suite.game)
	suite.player.BuyChips(1)

	_ = suite.player.Bet(Bet{Amount: 1, Score: 1})

	suite.Equal(1-1, suite.player.AvailableChips())
	suite.Equal(0+1, suite.player.GetBetOn(1))
}

func (suite *PlayerTestSuite) TestPlayer_InGame_AddWinChips_Success() {
	suite.player.Join(suite.game)

	suite.player.Win(1)

	suite.Equal(1, suite.player.AvailableChips())
}

func (suite *PlayerTestSuite) TestPlayer_InGame_GetBetOn_Success() {
	suite.player.Join(suite.game)
	suite.player.Bet(Bet{Amount: 0, Score: 1})

	value := suite.player.GetBetOn(1)

	suite.Equal(0, value)
}

func (suite *PlayerTestSuite) TestPlayer_InGame_Lose_Success() {
	suite.player.Join(suite.game)
	suite.player.Bet(Bet{Amount: 1, Score: 1})

	suite.player.Lose()

	suite.Equal(0, suite.player.GetBetOn(1))
}

func TestPlayerTestSuite(t *testing.T) {
	suite.Run(t, new(PlayerTestSuite))
}
