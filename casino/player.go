package casino

type Player struct {
	isInGame bool
}

func (p *Player) IsInGame() bool {
	return p.isInGame
}

func (p *Player) Join() {
	p.isInGame = true
}
