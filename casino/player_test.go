package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_ByDefault_HasNoChips(t *testing.T) {
	player := NewPlayer()

	assert.Equal(t, 0, player.AvailableChips())
}

func TestPlayer_CanBuyChips(t *testing.T) {
	player := NewPlayer()

	player.BuyChips(10)

	assert.Equal(t, 10, player.AvailableChips())
}

func TestPlayer_CanBetChips(t *testing.T) {
	player := NewPlayer()
	player.BuyChips(10)

	player.Bet(Bet{Score: 1, Amount: 2})

	assert.Equal(t, 10-2, player.AvailableChips())
}
