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

	player.Leave()

	c.Assert(game.HasPlayer(player), Equals, false)
}

func (s *CasinoTestsSuite) Test_Player_CanNotLeave_Game_UntilJoined(c *C) {
	player := &Player{}

	err := player.Leave()

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please join the game before leaving")
}

func (s *CasinoTestsSuite) Test_Player_CanNotJoin_AnotherGame_UntilLeftFirstGame(c *C) {
	game := &Game{}
	player := &Player{}
	player.Join(game)

	anotherGame := &Game{}
	err := player.Join(anotherGame)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Player is already in game")
}

func (s *CasinoTestsSuite) Test_Player_HasNoChips_ByDefailt(c *C) {
	player := &Player{}

	c.Assert(player.Balance(), Equals, Chips(0))
}

func (s *CasinoTestsSuite) Test_Player_Buy1Chip_Has1Chip(c *C) {
	player := &Player{}

	player.Buy(Chips(1))

	c.Assert(player.Balance(), Equals, Chips(1))
}

func (s *CasinoTestsSuite) Test_Player_BuyChipsTwice_HasSumOfChips(c *C) {
	player := &Player{}

	player.Buy(Chips(1))
	player.Buy(Chips(2))

	c.Assert(player.Balance(), Equals, Chips(1+2))
}

func (s *CasinoTestsSuite) Test_Player_CanBet(c *C) {
	player := &Player{}
	player.Buy(Chips(1))
}
