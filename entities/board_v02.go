package entities

// BoardV02 - 盤。
type BoardV02 struct {
	BoardV01
	uctChildrenSize int
}

// NewBoard - 盤を作成します。
// TODO BoardV01 の NewBoard を呼び出す方法がない？
func NewBoard(data []int, boardSize int, sentinelBoardArea int, komi float64, maxMoves int) *BoardV02 {
	var board = new(BoardV02)
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

// UctChildrenSize - UCTの最大手数
func (board BoardV02) UctChildrenSize() int {
	return board.uctChildrenSize
}
