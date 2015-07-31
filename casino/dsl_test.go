package casino_new

var dsl = &DSL{}

type DSL struct {
}

func (d *DSL) CreateGame() *GameBuilder {
	return &GameBuilder{
		nextRoll: 1,
	}
}

func (d *DSL) CreatePlayer() *PlayerBuilder {
	return &PlayerBuilder{
		chips: 1000000,
	}
}

// GameBuilder
type GameBuilder struct {
	nextRoll int
}

func (b *GameBuilder) NextRoll(v int) *GameBuilder {
	b.nextRoll = v

	return b
}

func (b *GameBuilder) Build() *RollDiceGame {
	game := NewRollDiceGame()
	game.dice = &testDice{b.nextRoll}

	return game
}

// PlayerBuilder
type PlayerBuilder struct {
	chips int
	game  *RollDiceGame
	bets []Bet
}

func (b *PlayerBuilder) WithChips(v int) *PlayerBuilder {
	b.chips = v

	return b
}

func (b *PlayerBuilder) JoinGame(g *RollDiceGame) *PlayerBuilder {
	b.game = g

	return b
}

func (b *PlayerBuilder) MakeBet(score int, amount int) *PlayerBuilder {
	b.bets = append(b.bets, Bet{score, amount})

	return b
}

func (b *PlayerBuilder) Build() *Player {
	player := NewPlayer()

	player.BuyChips(b.chips)

	if b.game != nil {
		player.Join(b.game)
	}

	for _ , bet := range b.bets {
		player.Bet(bet)
	}

	return player
}
