package casino

import "errors"

//import "fmt"

type Player struct {
	currentGame *Game
	balance     Chips
}

type Chips uint

func (player *Player) IsInGame() bool {
	return player.currentGame != nil
}

func (player *Player) Join(game *Game) error {
	if player.IsInGame() {
		return errors.New("Please leave the game before joining another game")
	}

	if err := game.Add(); err != nil {
		return err
	}

	player.currentGame = game
	return nil
}

func (player *Player) Leave() error {
	if !player.IsInGame() {
		return errors.New("Please join the game before leaving")
	}

	player.currentGame.Remove()
	player.currentGame = nil
	return nil
}

func (player *Player) Buy(chips Chips) {
	player.balance += chips
}

func (player *Player) Balance() Chips {
	return player.balance
}
