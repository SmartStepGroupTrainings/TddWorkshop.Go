package casino

import (
	. "gopkg.in/check.v1"
)

type GameTestsSuite struct{}

var _ = Suite(&GameTestsSuite{})

func (s *GameTestsSuite) Test_Game_HasPlayer_TrueForJoinedPlayer(c *C) {
	game := create.Game().Please()
	playerInGame := create.Player().Joined(game).Please()

	c.Assert(game.HasPlayer(playerInGame), Equals, true)
}

func (s *GameTestsSuite) Test_Game_HasPlayer_FalseForNotJoinedPlayer(c *C) {
	game := create.Game().Please()
	playerNotInGame := create.Player().InGame().Please()

	c.Assert(game.HasPlayer(playerNotInGame), Equals, false)
}

func (s *GameTestsSuite) Test_Game_HasPlayer_FalseForJoinedAndLeftPlayer(c *C) {
	game := create.Game().Please()
	player := create.Player().Joined(game).Please()

	game.Remove(player)

	c.Assert(game.HasPlayer(player), Equals, false)
}

func (s *GameTestsSuite) Test_Game_HasPlayer_HandlesMultiplePlayers(c *C) {
	game := create.Game().Please()
	playerInGame1 := create.Player().Joined(game).Please()
	playerLeftGame := create.Player().Joined(game).Please()
	playerInGame2 := create.Player().Joined(game).Please()

	game.Remove(playerLeftGame)

	c.Assert(game.HasPlayer(playerInGame1), Equals, true)
	c.Assert(game.HasPlayer(playerInGame2), Equals, true)
	c.Assert(game.HasPlayer(playerLeftGame), Equals, false)
}

func (s *GameTestsSuite) Test_Game_Leave_FailsWhenLeaveTheSamePlayerTwice(c *C) {
	game := create.Game().Please()
	player := create.Player().Joined(game).Please()

	game.Remove(player)
	err := game.Remove(player)

	c.Assert(err, Not(IsNil))
}

func (s *GameTestsSuite) Test_Game_Add_ShouldNotAllowPlayerToJoinAnotherGame(c *C) {
	game := create.Game().Please()
	player := create.Player().Please()
	game.Add(player)

	anotherGame := create.Game().Please()
	err := anotherGame.Add(player)

	c.Assert(err, Not(IsNil))
}

func (s *GameTestsSuite) Test_Game_Play_PlayerWithWrongBetLoses(c *C) {
	game := create.Game().WithLuckyScore(6).Please()
	player := create.Player().Joined(game).Please()
	startBalance := player.Balance()
	player.Bet(Chips(10), Score(1))

	game.Play()

	c.Assert(int(player.Balance()), Equals, int(Chips(startBalance-10)))
}

func (s *GameTestsSuite) Test_Game_Play_PlayerWithLuckyBetWins6Bets(c *C) {
	game := create.Game().WithLuckyScore(3).Please()
	player := create.Player().Joined(game).Please()
	startBalance := player.Balance()
	player.Bet(Chips(10), Score(3))

	game.Play()

	c.Assert(int(player.Balance()), Equals, int(Chips(startBalance-10+10*6)))
}

func (s *GameTestsSuite) Test_Game_Play_PlayerWithTwoBets_WinOnlyLuckyBet(c *C) {
	game := create.Game().WithLuckyScore(2).Please()
	player := create.Player().Joined(game).Please()
	startBalance := player.Balance()
	player.Bet(Chips(10), Score(1))
	player.Bet(Chips(15), Score(2))

	game.Play()

	c.Assert(int(player.Balance()), Equals, int(Chips(startBalance-10-15+15*6)))
}
