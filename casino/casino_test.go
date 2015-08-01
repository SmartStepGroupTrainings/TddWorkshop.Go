package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player_CanJoinInGame(t *testing.T) {
	player := NewPlayer()
	game := &Game{}

	player.Join(game)

	assert.True(t, player.IsInGame())
}

func Test_PlayerInGame_CanLeave(t *testing.T) {
	player := NewPlayer()
	game := &Game{}
	player.Join(game)

	player.Leave(game)

	assert.False(t, player.IsInGame())
}

func Test_PlayerNotInGame_CanNotLeave(t *testing.T) {
	player := NewPlayer()
	game := &Game{}

	err := player.Leave(game)

	assert.NotNil(t, err)
	assert.Equal(t, "You can not leave game", err.Error())
}

func Test_PlayerInGame_CanNotJoinInGame(t *testing.T) {
	player := NewPlayer()
	game := &Game{}
	player.Join(game)

	err := player.Join(game)

	assert.Error(t, err)
	assert.Equal(t, "You can not join game", err.Error())
}

func Test_PlayerNotInGame_CanNotJoinInFullGame(t *testing.T) {
	game := createGameWith6Player()

	err := NewPlayer().Join(game)

	assert.Error(t, err)
	assert.Equal(t, "Game is full", err.Error())
}

func Test_Player_CanBuyChips(t *testing.T) {
	player := NewPlayer()

	player.BuyChips(1)

	assert.Equal(t, 1, player.GetChipsCount())
}

func Test_PlayerInGame_CanBet(t *testing.T) {
	player := NewPlayer()
	game := &Game{}
	player.Join(game)
	player.BuyChips(1)

	player.Bet(1, 1)

	assert.Equal(t, 0, player.GetChipsCount())
}

func Test_PlayerInGame_CanNotBetMoreChipsThanAvailable(t *testing.T) {
	player := NewPlayer()
	game := &Game{}
	player.Join(game)

	err := player.Bet(1, 1)

	assert.Error(t, err)
	assert.Equal(t, "Not enough chips", err.Error())
}

func Test_PlayerInGameAndHaveChips_CanBetSeveralScore(t *testing.T) {
	player := NewPlayer()
	game := &Game{}
	player.Join(game)
	player.BuyChips(1 + 1)

	player.Bet(1, 1)
	player.Bet(1, 2)

	assert.True(t, player.HasBet(1))
	assert.True(t, player.HasBet(2))
}

func createGameWith6Player() *Game {
	game := &Game{}

	NewPlayer().Join(game)
	NewPlayer().Join(game)
	NewPlayer().Join(game)
	NewPlayer().Join(game)
	NewPlayer().Join(game)
	NewPlayer().Join(game)

	return game
}
