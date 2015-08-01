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
