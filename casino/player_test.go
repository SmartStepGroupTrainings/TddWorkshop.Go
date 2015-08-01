package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_CanJoinGame(t *testing.T) {
	player := Player{}
	game := Game{}

	err := game.Add(player)

	assert.Nil(t, err)
}

func TestPlayer_CanLeaveGame(t *testing.T) {
	player := Player{}
	game := Game{}
	game.Add(player)

	err := game.Remove(player)

	assert.Nil(t, err)
}

