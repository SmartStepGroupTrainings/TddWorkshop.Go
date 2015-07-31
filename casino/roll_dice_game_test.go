package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type diceStub struct {
	score int
}


func newDiceStub(score int) *diceStub {
	return &diceStub{
		score: score,
	}
}

func (self *diceStub) Roll() int {
	return self.score
}

type RollDiceGameSuite struct {
	CasinoBasicSuite
}

func TestRollDiceGameSuite(t *testing.T) {
	suite.Run(t, new(RollDiceGameSuite))
}

type playerInGame struct {
	score int
	winingScore int
	chips int
	bet int
}

func (suite *RollDiceGameSuite) getGame(data playerInGame) (*RollDiceGame, *Player) {
	player := NewPlayer()
	player.BuyChips(data.chips)
	bet := suite.bet(data.bet, data.score)
	player.Bet(bet)
	game := NewRollDiceGameWithDice(newDiceStub(data.winingScore))
	player.Join(game)
	return game, player
}

func (suite *RollDiceGameSuite) TestGame_PlayerInGameWitBet_Win_IncreaseChipsAmount() {
	game, player := suite.getGame(playerInGame{
		score: 5,
		winingScore: 5,
		chips: 100,
		bet: 50,
	})

	game.Play()

	suite.AssertEquals(50 + 50*6, player.AvailableChips())
}

func (suite *RollDiceGameSuite) TestGame_PlayerInGameWitBet_Loose_NoMoney() {
	game, player := suite.getGame(playerInGame{
		score: 5,
		winingScore: 3,
		chips: 100,
		bet: 50,
	})

	game.Play()

	suite.AssertEquals(50, player.AvailableChips())
}