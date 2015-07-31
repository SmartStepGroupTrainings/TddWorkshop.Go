package casino_new

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DiceSuite struct {
	CasinoBasicSuite
}

func TestDiceSuite(t *testing.T) {
	suite.Run(t, new(DiceSuite))
}

func (suite *RollDiceGameSuite) TestDice_Roll_ReturnScoreGraterThanZero() {
	dice := NewDice()
	score := dice.Roll()

	suite.AssertTrue(score > 0)
}

func (suite *RollDiceGameSuite) TestDice_Roll_ReturnScoreLessThanSix() {
	dice := NewDice()
	score := dice.Roll()

	suite.AssertTrue(score <= 6)
}