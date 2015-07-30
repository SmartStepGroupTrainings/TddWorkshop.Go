package casino_new

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestNewPlayerNotInGame_Leave_Error(t *testing.T) {
	player := NewPlayer()

	err := player.Leave()
    assert.NotNil(t, err, "Return value is not null")
    assert.Equal(t, "Unable to leave the game before joining", err.Error())
}

func TestNewPlayerInGame_Leave_Success(t *testing.T) {

	player := NewPlayer()
	player.Join(NewRollDiceGame())

	err := player.Leave()
    assert.Nil(t, err, "Player error is not null")
}

func TestPlayerIsInGame(t *testing.T) {
	var isInGame bool

	player := NewPlayer()

	isInGame = player.IsInGame()
	if isInGame {
		t.Fatal("Player is in game")
	}

	player.Join(NewRollDiceGame())

	isInGame = player.IsInGame()
	if !isInGame {
		t.Fatal("Player is not in game")
	}
}

func TestPlayerBuyChips(t *testing.T) {
	var err error

	player := NewPlayer()
	player.Join(NewRollDiceGame())

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
