package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_CanJoin(t *testing.T) {
	player := Player{}
	game := &Game{}

	player.Join(game)

	assert.True(t, player.IsInGame(), "Player had to join game")
}

func TestPlayer_PlayerInGame_CanLeave(t *testing.T) {
	player := &Player{}
	game := &Game{}

	player.Join(game)
	player.Leave()

	assert.False(t, player.IsInGame(), "Player had to leave game")
}

func TestPlayer_PlayerNotInGame_CantLeave(t *testing.T) {
	player := &Player{}

	err := player.Leave()

	assert.Error(t, err, "Player not in game cant leave game")
}

func TestPlayer_Player_CanPlayOnlyOneGame(t *testing.T) {
	player := &Player{}
	game := &Game{}

	player.Join(game)
	err := player.Join(game)

	assert.Error(t, err, "Player already in game cant join game")
}

func TestPlayer_Player_CantJoinFullGame(t *testing.T) {

	extraPlayer := &Player{}
	game := &Game{}

	for i := 0; i < 6; i++ {
		player := Player{}
		player.Join(game)
	}

	err := extraPlayer.Join(game)

	assert.Error(t, err, "Game is full")
}

func TestPlayer_Player_CanBuyChips(t *testing.T) {
	player := Player{}

	player.BuyChips(10)

	assert.Equal(t, 10, player.AvailableChips())
}

/*
	DSL example
	var Jaffar geene
	Jaffar.Create().Game().With().Players(6).Wish()
*/

type geene struct {
	needGame  bool
	playerNum int

	game Game
}

func (g geene) Create() geene {
	// Nothing here :)
	return g
}

func (g geene) Game() geene {
	return g
}

func (g geene) With() geene {
	// Nothing here :)
	return g
}

func (g geene) Players(num int) geene {
	return g
}

func (g geene) Wish() geene {
	return g
}
