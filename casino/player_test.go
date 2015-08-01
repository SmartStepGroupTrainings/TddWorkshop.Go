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


