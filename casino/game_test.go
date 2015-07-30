package casino_new

type GameTest struct {
    PlayerTest
    Game *RollDiceGame
}

func (self *GameTest) SetupTest() {
    self.Player = NewPlayer()
    self.Game = NewRollDiceGame()
}

func (self *GameTest) TestGame_NoPlayers_JoinPlayer_Success() {
    self.Player.Join(self.Game)

    self.Equal(true, self.Player.IsInGame())
}

func (self *GameTest) TestGame_HasPlayer_Leave_Success() {
    self.Player.Join(self.Game)

    self.Player.Leave()

    self.Equal(false, self.Player.IsInGame())
}