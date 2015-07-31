package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestGameSuite struct {
	suite.Suite
	game   *RollDiceGame
	player *Player
	dice   CheatersDice
}

type CheatersDice struct {
	Score int
}

const (
	anyAmount    = 1
	anyScore     = 1
	anotherScore = 2
)

func (dice CheatersDice) Roll() int {
	return dice.Score
}

func (self *TestGameSuite) SetupTest() {
	self.game = NewRollDiceGame()

	self.dice = CheatersDice{anyScore}
	self.game.dice = self.dice

	self.player = NewPlayer()
	self.player.BuyChips(anyAmount)
	self.game.Add(self.player)
}

func Test_Game(t *testing.T) {
	suite.Run(t, new(TestGameSuite))
}

func (self *TestGameSuite) TestGame_Play_PlayerWins6Chips_WhenBetsScoreOne() {
	bet := Bet{Score: self.dice.Score, Amount: anyAmount}
	self.player.Bet(bet)

	self.game.Play()

	assert.Equal(self.T(), bet.Amount*6, self.player.AvailableChips())
}

func (self *TestGameSuite) TestGame_Play_PlayerLosesEverything_andHisWifeGoesToNeighbor_OnWrongBetInCasino() {
	bet := Bet{Score: anotherScore, Amount: self.player.AvailableChips()}
	self.player.Bet(bet)

	self.game.Play()

	assert.Empty(self.T(), self.player.AvailableChips())
}

func (self *TestGameSuite) TestGame_Play_ResetBets_WhenPlayerWinsBet() {
	bet := Bet{Score: self.dice.Score, Amount: anyAmount}
	self.player.Bet(bet)

	self.game.Play()

	assert.Empty(self.T(), self.player.GetBetOn(anotherScore))
}

func (self *TestGameSuite) TestGame_Play_ResetBets_WhenPlayerLosesBet() {
	bet := Bet{Score: anotherScore, Amount: anyAmount}
	self.player.Bet(bet)

	self.game.Play()

	assert.Empty(self.T(), self.player.GetBetOn(anotherScore))
}
