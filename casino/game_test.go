package casino_new

type GameTest struct {
    PlayerTest
    Game *RollDiceGame
}

func (self *GameTest) SetupTest() {
    self.Player = NewPlayer()
    self.Game = NewRollDiceGame()
}

func (self *GameTest) TestPlay() {
    self.Player.Join(self.Game)
    self.Player.BuyChips(10)
    const Score = 1
    self.Player.Bet(Bet{Amount:1, Score:Score})
    self.Game.SetForceScore(Score)

    self.Game.Play()

    self.Equal((10-1)+1*6, self.Player.AvailableChips())

}