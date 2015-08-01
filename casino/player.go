package casino

import "errors"

type Player struct {
	isInGame bool
}

func (p *Player) IsInGame() bool {
	return p.isInGame
}

func (p *Player) Join() {
	p.isInGame = true
}

func (p *Player) Leave() error {
	if !p.isInGame {
		return errors.New("Player not in game")
	}
	p.isInGame = false

	return nil
}
