package casino

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

func (self *Player) Bet(bet Bet) {
	self.currentGame.bet = bet
}
