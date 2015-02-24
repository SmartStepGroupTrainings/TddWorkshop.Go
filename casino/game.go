package casino

type Game struct {
	players []*Player
}

func (game *Game) addPlayer(player *Player) {
	game.players = append(game.players, player)
}

func (game *Game) removePlayer(player *Player) {
	index := 0
	for i, p := range game.players {
		if p == player {
			index = i
		}
	}
	game.players = append(game.players[:index], game.players[index+1:]...)
}

func (game *Game) HasPlayer(player *Player) bool {
	for _, p := range game.players {
		if p == player {
			return true
		}
	}
	return false
}
