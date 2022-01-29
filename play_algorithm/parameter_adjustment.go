package play_algorithm

import e "github.com/muzudho/kifuwarabe-uec13/entities"

// プレイアウトする回数（あとで設定されます）
var PlayoutTrialCount = 0

// UCTをループする回数
var UctLoopCount = 4500

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
}
