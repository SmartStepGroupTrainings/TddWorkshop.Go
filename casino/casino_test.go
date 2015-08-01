package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_ByDefault_IsNotInGame(t *testing.T) {
	player := Player{}

	assert.False(t, player.IsInGame())
}

func TestPlayer_ByDefault_CanJoinGame(t *testing.T) {
	player := Player{}
	game := Game{}

	player.Join(game)

	assert.True(t, player.IsInGame())
}

func TestPlayer_InGame_CanLeaveGame(t *testing.T) {
	player := Player{}
	game := Game{}
	player.Join(game)

	player.Leave(game)

	assert.False(t, player.IsInGame())
}
