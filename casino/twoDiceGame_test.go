package casino

import (
	. "gopkg.in/check.v1"
)

type TwoDiceGameSuite struct{}

var _ = Suite(&TwoDiceGameSuite{})

func (s *TwoDiceGameSuite) Test_Player_Joins_TwoDiceGame(c *C) {
	game := create.TwoDiceGame().Please()

	player := create.Player().Joined(game).Please()

	c.Assert(game.HasPlayer(player), Equals, true)
}

func (s *TwoDiceGameSuite) Test_Player_CanBetOn7(c *C) {
	game := create.TwoDiceGame().Please()
	player := create.Player().Joined(game).Please()

	c.Assert(player.Bet(SOME_CHIPS, 7), IsNil)
}

func (s *TwoDiceGameSuite) Test_Player_CanBetOn12(c *C) {
	game := create.TwoDiceGame().Please()
	player := create.Player().Joined(game).Please()

	c.Assert(player.Bet(SOME_CHIPS, 12), IsNil)
}

func (s *TwoDiceGameSuite) Test_Player_CanNotBetOn13(c *C) {
	game := create.TwoDiceGame().Please()
	player := create.Player().Joined(game).Please()

	err := player.Bet(SOME_CHIPS, 13)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Bet only to numbers 2-12")
}

func (s *TwoDiceGameSuite) Test_Player_CanNotBetOn1(c *C) {
	game := create.TwoDiceGame().Please()
	player := create.Player().Joined(game).Please()

	err := player.Bet(SOME_CHIPS, 1)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Bet only to numbers 2-12")
}
