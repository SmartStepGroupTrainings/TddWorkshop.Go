package casino_new

import (
	"errors"
)

type Player struct {
	isInGame bool
	availableChips int
	bets []Bet
}

func (player *Player) Join(game *Game) error {
	if player.isInGame {
		return errors.New("Unable to join another game")
	}

	player.isInGame = true
	return nil
}

func (player *Player) Leave() error {
	if !player.isInGame {
		return errors.New("Unable to leave the game before joining")
	}

	player.isInGame = false
	return nil
}

func (player *Player) IsInGame() bool {
	return player.isInGame
}

func (player *Player) Bets() []Bet {
	return player.bets
}

func (player *Player) Bet(bet Bet) {
	player.bets = append(player.bets, bet)
}

func (player *Player) AvailableChips() int {
	return player.availableChips;
}

func (player *Player) BuyChips(chips int) error {
	if chips <= 0 {
		return errors.New("Please buy positive amount")
	}
	player.availableChips += chips
	return nil
}