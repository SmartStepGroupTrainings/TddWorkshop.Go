package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player_CanJoinGame(t *testing.T) {
	player := Player{}

	assert.True(t, player.CanJoinGame())
}

func Test_Player_CanLeaveGame(t *testing.T) {
	player := Player{}

	assert.True(t, player.CanLeaveGame())
}
