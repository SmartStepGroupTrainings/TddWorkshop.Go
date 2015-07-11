package casino_new

type DiceStub struct {
	winningScore int
}

func (self *DiceStub) WillRoll(score int) {
	self.winningScore = score
}

func (self DiceStub) Roll() int {
	return self.winningScore
}
