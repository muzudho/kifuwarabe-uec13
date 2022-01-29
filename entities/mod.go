package entities

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

func SetBoardSize(boardSize int) {
	BoardSize = boardSize
	SentinelWidth = BoardSize + 2
	SentinelBoardArea = SentinelWidth * SentinelWidth
}

// BoardSize - 何路盤
var BoardSize int

// SentinelWidth - 枠付きの盤の一辺の交点数
var SentinelWidth int

// SentinelBoardArea - 壁付き盤の面積
var SentinelBoardArea int

// Komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
var Komi float64

// MaxMovesNum - 上限手数
var MaxMovesNum int

// Dir4 - ４方向（右、下、左、上）の番地。初期値は仮の値。
var Dir4 = [4]int{1, 9, -1, 9}

// For count liberty.
var checkBoard = []int{}

// FlipColor - 白黒反転させます。
func FlipColor(col int) int {
	return 3 - col
}
