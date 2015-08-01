package casino

type Player struct {
	isInGame bool
}

func (p *Player) Join(game Game) {
	p.isInGame = true
}

func (p *Player) Leave(game Game) {
	p.isInGame = false
}
