package presenter

import (
	"fmt"

	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// GetCharZ - YX座標の文字表示？
func GetCharZ(board e.IBoard, z int) string {
	if z == 0 {
		return "pass"
	}

	var boardSize = board.BoardSize()

	var y = z / board.SentinelWidth()
	var x = z - y*board.SentinelWidth()
	var ax = 'A' + x - 1
	if ax >= 'I' {
		ax++
	}

	//return string(ax) + string(BoardSize+1-y+'0')
	return fmt.Sprintf("%d%d", ax, boardSize+1-y+'0')
}
