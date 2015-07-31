package casino_new

type GameBuilder struct {
	//winnerScore int
}

func (builder *GameBuilder) Please() *RollDiceGame {
	dice := &diceStub{}
	dice.On("Roll").Return(1)
	game := NewRollDiceGame(dice)

	return game
}
