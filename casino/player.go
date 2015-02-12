package casino

import "errors"

//import "fmt"

type Player struct {
	isInGame bool
}

func (player *Player) IsInGame() bool {
	return player.isInGame
}

func (player *Player) Join(game Game) error {
	if player.IsInGame() {
		return errors.New("Please leave the game before joining another game")
	}

	player.isInGame = true
	return nil
}

func (player *Player) Leave() error {
	if !player.IsInGame() {
		return errors.New("Please join the game before leaving")
	}

	player.isInGame = false
	return nil
}
