package casino_new

import (
	"testing"
)

type TestDice struct {
	nextValue int
}

func (d *TestDice) Roll() int {
	return d.nextValue
}

func TestNewPlayer(t *testing.T) {
	player := NewPlayer()

	if player.bets == nil {
		t.Fatalf("Map player bets is nil")
	}
}

func TestJoin(t *testing.T) {
	dice := &TestDice{1}
	game := NewRollDiceGame(dice)
	player := NewPlayer()

	if err := player.Join(game); err != nil {
		t.Error(err)
	}

	// First join
	if _, exists := game.players[player]; !exists {
		t.Errorf("Player wasn't added")
	}

	// The same player join twice
	if err := player.Join(game); err == nil {
		t.Errorf("Player can join twice")
	} else if err.Error() != "Unable to join another game" {
		t.Errorf("Invalid error message")
	}
}

func TestLeave(t *testing.T) {
	dice := &TestDice{1}
	game := NewRollDiceGame(dice)
	player := NewPlayer()

	if err := player.Join(game); err != nil {
		t.Error(err)
	}

	if err := player.Leave(); err != nil {
		t.Error(err)
	}

	// ....
}
