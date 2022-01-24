package entities

import (
	"fmt"
	"math/rand"
)

// CreatePrintingOfBoardDuringPlayoutIdling - プレイアウト中の盤の描画（何も描画しません）
func CreatePrintingOfBoardDuringPlayoutIdling() func(int, int, int, int) {
	var printBoardDuringPlayout = func(trial int, z4 int, color int, emptyNum int) {
		// 何もしません
	}

	return printBoardDuringPlayout
}

// CreatePrintingOfBoardDuringPlayout1 - プレイアウト中の盤の描画
func CreatePrintingOfBoardDuringPlayout1(board IBoardV01, printBoard func(IBoardV01, int)) func(int, int, int, int) {
	var printBoardDuringPlayout = func(trial int, z int, color int, emptyNum int) {
		var z4 = board.GetZ4(z)     // XXYY
		var koZ4 = board.GetZ4(KoZ) // XXYY
		printBoard(board, -1)
		fmt.Printf("trial=%d,z4=%04d,clr=%d,emptyNum=%d,koZ4=%04d\n",
			trial, z4, color, emptyNum, koZ4)
	}

	return printBoardDuringPlayout
}

// Playout - 最後まで石を打ちます。得点を返します
// * `printBoardDuringPlayout` - プレイアウト中の盤の描画
// * `getWinner` - 地計算
//
// # Returns
//
// 勝者（黒番が1なのか、白番が1なのか、手番が1なのかは設定によって異なります）
func Playout(
	board IBoardV01,
	turnColor int,
	printBoardDuringPlayout func(int, int, int, int),
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

			var err = PutStone(board, z, color, ExceptPutStoneDuringPlayout)

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

		printBoardDuringPlayout(trial, z, color, emptyNum)

		color = FlipColor(color)
	}

	return getWinner(turnColor)
}
