package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RollDiceGameSuite struct {
	CasinoBasicSuite
}

func TestRollDiceGameSuite(t *testing.T) {
	suite.Run(t, new(RollDiceGameSuite))
}

func (suite *RollDiceGameSuite) TestGame_PlayerInGameWitBet_Win_IncreaseChipsAmount() {
	game := NewRollDiceGameBuilder().addPlayerWithBet(100, 5).setWinningScore(5).Get()

	winners := game.Play()

	for _, player := range winners {
		suite.AssertEquals(100*6, player.AvailableChips())
	}
}

func (suite *RollDiceGameSuite) TestGame_PlayerInGameWitBet_Loose_NoMoney() {
	game := NewRollDiceGameBuilder().addPlayerWithBet(100, 5).setWinningScore(3).Get()

	game.Play()

	for _, player := range game.GetPlayers() {
		suite.AssertEquals(0, player.AvailableChips())
	}
}