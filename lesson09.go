package main

import (
	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-uec13/config_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
	p "github.com/muzudho/kifuwarabe-uec13/presenter"
)

// Lesson09 - レッスン９
func Lesson09() {
	code.Console.Trace("# GoGo Lesson09 プログラム開始☆（＾～＾）\n")
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

	e.ExceptPutStoneOnSearchUct = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
	e.ExceptPutStoneOnPrimitiveMonteCalro = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)
	e.ExceptPutStoneDuringPlayout = e.CreateExceptionForPutStoneLesson4(board, e.FillEyeErr)

	SelfplayLesson09(board, p.PrintBoard)
}

// SelfplayLesson09 - コンピューター同士の対局。
func SelfplayLesson09(board e.IBoardV02, printBoard func(e.IBoardV01, int)) {
	var color = 1

	for {
		e.GettingOfWinnerOnDuringUCTPlayout = e.WrapGettingOfWinner(board)
		var z = e.GetComputerMoveLesson09(board, color)

		var recItem = new(e.RecordItemV01)
		recItem.Z = z
		e.PutStoneOnRecord(board, z, color, recItem)
		printBoard(board, e.MovesNum)

		// パスで２手目以降で棋譜の１つ前（相手）もパスなら終了します。
		if z == 0 && 1 < e.MovesNum && e.Record[e.MovesNum-2].GetZ() == 0 {
			break
		}
		// 自己対局は300手で終了します。
		if 300 < e.MovesNum {
			break
		} // too long
		color = e.FlipColor(color)
	}

	p.PrintSgf(board, e.MovesNum, e.Record)
}
