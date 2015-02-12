package casino

import "errors"

//import "fmt"

type Player struct {
	isInGame    bool
	currentGame *Game
}

func (player *Player) IsInGame() bool {
	return player.isInGame
}

func (player *Player) Join(game *Game) error {
	if player.IsInGame() {
		return errors.New("Please leave the game before joining another game")
	}

	if game.NumberOfPlayer == 6 {
		return errors.New("Please join another game")
	}

	player.isInGame = true
	player.currentGame = game
	game.NumberOfPlayer++
	return nil
}

func (player *Player) Leave() error {
	if !player.IsInGame() {
		return errors.New("Please join the game before leaving")
	}

	player.isInGame = false
	player.currentGame.NumberOfPlayer--
	player.currentGame = nil
	return nil
}
