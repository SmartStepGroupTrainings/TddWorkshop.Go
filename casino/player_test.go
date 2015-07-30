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

func (test *GameTest) TestWhenLeaveFromTheGame_WithError() {
	err := test.player.Leave()

	test.NotNil(err)
}

func (test *GameTest) TestWhenLeaveFromTheGame_WithoutError() {
	test.player.Join(test.game)

	err := test.player.Leave()

	test.Nil(err)
}

func (test *GameTest) TestPlayer_BuyPositiveChips_WithoutError() {
	test.player.BuyChips(100)

	test.Equal(100, test.player.AvailableChips())
}

func (test *GameTest) TestPlayer_BuyNegativeChips_WithError() {
	err := test.player.BuyChips(-100)

	test.NotNil(err)
}

func (test *GameTest) TestPlayer_BuyChipsInMultiThread_WithoutError() {
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
	bet := Bet{Score: 1, Amount: 20}

	err := test.player.Bet(bet)

	test.NotNil(err)
}

func (test *GameTest) TestPlayer_CheckSetBet_Success() {
	bet := Bet{Score: 2, Amount: 20}
	test.player.BuyChips(bet.Amount)

	err := test.player.Bet(bet)

	test.Nil(err)
}

func (test *GameTest) TestPlayer_CheckBetScore_AfterBet() {
	bet := Bet{Score: 2, Amount: 20}
	test.player.BuyChips(bet.Amount)
	test.player.Bet(bet)

	betOn := test.player.GetBetOn(bet.Score)

	test.Equal(bet.Amount, betOn)
}

func (test *GameTest) TestPlayer_CheckBetScore_AfterLose() {
	bet := Bet{Score: 2, Amount: 20}
	test.player.BuyChips(bet.Amount)
	test.player.Bet(bet)

	test.player.Lose()

	test.Empty(test.player.GetBetOn(bet.Score))
}

func (test *GameTest) TestPlayer_CheckAvailableChips_AfterWin() {
	bet := Bet{Score: 2, Amount: 20}
	test.player.BuyChips(bet.Amount)

	test.player.Win(1)

	test.Equal(bet.Amount+1, test.player.AvailableChips())
}
