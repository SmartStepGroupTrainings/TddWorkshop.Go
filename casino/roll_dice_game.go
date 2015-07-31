package casino_new

type RollDiceGame struct {
	players map[*Player]struct{}
	dice IDice
}

func NewRollDiceGame() *RollDiceGame {
	dice := NewDice()
	return NewRollDiceGameWithDice(dice)
}

func NewRollDiceGameWithDice(dice IDice) *RollDiceGame {
	return &RollDiceGame{
		players: make(map[*Player]struct{}),
		dice: dice,
	}
}

func (self *RollDiceGame) Play() []*Player {
	winningScore := self.dice.Roll()
	winners := make([]*Player, 0, len(self.players))

	for player, _ := range self.players {
		if player.GetBetOn(winningScore) > 0 {
			winners = append(winners, player)
		}
		player.Win(player.GetBetOn(winningScore) * 6)
		player.Lose()
	}
	return winners
}

func (self *RollDiceGame) Add(player *Player) {
	self.players[player] = struct{}{}
}

func (self *RollDiceGame) Remove(player *Player) {
	delete(self.players, player)
}

func (self *RollDiceGame) GetPlayers() []*Player {
	players := make([]*Player, 0, len(self.players))

	for p := range self.players {
		players = append(players, p)
	}

	return players
}