package casino

import (
	"errors"
)

type Player struct {
}

func (player *Player) Join(game *Game) error {
	if game == nil {
		return errors.New("Joining nil game is not allowed")
	}
	if game.HasPlayer(player) {
		return errors.New("Player is already in game")
	}

	game.addPlayer(player)
	return nil
}

func (player *Player) Leave(game *Game) {
	game.removePlayer(player)
}
