package casino

import (
	. "gopkg.in/check.v1"
	"testing"
)

type CasinoTestsSuite struct{}

var _ = Suite(&CasinoTestsSuite{})

func Test(t *testing.T) { TestingT(t) }

var SOME_SCORE = Score(1)
var SOME_CHIPS = Chips(1)

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
	player := create.Player().Please()

	err := player.Join(nil)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Joining nil game is not allowed")
}

func (s *CasinoTestsSuite) Test_Player_JoinTwice_Fails(c *C) {
	game := create.Game().Please()
	player := create.Player().Joined(game).Please()

	player.Join(game)
	err := player.Join(game)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Player is already in game")
}

func (s *CasinoTestsSuite) Test_Player_Leaves_Game(c *C) {
	game := create.Game().Please()
	player := create.Player().Joined(game).Please()

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

func (s *CasinoTestsSuite) Test_Player_CanNotBuyNegativeChips(c *C) {
	player := &Player{}

	err := player.Buy(Chips(-1))

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Buying negative chips is not allowed")
}

func (s *CasinoTestsSuite) Test_Player_BuyChipsTwice_HasSumOfChips(c *C) {
	player := &Player{}

	player.Buy(Chips(1))
	player.Buy(Chips(2))

	c.Assert(player.Balance(), Equals, Chips(1+2))
}

func (s *CasinoTestsSuite) Test_Player_CanBetChips(c *C) {
	player := &Player{}
	player.Buy(Chips(1))

	player.Bet(Chips(1), SOME_SCORE)

	c.Assert(player.Balance(), Equals, Chips(0))
}

func (s *CasinoTestsSuite) Test_Player_CanBetChipsOn1(c *C) {
	player := &Player{}
	player.Buy(SOME_CHIPS)

	err := player.Bet(SOME_CHIPS, Score(2))

	c.Assert(err, IsNil)
}

func (s *CasinoTestsSuite) Test_Player_CanBetOn6(c *C) {
	player := &Player{}
	player.Buy(SOME_CHIPS)

	err := player.Bet(SOME_CHIPS, Score(6))

	c.Assert(err, IsNil)
}

func (s *CasinoTestsSuite) Test_Player_CanNotBetOnNumbersLessThan1(c *C) {
	player := &Player{}
	player.Buy(SOME_CHIPS)

	err := player.Bet(SOME_CHIPS, Score(0))

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Bet only to numbers 1-6")
}

func (s *CasinoTestsSuite) Test_Player_CanNotBetOnNumbersMoreThan6(c *C) {
	player := &Player{}
	player.Buy(SOME_CHIPS)

	err := player.Bet(SOME_CHIPS, Score(7))

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Bet only to numbers 1-6")
}

func (s *CasinoTestsSuite) Test_Player_CanMakeMultipleBets(c *C) {
	player := &Player{}
	player.Buy(Chips(25))

	player.Bet(10, Score(1))
	player.Bet(11, Score(2))

	c.Assert(player.Balance(), Equals, Chips(25-10-11))
}
