package casino_new

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PlayerTest struct {
	suite.Suite
	player *Player
	game   *RollDiceGame
}

func (test *PlayerTest) SetupTest() {
	test.player = NewPlayer()
	test.game = NewRollDiceGame()
}

func Test_Game(t *testing.T) {
	suite.Run(t, &PlayerTest{})
}

func (test *PlayerTest) NewPlayer_IsNotNil() {
	test.NotNil(test.player)
}

func (test *PlayerTest) НовыйИгрок_НеИмеетЧипсов() {
	test.Empty(test.player.AvailableChips())
}

func (test *PlayerTest) NewPlayer_NotIsInGame() {
	test.False(test.player.IsInGame())
}

func (test *PlayerTest) NewPlayer_Join_IsInGame() {
	test.player.Join(test.game)

	test.True(test.player.IsInGame())
}

func (test *PlayerTest) WhenLeaveFromNonYourGame_WithError() {
	err := test.player.Leave()

	test.NotNil(err)
}

func (test *PlayerTest) WhenLeaveFromYourGame_WithoutError() {
	test.player.Join(test.game)

	err := test.player.Leave()

	test.Nil(err)
}

func (test *PlayerTest) Player_BuyPositiveChips_WithoutError() {
	test.player.BuyChips(100)

	test.Equal(100, test.player.AvailableChips())
}

func (test *PlayerTest) Player_BuyNegativeChips_WithError() {
	err := test.player.BuyChips(-100)

	test.NotNil(err)
}

func (test *PlayerTest) Player_BuyChipsInMultiThread_WithoutError() {
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
