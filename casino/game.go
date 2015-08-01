package casino

import "errors"

type Game struct {
	players []*Player
	numOfPlayers int
}

const maxPlayers = 6

func (self *Game) Add(player *Player) error {
	if self.numOfPlayers >= maxPlayers {
		return errors.New("Cant add player to full game.")
	}

	for _, p := range self.players {
		if p == player {
			return errors.New("Player cannot join game twice")
		}
	}

	self.players = append(self.players, player)
	self.numOfPlayers++
	return nil
}

func (self *Game) Remove(player *Player) error {
	var removed bool

	players := self.players[:]
	for _, p := range self.players {
		if p == player {
			removed = true
		} else {
//			players = append(players, p)
		}
	}
	self.players = players

	if !removed {
		return errors.New("Player cannot leave game if he's not in game")
	}
	self.numOfPlayers--
	return nil
}
