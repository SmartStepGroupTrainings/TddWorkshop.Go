package casino_new

type Creator struct {
	Game          *GameBuilder
	Player        *PlayerBuilder
	SeveralPlayer *SeveralPlayerBuilder
}

var (
	create Creator
)

func init() {
	create = Creator{
		Game:          &GameBuilder{ctx: context{}, game: NewRollDiceGame()},
		Player:        &PlayerBuilder{ctx: context{}, player: NewPlayer()},
		SeveralPlayer: &SeveralPlayerBuilder{ctx: context{}},
	}

}

type context struct {
	registr int
	filled  bool
}

type GameBuilder struct {
	ctx  context
	game *RollDiceGame
}

type PlayerBuilder struct {
	ctx    context
	player *Player
}

type SeveralPlayerBuilder struct {
	ctx     context
	players []Player
}

func (ctx *context) Fill(i int) {
	ctx.filled = true
	ctx.registr = i
}

func (ctx *context) Get() int {
	if !ctx.filled {
		panic("put before get!")
	}
	ctx.filled = false
	return ctx.registr
}

func (gameBuilder GameBuilder) With(i int) GameBuilder {
	gameBuilder.ctx.Fill(i)
	return gameBuilder
}
func (playerBuilder PlayerBuilder) With(i int) PlayerBuilder {
	playerBuilder.ctx.Fill(i)
	return playerBuilder
}
func (severalPlayerBuilder SeveralPlayerBuilder) With(i int) SeveralPlayerBuilder {
	severalPlayerBuilder.ctx.Fill(i)
	return severalPlayerBuilder
}

func (gameBuilder GameBuilder) WinningScore() GameBuilder {
	score := gameBuilder.ctx.Get()

	dice := &DiceStub{}
	dice.On("Roll").Return(score)
	gameBuilder.ctx.filled = false
	gameBuilder.game.setDice(dice)

	return gameBuilder
}

func (gameBuilder GameBuilder) Build() *RollDiceGame {
	return gameBuilder.game
}

func (playerBuilder PlayerBuilder) Chips() PlayerBuilder {
	chips := playerBuilder.ctx.Get()
	playerBuilder.player.BuyChips(chips)
	return playerBuilder
}

func (playerBuilder PlayerBuilder) Build() *Player {
	return playerBuilder.player
}

func (playerBuilder PlayerBuilder) BetOn(score, amount int) PlayerBuilder {
	playerBuilder.player.Bet(Bet{Score: score, Amount: amount})
	return playerBuilder
}

func (playerBuilder PlayerBuilder) JointTo(game *RollDiceGame) PlayerBuilder {
	playerBuilder.player.Join(game)

	return playerBuilder
}
