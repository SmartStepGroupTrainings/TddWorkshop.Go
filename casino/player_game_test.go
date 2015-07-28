package casino_new

import (
	"github.com/stretchr/testify/assert"
  "testing"
)

func Test_PlayerCan_JoinGame(t *testing.T) {
	game := NewRollDiceGame(DiceStub{})
	player := &Player{}

	player.Join(game)

	assert.True(t, player.IsInGame())
}

func Test_PlayerCan_LeaveGame(t *testing.T) {
	game := NewRollDiceGame(DiceStub{})
	player := &Player{}

	player.Join(game)
	player.Leave()

	assert.False(t, player.IsInGame())
}

func Test_PlayerCanNot_JoinAnotherGame(t *testing.T) {
	game := NewRollDiceGame(DiceStub{})
	anotherGame := NewRollDiceGame(DiceStub{})
	player := &Player{}

	player.Join(game)
	err := player.Join(anotherGame)

	assert.NotNil(t, err)
	assert.Equal(t, "Unable to join another game", err.Error())
}

func Test_PlayerCanNot_LeaveGameBeforeJoining(t *testing.T) {
	player := &Player{}

	err := player.Leave()

	assert.NotNil(t, err)
	assert.Equal(t, "Unable to leave the game before joining", err.Error())
}
