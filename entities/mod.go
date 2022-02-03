package entities

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

// Stone - 石の色
type Stone int

const (
	// Empty - 空点
	Empty Stone = iota
	// Black - 黒石
	Black
	// White - 白石
	White
	// Wall - 壁
	Wall
)

// FlipColor - 白黒反転させます。
func FlipColor(color Stone) Stone {
	return 3 - color
}

func SetBoardSize(boardSize int) {
	BoardSize = boardSize
	BoardArea = BoardSize * BoardSize
	SentinelWidth = BoardSize + 2
	SentinelBoardArea = SentinelWidth * SentinelWidth
}

// BoardSize - 何路盤
var BoardSize int

// BoardArea - 壁無し盤の面積
var BoardArea int

// SentinelWidth - 枠付きの盤の一辺の交点数
var SentinelWidth int

// SentinelBoardArea - 壁付き盤の面積
var SentinelBoardArea int

type KomiType float64

// Komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
var Komi KomiType

type MovesNumType int

// MaxMovesNum - 上限手数
var MaxMovesNum MovesNumType

// Point - 交点の座標。壁を含む盤の左上を 0 とします
type Point int

// Pass - パス
const Pass Point = 0

// Dir4 - ４方向（東、北、西、南）の番地。初期値は仮の値。 2015年講習会サンプル、GoGo とは順序が違います
var Dir4 = [4]Point{1, -9, -1, 9}

type Direction4 int

// Dir4に対応
const (
	East Direction4 = iota
	North
	West
	South
)
