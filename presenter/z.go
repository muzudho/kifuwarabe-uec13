package presenter

import (
	"fmt"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// GetGtpZ - XY座標をアルファベット、数字で表したもの。 例: Q10
func GetGtpZ(board e.IBoard, z int) string {
	if z == 0 {
		return "pass"
	}

	var y = z / board.SentinelWidth()
	var x = z % board.SentinelWidth()

	// 筋が25（'Z'）より大きくなることは想定していません
	var alphabet_x = 'A' + x - 1
	if alphabet_x >= 'I' {
		alphabet_x++
	}

	code.Console.Debug("y=%d x=%d z=%d alphabet_x=%d alphabet_x(char)=%c\n", y, x, z, alphabet_x, alphabet_x)

	return fmt.Sprintf("%c%d", alphabet_x, y)
}
