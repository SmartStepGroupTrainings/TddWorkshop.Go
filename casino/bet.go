package casino

type Bet struct {
	Amount int
}

func (b *Bet) GetAmount() int {
	return b.Amount
}
