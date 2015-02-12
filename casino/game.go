package casino

type Game struct {
	NumberOfPlayers int
}

func (game *Game) IncrementNumberOfPlayers() {
	game.NumberOfPlayers++
}

func (game *Game) DecrementNumberOfPlayers() {
	game.NumberOfPlayers--
}
