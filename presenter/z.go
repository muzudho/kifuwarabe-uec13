package presenter

import (
	"fmt"
	"strings"

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

// GetZFromGtp - GTPの座標符号を z に変換します
// * `gtp_z` - 最初の１文字はアルファベット、２文字目（あれば３文字目）は数字と想定。 例: q10
func GetZFromGtp(board e.IBoard, gtp_z string) int {
	gtp_z = strings.ToLower(gtp_z)
	code.Console.Trace("# gtp_z=%s\n", gtp_z)

	if gtp_z == "pass" {
		return 0
	}

	// 筋
	var x = gtp_z[0] - 'a' + 1
	if gtp_z[0] >= 'i' {
		x--
	}

	// 段
	var y = int(gtp_z[1] - '0')
	if 2 < len(gtp_z) {
		y *= 10
		y += int(gtp_z[2] - '0')
	}

	// インデックス
	var z = board.GetZFromXy(int(x)-1, y-1)
	code.Console.Trace("# x=%d y=%d z=%d z4=%04d\n", x, y, z, board.GetZ4(z))
	return z
}
