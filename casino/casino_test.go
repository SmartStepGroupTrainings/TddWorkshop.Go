package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player_CanJoin(t *testing.T) {
	player := Player{}
	game := Game{}

	player.Join(game)

	assert.True(t, player.isInGame)
}

func Test_PlayerInGame_CanLeave(t *testing.T) {
	player := Player{}
	game := Game{}
	player.Join(game)

	player.Leave(game)

	assert.False(t, player.isInGame)
}
