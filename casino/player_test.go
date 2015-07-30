package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockDice struct {
	StaticRoll int
}

func (mockDice *MockDice) Roll() int {
	return mockDice.StaticRoll
}

var (
	player *Player
	rdGame *RollDiceGame
)

func init() {
	player = NewPlayer()
	rdGame = NewRollDiceGame(&MockDice{StaticRoll: 4})
}

func TestNewPlayer(t *testing.T) {
	assert.True(t, player != nil)
	assert.True(t, rdGame != nil)
}

func TestJoin(t *testing.T) {
	player.Join(rdGame)
	playerStruct, ok := rdGame.players[player]
	assert.True(t, ok)
	assert.Equal(t, struct{}{}, playerStruct)
}
func TestLeave(t *testing.T) {
	player := NewPlayer()
	err := player.Leave()
	assert.NotNil(t, err)

	player.Join(rdGame)
	err = player.Leave()
	assert.Nil(t, err)

	playerStruct, ok := rdGame.players[player]
	assert.False(t, ok)
	assert.Equal(t, struct{}{}, playerStruct)
}
