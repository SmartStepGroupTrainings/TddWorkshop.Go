package casino_new

type Father struct {}

func (self *Father) RollDiceGame() *RollDiceGameFather {
	return &RollDiceGameFather{
		dice: new(DiceStub),
	}
}

func (self *Father) Player() *PlayerFather {
	return &PlayerFather{
		player: NewPlayer(),
	}
}



