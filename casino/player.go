package casino

type Player struct {
	isInGame bool
}

func (p *Player) Join(game Game) {
	p.isInGame = true
}
