package casino

type TwoDiceGameBuilder struct {
	game TwoDiceGame
}

func (builder *TwoDiceGameBuilder) Please() *TwoDiceGame {
	return &builder.game
}
