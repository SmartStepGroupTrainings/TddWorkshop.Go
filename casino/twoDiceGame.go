package casino

import "errors"

type TwoDiceGame struct {
	wallet Wallet
	table  Table
}

func (game *TwoDiceGame) Bet(bet Bet, player *Player) error {
	if bet.Score < Score(2) || Score(12) < bet.Score {
		return errors.New("Please make a bet only to score 2 - 12")
	}

	game.wallet.Add(bet, player)
	return nil
}

func (game *TwoDiceGame) Add(player *Player) error {
	return game.table.Add(player, game)
}
