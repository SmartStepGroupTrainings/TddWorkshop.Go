package casino_new

import (
	"testing"
)

func TestPlayer_Join_WhenNewPlayer_Success(t *testing.T) {
	p := NewPlayer()
	if err := p.Join(NewRollDiceGame()); err != nil {
		t.Fatal()
	}
}

func TestPlayer_InGame_WhenJoinGame_Success(t *testing.T) {
	p := NewPlayer()
	p.Join(NewRollDiceGame())
	if !p.IsInGame() {
		t.Fatal()
	}
}

func TestPlayer_Join_WhenInGame_Fail(t *testing.T) {
	p := NewPlayer()
	p.Join(NewRollDiceGame())
	if err := p.Join(NewRollDiceGame()); err == nil {
		t.Fatal()
	}
}

func TestPlayer_Leave_Fails_WhenNotInGame(t *testing.T) {
	p := NewPlayer()
	if err := p.Leave(); err == nil {
		t.Fatal()
	}
}

func TestPlayer_Leave_Success_WhenInGame(t *testing.T) {
	p := NewPlayer()
	p.Join(NewRollDiceGame())
	if err := p.Leave(); err != nil {
		t.Fatal()
	}
}

func TestPlayer_Leave_RefundChips_WhenInGame(t *testing.T) {
	p := NewPlayer()
	p.Join(NewRollDiceGame())
	p.BuyChips(3)
	p.Bet(Bet{Score: 1, Amount: 1})
	p.Leave()
	if p.AvailableChips() != 3 {
		t.Fatal()
	}
}

func TestPlayer_NotInGame_WhenNewPlayer(t *testing.T) {
	p := NewPlayer()
	if p.IsInGame() {
		t.Fatal()
	}
}

func TestPlayer_HasNoChips_WhenNewPlayer(t *testing.T) {
	p := NewPlayer()
	if p.AvailableChips() != 0 {
		t.Fatal()
	}
}

func TestPlayer_BuyChips_ShouldHaveOneChip(t *testing.T) {
	p := NewPlayer()
	err := p.BuyChips(1)
	if err != nil || p.AvailableChips() != 1 {
		t.Fatal()
	}
}

func TestPlayer_BuyChips_CannotBuyZeroChips(t *testing.T) {
	p := NewPlayer()
	err := p.BuyChips(0)
	if err == nil {
		t.Fatal()
	}
}

func TestPlayer_BuyChips_CannotBuyNegativeChips(t *testing.T) {
	p := NewPlayer()
	err := p.BuyChips(-1)
	if err == nil {
		t.Fatal()
	}
}

func TestPlayer_BuyChips_ChangesAvailableChips(t *testing.T) {
	p := NewPlayer()
	p.BuyChips(1)
	err := p.BuyChips(5)
	if err != nil || p.AvailableChips() != 6 {
		t.Fatal()
	}
}

func TestPlayer_Bet_Fails_WhenHasNoChips(t *testing.T) {
	p := NewPlayer()
	if err := p.Bet(Bet{1, 1}); err == nil {
		t.Fatal()
	}
}

func TestPlayer_Bet_Fails_WhenNotEnoughChips(t *testing.T) {
	p := NewPlayer()
	p.BuyChips(10)
	p.Bet(Bet{Score: 3, Amount: 5})
	p.Bet(Bet{Score: 3, Amount: 3})

	if err := p.Bet(Bet{Score: 3, Amount: 7}); err == nil {
		t.Fatal()
	}
}
func TestPlayer_Bet_Fails_WhenBetHasZeroScoreAndAmount(t *testing.T) {
	p := NewPlayer()
	p.BuyChips(10)
	if err := p.Bet(Bet{0, 0}); err == nil {
		t.Fatal()
	}
}

func TestPlayer_Bet_DecreasesAvailableChips(t *testing.T) {
	p := NewPlayer()
	p.BuyChips(4)
	p.Bet(Bet{Score: 3, Amount: 2})
	p.Bet(Bet{Score: 1, Amount: 1})
	if p.AvailableChips() != 1 {
		t.Fatal()
	}
}

func TestPlayer_GetBetOn_SumsBets(t *testing.T) {
	p := NewPlayer()
	p.BuyChips(10)
	p.Bet(Bet{Score: 3, Amount: 4})
	p.Bet(Bet{Score: 1, Amount: 1})
	p.Bet(Bet{Score: 3, Amount: 2})
	if p.GetBetOn(3) != 6 {
		t.Fatal()
	}
}