package presenter

import (
	"fmt"

	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// PrintSgf - SGF形式の棋譜表示。
func PrintSgf(board e.IBoardV01, movesNum int, record []e.IRecordItemV01) {
	var boardSize = board.BoardSize()

	fmt.Printf("(;GM[1]SZ[%d]KM[%.1f]PB[]PW[]\n", boardSize, board.Komi())
	for i := 0; i < movesNum; i++ {
		var z = record[i].GetZ()
		var y = z / board.SentinelWidth()
		var x = z - y*board.SentinelWidth()
		var sStone = [2]string{"B", "W"}
		fmt.Printf(";%s", sStone[i&1])
		if z == 0 {
			fmt.Printf("[]")
		} else {
			fmt.Printf("[%c%c]", x+'a'-1, y+'a'-1)
		}
		if ((i + 1) % 10) == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf(")\n")
}
