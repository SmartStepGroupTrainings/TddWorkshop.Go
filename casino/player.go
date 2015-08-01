package casino

import "github.com/bronze1man/kmg/errors"

type Player struct {
	currentGame *Game
	chipsCount  int
}

func (self *Player) BuyChips(count int) {
	self.chipsCount += count
}

func (self *Player) ChipsCount() int {
	return self.chipsCount
}

func (self *Player) Bet(bet Bet) error {
	if self.chipsCount < bet.Amount {
		return errors.New("cheater cannot buy more than he has")
	}
	self.currentGame.bet = bet
	return nil
}
