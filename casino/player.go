package casino

import "errors"

type Player struct {
	isInGame bool
	chips    int
}

func (p *Player) Join(game *Game) error {
	if p.isInGame {
		return errors.New("You can not join game")
	}

	if err := game.AddPlayer(); err != nil {
		return err
	}

	p.isInGame = true
	return nil
}

func (p *Player) Leave(game *Game) error {

	if !p.isInGame {
		return errors.New("You can not leave game")
	}

	p.isInGame = false
	return nil
}

func (p *Player) IsInGame() bool {
	return p.isInGame
}

func (p *Player) BuyChips(count int) {
	p.chips += count
}

func (p *Player) GetChipsCount() int {
	return p.chips
}

func (p *Player) Bet(count int) error {
	if p.GetChipsCount() < count {
		return errors.New("Not enough chips")
	}

	p.chips -= count

	return nil
}
