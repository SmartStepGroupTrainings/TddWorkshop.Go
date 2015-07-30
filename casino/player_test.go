package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTest() (*Player, *RollDiceGame) {
	return NewPlayer(), NewRollDiceGame()
}

func TestPlayer_NewPlayer_IsNotNil(t *testing.T) {
	player, _ := setupTest()

	assert.NotNil(t, player)
}

func TestPlayer_NewPlayer_NotIsInGame(t *testing.T) {
	player, _ := setupTest()

	assert.False(t, player.IsInGame())
}

func TestPlayer_NewPlayer_Join_IsInGame(t *testing.T) {
	player, game := setupTest()

	player.Join(game)

	assert.True(t, player.IsInGame())
}

func TestPlayer_WhenLeaveFromNonYourGame_WithError(t *testing.T) {
	player, _ := setupTest()

	err := player.Leave()

	assert.NotNil(t, err)
}

func TestPlayer_WhenLeaveFromYourGame_WithoutError(t *testing.T) {
	player, game := setupTest()
	player.Join(game)

	err := player.Leave()

	assert.Nil(t, err)
}
