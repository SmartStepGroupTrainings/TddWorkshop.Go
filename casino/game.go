package casino
import "github.com/bronze1man/kmg/errors"

type Game struct{
	player *Player
}

func (self *Game) Add(player *Player) error {
	self.player = player
	return nil
}

func (self *Game) Remove(player *Player) error  {
	if self.player != player {
		return errors.New("Player cannot leave game if he's not in game")
	}
	self.player = nil
	return nil
}