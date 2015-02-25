package casino

type Table struct {
	players []*Player
}

func (table *Table) Add(player *Player) {
	table.players = append(table.players, player)
}

func (table *Table) Remove(player *Player) {
	index := 0
	for i, p := range table.players {
		if p == player {
			index = i
		}
	}
	table.players = append(table.players[:index], table.players[index+1:]...)
}

func (table *Table) HasPlayer(player *Player) bool {
	for _, p := range table.players {
		if p == player {
			return true
		}
	}
	return false
}
