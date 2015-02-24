package casino

type IGame interface {
	Add(*Player) error
	Bet(Bet, *Player) error
}
