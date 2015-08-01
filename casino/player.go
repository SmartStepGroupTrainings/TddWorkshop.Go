package casino
import "github.com/dropbox/godropbox/errors"


type Player struct {
	game *Game
}

func (player *Player) JoinTo(game *Game) error {
	if player.IsInGame() {
		return errors.New("Player not in game")
	}
	if game.countUsers >= 6 {
		return errors.New("Player can not join to game with 6 players")
	}
	player.game = game
	game.countUsers++
	return nil
}

func (player *Player) LeaveGame() error {
	if !player.IsInGame() {
		return errors.New("Player not in game")
	}
	player.game.countUsers--
	player.game = nil
	return nil
}

func (player *Player) IsInGame() bool {
	if player.game != nil {
		return true
	}
	return false
}
