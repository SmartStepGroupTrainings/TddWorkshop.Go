package casino_new

import (
	"testing"
)

func TestPlayer_IsInGameCheck_EmptyPlayer_ShouldFail(t *testing.T) {
	p := Player{}
	if p.IsInGame() {
		t.Error("New player should not be in game")
	}
}

func TestPlayer_IsInGameCheck_NotEmptyPlayer_ShouldSuccess(t *testing.T) {
	g := NewRollDiceGame(nil)
	p := Player{currentGame: g}
	if !p.IsInGame() {
		t.Error("This player should be in game")
	}
}

func TestPlayer_GetAvailableChips_PlayerWithChips_ShouldNotNull(t *testing.T) {
	chips := 100
	p := Player{availableChips: chips}
	if p.AvailableChips() != chips {
		t.Error("Player has wrong number of chips")
	}
}

func TestPlayer_BuyChipsNegativeCount_Player_ExpectError(t *testing.T) {
	p := Player{}
	initialChips := p.AvailableChips()
	if err := p.BuyChips(-1); err == nil {
		t.Error("Didn't get expected errory when buing -1 chip")
	}

	if p.AvailableChips() != initialChips {
		t.Error("Player has wrong number of chips")
	}
}

func TestPlayer_BuyChipsZeroCount_Player_ExpectError(t *testing.T) {
	p := Player{}
	initialChips := p.AvailableChips()

	if err := p.BuyChips(0); err == nil {
		t.Error("Didn't get expected error when buing 0 chip")
	}

	if p.AvailableChips() != initialChips {
		t.Error("Player has wrong number of chips")
	}
}

func TestPlayer_BuyChipsPositiveValue_DefaultPlayer_ShouldIncreaseByCorrectValue(t *testing.T) {
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
