package entities

import (
	"time"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
)

// IBoardV02 - 盤。
type IBoardV02 interface {
	IBoardV01
	// UctChildrenSize - UCTの最大手数
	UctChildrenSize() int
}

// GetComputerMoveLesson09 - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveLesson09(board IBoardV02, color int) int {

	var z int
	var start = time.Now()
	AllPlayouts = 0

	z = GetBestZByUct(
		board,
		color,
		WrapSearchUct(board))

	var sec = time.Since(start).Seconds()
	code.Console.Info("(GetComputerMoveLesson09) %.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(AllPlayouts)/sec, board.GetZ4(z), MovesNum, color, AllPlayouts)
	return z
}
