package casino

type Player struct {
	game *Game
}

func (player *Player) Join(game *Game) {
	player.game = game
}

func (player *Player) IsInGame() bool {
	return true
}

