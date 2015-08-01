package casino
import "github.com/dropbox/godropbox/errors"


type Player struct {
	game *Game
}

func (player *Player) Join(game *Game) error {
	if player.IsInGame() {
		return errors.New("Player not in game")
	}
	player.game = game
	return nil
}

func (player *Player) LeaveGame() error {
	if !player.IsInGame() {
		return errors.New("Player not in game")
	}
	player.game = nil
	return nil
}

func (player *Player) IsInGame() bool {
	if player.game != nil {
		return true
	}
	return false
}
