package casino_new

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Player_HasNoBets_ByDefault(t *testing.T) {
	player := &Player{}

	assert.Equal(t, 0, len(player.Bets()))
}

func Test_Player_Bet_AddsBet(t *testing.T) {
	player := &Player{}
	bet := Bet{ Score: 1, Amount: 10 }

	player.Bet(bet)

	assert.Equal(t, 1, len(player.Bets()))
	assert.Contains(t, player.Bets(), bet)
}

func Test_Player_HasNoAvailableChips_ByDefault(t *testing.T) {
	player := &Player{}

	assert.Equal(t, 0, player.AvailableChips())
}

func Test_Player_BuyChips_AddsAvailableChips(t *testing.T) {
	player := &Player{}

	player.BuyChips(1)

	assert.Equal(t, 1, player.AvailableChips())
}

func Test_Player_BuyMoreChips_AddsAvailableChips(t *testing.T) {
	player := &Player{}
	player.BuyChips(1)

	player.BuyChips(1)

	assert.Equal(t, 1 + 1, player.AvailableChips())
}

func Test_Player_BuyZeroChipa_NotAllowed(t *testing.T) {
	player := &Player{}

	err := player.BuyChips(0)

	assert.Equal(t, 0, player.AvailableChips())
	assert.NotNil(t, err)
	assert.Equal(t, "Please buy positive amount", err.Error())
}

func Test_Player_BuyNegativeChipa_NotAllowed(t *testing.T) {
	player := &Player{}

	err := player.BuyChips(-1)

	assert.Equal(t, 0, player.AvailableChips())
	assert.NotNil(t, err)
	assert.Equal(t, "Please buy positive amount", err.Error())
}
