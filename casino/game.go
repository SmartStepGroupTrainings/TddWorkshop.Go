package casino

import (
	"errors"
)

type Game struct {
	players []*Player
}

func (game *Game) Add(player *Player) error {
	if player.IsInGame() {
		return errors.New("Please leave the game before joining another game")
	}

	if len(game.players) == 6 {
		return errors.New("Please join another game")
	}

	game.players = append(game.players, player)
	player.isInGame = true
	return nil
}

// func (game *Game) Remove(player *Player) error {
// 	for i, p := range game.players {
// 		if p == player {
// 			game.players = append(game.players[:i], game.players[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return errors.New("Please join the game before leaving")
// }

func (game *Game) HasPlayer(player *Player) bool {
	for _, p := range game.players {
		if p == player {
			return true
		}
	}
	return false
}

func (game *Game) Play() {

}
