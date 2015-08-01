package casino

import "errors"

var errGameIsFull = errors.New("Game is Full")

type Game struct {
	playerCnt int
}

func (g *Game) AddPlayer(player *Player) error {
	if g.playerCnt >= 6 {
		return errGameIsFull
	}
	g.playerCnt++
	return nil
}
