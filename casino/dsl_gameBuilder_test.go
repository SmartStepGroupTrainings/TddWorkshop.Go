package casino

type GameBuilder struct {
	game Game
}

func (gameBuilder *GameBuilder) WithLuckyScore(luckyScore uint) *GameBuilder {
	dice := DiceFake{}
	dice.score = Score(luckyScore)
	gameBuilder.game.dice = dice
	return gameBuilder
}

func (gameBuilder *GameBuilder) Please() *Game {
	return &gameBuilder.game
}
