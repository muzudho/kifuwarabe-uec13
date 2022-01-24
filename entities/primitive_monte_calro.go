package entities

import (
	"fmt"
)

func InitBestValueForPrimitiveMonteCalroV6(color int) float64 {
	var bestValue float64

	if color == 1 {
		bestValue = -100.0
	} else {
		bestValue = 100.0
	}

	return bestValue
}

func InitBestValueForPrimitiveMonteCalroV7(color int) float64 {
	return -100.0
}

// CreateCalcWinnerForPrimitiveMonteCalroV6 - 盤を束縛変数として与えます
func CreateCalcWinnerForPrimitiveMonteCalroV6(board IBoardV01) func(int, func(int, int, int, int), func(int) int) int {
	var calcWinner = func(color int, printBoard func(int, int, int, int), getWinner func(int) int) int {
		return Playout(board, FlipColor(color), printBoard, getWinner)
	}

	return calcWinner
}

// CreateCalcWinnerForPrimitiveMonteCalroV7 - 盤を束縛変数として与えます
func CreateCalcWinnerForPrimitiveMonteCalroV7(board IBoardV01) func(int, func(int, int, int, int), func(int) int) int {
	var calcWinner = func(color int, printBoard func(int, int, int, int), getWinner func(int) int) int {
		return -Playout(board, FlipColor(color), printBoard, getWinner)
	}

	return calcWinner
}

func IsBestUpdateForPrimitiveMonteCalroV6(color int, bestValue float64, winRate float64) bool {
	var isBestUpdate = (color == 1 && bestValue < winRate) ||
		(color == 2 && winRate < bestValue)
	return isBestUpdate
}

func IsBestUpdateForPrimitiveMonteCalroV7(color int, bestValue float64, winRate float64) bool {
	var isBestUpdate = bestValue < winRate
	return isBestUpdate
}

// CreatePrintingOfInfoForPrimitiveMonteCalroV6 - 盤を束縛変数として与えます
func CreatePrintingOfInfoForPrimitiveMonteCalroV6(board IBoardV01) func(color int, tryNum int, bestZ int, bestValue float64) {
	var printInfo = func(color int, tryNum int, bestZ int, bestValue float64) {
		var bestZ4 = board.GetZ4(bestZ)
		fmt.Printf("(PrimitiveMonteCalro) bestZ4=%04d,color=%d,v=%5.3f,tryNum=%d\n", bestZ4, color, bestValue, tryNum)
	}

	return printInfo
}

func CreatePrintingOfInfoForPrimitiveMonteCalroIdling() func(color int, tryNum int, bestZ int, bestValue float64) {
	var printInfo = func(color int, tryNum int, bestZ int, bestValue float64) {
		// 何もしません
	}

	return printInfo
}

func PrimitiveMonteCalro(
	board IBoardV01,
	color int,
	initBestValue func(int) float64,
	calcWinner func(turnColor int, printBoard func(int, int, int, int), getWinner func(int) int) int,
	isBestUpdate func(color int, bestValue float64, winRate float64) bool,
	printInfo func(color int, trialCount int, bestZ int, bestValue float64),
	printBoard func(int, int, int, int)) int {

	// 一時記憶
	var copiedBoard1 = board.CopyData()
	var copiedKoZ1 = KoZ

	// 初期化
	var bestZ = 0
	var winRate float64

	var bestValue = initBestValue(color)

	var onPoint = func(z int) {
		if !board.Exists(z) {
			var err = PutStone(board, z, color, ExceptPutStoneOnPrimitiveMonteCalro)

			if err == 0 {
				var winSum = 0
				for i := 0; i < PrimitiveMonteCalroTrialCount; i++ {
					var boardCopy2 = board.CopyData()
					var koZCopy2 = KoZ

					// 手番の勝ちが1、引分けが0、相手の勝ちが-1 としてください
					var winner = calcWinner(FlipColor(color), printBoard, GettingOfWinnerOnDuringUCTPlayout)

					winSum += winner
					KoZ = koZCopy2
					board.ImportData(boardCopy2)
				}

				winRate = float64(winSum) / float64(PrimitiveMonteCalroTrialCount)
				if isBestUpdate(color, bestValue, winRate) {
					bestValue = winRate
					bestZ = z
					printInfo(color, PrimitiveMonteCalroTrialCount, bestZ, bestValue)
				}
			}
		}
	}

	board.IterateWithoutWall(onPoint)

	// 復元
	KoZ = copiedKoZ1
	board.ImportData(copiedBoard1)

	return bestZ
}
