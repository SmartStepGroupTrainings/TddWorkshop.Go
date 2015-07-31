package casino_new

type DiceFake struct {
	score int
}

func (dice DiceFake) Roll() int {
	return dice.score
}

//-------------- Game Methods -------------------
func (test *GameTest) NewGame() *GameTest {
	test.game = NewRollDiceGame()
	return test
}

func (test *GameTest) WithDiceWinningScore(winScore int) *GameTest {
	test.game.dice = DiceFake{score: winScore}
	return test
}

func (test *GameTest) AddPlayer() *GameTest {
	test.player = NewPlayer()
	return test
}

func (test *GameTest) Play() *GameTest {
	test.game.Play()
	return test
}

//-------------- Player Methods -------------------
func (test *GameTest) WithCache(amount int) *GameTest {
	test.NotNil(test.player)
	test.player.BuyChips(amount)
	return test
}

func (test *GameTest) WithUnlimitedCache() *GameTest {
	return test.WithCache(100500e6)
}

func (test *GameTest) WhoJoinsToCurrentGame() *GameTest {
	test.NotNil(test.game)
	test.NotNil(test.player)
	test.game.Add(test.player)
	return test
}

func (test *GameTest) Bet(amount int) *GameTest {
	test.bet = Bet{Amount: amount}
	return test
}

func (test *GameTest) On(score int) *GameTest {
	test.bet.Score = score
	test.player.Bet(test.bet)
	return test
}
func (test *GameTest) LeaveGame() *GameTest {
	test.player.Leave()
	return test
}

func (test *GameTest) CurrentPlayer() *Player {
	return test.player
}
