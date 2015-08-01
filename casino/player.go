package casino
import "errors"

type Player struct {
	isInGame bool
}

func (p *Player) IsInGame() bool {
	return p.isInGame
}

func (p *Player) Join(game Game) error {
    if p.IsInGame() {
        return errors.New("Player is already in game")
    }
	p.isInGame = true
    return nil
}

func (p *Player) Leave(game Game) {
	p.isInGame = false
}
