package casino

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type CasinoTestsSuite struct{}

var _ = Suite(&CasinoTestsSuite{})

func (s *CasinoTestsSuite) Test_Player_ByDefault_NotInGame(c *C) {
	player := Player{}

	c.Assert(player.IsInGame(), Equals, false)
}

func (s *CasinoTestsSuite) Test_Player_JoinedGame_IsInGame(c *C) {
	player := Player{}
	game := &Game{}

	player.Join(game)

	c.Assert(player.IsInGame(), Equals, true)
}

func (s *CasinoTestsSuite) Test_Player_LeaveGame_IsNotInGame(c *C) {
	player := Player{}
	player.Join(&Game{})

	player.Leave()

	c.Assert(player.IsInGame(), Equals, false)
}

func (s *CasinoTestsSuite) Test_Player_CanNotLeaveGame_UntilJoin(c *C) {
	player := Player{}

	err := player.Leave()

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please join the game before leaving")
}

func (s *CasinoTestsSuite) Test_Player_CanNotJoinAnotherGame_WhileHeIsInGame(c *C) {
	player := Player{}
	player.Join(&Game{})

	err := player.Join(&Game{})

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please leave the game before joining another game")
}

func (s *CasinoTestsSuite) Test_6Players_Join_Successfully(c *C) {
	game := &Game{}

	player1 := Player{}
	player1.Join(game)
	player2 := Player{}
	player2.Join(game)
	player3 := Player{}
	player3.Join(game)
	player4 := Player{}
	player4.Join(game)
	player5 := Player{}
	player5.Join(game)
	player6 := Player{}
	player6.Join(game)

	c.Assert(player6.IsInGame(), Equals, true)
}

func (s *CasinoTestsSuite) Test_7thPlayers_Join_Fails(c *C) {
	game := &Game{}
	player1 := Player{}
	player1.Join(game)
	player2 := Player{}
	player2.Join(game)
	player3 := Player{}
	player3.Join(game)
	player4 := Player{}
	player4.Join(game)
	player5 := Player{}
	player5.Join(game)
	player6 := Player{}
	player6.Join(game)

	player7 := Player{}
	player7.Join(game)

	c.Assert(player7.IsInGame(), Equals, false)
}

func (s *CasinoTestsSuite) Test_7thPlayers_Join_FailsWithError(c *C) {
	game := &Game{}
	player1 := Player{}
	player1.Join(game)
	player2 := Player{}
	player2.Join(game)
	player3 := Player{}
	player3.Join(game)
	player4 := Player{}
	player4.Join(game)
	player5 := Player{}
	player5.Join(game)
	player6 := Player{}
	player6.Join(game)

	player7 := Player{}
	err := player7.Join(game)

	c.Assert(err, Not(IsNil))
	c.Assert(err.Error(), Equals, "Please join another game")
}

func (s *CasinoTestsSuite) Test_Player_BuyChips_IncreasesBalance(c *C) {
	player := Player{}

	player.Buy(Chips(10))

	c.Assert(player.Balance(), Equals, Chips(10))
}
