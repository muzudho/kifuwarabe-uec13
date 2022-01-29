package play_algorithm

import e "github.com/muzudho/kifuwarabe-uec13/entities"

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

var GettingOfWinnerOnDuringUCTPlayout *func(int) int
var SearchingOfUct *func(color int, nodeN int) int

// FlagTestPlayout - ？。
var FlagTestPlayout int

func InitPosition(position *e.Position) {
	// 盤サイズが変わっていることもあるので、先に初期化します
	position.InitPosition()

	GettingOfWinnerOnDuringUCTPlayout = WrapGettingOfWinner(position)
	SearchingOfUct = WrapSearchUct(position)
	AdjustParameters(position)
}
