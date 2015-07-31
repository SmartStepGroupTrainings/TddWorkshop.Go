package casino_new

var create = Builder{}

type Builder struct{}

type GameBuilder struct {
	winScore int
}

type PlayerBuilder struct {
	player    *Player
	betScore  int
	betAmount int
}

func (b Builder) Game() *GameBuilder {
	return &GameBuilder{}
}

func (b Builder) Player() *PlayerBuilder {
	return &PlayerBuilder{player: NewPlayer()}
}

func (b *GameBuilder) WinningScore(score int) *GameBuilder {
	b.winScore = score
	return b
}

func (b *GameBuilder) Please() *RollDiceGame {
	game := NewRollDiceGame()

	if b.winScore != 0 {
		dice := &DiceStub{}
		dice.On("Roll").Return(b.winScore)
		game.dice = dice
	}

	return game
}

func (b *PlayerBuilder) WithChips(amount int) *PlayerBuilder {
	b.player.BuyChips(amount)
	return b
}

func (b *PlayerBuilder) InGame(game *RollDiceGame) *PlayerBuilder {
	if game == nil {
		game = create.Game().Please()
	}
	b.player.Join(game)
	return b
}

func (b *PlayerBuilder) OnScore(score int) *PlayerBuilder {
	b.betScore = score
	return b
}

func (b *PlayerBuilder) Bets(amount int) *PlayerBuilder {
	b.betAmount = amount
	return b
}

func (b *PlayerBuilder) MakesBet() *PlayerBuilder {
	// bet amount does not matter
	b.betAmount = 1
	return b
}

func (b *PlayerBuilder) Please() *Player {
	if b.betScore != 0 && b.betAmount != 0 {
		b.player.Bet(Bet{
			Score:  b.betScore,
			Amount: b.betAmount,
		})
	}
	return b.player
}
