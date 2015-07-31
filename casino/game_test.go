package casino_new

type TestRoll struct {
}

func (self *TestRoll) Roll() int {
	return 1
}

type GameTest struct {
	PlayerTest
	Game *RollDiceGame
}

func (self *GameTest) SetupTest() {
	self.Player = NewPlayer()
	self.Game = NewRollDiceGame(new(TestRoll))
}

func (self *GameTest) TestPlay_MakeBet_WinOk() {
	self.Player.Join(self.Game)
	self.Player.BuyChips(10)
	const Score = 1
	self.Player.Bet(Bet{Amount: 1, Score: Score})

	self.Game.Play()

	self.Equal((10-1)+1*6, self.Player.AvailableChips())
}
