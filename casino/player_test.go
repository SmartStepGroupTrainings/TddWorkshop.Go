package casino_new

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestDice struct {
	nextValue int
}

func (d *TestDice) Roll() int {
	return d.nextValue
}

func TestPlayer_New_BetsIsNotNull(t *testing.T) {
	player := NewPlayer()

	assert.NotNil(t, player.bets)
}

func TestPlayer_Join_FirstJoin_Success(t *testing.T) {
	dice := &TestDice{1}
	game := NewRollDiceGame(dice)
	player := NewPlayer()

	assert.NoError(t, player.Join(game))
}

func TestPlayer_Join_PlayerInGame_Exist(t *testing.T) {
	dice := &TestDice{1}
	game := NewRollDiceGame(dice)
	player := NewPlayer()
	player.Join(game)

	_, exists := game.players[player]
	assert.True(t, exists)
}

func TestPlayer_Join_TwiceJoin_Failed(t *testing.T) {
	dice := &TestDice{1}
	game := NewRollDiceGame(dice)
	player := NewPlayer()
	player.Join(game)

	assert.Error(t, player.Join(game))
}

func TestPlayer_Leave_JoinedGame_Success(t *testing.T) {
	dice := &TestDice{1}
	game := NewRollDiceGame(dice)
	player := NewPlayer()
	player.Join(game)

	assert.NoError(t, player.Leave())
}

func TestPlayer_Leave_NotJoinedGame_Failed(t *testing.T) {
	player := NewPlayer()

	assert.Error(t, player.Leave())
}

func TestPlayer_IsInGame_JoinedGame_Success(t *testing.T) {
	dice := &TestDice{1}
	game := NewRollDiceGame(dice)
	player := NewPlayer()
	player.Join(game)

	assert.True(t, player.IsInGame())
}
