package casino_new

import (
    "github.com/stretchr/testify/suite"
    "testing"
)

type DiceTestSuite struct {
    suite.Suite
    game *RollDiceGame
    player *Player
}

func (suite *DiceTestSuite) SetupTest() {
    suite.game = NewRollDiceGame()
    suite.player = NewPlayer()
}

func TestDiceTestSuite(t *testing.T) {
    suite.Run(t, new(DiceTestSuite))
}

func (suite *DiceTestSuite) TestDice_EmptyPlayers_NoPlayer_Success() {
    suite.Empty(suite.game.GetPlayers())
}

func (suite *DiceTestSuite) TestDice_EmptyPlayers_AddPlayer_Success() {
    suite.game.Add(suite.player)

    suite.NotNil(suite.game.GetPlayers())
}

func (suite *DiceTestSuite) TestDice_PlayerInGame_Play_NullAvailableChipsSuccess() {
    suite.game.Add(suite.player)
    suite.player.BuyChips(6)
    suite.player.Bet(Bet{Amount: 1, Score: 1})
    suite.player.Bet(Bet{Amount: 1, Score: 2})
    suite.player.Bet(Bet{Amount: 1, Score: 3})
    suite.player.Bet(Bet{Amount: 1, Score: 4})
    suite.player.Bet(Bet{Amount: 1, Score: 5})
    suite.player.Bet(Bet{Amount: 1, Score: 6})

    suite.game.Play()

    suite.Equal(6, suite.player.AvailableChips())
}