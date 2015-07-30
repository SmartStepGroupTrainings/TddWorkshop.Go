package casino_new

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestPlayerSuite struct {
	suite.Suite
}

func TestRun(t *testing.T) {
	suite.Run(t, new(TestPlayerSuite))
}

func (suite *TestPlayerSuite) Test_CreateNewPlayer_Success() {
	player := NewPlayer()

	assert.False(suite.T(), player.IsInGame())
	assert.Equal(suite.T(), 0, player.AvailableChips())
}

func (suite *TestPlayerSuite) Test_JoinGame_Success() {
	player := NewPlayer()
	game := NewRollDiceGame()

	player.Join(game)

	assert.True(suite.T(), player.IsInGame())
}

func (suite *TestPlayerSuite) Test_JoinSimultaneouslySecondGame_Fail() {
	player := NewPlayer()
	game_one := NewRollDiceGame()
	game_two := NewRollDiceGame()
	player.Join(game_one)

	err := player.Join(game_two)

	assert.NotNil(suite.T(), err)
}

func (suite *TestPlayerSuite) Test_LeaveGame_Success() {
	player := NewPlayer()
	game := NewRollDiceGame()
	player.Join(game)

	player.Leave()

	assert.False(suite.T(), player.IsInGame())
}

func (suite *TestPlayerSuite) Test_LeaveGameBeforeJoin_Fail() {
	player := NewPlayer()

	err := player.Leave()

	assert.NotNil(suite.T(), err)
}
