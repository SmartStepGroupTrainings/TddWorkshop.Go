package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestPlayerSuite struct {
	suite.Suite
	p *Player
}

func (self *TestPlayerSuite) SetupTest() {
	self.p = NewPlayer()
}

func Test_Player(t *testing.T) {
	suite.Run(t, new(TestPlayerSuite))
}

func (self *TestPlayerSuite) Test_PlayerCan_JoinNewGame() {
	err := self.p.Join(NewRollDiceGame())

	assert.Nil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_InGame_WhenJoinGame_Success() {
	self.p.Join(NewRollDiceGame())

	assert.True(self.T(), self.p.IsInGame())
}

func (self *TestPlayerSuite) TestPlayer_JoinGame_Twice_Fail() {
	self.p.Join(NewRollDiceGame())

	err := self.p.Join(NewRollDiceGame())

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Leave_Fails_WhenNotInGame() {
	err := self.p.Leave()

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Leave_Success_WhenInGame() {
	self.p.Join(NewRollDiceGame())

	err := self.p.Leave()

	assert.Nil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Leave_RefundChips_WhenInGame() {
	self.p.Join(NewRollDiceGame())
	self.p.BuyChips(2)
	self.p.Bet(Bet{Score: 1, Amount: 1})

	self.p.Leave()

	assert.Equal(self.T(), 2-1, self.p.AvailableChips())
}

func (self *TestPlayerSuite) TestPlayer_NotInGame_WhenNewPlayer() {
	assert.False(self.T(), self.p.IsInGame())
}

func (self *TestPlayerSuite) TestPlayer_HasNoChips_WhenNewPlayer() {
	assert.Equal(self.T(), 0, self.p.AvailableChips())
}

func (self *TestPlayerSuite) TestPlayer_BuyChips_ShouldHaveOneChip() {
	self.p.BuyChips(1)

	assert.Equal(self.T(), 1, self.p.AvailableChips())
}

func (self *TestPlayerSuite) TestPlayer_BuyChips_CannotBuyZeroChips() {
	err := self.p.BuyChips(0)

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_BuyChips_CannotBuyNegativeChips() {
	err := self.p.BuyChips(-1)

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Buy1And5Chips_ChangesAvailableChipsTo6() {
	self.p.BuyChips(1)
	self.p.BuyChips(2)

	assert.Equal(self.T(), 1+2, self.p.AvailableChips())
}

func (self *TestPlayerSuite) TestPlayer_Bet_Fails_WhenHasNoChips() {
	err := self.p.Bet(Bet{1, 1})

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Bet_Fails_WhenNotEnoughChips() {
	self.p.BuyChips(1)

	err := self.p.Bet(Bet{Score: 1, Amount: 3})

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Bet_Fails_WhenBetHasZeroScoreAndAmount() {
	invalidBet := Bet{0, 0}
	self.p.BuyChips(10)

	err := self.p.Bet(invalidBet)

	assert.NotNil(self.T(), err)
}

func (self *TestPlayerSuite) TestPlayer_Bet_DecreasesAvailableChips() {
	self.p.BuyChips(2)

	self.p.Bet(Bet{Score: 1, Amount: 1})

	assert.Equal(self.T(), 2-1, self.p.AvailableChips())
}

func (self *TestPlayerSuite) TestPlayer_GetBetOn_SumsBets() {
	anotherBet := Bet{Score: 1, Amount: 2}
	self.p.BuyChips(10)

	self.p.Bet(Bet{Score: 2, Amount: 1})
	self.p.Bet(anotherBet)
	self.p.Bet(Bet{Score: 2, Amount: 3})

	assert.Equal(self.T(), 1+3, self.p.GetBetOn(2))
}
