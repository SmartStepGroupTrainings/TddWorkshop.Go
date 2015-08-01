package casino

import "errors"

// Player model for casino players
type Player struct {
	inTheGame bool
}

// CanJoinGame check if player can join to game
func (player *Player) CanJoinGame() bool {
	return true
}

// CanLeaveGame check if player can join to game
func (player *Player) CanLeaveGame() bool {
	return player.inTheGame
}

func (player *Player) Leave() error {
	if !player.CanLeaveGame() {
		return errors.New("You cannot leave from the game")
	}
	player.inTheGame = false
	return nil
}

func (player *Player) Join(game Game) {
	player.inTheGame = true
}
