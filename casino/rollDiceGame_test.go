package casino_new

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type diceStub struct {
	score int
}

func (d diceStub) Roll() int {
	if d.score == 0 {
		return 1
	}
	return d.score
}

func (d diceStub) Set(score int) error {
	if score < 1 || score > 6 {
		return errors.New("Can't set score")
	}
	d.score = score
	return nil
}

type TestRollDiceGameSuite struct {
	suite.Suite
	player *Player
	game   *RollDiceGame
	dice   *diceStub
}

func (s *TestRollDiceGameSuite) SetupTest() {
	s.player = NewPlayer()
	s.dice = &diceStub{}
	s.game = NewRollDiceGame(s.dice)

	if err := s.dice.Set(1); err != nil {
		panic("Couldn't set dice stub score")
	}
}

func TestRollDiceGame(t *testing.T) {
	suite.Run(t, new(TestRollDiceGameSuite))
}

func (s *TestRollDiceGameSuite) TestRollDiceGame_Play_PlayerSetsWiinnigBet_ExpectIncreasingChips() {
	err := s.player.Join(s.game)
	s.Nil(err, "Player has to join game")

	err = s.player.BuyChips(10)
	s.Nil(err, "Player has to buy chips")

	bet := Bet{Score: 1, Amount: 10}
	err = s.player.Bet(bet)
	s.Nil(err, "Player has to make a bet")

	s.game.Play()

	availableChips := s.player.AvailableChips()
	s.Equal(10*6, availableChips, "Player must have 60 chips")

}
