package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_CanJoinGame(t *testing.T) {
	player := &Player{}
	game := Game{}

	err := game.Add(player)

	assert.Nil(t, err)
}

func TestPlayer_CanLeaveGame(t *testing.T) {
	player := &Player{}
	game := Game{}
	game.Add(player)

	err := game.Remove(player)

	assert.Nil(t, err)
}

func TestPlayer_CantLeaveGameWhenNotJoined(t *testing.T) {
	player := &Player{}
	game := Game{}

	err := game.Remove(player)

	assert.Error(t, err)
}

func TestPlayer_CantJoinGameTwice(t *testing.T) {
	player := &Player{}
	game := Game{}
	game.Add(player)

	err := game.Add(player)

	assert.Error(t, err)
}

func TestPlayer_CanBuyChips(t *testing.T) {
	player := &Player{}
	player.BuyChips(1)

	assert.Equal(t, 1, player.ChipsCount())
}

func TestPlayer_HasNoChipsByDefault(t *testing.T) {
	player := &Player{}

	assert.Equal(t, 0, player.ChipsCount())
}

func TestPlayer_CanBet(t *testing.T) {
	const score = 1
	const amount = 1
	player := &Player{}
	game := &Game{}
	game.Add(player)
	player.BuyChips(amount)

	player.Bet(Bet{Score: score, Amount: amount})

	assert.Equal(t, amount, game.BetsOn(score))
}

func TestPlayer_CantBetMoreChipsThanHeHas(t *testing.T) {
	player := &Player{}
	game := &Game{}
	game.Add(player)
	const score = 1

	err := player.Bet(Bet{Score: score, Amount: 1})

	assert.Error(t, err)
}

func TestPlayer_ChipsCountDecreasesAfterBet(t *testing.T) {
	player := &Player{}
	game := &Game{}
	game.Add(player)
	const score = 1
	player.BuyChips(1)

	player.Bet(Bet{Score: score, Amount: 1})

	assert.Equal(t, 0, player.ChipsCount())
}

func TestGame_2PlayersCanJoinGame(t *testing.T) {
	game := Game{}
	game.Add(&Player{})

	err := game.Add(&Player{})

	assert.Nil(t, err)
}

func TestGame_7PlayersCantJoinGame(t *testing.T) {
	game := Game{}
	game.Add(&Player{})
	game.Add(&Player{})
	game.Add(&Player{})
	game.Add(&Player{})
	game.Add(&Player{})
	game.Add(&Player{})

	err := game.Add(&Player{})

	assert.Error(t, err)
}
