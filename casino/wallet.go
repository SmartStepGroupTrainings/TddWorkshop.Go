package casino

type Wallet struct {
	bets    []PlayerBet
	balance Chips
}

func (wallet *Wallet) Balance() Chips {
	return wallet.balance
}

func (wallet *Wallet) Add(bet Bet, player *Player) {
	wallet.bets = append(wallet.bets, PlayerBet{bet: bet, player: player})
}

func (wallet *Wallet) Play(winningScore Score) {
	for _, pb := range wallet.bets {
		if pb.bet.Score == winningScore {
			pb.player.Win(pb.bet.Chips * 6)
		} else {
			wallet.balance += pb.bet.Chips
		}
	}

	wallet.bets = nil
}
