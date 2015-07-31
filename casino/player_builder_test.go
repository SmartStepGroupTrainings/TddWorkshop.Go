package casino_new

type PlayerBuilder struct {
	player *Player
}

func NewPlayerBuilder() *PlayerBuilder {
	return &PlayerBuilder {
		player: NewPlayer(),
	}
}

func (playerBuilder *PlayerBuilder) WithBet(amount int, score int) *PlayerBuilder {
	playerBuilder.player.BuyChips(100)
	playerBuilder.player.Bet(playerBuilder.bet(amount, score))
	return playerBuilder
}

func (playerBuilder *PlayerBuilder) bet(amount int, score int) Bet {
	return Bet {
		Amount: amount,
		Score: score,
	}
}

func (playerBuilder *PlayerBuilder) Get() *Player {
	return playerBuilder.player
}
