package casino
import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type GameTest struct {
	suite.Suite
}

func (test *GameTest) SetupTest() {
}

func TestInit(t *testing.T) {
	suite.Run(t, &GameTest{})
}

func (test *GameTest) Test_Player_JoinToGame_Success() {
	player := &Player{}
	game := &Game{}

	player.Join(game)

	test.True(player.IsInGame())
}

func (test *GameTest) Test_Player_Leave_Success() {
	player := &Player{}
	game := &Game{}
	player.Join(game)

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
	player.Join(game1)

	err := player.Join(game2)

	test.NotNil(err)
}


