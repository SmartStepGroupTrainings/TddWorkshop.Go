package casino

import (
	. "gopkg.in/check.v1"
	"testing"
)

type CasinoTestsSuite struct{}

var _ = Suite(&CasinoTestsSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *CasinoTestsSuite) Test_Player_Join_Game(c *C) {
	game := &Game{}
	player := &Player{}

	err := player.Join(game)

	c.Assert(err, IsNil)
	c.Assert(game.HasPlayer(player), Equals, true)
}

func (s *CasinoTestsSuite) Test_Game_HasPlayer_FalseForNotJoinedPlayer(c *C) {
	game := &Game{}
	player := &Player{}

	c.Assert(game.HasPlayer(player), Equals, false)
}

func (s *CasinoTestsSuite) Test_Player_CanNotJoin_NilGame(c *C) {
	player := &Player{}

	err := player.Join(nil)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Joining nil game is not allowed")
}

func (s *CasinoTestsSuite) Test_Player_JoinTwice_Fails(c *C) {
	game := &Game{}
	player := &Player{}

	player.Join(game)
	err := player.Join(game)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Player is already in game")
}

func (s *CasinoTestsSuite) Test_Player_Leaves_Game(c *C) {
	game := &Game{}
	player := &Player{}
	player.Join(game)

	player.Leave(game)

	c.Assert(game.HasPlayer(player), Equals, false)
}
