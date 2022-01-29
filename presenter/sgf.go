package presenter

import (
	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// PrintSgf - SGF形式の棋譜表示。
func PrintSgf(board *e.Board, movesNum int, record []e.IRecordItemV01) {
	var boardSize = board.BoardSize()

	code.Console.Print("(;GM[1]SZ[%d]KM[%.1f]PB[]PW[]\n", boardSize, board.Komi())
	for i := 0; i < movesNum; i++ {
		var z = record[i].GetZ()
		var y = z / board.SentinelWidth()
		var x = z - y*board.SentinelWidth()
		var sStone = [2]string{"B", "W"}
		code.Console.Print(";%s", sStone[i&1])
		if z == 0 {
			code.Console.Print("[]")
		} else {
			code.Console.Print("[%c%c]", x+'a'-1, y+'a'-1)
		}
		if ((i + 1) % 10) == 0 {
			code.Console.Print("\n")
		}
	}
	code.Console.Print(")\n")
}
