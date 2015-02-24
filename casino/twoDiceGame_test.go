package casino

import (
	. "gopkg.in/check.v1"
)

type TwoDiceGameSuite struct{}

var _ = Suite(&TwoDiceGameSuite{})

func (s *TwoDiceGameSuite) Test_Game_Bet_AcceptsBets2to12(c *C) {
	game := create.TwoDiceGame().Please()
	player := create.Player().Joined(game).Please()
	startBalance := player.Balance()

	player.Bet(Chips(1), Score(2))
	player.Bet(Chips(1), Score(3))
	player.Bet(Chips(1), Score(4))
	player.Bet(Chips(1), Score(5))
	player.Bet(Chips(1), Score(6))
	player.Bet(Chips(1), Score(7))
	player.Bet(Chips(1), Score(8))
	player.Bet(Chips(1), Score(9))
	player.Bet(Chips(1), Score(10))
	player.Bet(Chips(1), Score(11))
	player.Bet(Chips(1), Score(12))

	c.Assert(int(player.Balance()), Equals, int(startBalance-Chips(11)))
}

func (s *TwoDiceGameSuite) Test_Game_Bet_DoNotAcceptBetsLessThan2(c *C) {
	game := create.TwoDiceGame().Please()
	player := create.Player().Joined(game).Please()

	err := player.Bet(SOME_CHIPS, Score(1))

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please make a bet only to score 2 - 12")
}

func (s *TwoDiceGameSuite) Test_Game_Bet_DoNotAcceptBetsMoreThan12(c *C) {
	game := create.TwoDiceGame().Please()
	player := create.Player().Joined(game).Please()

	err := player.Bet(SOME_CHIPS, Score(13))

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please make a bet only to score 2 - 12")
}
