package casino_new

type RollDiceGameFather struct {
	dice *DiceStub
}

func (self *RollDiceGameFather) WithWinningScore(score int) *RollDiceGameFather {
	self.dice.On("Roll").Return(score)
	return self
}

func (self *RollDiceGameFather) Please() *RollDiceGame {
	return NewRollDiceGame(self.dice)
}

