package casino_new

import (
	"testing"

	. "gopkg.in/check.v1"
	//	"fmt"
)

type CasinoTest struct {
}

var _ = Suite(&CasinoTest{})

func TestStart(t *testing.T) {
	TestingT(t)
}

func (suite *CasinoTest) TestCreateNewPlayer(c *C) {
	player := NewPlayer()
	c.Assert(player.currentGame, IsNil)
	c.Assert(player.availableChips, Equals, 0)
}

func (suite *CasinoTest) TestCreateGameWithPlayer(c *C) {
	player := NewPlayer()
	game := NewRollDiceGame()
	game.Add(player)

	current_player, ok := game.players[player]
	c.Assert(ok, Equals, true)
	c.Assert(current_player, Equals, player)
}
