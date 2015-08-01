package casino

import "errors"

// Player model for casino players
type Player struct {
	inTheGame bool
}

// CanJoinGame check if player can join to game
func (player *Player) CanJoinGame() bool {
	return !player.inTheGame
}

// CanLeaveGame check if player can join to game
func (player *Player) CanLeaveGame() bool {
	return !player.CanJoinGame()
}

func (player *Player) Leave() error {
	if !player.CanLeaveGame() {
		return errors.New("You cannot leave from the game")
	}
	player.inTheGame = false
	return nil
}

func (player *Player) Join(game Game) error {
	if !player.CanJoinGame() {
		return errors.New("player cannot join game - because already in some game")
	}
	player.inTheGame = true
	return nil
}
