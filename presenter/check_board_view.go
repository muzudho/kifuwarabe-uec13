package presenter

import (
	"strings"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// " 0" - 空点
// " 1" - 黒石
var numberLabels = [2]string{" 0", " 1"}

// PrintCheckBoard - チェックボードを描画。
func PrintCheckBoard(position *e.Position) {

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
			var z = position.GetZFromXy(x, y)
			var number = position.CheckAt(z)
			b.WriteString(numberLabels[number])
		}
		b.WriteString(" |\n")
	}

	// Footer
	b.WriteString("  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	code.Console.Print(b.String())
}
