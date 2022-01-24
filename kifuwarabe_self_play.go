package main

import (
	"time"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-uec13/config_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
	p "github.com/muzudho/kifuwarabe-uec13/presenter"
)

// SelfPlay - コンピューター同士の対局。
func SelfPlay() {
	code.Console.Trace("# GoGo SelfPlay 自己対局開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/game_conf.toml", OnFatal)
	var board = e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())

	e.AdjustParameters(board) // パラーメーター調整

	var color = 1

	for {
		e.GettingOfWinnerOnDuringUCTPlayout = e.WrapGettingOfWinner(board)
		var z = GetComputerMoveDuringSelfPlay(board, color)

		var recItem = new(e.RecordItemV01)
		recItem.Z = z
		e.PutStoneOnRecord(board, z, color, recItem)
		p.PrintBoard(board, e.MovesNum)

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

// GetComputerMoveDuringSelfPlay - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveDuringSelfPlay(board e.IBoard, color int) int {

	var z int
	var start = time.Now()
	e.AllPlayouts = 0

	z = e.GetBestZByUct(
		board,
		color,
		e.WrapSearchUct(board))

	var sec = time.Since(start).Seconds()
	code.Console.Info("(GetComputerMoveDuringSelfPlay) %.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(e.AllPlayouts)/sec, board.GetZ4(z), e.MovesNum, color, e.AllPlayouts)
	return z
}
