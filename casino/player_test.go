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
	game := Game{}

	player.Join(game)

	assert.True(t, player.CanLeaveGame())
}

func Test_CannotLeaveFromTheGame_IfTheyNotJoin(t *testing.T) {
	player := Player{}

	err := player.Leave()

	assert.NotNil(t, err)
}

func Test_Player_CanPlayOnlyOneGameInTheSameTime(t *testing.T) {
	player := Player{}
	game := Game{}
	player.Join(game)

	err := player.Join(game)

	assert.NotNil(t, err)
}
