package casino

type IGame interface {
	addPlayer(*Player)
	removePlayer(*Player)
	HasPlayer(*Player) bool
}
