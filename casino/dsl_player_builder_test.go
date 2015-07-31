package casino_new

type PlayerBuilder struct {
	game        *RollDiceGame
	betScore    int
	betAmount   int
	chipsAmount int
}

func (this *PlayerBuilder) InGame(game *RollDiceGame) *PlayerBuilder {
	this.game = game
	return this
}

func (this *PlayerBuilder) BetOn(score int) *PlayerBuilder {
	this.betScore = score
	return this
}

func (this *PlayerBuilder) BetAmount(amount int) *PlayerBuilder {
	this.betAmount = amount
	return this
}

func (this *PlayerBuilder) Rich() *PlayerBuilder {
	this.chipsAmount = 1000
	return this
}

func (this *PlayerBuilder) Please() *Player {
	player := NewPlayer()
	player.BuyChips(this.chipsAmount)
	if this.betScore > 0 || this.betAmount > 0 {
		player.Bet(Bet{Score: this.betScore, Amount: this.betAmount})
	}
	if this.game != nil {
		player.Join(this.game)
	}

	return player
}
