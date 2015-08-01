package casino

import "errors"

type Game struct {
	numberOfPlayers int
}

func (self *Game) AddPlayer() error {
	if self.IsFull() {
		return errors.New("Game can not accept more than 6 players")
	}

	self.numberOfPlayers++
	return nil
}

func (self *Game) IsFull() bool {
	return self.numberOfPlayers >= 6
}
