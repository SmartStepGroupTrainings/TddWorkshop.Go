package casino_new

import (
	"errors"
	"math/rand"
	"time"
)

type IDice interface {
	Roll() int
}

type Dice struct {
}

func (dice *Dice) Roll() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Int()%6 + 1
}

type RollDiceGame struct {
	players map[*Player]struct{}
	dice    IDice
}

func NewRollDiceGame() *RollDiceGame {
	return &RollDiceGame{
		players: make(map[*Player]struct{}),
		dice:    &Dice{},
	}
}

func (self *RollDiceGame) setDice(dice IDice) {
	self.dice = dice
}

func (self *RollDiceGame) Play() error {
	if self.PlayersCount() == 0 {
		return errors.New("Cannot start game without any player")
	}

	winningScore := self.dice.Roll()

	for player, _ := range self.players {
		player.Win(player.GetBetOn(winningScore) * 6)
		player.Lose()
	}
	return nil
}

func (self *RollDiceGame) Add(player *Player) error {
	// !!! Added after testing
	if player == nil {
		return errors.New("Player for adding to game cannot be nil")
	}
	if _, ok := self.players[player]; ok {
		return errors.New("Player already in this game")
	}
	player.setCurrentGame(self)
	// ====

	self.players[player] = struct{}{}
	return nil
}

func (self *RollDiceGame) PlayersCount() int {
	return len(self.players)
}

func (self *RollDiceGame) Remove(player *Player) error {
	curGame, err := player.getCurrentGame()
	if err != nil {
		return err
	}
	if curGame != self {
		return errors.New("this player in another game")
	}
	delete(self.players, player)
	return nil
}
