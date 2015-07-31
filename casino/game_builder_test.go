package casino_new

import (
	"github.com/stretchr/testify/mock"
)

/*
playerBuilder := NewPlayerBuilder()
player := playerBuilder.addMoney(100).makeBet(50, 5).Get()
game.addPlyaer(plyaer).Get().Play()



player := builder.PlayerWitMoney(100).makeBet(50, 5)
game := builder.NewGame().Join(player).Get()

*/


type diceMock struct{
	mock.Mock
}

func (m diceMock) Roll() int {
	args := m.Called()
	return args.Int(0)
}

type RollDiceGameBuilder struct {
	players map[*Player]struct{}
	winningScore int
}

func NewRollDiceGameBuilder() *RollDiceGameBuilder {
	return &RollDiceGameBuilder {
		players: make(map[*Player]struct{}),
	}
}

func (builder *RollDiceGameBuilder) addPlayerWithBet(amount int, score int) *RollDiceGameBuilder {
	player := NewPlayerBuilder().WithBet(amount, score).Get()
	builder.players[player] = struct{}{}
	return builder
}

func (builder *RollDiceGameBuilder) setWinningScore(winningScore int) *RollDiceGameBuilder {
	builder.winningScore = winningScore
	return builder
}

func (builder *RollDiceGameBuilder) Get() *RollDiceGame {
	dice := diceMock{}
	dice.On("Roll").Return(builder.winningScore)

	game := NewRollDiceGameWithDice(dice)

	for player := range builder.players {
		player.Join(game)
	}

	return game
}
