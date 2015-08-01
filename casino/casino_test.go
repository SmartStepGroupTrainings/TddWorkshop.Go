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

	assert.Equal(t, errPlayerNotInGame, err)
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

func TestPlayer_HasChips_CanCreateBet(t *testing.T) {
	player := Player{}
	player.BuyChips(5)
	game := &Game{}
	player.Join(game)

	player.DoBet(&Bet{Amount: 5, Score: 1})

	assert.NotNil(t, player.GetBetByScore(1))
	assert.Equal(t, 5, player.GetBetByScore(1).GetAmount())
}

func TestPlayer_HasChips_CanNotDoBetBecauseNotEnoughChips(t *testing.T) {
	player := Player{}
	player.BuyChips(1)
	game := &Game{}
	player.Join(game)

	err := player.DoBet(&Bet{Amount: 10, Score: 1})

	assert.Equal(t, errNotEnoughChips, err)

}

func TestPlayer_HasChips_CanDoSeveralBets(t *testing.T) {
	player := Player{}
	player.BuyChips(10)
	game := &Game{}
	player.Join(game)

	player.DoBet(&Bet{Amount: 5, Score: 1})
	player.DoBet(&Bet{Amount: 5, Score: 2})

	assert.Equal(t, 2, player.GetBetCount())

}

func TestPlayer_NotInGame_CanNotDoBets(t *testing.T) {
	player := Player{}
	player.BuyChips(5)

	err := player.DoBet(&Bet{Amount: 5, Score: 1})

	assert.Equal(t, errPlayerNotInGame, err)
}

func TestGame_ByDefault_CanNotGetBetsNotAliquot5(t *testing.T) {
	player := Player{}
	player.BuyChips(6)
	game := &Game{}
	player.Join(game)

	err := player.DoBet(&Bet{Amount: 6, Score: 1})

	assert.Equal(t, errBetNotAliquot5, err)
}
