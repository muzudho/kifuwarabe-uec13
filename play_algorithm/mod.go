package play_algorithm

import (
	e "github.com/muzudho/kifuwarabe-uec13/entities"
	gd "github.com/muzudho/kifuwarabe-uec13/game_domain"
)

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

var GettingOfWinnerOnDuringUCTPlayout *func(e.Stone) int
var IsDislike *func(e.Stone, e.Point) bool

// FlagTestPlayout - ？。
var FlagTestPlayout int

func InitPosition(position *e.Position) {
	// 盤サイズが変わっていることもあるので、先に初期化します
	position.InitPosition()

	GettingOfWinnerOnDuringUCTPlayout = WrapGettingOfWinner(position)
	IsDislike = gd.WrapIsDislike(position)
	AdjustParameters(position)
}
