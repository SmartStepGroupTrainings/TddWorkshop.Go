package casino

import "errors"

var (
	errPlayerAlreadyInGame    = errors.New("Player is already in game")
	errPlayerAlreadyNotInGame = errors.New("Player is already not in game")
	errBuyNegativeChips       = errors.New("Player can't by negative chips amount")
)

type Player struct {
	isInGame       bool
	availableChips int
}

func (p *Player) IsInGame() bool {
	return p.isInGame
}

func (p *Player) Join(game *Game) error {
	if p.IsInGame() {
		return errPlayerAlreadyInGame
	}

	if err := game.AddPlayer(p); err != nil {
		return err
	}

	p.isInGame = true
	return nil
}

func (p *Player) Leave(game *Game) error {
	if !p.IsInGame() {
		return errPlayerAlreadyNotInGame
	}
	p.isInGame = false
	return nil
}

func (p *Player) GetAvailableChips() int {
	return p.availableChips
}

func (p *Player) BuyChips(count int) error {
	if count < 0 {
		return errBuyNegativeChips
	}
	p.availableChips += count
	return nil
}
