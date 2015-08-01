package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_ByDefault_IsNotInGame(t *testing.T) {
	player := Player{}

	assert.False(t, player.IsInGame())
}

func TestPlayer_ByDefault_CanJoinGame(t *testing.T) {
	player := Player{}
	game := &Game{}

	player.Join(game)

	assert.True(t, player.IsInGame())
}

func TestPlayer_InGame_CanLeaveGame(t *testing.T) {
	player := Player{}
	game := &Game{}
	player.Join(game)

	player.Leave(game)

	assert.False(t, player.IsInGame())
}

func TestPlayer_NotInGame_CanNotLeaveGame(t *testing.T) {
	player := Player{}
	game := &Game{}

	err := player.Leave(game)

	assert.Equal(t, errPlayerAlreadyNotInGame, err)
}

func TestPlayer_InGame_CanNotJoinGame(t *testing.T) {
	player := Player{}
	game := &Game{}
	player.Join(game)

	err := player.Join(game)

	assert.Equal(t, errPlayerAlreadyInGame, err)
}

func TestPlayer_ByDefault_CanNotJoinFullGame(t *testing.T) {
	game := getFullGame()
	extraPlayer := Player{}

	err := extraPlayer.Join(game)

	assert.Equal(t, errGameIsFull, err)
}

func getFullGame() *Game {
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

func TestPlayer_ByDefault_HasNoChips(t *testing.T) {
	player := Player{}

	assert.Equal(t, 0, player.GetAvailableChips())
}

func TestPlayer_ByDefault_CanBuyChips(t *testing.T) {
	player := Player{}

	player.BuyChips(1)

	assert.Equal(t, 1, player.GetAvailableChips())
}

func TestPlayer_HasChips_CanBuyMoreChips(t *testing.T) {
	player := Player{}
	player.BuyChips(1)

	player.BuyChips(1)

	assert.Equal(t, 2, player.GetAvailableChips())
}

func TestPlayer_ByDefault_CanNotBuyNegativeChips(t *testing.T) {
	player := Player{}

	err := player.BuyChips(-1)

	assert.Equal(t, errBuyNegativeChips, err)
}

func TestPlayer_ByDefault_CanCreateBet(t *testing.T) {
	player := Player{}

	player.DoBet(&Bet{Amount: 1})

	assert.NotNil(t, player.GetBet())
	assert.Equal(t, 1, player.GetBet().GetAmount())
}

func TestPlayer_HasChips_CanNotDoBetBecauseNotEnoughChips(t *testing.T) {
    player := Player{}
    player.BuyChips(1)

    err := player.DoBet(&Bet{Amount:2})

    assert.Equal(t, errNotEnoughChips, err)

}