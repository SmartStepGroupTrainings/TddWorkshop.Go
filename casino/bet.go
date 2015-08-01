package casino

type Bet struct {
	Amount int
    Score int
}

func (b *Bet) GetAmount() int {
	return b.Amount
}
