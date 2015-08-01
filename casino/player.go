package casino

type Player struct {
	isInGame bool
}

func (self *Player) IsIn(game *Game) bool {
	return self.isInGame
}

func (self *Player) Join(game *Game) error {
	if err := game.AddPlayer(); err != nil {
		return err
	}

	self.isInGame = true
	return nil
}

func (self *Player) Leave() {
	self.isInGame = false
}
