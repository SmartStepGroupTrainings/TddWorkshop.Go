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

func (test *PlayerTest) SetupTest() {
	test.player = &Player{}
	test.game = &Game{}
}

func (test *PlayerTest) Test_Player_CanJoinGame() {

	test.True(test.player.CanJoinGame())
}

func (test *PlayerTest) Test_Player_CanLeaveGame() {

	test.player.Join(test.game)

	test.True(test.player.CanLeaveGame())
}

func (test *PlayerTest) Test_Player_CannotLeaveGame_AfterLeave() {

	test.player.Join(test.game)

	test.player.Leave()

	test.False(test.player.CanLeaveGame())
}

func (test *PlayerTest) Test_CannotLeaveFromTheGame_IfTheyNotJoin() {

	err := test.player.Leave()

	test.NotNil(err)
}

func (test *PlayerTest) Test_Player_CanPlayOnlyOneGameInTheSameTime() {

	test.player.Join(test.game)

	err := test.player.Join(test.game)

	test.NotNil(err)
}

func (test *PlayerTest) Test_GameNotPlayWithMoreThan6Players() {
	player1 := &Player{}
	player2 := &Player{}
	player3 := &Player{}
	player4 := &Player{}
	player5 := &Player{}
	player6 := &Player{}

	player1.Join(test.game)
	player2.Join(test.game)
	player3.Join(test.game)
	player4.Join(test.game)
	player5.Join(test.game)
	player6.Join(test.game)

	extraPlayer := &Player{}
	err := extraPlayer.Join(test.game)

	test.NotNil(err)
}

func (test *PlayerTest) Test_Player_BuyCoin_Succes() {

	test.player.BuyCoin(1)

	test.Equal(1, test.player.Coins())
}

func (test *PlayerTest) Test_Player_BetOnlyIfHasCoin() {
	bet := Bet{Coins: 1}

	err := test.player.Bet(bet)

	test.NotNil(err)
}

func (test *PlayerTest) Test_Player_WithEnoughCoinBet_Success() {

	bet := Bet{Coins: 1}
	test.player.BuyCoin(1)

	err := test.player.Bet(bet)

	test.Nil(err)
}

func (test *PlayerTest) Test_Player_DecreaseCoins_AfterBet() {

	bet := Bet{Coins: 1}
	test.player.BuyCoin(1)

	test.player.Bet(bet)

	test.Equal(1-1, test.player.Coins())
}

func (test *PlayerTest) Test_Player_CanBetToDifferentScore() {

	test.player.BuyCoin(2)

	test.player.Bet(Bet{Coins: 1, Score: 1})
	err := test.player.Bet(Bet{Coins: 1, Score: 2})

	test.Nil(err)
}

func (test *PlayerTest) Test_Bet_CanBeOnly5th_FailIfNot() {
	bet := Bet{Coins: 1}

	valid := test.game.isValid(bet)

	test.False(valid)
}

func (test *PlayerTest) Test_Player_Bet5_Valid() {
	bet := Bet{Coins: 5}

	valid := test.game.isValid(bet)

	test.True(valid)
}

func (test *PlayerTest) Test_Player_CanBetOnlyScore1to6() {
	bet := Bet{Coins: 5, Score: 7}
	test.player.BuyCoin(5)

	err := test.player.Bet(bet)

	test.NotNil(err)
}

func (test *PlayerTest) Test_Player_BetValid_Success() {
	bet := Bet{Coins: 5, Score: 1}
	test.player.BuyCoin(5)

	err := test.player.Bet(bet)

	test.Nil(err)
}
