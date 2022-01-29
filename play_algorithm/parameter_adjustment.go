package play_algorithm

import (
	"math"

	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// プレイアウトする回数（あとで設定されます）
var PlayoutTrialCount = 0

// UCTをループする回数（あとで設定されます）
var UctLoopCount = 0

// var UctLoopCount = 2500 // 少なめ（ペンキ塗り）
// var UctLoopCount = 4500 // 盤の上の方に集まりがち
// var UctLoopCount = 5000 // ペンキ塗りはしなくなる
// var UctLoopCount = 10000 // 多め（遅い感じ）

func AdjustParameters(position *e.Position) {
	var boardSize = e.BoardSize
	if boardSize < 10 {
		// 10路盤より小さいとき
		PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		PlayoutTrialCount = boardSize * boardSize
	}

	// 面積 * 2 pi e で、盤面全体をだいぶカバーします（自分調べ）
	UctLoopCount = int(math.Ceil(float64(e.BoardArea) * 2 * math.Pi * math.E))
}
