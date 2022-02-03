package play_algorithm

import (
	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// WrapGettingOfWinner - 盤を束縛変数として与えます
func WrapGettingOfWinner(position *e.Position) *func(turnColor e.Stone) int {
	// 「手番の勝ちなら1、引き分けなら0、手番の負けなら-1を返す関数（自分視点）」を作成します
	// * `turnColor` - 手番の石の色
	var getWinner = func(turnColor e.Stone) int {
		return getWinner(position, turnColor)
	}

	return &getWinner
}

// 手番の勝ちなら1、引き分けなら0、手番の負けなら-1（自分視点）
// * `turnColor` - 手番の石の色
func getWinner(position *e.Position, turnColor e.Stone) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int

	var onPoint = func(z e.Point) {
		var color2 = position.ColorAt(z)
		kind[color2]++
		if color2 == 0 {
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[position.ColorAt(z+e.Dir4[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}

	position.IterateWithoutWall(onPoint)

	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	var win = 0
	if 0 < e.KomiType(score)-e.Komi {
		win = 1
	}
	if turnColor == 2 {
		win = -win
	} // gogo07

	return win
}
