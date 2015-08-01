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
