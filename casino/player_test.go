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

func (suite *TestPlayerSuite) Test_NewPlayer_Create_Success() {
	player := NewPlayer()

	assert.False(suite.T(), player.IsInGame())
	assert.Equal(suite.T(), 0, player.AvailableChips())
}

func (suite *TestPlayerSuite) Test_Player_JoinGame_Success() {
	player := NewPlayer()
	game := NewRollDiceGame()

	player.Join(game)

	assert.True(suite.T(), player.IsInGame())
}
