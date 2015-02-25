package casino

type GameBuilder struct {
	game Game
}

func (gameBuilder *GameBuilder) Please() *Game {
	return &gameBuilder.game
}
