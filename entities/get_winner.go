package entities

import (
	"fmt"
)

// GettingOfWinnerForPlayoutEverDraw - Lesson04以前に使用。「常に引分け（0）を返す関数」を作成します。つまり勝者判定を行いません
func GettingOfWinnerForPlayoutEverDraw(int) int {
	return 0
}

// WrapGettingOfWinnerV05BlackSideView - Lesson05で使用。盤を束縛変数として与えます
func WrapGettingOfWinnerV05BlackSideView(board IBoardV01) func(colorIsNotUsed int) int {
	// 「黒勝ちなら1、引き分け、または白勝ちなら0を返す関数」を作成します
	// * `colorIsNotUsed` - 使っていません
	var getWinner = func(colorIsNotUsed int) int {
		return GetWinnerV05BlackSideView(board, colorIsNotUsed)
	}

	return getWinner
}

// WrapGettingOfWinnerForPlayoutLesson06BlackSideView - Lesson06で使用
func WrapGettingOfWinnerForPlayoutLesson06BlackSideView(board IBoardV01) func(colorIsNotUsed int) int {
	// 「黒勝ちなら1、引き分け、または白勝ちなら0を返す関数（黒側の視点）」を作成します
	// * `colorIsNotUsed` - 使っていません
	var getWinner = func(colorIsNotUsed int) int {
		return GetWinnerV06BlackSideView(board, colorIsNotUsed)
	}

	return getWinner
}

// WrapGettingOfWinnerForPlayoutLesson07SelfView
func WrapGettingOfWinnerForPlayoutLesson07SelfView(board IBoardV01) func(turnColor int) int {
	// 「手番の勝ちなら1、引き分けなら0、手番の負けなら-1を返す関数（自分視点）」を作成します
	// * `turnColor` - 手番の石の色
	var getWinner = func(turnColor int) int {
		return GetWinnerV07SelfView(board, turnColor)
	}

	return getWinner
}

// WrapGetWinnerV06BlackSideView - 盤を束縛変数として与えます
func WrapGetWinnerV06BlackSideView(board IBoardV01) func(colorIsNotUsed int) int {
	var getWinner = func(colorIsNotUsed int) int {
		return GetWinnerV06BlackSideView(board, colorIsNotUsed)
	}

	return getWinner
}

// WrapGetWinnerV07SelfView - 盤を束縛変数として与えます
func WrapGetWinnerV07SelfView(board IBoardV01) func(colorIsNotUsed int) int {
	var getWinner = func(colorIsNotUsed int) int {
		return GetWinnerV07SelfView(board, colorIsNotUsed)
	}

	return getWinner
}

// GetWinnerV05BlackSideView - Lesson05。 黒勝ちなら1、引き分け、または白勝ちなら0（黒側の視点）
// * `colorIsNotUsed` - 使っていません
func GetWinnerV05BlackSideView(board IBoardV01, colorIsNotUsed int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int

	var onPoint = func(z int) {
		var color2 = board.ColorAt(z)
		kind[color2]++
		if color2 == 0 {
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(z+Dir4[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}

	board.IterateWithoutWall(onPoint)

	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	var win = 0
	if 0 < float64(score)-board.Komi() { // float32 → float64
		win = 1
	}
	fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

// GetWinnerV06BlackSideView - Lesson06。 黒勝ちなら1、引き分け、または白勝ちなら0（黒側の視点）
// * `colorIsNotUsed` - 使っていません
func GetWinnerV06BlackSideView(board IBoardV01, colorIsNotUsed int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int

	var onPoint = func(z int) {
		color2 := board.ColorAt(z)
		kind[color2]++
		if color2 == 0 {
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(z+Dir4[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}

	board.IterateWithoutWall(onPoint)

	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	var win = 0
	if 0 < float64(score)-board.Komi() { // float32 → float64
		win = 1
	}
	// fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	// fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	// fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}

// GetWinnerV07SelfView - 手番の勝ちなら1、引き分けなら0、手番の負けなら-1（自分視点）
// * `turnColor` - 手番の石の色
func GetWinnerV07SelfView(board IBoardV01, turnColor int) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int

	var onPoint = func(z int) {
		var color2 = board.ColorAt(z)
		kind[color2]++
		if color2 == 0 {
			mk[1] = 0
			mk[2] = 0
			for dir := 0; dir < 4; dir++ {
				mk[board.ColorAt(z+Dir4[dir])]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}

	board.IterateWithoutWall(onPoint)

	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	var win = 0
	if 0 < float64(score)-board.Komi() {
		win = 1
	}
	if turnColor == 2 {
		win = -win
	} // gogo07

	// fmt.Printf("blackSum=%2d, (stones=%2d, area=%2d)\n", blackSum, kind[1], blackArea)
	// fmt.Printf("whiteSum=%2d, (stones=%2d, area=%2d)\n", whiteSum, kind[2], whiteArea)
	// fmt.Printf("score=%d, win=%d\n", score, win)
	return win
}
