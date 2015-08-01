package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PlayerByDefaultNotInGame(t *testing.T) {
	game := &Game{}
	player := Player{}

	assert.False(t, player.IsIn(game))
}

func Test_PlayerCanJoinGame(t *testing.T) {
	game := &Game{}
	player := Player{}

	player.Join(game)

	assert.True(t, player.IsIn(game))
}

func Test_PlayerCanLeaveGame(t *testing.T) {
	game := &Game{}
	player := Player{}
	player.Join(game)

	player.Leave()

	assert.False(t, player.IsIn(game))
}

func Test_PlayerCanNotJoinGameWith6Players(t *testing.T) {
	game := CreateGameWith6Players()
	extraPlayer := Player{}

	err := extraPlayer.Join(game)

	assert.NotNil(t, err)
	assert.Equal(t, "Game can not accept more than 6 players", err.Error())
}

func CreateGameWith6Players() *Game {
	game := &Game{}
	player1 := Player{}
	player2 := Player{}
	player3 := Player{}
	player4 := Player{}
	player5 := Player{}
	player6 := Player{}
	player1.Join(game)
	player2.Join(game)
	player3.Join(game)
	player4.Join(game)
	player5.Join(game)
	player6.Join(game)
	return game
}
