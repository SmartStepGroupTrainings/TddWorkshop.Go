package casino_new

import (
	"testing"
)

func TestIsInGame_EmptyPlayer_ExpectNotInGame(t *testing.T) {
	p := Player{}
	if p.IsInGame() {
		t.Error("New player should not be in game")
	}
}

func TestIsInGame_PlayerInGame_ExpectInGame(t *testing.T) {
	g := NewRollDiceGame(nil)
	p := Player{currentGame: g}
	if !p.IsInGame() {
		t.Error("This player should be in game")
	}
}

func TestAvailableChips(t *testing.T) {
	chips := 100
	p := Player{availableChips: chips}
	if p.AvailableChips() != chips {
		t.Error("Player has wrong number of chips")
	}
}

func TestBuyChips_NegativeValue_ExpectError(t *testing.T) {
	p := Player{}
	initialChips := p.AvailableChips()
	if err := p.BuyChips(-1); err == nil {
		t.Error("Didn't get expected errory when buing -1 chip")
	}

	if err := p.BuyChips(0); err == nil {
		t.Error("Didn't get expected error when buing 0 chip")
	}

	if p.AvailableChips() != initialChips {
		t.Error("Player has wrong number of chips")
	}
}
