package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PlayerByDefaultNotInGame(t *testing.T) {
	game := Game{}
	player := Player{}

	assert.False(t, player.IsIn(game))
}

type Player struct{}

func (self *Player) IsIn(game Game) bool {
	return false
}

type Game struct{}
