package casino

import "errors"

type Game struct {
	playersCount int64
}

func (g *Game) AddPlayer() error {
	if g.playersCount >= 6 {
		return errors.New("Game is full")
	}

	g.playersCount++

	return nil
}

func (g *Game) Bet(count int, score int) error {
	if count%5 != 0 {
		return errors.New("Not divide on 5")
	}
	return nil
}
