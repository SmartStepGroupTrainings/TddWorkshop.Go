package casino

import "errors"

var (
	errGameIsFull     = errors.New("game is full")
	errBetNotAliquot5 = errors.New("bet not aliquot 5")
)

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

func (g *Game) IsBetAmountValid(amount int) error {
	if amount%5 != 0 {
		return errBetNotAliquot5
	}
	return nil
}
