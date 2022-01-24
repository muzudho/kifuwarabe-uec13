package main

import (
	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-uec13/config_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
	p "github.com/muzudho/kifuwarabe-uec13/presenter"
)

// Lesson04 - レッスン４
func Lesson04() {
	code.Console.Trace("# GoGo Lesson04 プログラム開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/example-v3.gameConf.toml", OnFatal)

	var board = e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	var boardSize = board.BoardSize()
	if boardSize < 10 {
		// 10路盤より小さいとき
		e.PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		e.PlayoutTrialCount = boardSize * boardSize
	}

	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	var color = 1

	var printBoardDuringPlayout = e.CreatePrintingOfBoardDuringPlayout1(board, p.PrintBoard)
	e.Playout(board, color, printBoardDuringPlayout, e.GettingOfWinnerForPlayoutEverDraw)
}
