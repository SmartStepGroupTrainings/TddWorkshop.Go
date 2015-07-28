package casino_new

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PlayerTestSuite struct {
	suite.Suite
	player *Player
}

func (suite *PlayerTestSuite) SetupTest() {
	suite.player = NewPlayer()
}

func (s *PlayerTestSuite) TestSuite_Player_HasNoBets_ByDefault() {
	assert.Equal(s.T(), 0, s.player.GetBetOn(1))
	assert.Equal(s.T(), 0, s.player.GetBetOn(2))
	assert.Equal(s.T(), 0, s.player.GetBetOn(3))
	assert.Equal(s.T(), 0, s.player.GetBetOn(4))
	assert.Equal(s.T(), 0, s.player.GetBetOn(5))
	assert.Equal(s.T(), 0, s.player.GetBetOn(6))
}

func (s *PlayerTestSuite) TestSuite_Player_Bet_AddsBet() {
	s.player.BuyChips(10)
	bet := Bet{ Score: 1, Amount: 10 }

	s.player.Bet(bet)

	assert.Equal(s.T(), 10, s.player.GetBetOn(1))
}

func (s *PlayerTestSuite) Test_Player_HasNoAvailableChips_ByDefault() {
	assert.Equal(s.T(), 0, s.player.AvailableChips())
}

func (s *PlayerTestSuite) Test_Player_BuyChips_AddsAvailableChips() {
	s.player.BuyChips(1)

	assert.Equal(s.T(), 1, s.player.AvailableChips())
}

func (s *PlayerTestSuite) Test_Player_BuyMoreChips_AddsAvailableChips() {
	s.player.BuyChips(1)

	s.player.BuyChips(1)

	assert.Equal(s.T(), 1 + 1, s.player.AvailableChips())
}

func (s *PlayerTestSuite) Test_Player_BuyZeroChipa_NotAllowed() {
	err := s.player.BuyChips(0)

	assert.Equal(s.T(), 0, s.player.AvailableChips())
	require.NotNil(s.T(), err)
	assert.Equal(s.T(), "Please buy positive amount", err.Error())
}

func (s *PlayerTestSuite) Test_Player_BuyNegativeChips_NotAllowed() {
	err := s.player.BuyChips(-1)

	assert.Equal(s.T(), 0, s.player.AvailableChips())
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), "Please buy positive amount", err.Error())
}

func (s *PlayerTestSuite) Test_Player_BetChips_ReducesAvailableChips() {
	s.player.BuyChips(10)

	s.player.Bet(Bet { Amount: 1, Score: 1 })

	assert.Equal(s.T(), 10 - 1, s.player.AvailableChips())
}

func (s *PlayerTestSuite) Test_Player_BetChipsMoreThanAvailable_NotAllowed() {
	s.player.BuyChips(1)

	err := s.player.Bet(Bet { Amount: 2 })

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "Unable to bet chips more than available")
}

func (s *PlayerTestSuite) Test_Player_BetChipsOnSameScore_SumsUpAmount() {
	s.player.BuyChips(10)

	s.player.Bet(Bet { Amount: 1, Score: 5 })
	s.player.Bet(Bet { Amount: 2, Score: 5 })

	assert.Equal(s.T(), 1 + 2, s.player.GetBetOn(5))
}

func (s *PlayerTestSuite) Test_Player_BetOn1_Allowed() {
	s.player.BuyChips(10)

	s.player.Bet(Bet { Score: 1, Amount: 10 })

	assert.Equal(s.T(), 10, s.player.GetBetOn(1))
}

func (s *PlayerTestSuite) Test_Player_BetOn0_NotAllowed() {
	s.player.BuyChips(1)

	err := s.player.Bet(Bet { Score: 0, Amount: 1})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "Bets on 1..6 only are allowed")
}

func (s *PlayerTestSuite) Test_Player_BetOn7_NotAllowed() {
	s.player.BuyChips(1)

	err := s.player.Bet(Bet { Score: 7, Amount: 1})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), err.Error(), "Bets on 1..6 only are allowed")
}

func TestPlayerTestSuite(t *testing.T) {
	suite.Run(t, new(PlayerTestSuite))
}
