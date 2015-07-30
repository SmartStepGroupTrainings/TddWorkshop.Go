package casino_new

import (
	"testing"
	"sync"

	"github.com/stretchr/testify/assert"
)

func setupTest() (*Player, *RollDiceGame) {
	return NewPlayer(), NewRollDiceGame()
}

func TestPlayer_NewPlayer_IsNotNil(t *testing.T) {
	player, _ := setupTest()

	assert.NotNil(t, player)
}

func TestPlayer_НовыйИгрок_НеИмеетЧипсов(t *testing.T) {
	player, _ := setupTest()

	assert.Empty(t, player.AvailableChips())
}

func TestPlayer_NewPlayer_NotIsInGame(t *testing.T) {
	player, _ := setupTest()

	assert.False(t, player.IsInGame())
}

func TestPlayer_NewPlayer_Join_IsInGame(t *testing.T) {
	player, game := setupTest()

	player.Join(game)

	assert.True(t, player.IsInGame())
}

func TestPlayer_WhenLeaveFromNonYourGame_WithError(t *testing.T) {
	player, _ := setupTest()

	err := player.Leave()

	assert.NotNil(t, err)
}

func TestPlayer_WhenLeaveFromYourGame_WithoutError(t *testing.T) {
	player, game := setupTest()
	player.Join(game)

	err := player.Leave()

	assert.Nil(t, err)
}

func TestPlayer_Player_BuyPositiveChips_WithoutError(t *testing.T) {
	player, _ := setupTest()
	const amountChips = 100

	player.BuyChips(amountChips)

	assert.Equal(t, amountChips, player.AvailableChips())
}

func TestPlayer_Player_BuyNegativeChips_WithError(t *testing.T) {
	player, _ := setupTest()
	const amountChips = -100

	err := player.BuyChips(amountChips)

	assert.NotNil(t, err)
}

func TestPlayer_Player_BuyChipsInMultiThread_WithoutError(t *testing.T) {
	player, _ := setupTest()
	const amountChips = 100

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < amountChips; i++ {
		go func() {
			player.BuyChips(1)
			wg.Done()
		}()
	}
	wg.Wait()

	assert.Equal(t, amountChips, player.AvailableChips())
}
