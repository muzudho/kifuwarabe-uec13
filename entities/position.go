package entities

import (
	"math/rand"
)

// Position - 盤
type Position struct {
	// 盤
	board []Stone
	// 呼吸点を数えるための一時盤
	checkBoard []int
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ Point
	// MovesNum - 手数
	MovesNum int
	// Record - 棋譜
	Record []*RecordItem
	// 二重ループ
	iteratorWithoutWall func(func(Point))
	// UCT計算中の子の数
	uctChildrenSize int
}

// TemporaryPosition - 盤をコピーするときの一時メモリーとして使います
type TemporaryPosition struct {
	// 盤
	Board []Stone
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ Point
}

// CopyPosition - 盤データのコピー。
func (position *Position) CopyPosition() *TemporaryPosition {
	var temp = new(TemporaryPosition)
	temp.Board = make([]Stone, SentinelBoardArea)
	copy(temp.Board[:], position.board[:])
	temp.KoZ = position.KoZ
	return temp
}

// ImportPosition - 盤データのコピー。
func (position *Position) ImportPosition(temp *TemporaryPosition) {
	copy(position.board[:], temp.Board[:])
	position.KoZ = temp.KoZ
}

// NewPosition - 空っぽの局面を生成します
// あとで InitPosition() を呼び出してください
func NewPosition() *Position {
	return new(Position)
}

// InitPosition - 局面の初期化。
func (position *Position) InitPosition() {
	position.Record = make([]*RecordItem, MaxMovesNum)
	position.uctChildrenSize = BoardArea + 1

	// サイズが変わっているケースに対応するため、配列の作り直し
	var boardMax = SentinelBoardArea
	position.board = make([]Stone, boardMax)
	position.checkBoard = make([]int, boardMax)
	position.iteratorWithoutWall = CreateBoardIteratorWithoutWall(position)
	Dir4 = [4]Point{1, Point(-SentinelWidth), -1, Point(SentinelWidth)}

	// 枠線
	for z := Point(0); z < Point(boardMax); z++ {
		position.SetColor(z, Wall)
	}

	// 盤上
	var onPoint = func(z Point) {
		position.SetColor(z, 0)
	}
	position.iteratorWithoutWall(onPoint)

	position.MovesNum = 0
	position.KoZ = 0 // コウの指定がないので消します
}

// SetBoard - 盤面を設定します
func (position *Position) SetBoard(board []Stone) {
	// TODO 消す
	// fmt.Print("[[")
	// for z := 0; z < SentinelBoardArea; z++ {
	// 	fmt.Printf("%d,", board[z])
	// 	position.SetColor(Point(z), board[z])
	// }
	// fmt.Print("]]")
	position.board = board
}

// ColorAt - 指定した交点の石の色
func (position *Position) ColorAt(z Point) Stone {
	return position.board[z]
}

// CheckAt - 指定した交点のチェック
func (position *Position) CheckAt(z Point) int {
	return position.checkBoard[z]
}

// ColorAtXy - 指定した交点の石の色
func (position *Position) ColorAtXy(x int, y int) Stone {
	return position.board[(y+1)*SentinelWidth+x+1]
}

// IsEmpty - 指定の交点は空点か？
func (position *Position) IsEmpty(z Point) bool {
	return position.board[z] == Empty
}

// SetColor - 盤データ。
func (position *Position) SetColor(z Point, color Stone) {
	// TODO 消す
	// if z == Point(11) && color == Empty { // テスト
	// 	panic(fmt.Sprintf("z=%d color=%d SentinelWidth=%d", z, color, SentinelWidth))
	// }

	position.board[z] = color
}

// GetZ4 - z（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func (position *Position) GetZ4(z Point) int {
	if z == 0 {
		return 0
	}
	var y = int(z) / SentinelWidth
	var x = int(z) - y*SentinelWidth
	return x*100 + y
}

// GetZFromXy - x,y 形式の座標を、 z （配列のインデックス）へ変換します。
// x,y は壁を含まない領域での 0 から始まる座標です。 z は壁を含む盤上での座標です
func (position *Position) GetZFromXy(x int, y int) Point {
	return Point((y+1)*SentinelWidth + x + 1)
}

// GetEmptyZ - 空点の z （配列のインデックス）を返します。
func (position *Position) GetEmptyZ() Point {
	var x, y int
	var z Point
	for {
		// ランダムに交点を選んで、空点を見つけるまで繰り返します。
		x = rand.Intn(9)
		y = rand.Intn(9)
		z = position.GetZFromXy(x, y)
		if position.IsEmpty(z) { // 空点
			break
		}
	}
	return z
}

// CountLiberty - 呼吸点を数えます。
// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (position *Position) CountLiberty(z Point, libertyArea *int, renArea *int) {
	*libertyArea = 0
	*renArea = 0

	// チェックボードの初期化
	var onPoint = func(z Point) {
		position.checkBoard[z] = 0
	}
	position.iteratorWithoutWall(onPoint)

	position.countLibertySub(z, position.board[z], libertyArea, renArea)
}

// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (position *Position) countLibertySub(z Point, color Stone, libertyArea *int, renArea *int) {
	position.checkBoard[z] = 1
	*renArea++
	for i := 0; i < 4; i++ {
		var adjZ = z + Dir4[i]
		if position.checkBoard[adjZ] != 0 {
			continue
		}
		if position.IsEmpty(adjZ) { // 空点
			position.checkBoard[adjZ] = 1
			*libertyArea++
		} else if position.board[adjZ] == color {
			position.countLibertySub(adjZ, color, libertyArea, renArea) // 再帰
		}
	}
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (position *Position) TakeStone(z Point, color Stone) {
	position.board[z] = Empty // 石を消します

	for dir := 0; dir < 4; dir++ {
		var adjZ = z + Dir4[dir]

		if position.board[adjZ] == color { // 再帰します
			position.TakeStone(adjZ, color)
		}
	}
}

// IterateWithoutWall - 盤イテレーター
func (position *Position) IterateWithoutWall(onPoint func(Point)) {
	position.iteratorWithoutWall(onPoint)
}

// UctChildrenSize - UCTの最大手数
func (position *Position) UctChildrenSize() int {
	return position.uctChildrenSize
}

// CreateBoardIteratorWithoutWall - 盤の（壁を除く）全ての交点に順にアクセスする boardIterator 関数を生成します
func CreateBoardIteratorWithoutWall(position *Position) func(func(Point)) {

	var boardSize = BoardSize
	var boardIterator = func(onPoint func(Point)) {

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
