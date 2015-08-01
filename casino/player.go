package casino

type Player struct {
	game *Game
}

func (player *Player) Join(game *Game) {
	player.game = game
}

func (player *Player) LeaveGame() {
	player.game = nil
}

func (player *Player) IsInGame() bool {
	if player.game != nil {
		return true
	}
	return false
}
