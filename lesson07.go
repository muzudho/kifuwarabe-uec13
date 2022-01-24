package main

import (
	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-uec13/config_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
	p "github.com/muzudho/kifuwarabe-uec13/presenter"
)

// Lesson07 - レッスン７
func Lesson07() {
	code.Console.Trace("# GoGo Lesson07 プログラム開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/example-v3.gameConf.toml", OnFatal)

	var board = e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	var boardSize = board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
		e.PrimitiveMonteCalroTrialCount = 30
	} else {
		e.PlayoutTrialCount = boardSize * boardSize
		e.PrimitiveMonteCalroTrialCount = 3
	}

	var exceptPutStoneLesson07 = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeOk)
	e.ExceptPutStoneOnPrimitiveMonteCalro = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	var printBoard = e.CreatePrintingOfBoardDuringPlayoutIdling()

	var color = 1
	for i := 0; i < 2; i++ {
		e.GettingOfWinnerOnDuringUCTPlayout = e.WrapGettingOfWinnerForPlayoutLesson07SelfView(board)
		z := e.PrimitiveMonteCalro(
			board,
			color,
			e.InitBestValueForPrimitiveMonteCalroV7,
			e.CreateCalcWinnerForPrimitiveMonteCalroV7(board),
			e.IsBestUpdateForPrimitiveMonteCalroV7,
			e.CreatePrintingOfInfoForPrimitiveMonteCalroV6(board),
			printBoard)

		e.PutStone(board, z, color, exceptPutStoneLesson07)

		p.PrintBoard(board, -1)

		color = e.FlipColor(color)
	}
}
