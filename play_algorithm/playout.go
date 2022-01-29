package play_algorithm

import (
	"math/rand"

	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// Playout - 最後まで石を打ちます。得点を返します
// * `getWinner` - 地計算
//
// # Returns
//
// 手番が勝ったら 1、引分けなら 0、 相手が勝ったら -1
func Playout(
	board *e.Board,
	turnColor int,
	getWinner func(int) int) int {

	AllPlayouts++

	var color = turnColor
	var previousZ = 0
	var boardMax = board.SentinelBoardArea()

	var playoutTrialCount = e.PlayoutTrialCount
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

			var err = e.PutStone(board, z, color)

			if err == 0 {
				break
			}
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}

		// テストのときは棋譜を残します
		if FlagTestPlayout != 0 {
			e.Record[e.MovesNum].SetZ(z)
			e.MovesNum++
		}

		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z

		color = e.FlipColor(color)
	}

	return getWinner(turnColor)
}
