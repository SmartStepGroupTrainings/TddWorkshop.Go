package casino_new

import "github.com/stretchr/testify/mock"

type DiceMock struct {
	mock.Mock
}

func (s *DiceMock) Roll() int {
	args := s.Called()
	return args.Int(0)
}

type GameTest struct {
	PlayerTest
	Game *RollDiceGame
	dice *DiceMock
}

func (self *GameTest) SetupTest() {
	self.Player = NewPlayer()
	self.dice = &DiceMock{}
	self.Game = NewRollDiceGame(self.dice)
}

func (self *GameTest) TestPlay_MakeBet_WinOk() {
	self.Player.Join(self.Game)
	self.Player.BuyChips(10)
	self.dice.On("Roll").Return(1)
	self.Player.Bet(Bet{Amount: 1, Score: 1})

	self.Game.Play()


}

func (self *GameTest) TestPlay_MakeBet_WinOk_DSL() {
	игра := Сконструировать{}.ИгруСФиксированнымСчетом(1).Поехали()
	игрок := Сконструировать{}.Игрока().СФишечками(10).ПрисоединитьКИгре(игра).ПоставитьБабки(1).ПоставитьНа(1).Поехали()

    игра.Play()

    self.Equal((10-1)+1*6, игрок.AvailableChips())
}

func (self *GameTest) TestPlay_MakeBet_LooseOk() {
	self.Player.Join(self.Game)
	self.Player.BuyChips(10)
	self.dice.On("Roll").Return(2)
	self.Player.Bet(Bet{Amount: 1, Score: 1})

	self.Game.Play()

	self.Equal(10-1, self.Player.AvailableChips())
}

func (self *GameTest) TestGameMethodWorksOk() {
	self.Game.XXX()
}

// Можно определить метод к объекту из основного кода
// Это кстати, аргумент класть тесты рядом с основным
// кодом, а не во внешние пакеты.
func (self *RollDiceGame) XXX() {
	self.testXXX = true
	return
}
