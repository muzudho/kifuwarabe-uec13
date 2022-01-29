package main

import (
	"time"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-uec13/config_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
	pl "github.com/muzudho/kifuwarabe-uec13/play_algorithm"
	p "github.com/muzudho/kifuwarabe-uec13/presenter"
)

// SelfPlay - コンピューター同士の対局。
func SelfPlay() {
	code.Console.Trace("# GoGo SelfPlay 自己対局開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/game_conf.toml", OnFatal)
	e.Komi = config.Komi()
	e.MaxMovesNum = config.MaxMovesNum()
	e.SetBoardSize(config.BoardSize())
	var position = e.NewPosition(config.GetBoardArray())

	e.AdjustParameters(position) // パラーメーター調整

	var color = 1

	for {
		pl.GettingOfWinnerOnDuringUCTPlayout = pl.WrapGettingOfWinner(position)
		var z = GetComputerMoveDuringSelfPlay(position, color)

		var recItem = new(e.RecordItemV01)
		recItem.Z = z
		e.PutStoneOnRecord(position, z, color, recItem)
		p.PrintBoard(position, e.MovesNum)

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

	p.PrintSgf(position, e.MovesNum, e.Record)
}

// GetComputerMoveDuringSelfPlay - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveDuringSelfPlay(position *e.Position, color int) int {

	var z int
	var start = time.Now()
	pl.AllPlayouts = 0

	z = pl.GetBestZByUct(
		position,
		color,
		pl.WrapSearchUct(position))

	var sec = time.Since(start).Seconds()
	code.Console.Info("(GetComputerMoveDuringSelfPlay) %.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(pl.AllPlayouts)/sec, position.GetZ4(z), e.MovesNum, color, pl.AllPlayouts)
	return z
}
