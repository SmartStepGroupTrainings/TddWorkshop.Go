package casino_new

import (
	"testing"

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

	s.Equal(false, res, "New player should not be in game")
}

func (s *TestPlayerSuite) TestPlayer_IsInGameCheck_NotEmptyPlayer_ShouldSuccess() {
	g := NewRollDiceGame(nil)
	p := Player{currentGame: g}

	res := p.IsInGame()

	s.Equal(true, res, "This player should be in game")
}

func (s *TestPlayerSuite) TestPlayer_GetAvailableChips_PlayerWithChips_ShouldNotNull() {
	anyChips := 100
	p := NewPlayer()
	p.BuyChips(anyChips)

	availableChips := p.AvailableChips()

	s.Equal(anyChips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsNegativeCount_Player_ExpectError() {
	p := NewPlayer()
	initialChips := p.AvailableChips()

	err := p.BuyChips(-1)
	availableChips := p.AvailableChips()

	s.Error(err, "Didn't get expected error when buying -1 chip")
	s.Equal(initialChips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsZeroCount_Player_ExpectError() {
	p := NewPlayer()
	initialChips := p.AvailableChips()

	err := p.BuyChips(0)
	availableChips := p.AvailableChips()

	s.Error(err, "Didn't get expected error when buying -1 chip")
	s.Equal(initialChips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsPositiveValue_DefaultPlayer_ShouldIncreaseByCorrectValue() {
	chips := 100
	p := NewPlayer()

	err := p.BuyChips(chips)
	availableChips := p.AvailableChips()

	s.Nil(err, "Got unexpected error when buying positive chips amount")
	s.Equal(chips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_Bet_Player_CantBetMoreThanAvailableChips() {
	player := NewPlayer()
	player.BuyChips(20)
	bet := Bet{Amount: 30}

	err := player.Bet(bet)

	s.Error(err, "Error should be not nil")
	s.Equal("Unable to bet chips more than available", err.Error(), "Error message is not valid")
}
