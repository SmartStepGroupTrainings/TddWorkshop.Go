package casino

import "errors"

type Player struct {
	game  *Game
	chips int
}

func (p *Player) IsInGame() bool {
	return p.game != nil
}

func (p *Player) Join(game *Game) error {
	if p.IsInGame() {
		return errors.New("Player already in game")
	}

	if err := game.Add(); err != nil {
		return err
	}

	p.game = game
	return nil
}

func (p *Player) Leave() error {
	if !p.IsInGame() {
		return errors.New("Player not in game")
	}
	p.game = nil

	return nil
}

func (p *Player) BuyChips(count int) {
	p.chips += count
}

func (p *Player) AvailableChips() int {
	return p.chips
}

func (p *Player) MakeBet(bet Bet) error {
	if p.AvailableChips() < bet.Amount {
		return errors.New("Not enouth chips for bet")
	}

	p.chips -= bet.Amount

	return nil
}
