package entities

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

func IsBestUpdateForPrimitiveMonteCalroV7(color int, bestValue float64, winRate float64) bool {
	var isBestUpdate = bestValue < winRate
	return isBestUpdate
}

func CreatePrintingOfInfoForPrimitiveMonteCalroIdling() func(color int, tryNum int, bestZ int, bestValue float64) {
	var printInfo = func(color int, tryNum int, bestZ int, bestValue float64) {
		// 何もしません
	}

	return printInfo
}
