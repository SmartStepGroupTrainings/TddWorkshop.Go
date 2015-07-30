package casino_new

import (
	"testing"
)

type TestDice struct{}

func (dice *TestDice) Roll() int {
	return 0
}

func TestPlayerJoin(t *testing.T) {
	p := NewPlayer()
	game := NewRollDiceGame(&TestDice{})

	err := p.Join(game)
	if err != nil {
		t.Fatal("Player.Join failed")
	}
	if !p.IsInGame() {
		t.Fatal("Joined player should be in game")
	}
	if p.Join(game) == nil {
		t.Fatal("Repeated joine shuld returned error, got nil")
	}
}

func TestNewPlayer(t *testing.T) {
	p := NewPlayer()
	if p.IsInGame() {
		t.Fatal("New paler should be w/o game")
	}

	if p.AvailableChips() != 0 {
		t.Fatal("New player must have no chips")
	}
}

func TestBuyChips(t *testing.T) {
	p := NewPlayer()

	if p.AvailableChips() != 0 {
		t.Fatal("New player must have no chips")
	}

	err := p.BuyChips(1)
	if err != nil || p.AvailableChips() != 1 {
		t.Fatal("Player can't buy chips")
	}

	err = p.BuyChips(0)
	if err == nil {
		t.Fatal("Player can't buy zero chips")
	}

	err = p.BuyChips(-1)
	if err == nil {
		t.Fatal("Player didn't return error for BuyChipos(-1)")
	}

	err = p.BuyChips(5)
	if err != nil || p.AvailableChips() != 6 {
		t.Fatal("Player had 1 chip and bought 5 should have 6 chips.")
	}
}

func TestPlayerIncorrectBets(t *testing.T) {
	p := NewPlayer()
	if err := p.Bet(Bet{1, 1}); err == nil {
		t.Fatal("Player w/o chips can't bet")
	}

	p.BuyChips(10)
	if err := p.Bet(Bet{0, 0}); err == nil {
		t.Fatal("Player can't make invalid bets")
	}
}
