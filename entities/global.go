package entities

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

// Record - 棋譜
var Record []IRecordItemV01

// Dir4 - ４方向（右、下、左、上）の番地。初期値は仮の値。
var Dir4 = [4]int{1, 9, -1, 9}

// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
var KoZ int

// For count liberty.
var checkBoard = []int{}

// MovesNum - 手数？
var MovesNum int

// FlagTestPlayout - ？。
var FlagTestPlayout int

// FlipColor - 白黒反転させます。
func FlipColor(col int) int {
	return 3 - col
}

var GettingOfWinnerOnDuringUCTPlayout func(int) int
