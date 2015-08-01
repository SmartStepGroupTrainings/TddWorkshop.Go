package casino

import "errors"

var errGameIsFull = errors.New("Game is Full")

type Game struct {
	totalJoins int
}

func (g *Game) AddPlayer(player *Player) error {
	if g.totalJoins >= 6 {
		return errGameIsFull
	}
	g.totalJoins++
	return nil
}
