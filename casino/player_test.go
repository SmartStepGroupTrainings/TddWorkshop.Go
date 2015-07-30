package casino_new

import (
	"testing"
)

type Dice struct {
}

func (d Dice) Roll() int {
	return 1
}

func TestPlayerLevelInGame(t *testing.T) {
	var err error

	player := NewPlayer()

	err = player.Leave()
	if err == nil {
		t.Fatal("Empty error")
	}

	player.Join(NewRollDiceGame(Dice{}))

	err = player.Leave()
	if err != nil {
		t.Fatal("Leave error not empty")
	}
}

func TestPlayerIsInGame(t *testing.T) {
	var isInGame bool

	player := NewPlayer()

	isInGame = player.IsInGame()
	if isInGame {
		t.Fatal("Player is in game")
	}

	player.Join(NewRollDiceGame(Dice{}))

	isInGame = player.IsInGame()
	if !isInGame {
		t.Fatal("Player is not in game")
	}
}

func TestPlayerBuyChips(t *testing.T) {
	var err error

	player := NewPlayer()
	player.Join(NewRollDiceGame(Dice{}))

	err = player.BuyChips(-1)
	if err == nil {
		t.Fatal("Chips not valid")
	}

	err = player.BuyChips(22)
	if err != nil {
		t.Fatal("Chips not valid")
	}

	if player.AvailableChips() != 22 {
		t.Fatal("Chips not apply")
	}
}

func TestPlayerEmptyAvailableChips(t *testing.T) {
	player := NewPlayer()

	if player.AvailableChips() != 0 {
		t.Fatal("Chips != 0")
	}
}
