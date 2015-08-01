package casino
import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type GameTest struct {
	suite.Suite
}

func TestInit(t *testing.T) {
	suite.Run(t, &GameTest{})
}

func (test *GameTest) Test_Player_JoinToGame_Success() {
	player := &Player{}
	game := &Game{}

	player.JoinTo(game)

	test.True(player.IsInGame())
}

func (test *GameTest) Test_Player_Leave_Success() {
	player := &Player{}
	game := &Game{}
	player.JoinTo(game)

	player.LeaveGame()

	test.False(player.IsInGame())
}

func (test *GameTest) Test_PlayerNotInGame_LeaveGame_Fail() {
	player := &Player{}

	err := player.LeaveGame()

	test.NotNil(err)
}

func (test *GameTest) Test_Player_JoinTwiceInGame_Fail() {
	player := &Player{}
	game1 := &Game{}
	game2 := &Game{}
	player.JoinTo(game1)

	err := player.JoinTo(game2)

	test.NotNil(err)
}

func (test *GameTest) Test_Game_JoinMoreThen6Players_Fail() {
	game := &Game{}
	player1 := &Player{}
	player2 := &Player{}
	player3 := &Player{}
	player4 := &Player{}
	player5 := &Player{}
	player6 := &Player{}
	player7 := &Player{}
	player1.JoinTo(game)
	player2.JoinTo(game)
	player3.JoinTo(game)
	player4.JoinTo(game)
	player5.JoinTo(game)
	player6.JoinTo(game)

	err := player7.JoinTo(game)

	test.NotNil(err)
}
