package casino_new

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestDice struct {
	nextValue int
}

func (d *TestDice) Roll() int {
	return d.nextValue
}

// Suite with player
type TestSuitePlayer struct {
	suite.Suite
	player *Player
}

func (s *TestSuitePlayer) SetupTest() {
	s.player = NewPlayer()
}

// Suite with game and player
type TestSuiteGameAndPlayer struct {
	suite.Suite
	player *Player
	game   *RollDiceGame
}

func (s *TestSuiteGameAndPlayer) SetupTest() {
	s.game = NewRollDiceGame(&TestDice{})
	s.player = NewPlayer()
}

// Run suits
func TestRunSuits(t *testing.T) {
	suite.Run(t, new(TestSuitePlayer))
	suite.Run(t, new(TestSuiteGameAndPlayer))
}

// Tests
func (s *TestSuiteGameAndPlayer) TestPlayer_FirstJoin_Success() {
	s.NoError(s.player.Join(s.game))
}

func (s *TestSuiteGameAndPlayer) TestPlayer_PlayerInGameAfterJoin() {
	s.player.Join(s.game)

	s.True(s.player.IsInGame())
}

func (s *TestSuiteGameAndPlayer) TestPlayer_TwiceJoin_Fail() {
	s.player.Join(s.game)

	s.Error(s.player.Join(s.game))
}

func (s *TestSuiteGameAndPlayer) TestPlayer_Leave_JoinedGame_Success() {
	s.player.Join(s.game)

	s.NoError(s.player.Leave())
}

func (s *TestSuitePlayer) TestPlayer_Leave_NotJoinedGame_Failed() {
	s.Error(s.player.Leave())
}

func (s *TestSuitePlayer) TestPlayer_BuyCheapsOnce_CheckError() {
	s.NoError(s.player.BuyChips(1))
}

func (s *TestSuitePlayer) TestPlayer_BuyCheapsOnce_CheckState() {
	s.player.BuyChips(1)
	s.Equal(1, s.player.AvailableChips())
}

func (s *TestSuiteGameAndPlayer) TestPlayer_AddBet_WhenNotEnoughtChips() {
	s.player.BuyChips(1)
	s.Error(s.player.Bet(Bet{6, 2}))
}

func (s *TestSuiteGameAndPlayer) TestPlayer_CheckAvailableChips_AfterBet() {
	s.player.BuyChips(10)
	s.player.Bet(Bet{6, 2})

	s.Equal(8, s.player.AvailableChips())
}

func (s *TestSuiteGameAndPlayer) TestPlayer_Play_Win_IncreasedChips() {
	s.player.Join(s.game)
	s.player.BuyChips(1)
	s.player.Bet(Bet{6, 1})
	s.game.dice.(*TestDice).nextValue = 6;
	s.game.Play()

	s.Equal(6, s.player.AvailableChips())
}

func (s *TestSuiteGameAndPlayer) TestPlayer_Play_Lose_LostChips() {
	s.player.Join(s.game)
	s.player.BuyChips(1)
	s.player.Bet(Bet{6, 1})
	s.game.dice.(*TestDice).nextValue = 5;
	s.game.Play()

	s.Equal(0, s.player.AvailableChips())
}