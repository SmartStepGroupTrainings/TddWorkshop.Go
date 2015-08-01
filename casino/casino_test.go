package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player_Join_InGameSuccess(t *testing.T) {
	player := Player{}
	game := Game{}

	player.Join(game)

	assert.True(t, player.isInGame)
}
