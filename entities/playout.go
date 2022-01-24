package entities

import (
	"math/rand"
)

// Playout - 最後まで石を打ちます。得点を返します
// * `getWinner` - 地計算
//
// # Returns
//
// 勝者（黒番が1なのか、白番が1なのか、手番が1なのかは設定によって異なります）
func Playout(
	board IBoard,
	turnColor int,
	getWinner func(int) int) int {

	AllPlayouts++

	var color = turnColor
	var previousZ = 0
	var boardMax = board.SentinelBoardArea()

	var playoutTrialCount = PlayoutTrialCount
	for trial := 0; trial < playoutTrialCount; trial++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, z int

		var onPoint = func(z int) {
			if !board.Exists(z) {
				empty[emptyNum] = z
				emptyNum++
			}
		}

		board.IterateWithoutWall(onPoint)

		r = 0
		for {
			if emptyNum == 0 {
				z = 0
			} else {
				r = rand.Intn(emptyNum)
				z = empty[r]
			}

			var err = PutStone(board, z, color)

			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}

		// テストのときは棋譜を残します
		if FlagTestPlayout != 0 {
			Record[MovesNum].SetZ(z)
			MovesNum++
		}

		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z

		color = FlipColor(color)
	}

	return getWinner(turnColor)
}
