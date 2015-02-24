package casino

import "errors"

type Table struct {
	players []*Player
}

func (table *Table) Add(player *Player, game IGame) error {
	if player.IsInGame() {
		return errors.New("Please leave the game before joining another game")
	}

	if len(table.players) == 6 {
		return errors.New("Please join another game")
	}

	table.players = append(table.players, player)
	player.game = game
	return nil
}
