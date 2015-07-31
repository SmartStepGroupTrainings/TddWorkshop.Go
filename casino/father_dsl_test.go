package casino_new

type Father struct{}

func (self *Father) RollDiceGame() *RollDiceGameFather {
	return &RollDiceGameFather{}
}

func (self *Father) Looser() *PlayerFather {
	return &PlayerFather{}
}

func (self *Father) Winner() *PlayerFather {
	return &PlayerFather{}
}

type RollDiceGameFather struct {
	self.dice = new(DiceStub)
}

func (self *RollDiceGameFather) Please() *RollDiceGame {
	game := NewRollDiceGame(self.dice)
	return game
}

type PlayerFather struct {
	bet *Bet
}

func (self *PlayerFather) WithBet(chips int) *PlayerFather {
	if self.bet == nil {
		self.bet = &Bet{Amount: chips, Score: 1}
	}
	self.bet.Amount = chips
	return self
}

func (self *PlayerFather) Chips() *PlayerFather {
	return self
}

func (self *PlayerFather) Please() *Player {
	player := &Player{}
	if self.bet != nil {
		player.Bet(*self.bet)
	}
	return player
}
