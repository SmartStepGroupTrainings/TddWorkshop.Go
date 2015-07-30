package casino_new

type GameTest struct {
    PlayerTest
    Game *RollDiceGame
}

func (self *GameTest) SetupTest() {
    self.Player = NewPlayer()
    self.Game = NewRollDiceGame()
}
