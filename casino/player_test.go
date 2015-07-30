package casino_new

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_NotInGame_Leave_Fail(t *testing.T) {
	player := NewPlayer()

	err := player.Leave()

	assert.NotNil(t, err, "Return value is not null")
	assert.Equal(t, "Unable to leave the game before joining", err.Error())
}

func TestPlayer_NotInGame_Leave_Success(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())

	err := player.Leave()

	assert.Nil(t, err, "Player error is not null")
}

func TestPlayer_NotInGame_IsInGame_Fail(t *testing.T) {
	player := NewPlayer()

	assert.Equal(t, false, player.IsInGame())
}
func TestPlayer_InGame_Join_Success(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())

	assert.Equal(t, true, player.IsInGame())
}

func TestPlayer_InGame_BuyChipsWithInvalidValue_Fail(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())

	err := player.BuyChips(-1)

	assert.NotNil(t, err, "Return value is not null")
	assert.Equal(t, "Please buy positive amount", err.Error())
}

func TestPlayer_InGame_BuyChipsWithValidValue_Success(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())

	err := player.BuyChips(22)

	assert.Nil(t, err, "Player error is not null")
	assert.Equal(t, 22, player.AvailableChips())
}

func TestPlayer_NotInGame_HasAvailableChipsIsZero(t *testing.T) {
	player := NewPlayer()

	assert.Equal(t, 0, player.AvailableChips())
}

func TestPlayer_InGame_BetAmountMoreThanAvailable_Fail(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())
	player.BuyChips(1)

	err := player.Bet(Bet{Amount: 2, Score: 2})

	assert.NotNil(t, err, "Return value is not null")
	assert.Equal(t, "Unable to bet chips more than available", err.Error())

}

func TestPlayer_InGame_BetScoreNotValid_Fail(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())
	player.BuyChips(1)

	err := player.Bet(Bet{Amount: 1, Score: 7})

	assert.NotNil(t, err, "Return value is not null")
	assert.Equal(t, "Bets on 1..6 only are allowed", err.Error())
}

func TestPlayer_InGame_BetScoreValid_Success(t *testing.T) {
	player := NewPlayer()
	player.Join(NewRollDiceGame())
	player.BuyChips(1)

	_ = player.Bet(Bet{Amount: 1, Score: 5})

	assert.Equal(t, 1-1, player.AvailableChips())
	assert.Equal(t, 0+1, player.GetBetOn(5))
}
