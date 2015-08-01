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
