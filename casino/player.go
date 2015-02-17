package casino

import "errors"
import "fmt"

type Player struct {
	IsInGame bool
	balance  Chips
	bets     []Bet
}

type Chips uint
type Score uint

func (player *Player) Buy(chips Chips) {
	player.balance += chips
}

func (player *Player) Balance() Chips {
	return player.balance
}

func (player *Player) Bet(chips Chips, score Score) error {
	if player.Balance() < chips {
		return fmt.Errorf("You can not bet more than %v chips", player.Balance())
	}

	if !player.IsInGame {
		return errors.New("You should join a game before making a bet")
	}

	if score < 1 || 6 < score {
		return errors.New("Please make a bet only to score 1 - 6")
	}

	player.bets = append(player.bets, Bet{Score: score, Chips: chips})
	player.balance -= chips
	return nil
}

func (player *Player) Bets() []Bet {
	return player.bets
}
