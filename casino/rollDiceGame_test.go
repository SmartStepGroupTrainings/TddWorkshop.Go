package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestGameSuite struct {
	suite.Suite
	game *RollDiceGame
}

func (self *TestGameSuite) SetupTest() {
	self.game = NewRollDiceGame()
}

func Test_Game(t *testing.T) {
	suite.Run(t, new(TestGameSuite))
}

func (self *TestGameSuite) TestGame_Play_injijminji() {

}
