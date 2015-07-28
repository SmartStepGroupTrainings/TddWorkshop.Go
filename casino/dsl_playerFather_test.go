package casino_new

type PlayerFather struct {
	bet Bet
	player *Player
}

func (self *PlayerFather) Rich() *PlayerFather {
	self.player.BuyChips(1000)
	return self
}

func (self *PlayerFather) WithBet(chips int) *PlayerFather {
	self.bet = Bet{}
	self.bet.Amount = chips
	return self
}

func (self *PlayerFather) On(score int) *PlayerFather {
	self.bet.Score = score
	self.player.Bet(self.bet)
	return self
}

func (self *PlayerFather) Joined(game *RollDiceGame) *PlayerFather {
	self.player.Join(game)
	return self
}

func (self *PlayerFather) Please() *Player {
	return self.player
}

