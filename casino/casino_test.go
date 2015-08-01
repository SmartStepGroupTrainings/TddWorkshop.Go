package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player_CanJoinInGame(t *testing.T) {
	player := Player{}
	game := &Game{}

	player.Join(game)

	assert.True(t, player.IsInGame())
}

func Test_PlayerInGame_CanLeave(t *testing.T) {
	player := Player{}
	game := &Game{}
	player.Join(game)

	player.Leave(game)

	assert.False(t, player.IsInGame())
}

func Test_PlayerNotInGame_CanNotLeave(t *testing.T) {
	player := Player{}
	game := &Game{}

	err := player.Leave(game)

	assert.NotNil(t, err)
	assert.Equal(t, "You can not leave game", err.Error())
}

func Test_PlayerInGame_CanNotJoinInGame(t *testing.T) {
	player := Player{}
	game := &Game{}
	player.Join(game)

	err := player.Join(game)

	assert.Error(t, err)
	assert.Equal(t, "You can not join game", err.Error())
}

func Test_PlayerNotInGame_CanNotJoinInFullGame(t *testing.T) {
	game := createGameWith6Player()

	err := (&Player{}).Join(game)

	assert.Error(t, err)
	assert.Equal(t, "Game is full", err.Error())
}

func Test_Player_CanBuyChips(t *testing.T) {
	player := Player{}

	player.BuyChips(1)

	assert.Equal(t, 1, player.GetChipsCount())
}

func createGameWith6Player() *Game {
	game := &Game{}

	(&Player{}).Join(game)
	(&Player{}).Join(game)
	(&Player{}).Join(game)
	(&Player{}).Join(game)
	(&Player{}).Join(game)
	(&Player{}).Join(game)

	return game
}
