package casino_new

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreatingNewPlayer (t *testing.T) {
	player := NewPlayer()

	assert.NotNil(t, player)

	assert.Nil(t, player.currentGame)
	assert.Equal(t, 0, player.availableChips)
	assert.Equal(t, map[int]int{}, player.bets)
}

type testDice struct {

}

func (dice testDice) Roll() int {
	return 0
}

func TestJoinGame(t *testing.T) {
	player := NewPlayer()
	game := NewRollDiceGame(testDice{})

	err := player.Join(game)

	assert.Nil(t, err)
	assert.Equal(t, game, player.currentGame)

	_, exists := game.players[player]
	assert.Equal(t, true, exists)
}


func TestJoinGameFail(t *testing.T) {
	player := NewPlayer()
	game := NewRollDiceGame(testDice{})

	err := player.Join(game)
	assert.Nil(t, err)

	err = player.Join(game)
	assert.NotNil(t, err)
	assert.Equal(t, "Unable to join another game", err.Error())

}