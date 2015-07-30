package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestPlayerSuite struct {
	suite.Suite
	player *Player
}

func (self *TestPlayerSuite) SetupTest() {
	self.player = NewPlayer()
}

func Test_Player(t *testing.T) {
	suite.Run(t, new(TestPlayerSuite))
}

func (self *TestPlayerSuite) Test_PlayerCan_JoinNewGame() {
	err := self.player.Join(NewRollDiceGame())

	assert.Nil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_InGame_WhenJoinGame_Success() {
	self.player.Join(NewRollDiceGame())

	assert.True(self.T(), self.player.IsInGame())
}

func (self *TestPlayerSuite) TestPlayer_JoinGame_Twice_Fail() {
	self.player.Join(NewRollDiceGame())

	err := self.player.Join(NewRollDiceGame())

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Leave_Fails_WhenNotInGame() {
	err := self.player.Leave()

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Leave_Success_WhenInGame() {
	self.player.Join(NewRollDiceGame())

	err := self.player.Leave()

	assert.Nil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Leave_RefundChips_WhenInGame() {
	self.player.Join(NewRollDiceGame())
	self.player.BuyChips(2)
	self.player.Bet(Bet{Score: 1, Amount: 1})

	self.player.Leave()

	assert.Equal(self.T(), 2, self.player.AvailableChips())
}

func (self *TestPlayerSuite) TestPlayer_NotInGame_WhenNewPlayer() {
	assert.False(self.T(), self.player.IsInGame())
}

func (self *TestPlayerSuite) TestPlayer_HasNoChips_WhenNewPlayer() {
	assert.Equal(self.T(), 0, self.player.AvailableChips())
}

func (self *TestPlayerSuite) TestPlayer_BuyChips_ShouldHaveOneChip() {
	self.player.BuyChips(1)

	assert.Equal(self.T(), 1, self.player.AvailableChips())
}

func (self *TestPlayerSuite) TestPlayer_BuyChips_CannotBuyZeroChips() {
	err := self.player.BuyChips(0)

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_BuyChips_CannotBuyNegativeChips() {
	err := self.player.BuyChips(-1)

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Buy1And5Chips_ChangesAvailableChipsTo6() {
	self.player.BuyChips(1)
	self.player.BuyChips(2)

	assert.Equal(self.T(), 1+2, self.player.AvailableChips())
}

func (self *TestPlayerSuite) TestPlayer_Bet_Fails_WhenHasNoChips() {
	err := self.player.Bet(Bet{1, 1})

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Bet_Fails_WhenNotEnoughChips() {
	self.player.BuyChips(1)

	err := self.player.Bet(Bet{Score: 1, Amount: 3})

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Bet_Fails_WhenBetHasZeroScoreAndAmount() {
	invalidBet := Bet{0, 0}
	self.player.BuyChips(10)

	err := self.player.Bet(invalidBet)

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Bet_DecreasesAvailableChips() {
	self.player.BuyChips(2)

	self.player.Bet(Bet{Score: 1, Amount: 1})

	assert.Equal(self.T(), 2-1, self.player.AvailableChips())
}

func (self *TestPlayerSuite) TestPlayer_GetBetOn_SumsBets() {
	anotherBet := Bet{Score: 1, Amount: 2}
	self.player.BuyChips(10)

	self.player.Bet(Bet{Score: 2, Amount: 1})
	self.player.Bet(anotherBet)
	self.player.Bet(Bet{Score: 2, Amount: 3})

	assert.Equal(self.T(), 1+3, self.player.GetBetOn(2))
}
