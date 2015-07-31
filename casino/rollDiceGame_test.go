package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type diceStub struct {
	score int
}

func (d *diceStub) Roll() int {
	return d.score
}

func (d *diceStub) SetWiningScore(score int) {
	d.score = score
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
}

func TestRollDiceGame(t *testing.T) {
	suite.Run(t, new(TestRollDiceGameSuite))
}

func (s *TestRollDiceGameSuite) TestRollDiceGame_Play_Player_Win() {
	s.dice.SetWiningScore(1)
	s.player.Join(s.game)
	s.player.BuyChips(10)
	bet := Bet{Score: 1, Amount: 10}
	s.player.Bet(bet)

	s.game.Play()

	s.Equal(10*6, s.player.AvailableChips(), "Player must have 60 chips")

}

func (s *TestRollDiceGameSuite) TestRollDiceGame_Play_Player_Loose() {
	s.dice.SetWiningScore(1)
	s.player.Join(s.game)
	s.player.BuyChips(20)
	bet := Bet{Score: 2, Amount: 15}
	s.player.Bet(bet)

	s.game.Play()

	s.Equal(5, s.player.AvailableChips(), "Player must have 5 chips")
}
