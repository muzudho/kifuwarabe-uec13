package entities

import (
	"math/rand"
)

// BoardV01 - 盤
type BoardV01 struct {
	data                []int
	boardSize           int
	sentinelWidth       int
	sentinelBoardArea   int
	komi                float64
	maxMoves            int
	iteratorWithoutWall func(func(int))
}

// InitBoard - 盤の初期化。
func (board *BoardV01) InitBoard() {
	boardMax := board.SentinelBoardArea()
	// G.Chat.Trace("# (^q^) boardMax=%d\n", boardMax)

	// 枠線
	for z := 0; z < boardMax; z++ {
		board.SetColor(z, 3)
	}

	// G.Chat.Trace("# (^q^) 盤を 3 で埋めた☆\n")

	// 盤上
	var onPoint = func(z int) {
		board.SetColor(z, 0)
	}
	board.iteratorWithoutWall(onPoint)

	// G.Chat.Trace("# (^q^) 石は置いた☆\n")

	MovesNum = 0
	KoZ = 0

	// G.Chat.Trace("# (^q^) 盤の初期化は終わったぜ☆\n")
}

// BoardSize - 何路盤か
func (board BoardV01) BoardSize() int {
	return board.boardSize
}

// SentinelWidth - 枠付きの盤の一辺の交点数
func (board BoardV01) SentinelWidth() int {
	return board.sentinelWidth
}

// SentinelBoardArea - 枠付きの盤の交点数
func (board BoardV01) SentinelBoardArea() int {
	return board.sentinelBoardArea
}

// Komi - コミ
func (board BoardV01) Komi() float64 {
	return board.komi
}

// MaxMovesNum - 最大手数
func (board BoardV01) MaxMovesNum() int {
	return board.maxMoves
}

// ColorAt - 指定した交点の石の色
func (board BoardV01) ColorAt(z int) int {
	return board.data[z]
}

// ColorAtXy - 指定した交点の石の色
func (board BoardV01) ColorAtXy(x int, y int) int {
	return board.data[(y+1)*board.sentinelWidth+x+1]
}

// Exists - 指定の交点に石があるか？
func (board BoardV01) Exists(z int) bool {
	return board.data[z] != 0
}

// SetColor - 盤データ。
func (board *BoardV01) SetColor(z int, color int) {
	board.data[z] = color
}

// CopyData - 盤データのコピー。
func (board BoardV01) CopyData() []int {
	boardArea := board.SentinelBoardArea()

	var boardCopy2 = make([]int, boardArea)
	copy(boardCopy2[:], board.data[:])
	return boardCopy2
}

// ImportData - 盤データのコピー。
func (board *BoardV01) ImportData(boardCopy2 []int) {
	copy(board.data[:], boardCopy2[:])
}

// GetZ4 - z（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func (board BoardV01) GetZ4(z int) int {
	if z == 0 {
		return 0
	}
	y := z / board.SentinelWidth()
	x := z - y*board.SentinelWidth()
	return x*100 + y
}

// GetZFromXy - x,y を z （配列のインデックス）へ変換します。
// x,y は壁を含まない領域での座標です。 z は壁を含む領域での座標です
func (board BoardV01) GetZFromXy(x int, y int) int {
	return (y+1)*board.SentinelWidth() + x + 1
}

// GetEmptyZ - 空点の z （配列のインデックス）を返します。
func (board BoardV01) GetEmptyZ() int {
	var x, y, z int
	for {
		// ランダムに交点を選んで、空点を見つけるまで繰り返します。
		x = rand.Intn(9)
		y = rand.Intn(9)
		z = board.GetZFromXy(x, y)
		if !board.Exists(z) {
			break
		}
	}
	return z
}

func (board BoardV01) countLibertySub(z int, color int, pLiberty *int, pStone *int) {
	checkBoard[z] = 1
	*pStone++
	for i := 0; i < 4; i++ {
		z := z + Dir4[i]
		if checkBoard[z] != 0 {
			continue
		}
		if !board.Exists(z) {
			checkBoard[z] = 1
			*pLiberty++
		}
		if board.data[z] == color {
			board.countLibertySub(z, color, pLiberty, pStone)
		}
	}

}

// CountLiberty - 呼吸点を数えます。
func (board BoardV01) CountLiberty(z int, pLiberty *int, pStone *int) {
	*pLiberty = 0
	*pStone = 0
	boardMax := board.SentinelBoardArea()
	// 初期化
	for z2 := 0; z2 < boardMax; z2++ {
		checkBoard[z2] = 0
	}
	board.countLibertySub(z, board.data[z], pLiberty, pStone)
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (board *BoardV01) TakeStone(z int, color int) {
	board.data[z] = 0
	for dir := 0; dir < 4; dir++ {
		z2 := z + Dir4[dir]
		if board.data[z2] == color {
			board.TakeStone(z2, color)
		}
	}
}

// IterateWithoutWall - 盤イテレーター
func (board BoardV01) IterateWithoutWall(onPoint func(int)) {
	board.iteratorWithoutWall(onPoint)
}
