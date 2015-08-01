package casino

import "errors"

type Player struct {
	isInGame bool
}

func (p *Player) Join(game Game) error {
	if p.isInGame {
		return errors.New("You can not join game")
	}

	p.isInGame = true
	return nil
}

func (p *Player) Leave(game Game) error {

	if !p.isInGame {
		return errors.New("You can not leave game")
	}

	p.isInGame = false
	return nil
}
