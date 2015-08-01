package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player_CanJoinGame(t *testing.T) {
	player := &Player{}

	assert.True(t, player.CanJoinGame())
}

func Test_Player_CanLeaveGame(t *testing.T) {
	player := &Player{}
	game := &Game{}

	player.Join(game)

	assert.True(t, player.CanLeaveGame())
}

func Test_CannotLeaveFromTheGame_IfTheyNotJoin(t *testing.T) {
	player := &Player{}

	err := player.Leave()

	assert.NotNil(t, err)
}

func Test_Player_CanPlayOnlyOneGameInTheSameTime(t *testing.T) {
	player := &Player{}
	game := &Game{}
	player.Join(game)

	err := player.Join(game)

	assert.NotNil(t, err)
}

func Test_GameNotPlayWithMoreThan6Players(t *testing.T) {
	player1 := &Player{}
	player2 := &Player{}
	player3 := &Player{}
	player4 := &Player{}
	player5 := &Player{}
	player6 := &Player{}
	game := &Game{}
	player1.Join(game)
	player2.Join(game)
	player3.Join(game)
	player4.Join(game)
	player5.Join(game)
	player6.Join(game)

	extraPlayer := &Player{}
	err := extraPlayer.Join(game)

	assert.NotNil(t, err)
}

func Test_Player_BuyCoin_Succes(t *testing.T) {
	player := Player{}

	player.BuyCoin(1)

	assert.Equal(t, 1, player.Coins())
}

func Test_Player_BetOnlyIfHasCoin(t *testing.T) {
	player := Player{}
	bet := Bet{Coins: 1}

	err := player.Bet(bet)

	assert.NotNil(t, err)
}

func Test_Player_WithEnoughCoinBet_Success(t *testing.T) {
	player := Player{}
	bet := Bet{Coins: 1}
	player.BuyCoin(1)

	err := player.Bet(bet)

	assert.Nil(t, err)
}

func Test_Player_DecreaseCoins_AfterBet(t *testing.T) {
	player := Player{}
	bet := Bet{Coins: 1}
	player.BuyCoin(1)

	player.Bet(bet)

	assert.Equal(t, 1-1, player.Coins())
}
