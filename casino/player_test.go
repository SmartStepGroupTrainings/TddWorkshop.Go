package casino_new

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GameTest struct {
	suite.Suite
	player *Player
	game   *RollDiceGame
}

func (test *GameTest) SetupTest() {
	test.player = NewPlayer()
	test.game = NewRollDiceGame()
}

const someScore = 1
const someAmount = 10

func Test_Game(t *testing.T) {
	suite.Run(t, &GameTest{})
}

func (test *GameTest) TestNewPlayer_IsNotNil() {
	test.NotNil(test.player)
}

func (test *GameTest) TestНовыйИгрок_НеИмеетЧипсов() {
	test.Empty(test.player.AvailableChips())
}

func (test *GameTest) TestNewPlayer_NotIsInGame() {
	test.False(test.player.IsInGame())
}

func (test *GameTest) TestNewPlayer_Join_IsInGame() {
	test.player.Join(test.game)

	test.True(test.player.IsInGame())
}

func (test *GameTest) TestWhenLeaveTheGame_Fail() {
	err := test.player.Leave()

	test.NotNil(err)
}

func (test *GameTest) TestWhenLeaveTheGame_Success() {
	test.player.Join(test.game)

	err := test.player.Leave()

	test.Nil(err)
}

func (test *GameTest) TestPlayer_BuyPositiveChips_Success() {
	test.player.BuyChips(100)

	test.Equal(100, test.player.AvailableChips())
}

func (test *GameTest) TestPlayer_BuyNegativeChips_Fail() {
	err := test.player.BuyChips(-100)

	test.NotNil(err)
}

func (test *GameTest) TestPlayer_BuyChipsInMultiThread_Success() {
	const amount = 100

	wg := sync.WaitGroup{}
	wg.Add(amount)
	for i := 0; i < amount; i++ {
		go func() {
			test.player.BuyChips(1)
			wg.Done()
		}()
	}
	wg.Wait()

	test.Equal(amount, test.player.AvailableChips())
}

func (test *GameTest) TestPlayer_CannotBetWithoutMoney() {
	bet := Bet{Score: someScore, Amount: 20}

	err := test.player.Bet(bet)

	test.NotNil(err)
}

func (test *GameTest) TestPlayer_CheckSetBet_Success() {
	bet := Bet{Score: someScore, Amount: 20}
	test.player.BuyChips(20)

	err := test.player.Bet(bet)

	test.Nil(err)
}

func (test *GameTest) TestPlayer_CheckBetScore_AfterBet() {
	bet := Bet{Score: 3, Amount: 666}
	test.player.BuyChips(bet.Amount)
	test.player.Bet(bet)

	betOn := test.player.GetBetOn(3)

	test.Equal(666, betOn)
}

func (test *GameTest) TestPlayer_CheckBetScore_AfterLose() {
	bet := Bet{Score: 3, Amount: someAmount}
	test.player.BuyChips(bet.Amount)
	test.player.Bet(bet)

	test.player.Lose()

	test.Empty(test.player.GetBetOn(3))
}

func (test *GameTest) TestPlayer_CheckAvailableChips_AfterWin() {
	bet := Bet{Score: someScore, Amount: 20}
	test.player.BuyChips(bet.Amount)

	test.player.Win(1)

	test.Equal(20+1, test.player.AvailableChips())
}

func (test *GameTest) TestGame_AddPlayer_Success() {
	err := test.game.Add(test.player)

	test.Nil(err)
}

//!!! Founded Error in original Code and fix after testing!!!

func (test *GameTest) TestGame_PlayerIsInGameAfterGameAdd() {
	test.game.Add(test.player)

	test.True(test.player.IsInGame())
}

func (test *GameTest) TestGame_AddPlayer_FailOnNilPlayer() {
	err := test.game.Add(nil)

	test.NotNil(err)
}

func (test *GameTest) TestGame_OnAddTheSamePlayerTwice_Fail() {
	test.game.Add(test.player)
	err := test.game.Add(test.player)

	test.NotNil(err)
}

func (test *GameTest) TestGame_CannotPlayWithoutPlayers() {
	err := test.game.Play()

	test.NotNil(err)
}

func (test *GameTest) TestGame_CountOfPlayerOnNewGame_IsZero() {
	test.Equal(0, test.game.PlayersCount())
}

func (test *GameTest) TestGame_CountOfPlayerAfterAddOnePlayer_IsOne() {
	test.game.Add(test.player)

	test.Equal(1, test.game.PlayersCount())
}
