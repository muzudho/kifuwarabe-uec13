package presenter

import (
	"strconv"
	"strings"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

var sz8k = 8 * 1024

// 案
//     A B C D E F G H J K L M N O P Q R S T
//   +---------------------------------------+
//  1| . . . . . . . . . . . . . . . . . . . |
//  2| . . . . . . . . . . . . . . . . . . . |
//  3| . . . . . . . . . . . . . . . . x . . |
//  4| . . . . . . . . . . . . . . . . . . . |
//  5| . . . . . . . . . . . . . . . . . . . |
//  6| . . . . . . . . . . . . . . . . . . . |
//  7| . . . . . . . . . . . . . . . . . . . |
//  8| . . . . . . . . . . . . . . . . . . . |
//  9| . . . . . . . . . . . . . . . . . . . |
// 10| . . . . . . . . . . . . . . . . . . . |
// 11| . . . . . . . . . . . . . . . . . . . |
// 12| . . . . . . . . . . . . . . . . . . . |
// 13| . . . . . . . . . . . . . . . . . . . |
// 14| . . . . . . . . . . . . . . . . . . . |
// 15| . . . . . . . . . . . . . . . . . . . |
// 16| . . . . . . . . . . . . . . . . . . . |
// 17| . . o . . . . . . . . . . . . . . . . |
// 18| . . . . . . . . . . . . . . . . . . . |
// 19| . . . . . . . . . . . . . . . . . . . |
//   +---------------------------------------+
//  KoZ=0,movesNum=999
//
// ASCII文字を使います（全角、半角の狂いがないため）
// 黒石は x 、 白石は o （ダークモードでもライトモードでも識別できるため）

// labelOfColumns - 各列の表示符号。
// 国際囲碁連盟のフォーマット
var labelOfColumns = [20]string{"xx", " A", " B", " C", " D", " E", " F", " G", " H", " J",
	" K", " L", " M", " N", " O", " P", " Q", " R", " S", " T"}

// labelOfRows - 各行の表示符号。
var labelOfRows = [20]string{" 0", " 1", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9",
	"10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}

// " ." - 空点
// " x" - 黒石
// " o" - 白石
// " #" - 壁（使いません）
var stoneLabels = [4]string{" .", " x", " o", " #"}

// PrintBoard - 盤を描画。
func PrintBoard(position *e.Position, movesNum int) {

	var b = &strings.Builder{}
	b.Grow(sz8k)

	var boardSize = e.BoardSize

	// Header
	b.WriteString("\n   ")
	for x := 0; x < boardSize; x++ {
		b.WriteString(labelOfColumns[x+1])
	}
	b.WriteString("\n  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	// Body
	for y := 0; y < boardSize; y++ {
		b.WriteString(labelOfRows[y+1])
		b.WriteString("|")
		for x := 0; x < boardSize; x++ {
			b.WriteString(stoneLabels[position.ColorAtXy(x, y)])
		}
		b.WriteString(" |\n")
	}

	// Footer
	b.WriteString("  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	// Info
	b.WriteString("  KoZ=")
	if position.KoZ == e.Pass {
		b.WriteString("_")
	} else {
		b.WriteString(GetGtpZ(position, position.KoZ))
		// b.WriteString(strconv.Itoa(position.GetZ4(position.KoZ)))
	}
	if movesNum != -1 {
		b.WriteString(",movesNum=")
		b.WriteString(strconv.Itoa(movesNum))
	}
	b.WriteString("\n")

	code.Console.Print(b.String())
}
