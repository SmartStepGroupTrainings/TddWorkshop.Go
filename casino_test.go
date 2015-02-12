package casino

import (
	"github.com/bevzuk/tdd/casino"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Player_ByDefault_NotInGame(t *testing.T) {
	player := casino.Player{}

	assert.False(t, player.IsInGame())
}

func Test_Player_JoinedGame_IsInGame(t *testing.T) {
	player := casino.Player{}
	game := casino.Game{}

	player.Join(game)

	assert.True(t, player.IsInGame())
}

func Test_Player_LeaveGame_IsNotInGame(t *testing.T) {
	player := casino.Player{}
	game := casino.Game{}
	player.Join(game)

	player.Leave()

	assert.False(t, player.IsInGame())
}

func Test_Player_CanNotLeaveGame_UntilJoin(t *testing.T) {
	player := casino.Player{}

	error := player.Leave()

	assert.Error(t, error)
	assert.Equal(t, "Please join the game before leaving", error.Error())
}

func Test_Player_CanNotJoinAnotherGame_WhileHeIsInGame(t *testing.T) {
	player := casino.Player{}
	player.Join(casino.Game{})

	error := player.Join(casino.Game{})

	assert.Error(t, error)
	assert.Equal(t, "Please leave the game before joining another game", error.Error())
}

func Ignored_Test_6Players_Join_Successfully(t *testing.T) {
	game := casino.Game{}

	player1 := casino.Player{}
	player1.Join(game)
	player2 := casino.Player{}
	player2.Join(game)
	player3 := casino.Player{}
	player3.Join(game)
	player4 := casino.Player{}
	player4.Join(game)
	player5 := casino.Player{}
	player5.Join(game)
	player6 := casino.Player{}
	player6.Join(game)

	assert.True(t, player6.IsInGame())
}
