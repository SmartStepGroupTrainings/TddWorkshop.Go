package casino

import "errors"

// Game model for casino Game
type Game struct {
	cntOfPlayers int
}

func (g *Game) Add(player *Player) error {
	if !player.CanJoinGame() {
		return errors.New("Player already in the game")
	}

	if g.isFull() {
		return errors.New("Game is full")
	}

	g.cntOfPlayers++
	player.inTheGame = true

	return nil
}

func (g *Game) isFull() bool {
	return g.cntOfPlayers >= 6
}

func (g *Game) isValid(bet Bet) bool {
	return bet.Coins%5 == 0
}
