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
	p := Player{availableChips: chips}

	availableChips := p.AvailableChips()

	assert.Equal(s.T(), chips, availableChips, "Player has wrong number of chips")
}

/*
func (s *TestPlayerSuite) TestPlayer_BuyChipsNegativeCount_Player_ExpectError() {
	p := Player{}
	initialChips := p.AvailableChips()
	err := p.BuyChips(-1)
		t.Error("Didn't get expected errory when buing -1 chip")

	if p.AvailableChips() != initialChips {
		t.Error("Player has wrong number of chips")
	}
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsZeroCount_Player_ExpectError() {
	p := Player{}
	initialChips := p.AvailableChips()

	if err := p.BuyChips(0); err == nil {
		t.Error("Didn't get expected error when buing 0 chip")
	}

	if p.AvailableChips() != initialChips {
		t.Error("Player has wrong number of chips")
	}
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsPositiveValue_DefaultPlayer_ShouldIncreaseByCorrectValue() {
	p := Player{}
	initialChips := p.AvailableChips()
	needToBuy := 400
	expectedResult := initialChips + needToBuy

	if err := p.BuyChips(needToBuy); err != nil {
		t.Errorf("Error ocured while buy chips for palyer: %s", err)
	}
	if current := p.AvailableChips(); current != expectedResult {
		t.Errorf("Total chips number is invalid for player after BuyChips(): expected: %d, have: %d", expectedResult, current)
	}
}
*/
