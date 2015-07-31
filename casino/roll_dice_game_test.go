package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/mock"
)

type diceMock struct{
	mock.Mock
}

func (m diceMock) Roll() int {
	args := m.Called()
	return args.Int(0)
}

type RollDiceGameSuite struct {
	CasinoBasicSuite
}

func TestRollDiceGameSuite(t *testing.T) {
	suite.Run(t, new(RollDiceGameSuite))
}

type gameCase struct {
	score int
	winingScore int
	chips int
	bet int
}

func (suite *RollDiceGameSuite) getGame(data gameCase) (*RollDiceGame, *Player) {
	player := NewPlayer()
	player.BuyChips(data.chips)
	bet := suite.bet(data.bet, data.score)
	player.Bet(bet)
	dice := diceMock{}
	dice.On("Roll").Return(data.winingScore)
	game := NewRollDiceGameWithDice(dice)
	player.Join(game)
	return game, player
}

func (suite *RollDiceGameSuite) TestGame_PlayerInGameWitBet_Win_IncreaseChipsAmount() {
	game, player := suite.getGame(gameCase{
		score: 5,
		winingScore: 5,
		chips: 100,
		bet: 50,
	})

	game.Play()

	suite.AssertEquals(50 + 50*6, player.AvailableChips())
}

func (suite *RollDiceGameSuite) TestGame_PlayerInGameWitBet_Loose_NoMoney() {
	game, player := suite.getGame(gameCase{
		score: 5,
		winingScore: 3,
		chips: 100,
		bet: 50,
	})

	game.Play()

	suite.AssertEquals(50, player.AvailableChips())
}