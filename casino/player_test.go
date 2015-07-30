package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestPlayer_ByDefault_HasNoChips(t *testing.T) {
	player := NewPlayer()

	assert.Equal(t, 0, player.AvailableChips())
}

func TestPlayer_CanBuyChips(t *testing.T) {
	player := NewPlayer()

	player.BuyChips(10)

	assert.Equal(t, 10, player.AvailableChips())
}

func TestPlayer_CanBetChipsOn1(t *testing.T) {
	player := NewPlayer()
	player.BuyChips(10)

	player.Bet(Bet{Amount: 2, Score: 1})

	assert.Equal(t, 10-2, player.AvailableChips())
}

func (self *PlayerTest) TestPlayer_CanNotBetChipsOn7() {
	self.player.BuyChips(10)

	err := self.player.Bet(Bet{Amount: 2, Score: 7})

	assert.Equal(self.T(), 10, self.player.AvailableChips())
	assert.NotNil(self.T(), err)
	assert.Equal(self.T(), "Bets on 1..6 only are allowed", err.Error())
}

func (self *PlayerTest) Test_PlayerCanLose() {
	self.player.BuyChips(10)
	self.player.Join(self.game)
	self.dice.On("Roll").Return(6)
	self.player.Bet(Bet{Amount: 2, Score: 1})

	self.game.Play()

	assert.Equal(self.T(), 10-2, self.player.AvailableChips())
	assert.Empty(self.T(), self.player.GetBetOn(1))
}

type PlayerTest struct {
	suite.Suite
	player *Player
	dice   *DiceStub
	game   *RollDiceGame
}

func (self *PlayerTest) SetupTest() {
	self.player = NewPlayer()
	self.dice = new(DiceStub)
	self.game = NewRollDiceGame(self.dice)
}

func Test_RollDiceGame(t *testing.T) {
	suite.Run(t, new(PlayerTest))
}

type DiceStub struct {
	mock.Mock
}

func (self DiceStub) Roll() int {
	args := self.Called()
	return args.Int(0)
}
