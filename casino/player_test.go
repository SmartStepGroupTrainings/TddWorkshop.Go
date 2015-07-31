package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestPlayerSuite struct {
	suite.Suite
	player *Player
	game   *RollDiceGame
}

func TestPlayerTest(t *testing.T) {
	suite.Run(t, new(TestPlayerSuite))
}

func (s *TestPlayerSuite) SetupTest() {
	s.player = NewPlayer()
	s.game = NewRollDiceGame(nil)
}

func (s *TestPlayerSuite) TestPlayer_IsInGameCheck_EmptyPlayer_ShouldFail() {
	res := s.player.IsInGame()

	s.Equal(false, res, "New player should not be in game")
}

func (s *TestPlayerSuite) TestPlayer_IsInGameCheck_NotEmptyPlayer_ShouldSuccess() {
	err := s.player.Join(s.game)
	s.Nil(err, "Player has to join game")

	res := s.player.IsInGame()

	s.Equal(true, res, "This player should be in game")
}

func (s *TestPlayerSuite) TestPlayer_GetAvailableChips_PlayerWithChips_ShouldNotNull() {
	anyChips := 100
	s.player.BuyChips(anyChips)

	availableChips := s.player.AvailableChips()

	s.Equal(anyChips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsNegativeCount_Player_ExpectError() {
	initialChips := s.player.AvailableChips()

	err := s.player.BuyChips(-1)
	availableChips := s.player.AvailableChips()

	s.Error(err, "Didn't get expected error when buying -1 chip")
	s.Equal(initialChips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsZeroCount_Player_ExpectError() {
	initialChips := s.player.AvailableChips()

	err := s.player.BuyChips(0)
	availableChips := s.player.AvailableChips()

	s.Error(err, "Didn't get expected error when buying -1 chip")
	s.Equal(initialChips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_BuyChipsPositiveValue_DefaultPlayer_ShouldIncreaseByCorrectValue() {
	anyChips := 100

	err := s.player.BuyChips(anyChips)
	availableChips := s.player.AvailableChips()

	s.Nil(err, "Got unexpected error when buying positive chips amount")
	s.Equal(anyChips, availableChips, "Player has wrong number of chips")
}

func (s *TestPlayerSuite) TestPlayer_Bet_Player_CantBetMoreThanAvailableChips() {
	s.player.BuyChips(20)
	bet := Bet{Amount: 30}

	err := s.player.Bet(bet)

	s.Error(err, "Error should be not nil")
	s.Equal("Unable to bet chips more than available", err.Error(), "Error message is not valid")
}

func (s *TestPlayerSuite) TestPlayer_Bet_Player_CantBetLessThanOneScore() {
	s.player.BuyChips(1)
	bet := Bet{
		Amount: 1,
		Score:  0,
	}

	err := s.player.Bet(bet)

	s.Error(err, "Error should be not nil")
	s.Equal("Bets on 1..6 only are allowed", err.Error(), "Error message is not valid")
}

func (s *TestPlayerSuite) TestPlayer_Bet_Player_CantBetMoreThanSixScore() {
	s.player.BuyChips(1)
	bet := Bet{
		Amount: 1,
		Score:  7,
	}

	err := s.player.Bet(bet)

	s.Error(err, "Error should be not nil")
	s.Equal("Bets on 1..6 only are allowed", err.Error(), "Error message is not valid")
}

func (s *TestPlayerSuite) TestPlayer_Bet_Player_CantBetBetweenOneAndSixScore() {
	s.player.BuyChips(1)
	bet := Bet{
		Amount: 1,
		Score:  3,
	}

	err := s.player.Bet(bet)

	s.Nil(err, "Error should be nil")
}

func (s *TestPlayerSuite) TestPlayer_Join_Player_AlreadyInGameShouldFail() {
	err := s.player.Join(s.game)
	s.Nil(err, "Player has to join game")

	// try to join game again
	err = s.player.Join(s.game)
	s.Error(err, "Error should be not nil")
}

func (s *TestPlayerSuite) TestPlayer_Leave_Player_WhenNotInGameShouldFail() {
	err := s.player.Leave()
	s.Error(err, "Error should be not nil")
	s.Equal("Unable to leave the game before joining", err.Error(), "Error message is not valid")
}

func (s *TestPlayerSuite) TestPlayer_Leave_PlayerWithNoBets_Success() {
	s.player.Join(s.game)
	s.player.Leave()

	s.False(s.player.IsInGame(), "Player state is invalid: is_in_game should be false")
}

func (s *TestPlayerSuite) TestPlayer_Leave_PlayerWithBetsOnTable_Success() {
	s.player.BuyChips(20)
	s.player.Join(s.game)
	s.player.Bet(Bet{Score: 1, Amount: 10})
	s.player.Leave()

	s.False(s.player.IsInGame(), "Player state is invalid: is_in_game should be false")
}

func (s *TestPlayerSuite) TestPlayer_Leave_PlayerWithBetsOnTable_BetReturnedToPlayer() {
	game := create.Game().Please()
	player := create.Player().InGame(game).BetOn(2).BetAmount(10).Please()
	observer := create.Player().Please()

	player.Leave()

	s.Equal(observer.AvailableChips(), player.AvailableChips(), "ALARM	!")
}
