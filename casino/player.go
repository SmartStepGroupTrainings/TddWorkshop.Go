package casino

import "errors"

var (
	errPlayerAlreadyInGame  = errors.New("player is already in game")
	errPlayerNotInGame      = errors.New("player is not in game")
	errBuyNegativeChips     = errors.New("player can't by negative chips amount")
	errNotEnoughChips       = errors.New("not enough chips")
	errBetScoreIsNotAllowed = errors.New("bet score is not allowed")
)

type Player struct {
	availableChips int
	bets           map[int]*Bet
	game           *Game
}

func (p *Player) IsInGame() bool {
	return p.game != nil
}

func (p *Player) Join(game *Game) error {
	if p.IsInGame() {
		return errPlayerAlreadyInGame
	}

	if err := game.AddPlayer(p); err != nil {
		return err
	}

	p.game = game
	return nil
}

func (p *Player) Leave(game *Game) error {
	if !p.IsInGame() {
		return errPlayerNotInGame
	}
	p.game = nil
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

func (p *Player) DoBet(bet *Bet) error {
	if bet.Score > 6 || bet.Score < 1 {
		return errBetScoreIsNotAllowed
	}

	if p.GetAvailableChips() < bet.Amount {
		return errNotEnoughChips
	}

	if !p.IsInGame() {
		return errPlayerNotInGame
	}

	if err := p.game.IsBetAmountValid(bet.Amount); err != nil {
		return err
	}

	if p.bets == nil {
		p.bets = make(map[int]*Bet)
	}

	p.bets[bet.Score] = bet
	return nil
}

func (p *Player) GetBetByScore(score int) *Bet {
	return p.bets[score]
}

func (p *Player) GetBetCount() int {
	return len(p.bets)
}
