package casino

import (
	"errors"
	//"log"
)

type PlayerBet struct {
	bet    Bet
	player *Player
}

type Game struct {
	dice    IDice
	players []*Player
	bets    []PlayerBet
}

type IDice interface {
	Roll() Score
}

func (game *Game) Add(player *Player) error {
	if player.IsInGame() {
		return errors.New("Please leave the game before joining another game")
	}

	if len(game.players) == 6 {
		return errors.New("Please join another game")
	}

	game.players = append(game.players, player)
	player.game = game
	return nil
}

func (game *Game) Remove(player *Player) error {
	for i, p := range game.players {
		if p == player {
			game.players = append(game.players[:i], game.players[i+1:]...)
			player.game = nil
			return nil
		}
	}
	return errors.New("Please join the game before leaving")
}

func (game *Game) Bet(bet Bet, player *Player) {
	game.bets = append(game.bets, PlayerBet{bet: bet, player: player})
}

func (game *Game) HasPlayer(player *Player) bool {
	for _, p := range game.players {
		if p == player {
			return true
		}
	}
	return false
}

func (game *Game) Play() {
	winningScore := game.dice.Roll()

	for _, pb := range game.bets {
		if pb.bet.Score == winningScore {
			pb.player.Win(pb.bet.Chips * 6)
		}
	}
}
