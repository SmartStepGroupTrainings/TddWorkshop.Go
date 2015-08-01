package casino

type Player struct {
	currentGame *Game
}

func (self *Player) BuyChips (count int) {

}

func (self *Player) ChipsCount() int {
	return 1
}