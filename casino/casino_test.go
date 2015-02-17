package casino

import (
	. "gopkg.in/check.v1"
	"testing"
)

const SOME_SCORE = Score(1)
const SOME_CHIPS = Chips(1)

func Test(t *testing.T) { TestingT(t) }

type CasinoTestsSuite struct{}

var _ = Suite(&CasinoTestsSuite{})

func (s *CasinoTestsSuite) Test_Player_ByDefault_NotInGame(c *C) {
	player := Player{}

	c.Assert(player.IsInGame, Equals, false)
}

func (s *CasinoTestsSuite) Test_Player_JoinedGame_IsInGame(c *C) {
	player := create.Player().Please()
	game := create.Game().Please()

	game.Add(player)

	c.Assert(player.IsInGame, Equals, true)
}

func (s *CasinoTestsSuite) Test_Player_LeaveGame_IsNotInGame(c *C) {
	game := create.Game().Please()
	player := create.Player().Joined(game).Please()

	game.Remove(player)

	c.Assert(player.IsInGame, Equals, false)
}

func (s *CasinoTestsSuite) Test_Player_CanNotLeaveGame_UntilJoin(c *C) {
	player := create.Player().Please()

	anotherGame := create.Game().Please()
	err := anotherGame.Remove(player)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please join the game before leaving")
}

func (s *CasinoTestsSuite) Test_Player_CanNotJoinAnotherGame_WhileHeIsInGame(c *C) {
	player := create.Player().Please()
	game := create.Game().Please()
	game.Add(player)

	anotherGame := create.Game().Please()
	err := anotherGame.Add(player)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please leave the game before joining another game")
}

func (s *CasinoTestsSuite) Test_6Players_Join_Successfully(c *C) {
	game := create.Game().Please()

	player1 := create.Player().Please()
	game.Add(player1)
	player2 := create.Player().Please()
	game.Add(player2)
	player3 := create.Player().Please()
	game.Add(player3)
	player4 := create.Player().Please()
	game.Add(player4)
	player5 := create.Player().Please()
	game.Add(player5)
	player6 := create.Player().Please()
	game.Add(player6)

	c.Assert(player6.IsInGame, Equals, true)
}

func (s *CasinoTestsSuite) Test_7thPlayers_Join_Fails(c *C) {
	game := create.Game().Please()
	player1 := create.Player().Please()
	game.Add(player1)
	player2 := create.Player().Please()
	game.Add(player2)
	player3 := create.Player().Please()
	game.Add(player3)
	player4 := create.Player().Please()
	game.Add(player4)
	player5 := create.Player().Please()
	game.Add(player5)
	player6 := create.Player().Please()
	game.Add(player6)

	player7 := create.Player().Please()
	game.Add(player7)

	c.Assert(player7.IsInGame, Equals, false)
}

func (s *CasinoTestsSuite) Test_7thPlayers_Join_FailsWithError(c *C) {
	game := create.Game().Please()
	player1 := create.Player().Please()
	game.Add(player1)
	player2 := create.Player().Please()
	game.Add(player2)
	player3 := create.Player().Please()
	game.Add(player3)
	player4 := create.Player().Please()
	game.Add(player4)
	player5 := create.Player().Please()
	game.Add(player5)
	player6 := create.Player().Please()
	game.Add(player6)

	player7 := create.Player().Please()
	err := game.Add(player7)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please join another game")
}

func (s *CasinoTestsSuite) Test_Player_Balance_IsZeroByDefault(c *C) {
	player := Player{}

	c.Assert(player.Balance(), Equals, Chips(0))
}

func (s *CasinoTestsSuite) Test_Player_BuyChips_IncreasesBalance(c *C) {
	player := Player{}

	player.Buy(Chips(10))

	c.Assert(player.Balance(), Equals, Chips(10))
}

func (s *CasinoTestsSuite) Test_Player_BetChips_DecreasesBalance(c *C) {
	player := create.Player().InGame().With(Chips(10)).Bet(1).Please()

	c.Assert(player.Balance(), Equals, Chips(9))
}

func (s *CasinoTestsSuite) Test_Player_CanNotBetChips_MoreThanHeHas(c *C) {
	player := create.Player().InGame().With(Chips(10)).Please()

	err := player.Bet(Chips(11), SOME_SCORE)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "You can not bet more than 10 chips")
}

func (s *CasinoTestsSuite) Test_Player_CanNotBetChips_UntilHeJoinedGame(c *C) {
	player := create.Player().With(Chips(1)).Please()

	err := player.Bet(Chips(1), SOME_SCORE)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "You should join a game before making a bet")
}

func (s *CasinoTestsSuite) Test_Player_CanNotBet_On7(c *C) {
	player := create.Player().InGame().Please()

	invalidScore := Score(7)
	err := player.Bet(SOME_CHIPS, invalidScore)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please make a bet only to score 1 - 6")
}

func (s *CasinoTestsSuite) Test_Player_CanNotBet_On0(c *C) {
	player := create.Player().InGame().Please()

	invalidScore := Score(0)
	err := player.Bet(SOME_CHIPS, invalidScore)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please make a bet only to score 1 - 6")
}

func (s *CasinoTestsSuite) Test_Player_BetChips_OnDifferentScores(c *C) {
	player := create.Player().InGame().Please()
	startBalance := player.Balance()

	player.Bet(Chips(1), Score(1))
	player.Bet(Chips(2), Score(2))
	player.Bet(Chips(3), Score(6))

	c.Assert(player.Balance(), Equals, Chips(startBalance-1-2-3))
}

func (s *CasinoTestsSuite) Test_Player_GetBets_ReturnsSingleBet(c *C) {
	player := create.Player().InGame().Please()
	player.Bet(Chips(10), Score(2))

	playerBets := player.Bets()

	c.Assert(len(playerBets), Equals, 1)
	c.Assert(playerBets[0].Chips, Equals, Chips(10))
	c.Assert(playerBets[0].Score, Equals, Score(2))
}

func (s *CasinoTestsSuite) Test_Player_GetBets_ReturnsTwoBets(c *C) {
	player := create.Player().InGame().Please()
	player.Bet(Chips(10), Score(2))
	player.Bet(Chips(11), Score(3))

	playerBets := player.Bets()

	c.Assert(len(playerBets), Equals, 2)
	c.Assert(playerBets[0].Chips, Equals, Chips(10))
	c.Assert(playerBets[0].Score, Equals, Score(2))
	c.Assert(playerBets[1].Chips, Equals, Chips(11))
	c.Assert(playerBets[1].Score, Equals, Score(3))
}

// func (s *CasinoTestsSuite) Ignore_Test_Player_Loses_WhenMakeWrongBet(c *C) {
// 	game := create.Game().WithLuckyScore(6).Please()
// 	player := create.Player().Joined(game).Bet(10).OnScore(1).Please()
// 	startBalance := player.Balance()

// 	game.Play()

// 	c.Assert(player.Balance(), Equals, Chips(startBalance-10))
// }
