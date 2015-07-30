package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestPlayerSuite struct {
	suite.Suite
}

func TestMoverTestSuite(t *testing.T) {
	suite.Run(t, new(TestPlayerSuite))
}

func (s *TestPlayerSuite) TestPlayer_IsInGameCheck_EmptyPlayer_ShouldFail() {
	p := Player{}

	res := p.IsInGame()

	assert.Equal(s.T(), false, res, "New player should not be in game")
}

func (s *TestPlayerSuite) TestPlayer_IsInGameCheck_NotEmptyPlayer_ShouldSuccess() {
	g := NewRollDiceGame(nil)
	p := Player{currentGame: g}

	res := p.IsInGame()

	assert.Equal(s.T(), true, res, "This player should be in game")
}

func (s *TestPlayerSuite) TestPlayer_GetAvailableChips_PlayerWithChips_ShouldNotNull() {
	chips := 100
	p := NewPlayer()
	p.BuyChips(chips)

	availableChips := p.AvailableChips()

	assert.Equal(s.T(), chips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsNegativeCount_Player_ExpectError() {
	p := NewPlayer()
	initialChips := p.AvailableChips()

	err := p.BuyChips(-1)
	availableChips := p.AvailableChips()

	assert.Error(s.T(), err, "Didn't get expected error when buying -1 chip")
	assert.Equal(s.T(), initialChips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsZeroCount_Player_ExpectError() {
	p := NewPlayer()
	initialChips := p.AvailableChips()

	err := p.BuyChips(0)
	availableChips := p.AvailableChips()

	assert.Error(s.T(), err, "Didn't get expected error when buying -1 chip")
	assert.Equal(s.T(), initialChips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsPositiveValue_DefaultPlayer_ShouldIncreaseByCorrectValue() {
	chips := 100
	p := NewPlayer()

	err := p.BuyChips(chips)
	availableChips := p.AvailableChips()

	assert.Nil(s.T(), err, "Got unexpected error when buying positive chips amount")
	assert.Equal(s.T(), chips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_Bet_Player_CantBetMoreThanAvailableChips() {
	player := NewPlayer()
	player.BuyChips(20)
	bet := Bet{Amount: 30}

	err := player.Bet(bet)

	assert.Error(s.T(), err, "Error should be not nil")
	assert.Equal(s.T(), "Unable to bet chips more than available", err.Error(), "Error message is not valid")
}
