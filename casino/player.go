package casino_new

import (
	"errors"
)

type Player struct {
	currentGame    *RollDiceGame
	availableChips int
	bets           map[int]int
}

func NewPlayer() *Player {
	return &Player{bets: make(map[int]int)}
}

func (player *Player) Join(game *RollDiceGame) error {
	if player.IsInGame() {
		return errors.New("Unable to join another game")
	}

	player.currentGame = game
	game.Add(player)
	return nil
}

func (player *Player) Leave() error {
	if !player.IsInGame() {
		return errors.New("Unable to leave the game before joining")
	}

	player.currentGame.Remove(player)
	player.currentGame = nil
	for _, chips := range player.bets {
		player.BuyChips(chips)
	}
	player.bets = make(map[int]int)

	return nil
}

func (player *Player) IsInGame() bool {
	return player.currentGame != nil
}

func (player *Player) Bet(bet Bet) error {
	if bet.Amount < 0 {
		return errors.New("Can't bet negative Amount")
	}

	if player.AvailableChips() < bet.Amount {
		return errors.New("Unable to bet chips more than available")
	}
	if bet.Score < 1 || 6 < bet.Score {
		return errors.New("Bets on 1..6 only are allowed")
	}

	player.availableChips -= bet.Amount
	player.bets[bet.Score] += bet.Amount

	return nil
}

func (player *Player) AvailableChips() int {
	return player.availableChips
}

func (player *Player) BuyChips(chips int) error {
	if chips <= 0 {
		return errors.New("Please buy positive amount")
	}
	player.availableChips += chips
	return nil
}

func (player *Player) GetBetOn(score int) int {
	return player.bets[score]
}

func (self *Player) Lose() {
	self.bets = make(map[int]int)
}

func (self *Player) Win(wonChips int) {
	self.availableChips += wonChips
}
