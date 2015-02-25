package casino

type DiceFake struct {
	score Score
}

func (dice DiceFake) Roll() Score {
	return dice.score
}
