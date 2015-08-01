package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_CanJoin(t *testing.T) {
	player := Player{}

	player.Join()

	assert.True(t, player.IsInGame(), "Player had to join game")
}

func TestPlayer_CanLeave(t *testing.T) {
	player := Player{}

	player.Leave()

	assert.False(t, player.IsInGame(), "Player had to leave game")
}

func TestPlayer_PlayerNotInGame_CantLeave(t *testing.T) {
	player := Player{}

	err := player.Leave()

	assert.Error(t, err, "Player not in game cant leave game")
}
