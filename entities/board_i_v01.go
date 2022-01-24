package entities

// IBoardV01 - 盤。
type IBoardV01 interface {
	// 指定した交点の石の色
	ColorAt(z int) int
	ColorAtXy(x int, y int) int
	SetColor(z int, color int)

	CopyData() []int
	ImportData(boardCopy2 []int)
	Exists(z int) bool

	CountLiberty(z int, pLiberty *int, pStone *int)
	TakeStone(z int, color int)
	GetEmptyZ() int

	// n路盤
	BoardSize() int

	// SentinelWidth - 枠付きの盤の一辺の交点数
	SentinelWidth() int
	SentinelBoardArea() int
	// 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
	Komi() float64
	MaxMovesNum() int
	// GetZFromXy - YX形式の座標を、z（配列のインデックス）へ変換します。
	GetZFromXy(x int, y int) int
	// GetZ4 - z（配列のインデックス）を XXYY形式へ変換します。
	GetZ4(z int) int

	IterateWithoutWall(func(int))
}

// CreateBoardIteratorWithoutWall - 盤の（壁を除く）全ての交点に順にアクセスする boardIterator 関数を生成します
func CreateBoardIteratorWithoutWall(
	board IBoardV01) func(func(int)) {

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
