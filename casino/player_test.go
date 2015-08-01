package casino

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PlayerTest struct {
	suite.Suite
	player *Player
	game   *Game
}

func TestPlayer(t *testing.T) {
	suite.Run(t, new(PlayerTest))
}

func (test *PlayerTest) Test_Player_CanJoinGame() {
	player := &Player{}

	test.True(player.CanJoinGame())
}

func (test *PlayerTest) Test_Player_CanLeaveGame() {
	player := &Player{}
	game := &Game{}

	player.Join(game)

	test.True(player.CanLeaveGame())
}

func (test *PlayerTest) Test_Player_CannotLeaveGame_AfterLeave() {
	player := &Player{}
	game := &Game{}
	player.Join(game)

	player.Leave()

	test.False(player.CanLeaveGame())
}

func (test *PlayerTest) Test_CannotLeaveFromTheGame_IfTheyNotJoin() {
	player := &Player{}

	err := player.Leave()

	test.NotNil(err)
}

func (test *PlayerTest) Test_Player_CanPlayOnlyOneGameInTheSameTime() {
	player := &Player{}
	game := &Game{}
	player.Join(game)

	err := player.Join(game)

	test.NotNil(err)
}

func (test *PlayerTest) Test_GameNotPlayWithMoreThan6Players() {
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

	test.NotNil(err)
}

func (test *PlayerTest) Test_Player_BuyCoin_Succes() {
	player := Player{}

	player.BuyCoin(1)

	test.Equal(1, player.Coins())
}

func (test *PlayerTest) Test_Player_BetOnlyIfHasCoin() {
	player := Player{}
	bet := Bet{Coins: 1}

	err := player.Bet(bet)

	test.NotNil(err)
}

func (test *PlayerTest) Test_Player_WithEnoughCoinBet_Success() {
	player := Player{}
	bet := Bet{Coins: 1}
	player.BuyCoin(1)

	err := player.Bet(bet)

	test.Nil(err)
}

func (test *PlayerTest) Test_Player_DecreaseCoins_AfterBet() {
	player := Player{}
	bet := Bet{Coins: 1}
	player.BuyCoin(1)

	player.Bet(bet)

	test.Equal(1-1, player.Coins())
}

func (test *PlayerTest) Test_Player_CanBetToDifferentScore() {
	player := Player{}
	player.BuyCoin(2)

	player.Bet(Bet{Coins: 1, Score: 1})
	err := player.Bet(Bet{Coins: 1, Score: 2})

	test.Nil(err)
}

func (test *PlayerTest) Test_Bet_CanBeOnly5th_FailIfNot() {
	game := Game{}
	bet := Bet{Coins: 1}

	valid := game.isValid(bet)

	test.False(valid)
}

func (test *PlayerTest) Test_Player_Bet5_Valid() {
	game := Game{}
	bet := Bet{Coins: 5}

	valid := game.isValid(bet)

	test.True(valid)
}
