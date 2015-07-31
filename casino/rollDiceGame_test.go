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

const (
	anyAmount    = 1
	anyScore     = 1
	anotherScore = 2
)

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
	self.game.Add(self.player)
}

func Test_Game(t *testing.T) {
	suite.Run(t, new(TestGameSuite))
}

func (self *TestGameSuite) TestGame_Play_PlayerWins6Chips_WhenBetsScoreOne() {
	self.dice.On("Roll").Return(anyScore)
	bet := Bet{Score: anyScore, Amount: anyAmount}
	self.player.Bet(bet)

	self.game.Play()

	assert.Equal(self.T(), bet.Amount*6, self.player.AvailableChips())
}

func (self *TestGameSuite) TestGame_Play_PlayerLosesEverything_andHisWifeGoesToNeighbor_OnWrongBetInCasino() {
	self.dice.On("Roll").Return(anyScore)
	bet := Bet{Score: anotherScore, Amount: self.player.AvailableChips()}
	self.player.Bet(bet)

	self.game.Play()

	assert.Empty(self.T(), self.player.AvailableChips())
}

func (self *TestGameSuite) TestGame_Play_ResetBets_WhenPlayerWinsBet() {
	self.dice.On("Roll").Return(anyScore)
	bet := Bet{Score: anyScore, Amount: anyAmount}
	self.player.Bet(bet)

	self.game.Play()

	assert.Empty(self.T(), self.player.GetBetOn(anotherScore))
}

func (self *TestGameSuite) TestGame_Play_ResetBets_WhenPlayerLosesBet() {
	self.dice.On("Roll").Return(anyScore)
	bet := Bet{Score: anotherScore, Amount: anyAmount}
	self.player.Bet(bet)

	self.game.Play()

	assert.Empty(self.T(), self.player.GetBetOn(anotherScore))
}
