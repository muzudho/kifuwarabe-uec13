package play_algorithm

import (
	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// プレイアウトする回数（あとで設定されます）
var PlayoutTrialCount = 0

// UCTをループする回数（あとで設定されます）
var UctLoopCount = 0

// ランダム鳩の巣仮説定数 a。およそ 18
// 面積 * 2 pi e 、つまり およそ 17 で、５００回に１回見落としがある程度、
// 面積 * (2 pi e + 1) 、 つまり およそ 18 で、１万回に１回見落としがある程度の精度（自分調べ）
var randomPigeonA = 17 // 2 * math.Pi * math.E

// ランダム鳩の巣仮説 試行回数 x
// 📖 [random-pigeon-nest-hypothesis](https://github.com/muzudho/random-pigeon-nest-hypothesis)
func GetRandomPigeonX(N int) int {
	return N * randomPigeonA
	// return int(math.Ceil(float64(N) * randomPigeonA))
}

func AdjustParameters(position *e.Position) {
	var boardSize = e.BoardSize
	if boardSize < 10 {
		// 10路盤より小さいとき
		PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		PlayoutTrialCount = boardSize * boardSize
	}

	// 盤面全体を１回は選ぶことを、完璧ではありませんが、ある程度の精度でカバーします
	UctLoopCount = GetRandomPigeonX(e.BoardArea)
}
