package entities

import (
	"math/rand"
)

// Board - 盤
type Board struct {
	data                []int
	boardSize           int
	sentinelWidth       int
	sentinelBoardArea   int
	komi                float64
	maxMoves            int
	iteratorWithoutWall func(func(int))

	uctChildrenSize int
}

// NewBoard - 盤を作成します。
// TODO Board の NewBoard を呼び出す方法がない？
func NewBoard(data []int, boardSize int, sentinelBoardArea int, komi float64, maxMoves int) *Board {
	var board = new(Board)
	board.data = data
	board.boardSize = boardSize
	board.sentinelWidth = boardSize + 2
	board.sentinelBoardArea = sentinelBoardArea
	board.komi = komi
	board.maxMoves = maxMoves
	board.uctChildrenSize = boardSize*boardSize + 1
	board.iteratorWithoutWall = CreateBoardIteratorWithoutWall(board)

	checkBoard = make([]int, board.SentinelBoardArea())
	Record = make([]IRecordItemV01, board.MaxMovesNum())
	Dir4 = [4]int{1, board.SentinelWidth(), -1, -board.SentinelWidth()}

	return board
}

// InitBoard - 盤の初期化。
func (board *Board) InitBoard() {
	boardMax := board.SentinelBoardArea()

	// 枠線
	for z := 0; z < boardMax; z++ {
		board.SetColor(z, 3)
	}

	// 盤上
	var onPoint = func(z int) {
		board.SetColor(z, 0)
	}
	board.iteratorWithoutWall(onPoint)

	MovesNum = 0
	KoZ = 0
}

// BoardSize - 何路盤か
func (board Board) BoardSize() int {
	return board.boardSize
}

// SentinelWidth - 枠付きの盤の一辺の交点数
func (board Board) SentinelWidth() int {
	return board.sentinelWidth
}

// SentinelBoardArea - 枠付きの盤の交点数
func (board Board) SentinelBoardArea() int {
	return board.sentinelBoardArea
}

// Komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
func (board Board) Komi() float64 {
	return board.komi
}

// MaxMovesNum - 最大手数
func (board Board) MaxMovesNum() int {
	return board.maxMoves
}

// ColorAt - 指定した交点の石の色
func (board Board) ColorAt(z int) int {
	return board.data[z]
}

// ColorAtXy - 指定した交点の石の色
func (board Board) ColorAtXy(x int, y int) int {
	return board.data[(y+1)*board.sentinelWidth+x+1]
}

// Exists - 指定の交点に石があるか？
func (board Board) Exists(z int) bool {
	return board.data[z] != 0
}

// SetColor - 盤データ。
func (board *Board) SetColor(z int, color int) {
	board.data[z] = color
}

// CopyData - 盤データのコピー。
func (board Board) CopyData() []int {
	boardArea := board.SentinelBoardArea()

	var boardCopy2 = make([]int, boardArea)
	copy(boardCopy2[:], board.data[:])
	return boardCopy2
}

// ImportData - 盤データのコピー。
func (board *Board) ImportData(boardCopy2 []int) {
	copy(board.data[:], boardCopy2[:])
}

// GetZ4 - z（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func (board Board) GetZ4(z int) int {
	if z == 0 {
		return 0
	}
	y := z / board.SentinelWidth()
	x := z - y*board.SentinelWidth()
	return x*100 + y
}

// GetZFromXy - x,y 形式の座標を、 z （配列のインデックス）へ変換します。
// x,y は壁を含まない領域での座標です。 z は壁を含む領域での座標です
func (board Board) GetZFromXy(x int, y int) int {
	return (y+1)*board.SentinelWidth() + x + 1
}

// GetEmptyZ - 空点の z （配列のインデックス）を返します。
func (board Board) GetEmptyZ() int {
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

func (board Board) countLibertySub(z int, color int, pLiberty *int, pStone *int) {
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
func (board Board) CountLiberty(z int, pLiberty *int, pStone *int) {
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
func (board *Board) TakeStone(z int, color int) {
	board.data[z] = 0
	for dir := 0; dir < 4; dir++ {
		z2 := z + Dir4[dir]
		if board.data[z2] == color {
			board.TakeStone(z2, color)
		}
	}
}

// IterateWithoutWall - 盤イテレーター
func (board Board) IterateWithoutWall(onPoint func(int)) {
	board.iteratorWithoutWall(onPoint)
}

// UctChildrenSize - UCTの最大手数
func (board Board) UctChildrenSize() int {
	return board.uctChildrenSize
}

// CreateBoardIteratorWithoutWall - 盤の（壁を除く）全ての交点に順にアクセスする boardIterator 関数を生成します
func CreateBoardIteratorWithoutWall(
	board *Board) func(func(int)) {

	var boardSize = board.BoardSize()
	var boardIterator = func(onPoint func(int)) {

		// x,y は壁無しの盤での0から始まる座標、 z は壁有りの盤での0から始まる座標
		for y := 0; y < boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				var z = board.GetZFromXy(x, y)
				onPoint(z)
			}
		}
	}

	return boardIterator
}
