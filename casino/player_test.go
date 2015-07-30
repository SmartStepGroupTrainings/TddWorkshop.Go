package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PlayerSuite struct {
	suite.Suite
	player *Player
}

func TestPlayerSuite(t *testing.T) {
	suite.Run(t, new(PlayerSuite))
}

func (suite *PlayerSuite) SetupTest() {
	suite.player = NewPlayer()
}

func (suite *PlayerSuite) TestPlayer_CreateNew_Success() {
	assert.False(suite.T(), suite.player.IsInGame())
	assert.Equal(suite.T(), 0, suite.player.AvailableChips())
}

func (suite *PlayerSuite) TestPlayer_JoinGame_Success() {
	game := NewRollDiceGame()

	suite.player.Join(game)

	assert.True(suite.T(), suite.player.IsInGame())
}

func (suite *PlayerSuite) TestPlayer_JoinSimultaneouslySecondGame_Fail() {
	game_one := NewRollDiceGame()
	game_two := NewRollDiceGame()
	suite.player.Join(game_one)

	err := suite.player.Join(game_two)

	assert.NotNil(suite.T(), err)
}

func (suite *PlayerSuite) TestPlayer_LeaveGame_Success() {
	game := NewRollDiceGame()
	suite.player.Join(game)

	suite.player.Leave()

	assert.False(suite.T(), suite.player.IsInGame())
}

func (suite *PlayerSuite) TestPlayer_LeaveGameBeforeJoin_Fail() {
	err := suite.player.Leave()

	assert.NotNil(suite.T(), err)
}
