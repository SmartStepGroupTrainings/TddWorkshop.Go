package casino

import "errors"

type Player struct {
	isInGame bool
	chips    int
}

func (p *Player) IsInGame() bool {
	return p.isInGame
}

func (p *Player) Join(game *Game) error {
	if p.isInGame {
		return errors.New("Player already in game")
	}

	if err := game.Add(); err != nil {
		return err
	}

	p.isInGame = true
	return nil
}

func (p *Player) Leave() error {
	if !p.isInGame {
		return errors.New("Player not in game")
	}
	p.isInGame = false

	return nil
}

func (p *Player) BuyChips(count int) {
	p.chips += count
}

func (p *Player) AvailableChips() int {
	return p.chips
}
