package casino

import (
	"errors"
	//"log"
)

type PlayerBet struct {
	bet    Bet
	player *Player
}

type Game struct {
	dice    IDice
	balance Chips
	wallet  Wallet
	table   Table
}

type IDice interface {
	Roll() Score
}

func (game *Game) Add(player *Player) error {
	return game.table.Add(player, game)
}

func (game *Game) Remove(player *Player) error {
	for i, p := range game.table.players {
		if p == player {
			game.table.players = append(game.table.players[:i], game.table.players[i+1:]...)
			player.game = nil
			return nil
		}
	}
	return errors.New("Please join the game before leaving")
}

func (game *Game) Bet(bet Bet, player *Player) error {
	if bet.Score < Score(1) || Score(6) < bet.Score {
		return errors.New("Please make a bet only to score 1 - 6")
	}

	game.wallet.Add(bet, player)
	return nil
}

func (game *Game) HasPlayer(player *Player) bool {
	for _, p := range game.table.players {
		if p == player {
			return true
		}
	}
	return false
}

func (game *Game) Play() {
	winningScore := game.dice.Roll()
	game.wallet.Play(winningScore)
}

func (game *Game) Balance() Chips {
	return game.wallet.Balance()
}
