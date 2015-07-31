package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestGameSuite struct {
	suite.Suite
	game   *RollDiceGame
	player *Player
	dice   *DiceStub
}

type DiceStub struct {
	mock.Mock
}

const anyAmount = 1

func (dice *DiceStub) Roll() int {
	args := dice.Called()
	return args.Int(0)
}

func (self *TestGameSuite) SetupTest() {
	self.game = NewRollDiceGame()

	self.dice = &DiceStub{}
	self.game.dice = self.dice

	self.player = NewPlayer()
	self.player.BuyChips(anyAmount)
	self.player.Join(self.game)
}

func Test_Game(t *testing.T) {
	suite.Run(t, new(TestGameSuite))
}

func (self *TestGameSuite) TestGame_Play_PlayerWins6Chips_WhenBetsScoreOne() {
	game := create.Game().WinningScore(2).Please()
	player := create.Player().WithChips(1).InGame(game).Bets(1).OnScore(2).Please()

	game.Play()

	assert.Equal(self.T(), 1*6, player.AvailableChips())
}

func (self *TestGameSuite) TestGame_Play_PlayerLosesEverything_andHisWifeGoesToNeighbor_OnWrongBetInCasino() {
	self.dice.On("Roll").Return(1)
	self.player.Bet(Bet{Score: 2, Amount: self.player.AvailableChips()})

	self.game.Play()

	assert.Empty(self.T(), self.player.AvailableChips())
}

func (self *TestGameSuite) TestGame_Play_ResetBets_WhenPlayerWinsBet() {
	self.dice.On("Roll").Return(1)
	self.player.Bet(Bet{Score: 1, Amount: anyAmount})

	self.game.Play()

	assert.Empty(self.T(), self.player.GetBetOn(1))
}

func (self *TestGameSuite) TestGame_Play_ResetBets_WhenPlayerLosesBet() {
	self.dice.On("Roll").Return(1)
	self.player.Bet(Bet{Score: 2, Amount: anyAmount})

	self.game.Play()

	assert.Empty(self.T(), self.player.GetBetOn(2))
}
