package casino

import (
	. "gopkg.in/check.v1"
)

type TwoDiceGameSuite struct{}

func (s *TwoDiceGameSuite) Test_Player_Joins_TwoDiceGame(c *C) {
	game := create.TwoDiceGame().Please()

	player := create.Player().Joined(game).Please()

	c.Assert(game.HasPlayer(player), Equals, true)
}
