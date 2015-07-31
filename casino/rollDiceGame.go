package casino_new

type RollDiceGame struct {
	dice    iDice
	players map[*Player]struct{}
}

func NewRollDiceGame(dice iDice) *RollDiceGame {
	return &RollDiceGame{
		dice:    dice,
		players: make(map[*Player]struct{}),
	}
}

func (self *RollDiceGame) Play() {
	var winningScore = self.dice.Roll()

	for player, _ := range self.players {
		player.Win(player.GetBetOn(winningScore) * 6)
		player.Lose()
	}
}

func (self *RollDiceGame) Add(player *Player) {
	self.players[player] = struct{}{}
}

func (self *RollDiceGame) Remove(player *Player) {
	delete(self.players, player)
}
