package entities

import (
	"math/rand"
)

// Position - 盤
type Position struct {
	// 盤
	board []int
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ int
	// MovesNum - 手数
	MovesNum int
	// Record - 棋譜
	Record []*RecordItem
	// 二重ループ
	iteratorWithoutWall func(func(int))
	// UCT計算中の子の数
	uctChildrenSize int
}

// NewPosition - 盤を作成します。
// TODO Position の NewPosition を呼び出す方法がない？
func NewPosition(board []int) *Position {
	var position = new(Position)
	position.board = board
	position.Record = make([]*RecordItem, MaxMovesNum)
	position.uctChildrenSize = BoardArea + 1
	position.iteratorWithoutWall = CreateBoardIteratorWithoutWall(position)

	checkBoard = make([]int, SentinelBoardArea)
	Dir4 = [4]int{1, SentinelWidth, -1, -SentinelWidth}

	return position
}

// InitPosition - 局面の初期化。
func (position *Position) InitPosition() {
	var boardMax = SentinelBoardArea

	// サイズが変わっているケースに対応
	position.iteratorWithoutWall = CreateBoardIteratorWithoutWall(position)

	// 枠線
	for z := 0; z < boardMax; z++ {
		position.SetColor(z, 3)
	}

	// 盤上
	var onPoint = func(z int) {
		position.SetColor(z, 0)
	}
	position.iteratorWithoutWall(onPoint)

	position.MovesNum = 0
	position.KoZ = 0
}

// ColorAt - 指定した交点の石の色
func (position *Position) ColorAt(z int) int {
	return position.board[z]
}

// ColorAtXy - 指定した交点の石の色
func (position *Position) ColorAtXy(x int, y int) int {
	return position.board[(y+1)*SentinelWidth+x+1]
}

// Exists - 指定の交点に石があるか？
func (position *Position) Exists(z int) bool {
	return position.board[z] != 0
}

// SetColor - 盤データ。
func (position *Position) SetColor(z int, color int) {
	position.board[z] = color
}

// CopyData - 盤データのコピー。
func (position *Position) CopyData() []int {
	boardArea := SentinelBoardArea

	var boardCopy2 = make([]int, boardArea)
	copy(boardCopy2[:], position.board[:])
	return boardCopy2
}

// ImportData - 盤データのコピー。
func (position *Position) ImportData(boardCopy2 []int) {
	copy(position.board[:], boardCopy2[:])
}

// GetZ4 - z（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func (position *Position) GetZ4(z int) int {
	if z == 0 {
		return 0
	}
	y := z / SentinelWidth
	x := z - y*SentinelWidth
	return x*100 + y
}

// GetZFromXy - x,y 形式の座標を、 z （配列のインデックス）へ変換します。
// x,y は壁を含まない領域での座標です。 z は壁を含む領域での座標です
func (position *Position) GetZFromXy(x int, y int) int {
	return (y+1)*SentinelWidth + x + 1
}

// GetEmptyZ - 空点の z （配列のインデックス）を返します。
func (position *Position) GetEmptyZ() int {
	var x, y, z int
	for {
		// ランダムに交点を選んで、空点を見つけるまで繰り返します。
		x = rand.Intn(9)
		y = rand.Intn(9)
		z = position.GetZFromXy(x, y)
		if !position.Exists(z) {
			break
		}
	}
	return z
}

func (position *Position) countLibertySub(z int, color int, pLiberty *int, pStone *int) {
	checkBoard[z] = 1
	*pStone++
	for i := 0; i < 4; i++ {
		z := z + Dir4[i]
		if checkBoard[z] != 0 {
			continue
		}
		if !position.Exists(z) {
			checkBoard[z] = 1
			*pLiberty++
		}
		if position.board[z] == color {
			position.countLibertySub(z, color, pLiberty, pStone)
		}
	}

}

// CountLiberty - 呼吸点を数えます。
func (position *Position) CountLiberty(z int, pLiberty *int, pStone *int) {
	*pLiberty = 0
	*pStone = 0
	boardMax := SentinelBoardArea
	// 初期化
	for z2 := 0; z2 < boardMax; z2++ {
		checkBoard[z2] = 0
	}
	position.countLibertySub(z, position.board[z], pLiberty, pStone)
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (position *Position) TakeStone(z int, color int) {
	position.board[z] = 0
	for dir := 0; dir < 4; dir++ {
		z2 := z + Dir4[dir]
		if position.board[z2] == color {
			position.TakeStone(z2, color)
		}
	}
}

// IterateWithoutWall - 盤イテレーター
func (position *Position) IterateWithoutWall(onPoint func(int)) {
	position.iteratorWithoutWall(onPoint)
}

// UctChildrenSize - UCTの最大手数
func (position *Position) UctChildrenSize() int {
	return position.uctChildrenSize
}

// CreateBoardIteratorWithoutWall - 盤の（壁を除く）全ての交点に順にアクセスする boardIterator 関数を生成します
func CreateBoardIteratorWithoutWall(
	position *Position) func(func(int)) {

	var boardSize = BoardSize
	var boardIterator = func(onPoint func(int)) {

		// x,y は壁無しの盤での0から始まる座標、 z は壁有りの盤での0から始まる座標
		for y := 0; y < boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				var z = position.GetZFromXy(x, y)
				onPoint(z)
			}
		}
	}

	return boardIterator
}
