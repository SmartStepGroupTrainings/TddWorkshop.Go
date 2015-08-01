package casino

import "errors"

type Game struct {
	numOfPlayers int
}

const maxPlayers = 6

func (self *Game) Add(player *Player) error {
	if self.numOfPlayers >= maxPlayers {
		return errors.New("Cant add player to full game.")
	}

	if player.currentGame != nil {
		return errors.New("Player cannot join game twice")
	}

	player.currentGame = self
	self.numOfPlayers++
	return nil
}

func (self *Game) Remove(player *Player) error {
	if player.currentGame != self {
		return errors.New("Player cannot leave game if he's not in this game")
	}

	player.currentGame = nil
	self.numOfPlayers--
	return nil
}
