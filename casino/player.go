package casino

import (
	"errors"
)

type Player struct {
	currentGame *Game
}

func (player *Player) Join(game *Game) error {
	if game == nil {
		return errors.New("Joining nil game is not allowed")
	}
	if game.HasPlayer(player) {
		return errors.New("Player is already in game")
	}

	game.addPlayer(player)
	player.currentGame = game
	return nil
}

func (player *Player) Leave() error {
	if player.currentGame == nil {
		return errors.New("Please join the game before leaving")
	}
	if !player.currentGame.HasPlayer(player) {
		return errors.New("Please join the game before leaving")
	}

	player.currentGame.removePlayer(player)
	player.currentGame = nil
	return nil
}
