package casino

import (
	"errors"
)

type Game struct {
	numberOfPlayer int
}

func (game *Game) Add() error {
	if game.numberOfPlayer == 6 {
		return errors.New("Please join another game")
	}

	game.numberOfPlayer++
	return nil
}

func (game *Game) Remove() {
	game.numberOfPlayer--
}
