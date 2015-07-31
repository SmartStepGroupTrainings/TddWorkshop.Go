package casino_new

type RollDiceGame struct {
	players map[*Player]struct{}
	dice IDice
}

func NewRollDiceGame() *RollDiceGame {
	dice := NewDice()
	return &RollDiceGame{
		players: make(map[*Player]struct{}),
		dice: dice,
	}
}

func (self *RollDiceGame) Play() {
	winningScore := self.dice.Roll()

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
