package main

import (
	"time"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
	pl "github.com/muzudho/kifuwarabe-uec13/play_algorithm"
	p "github.com/muzudho/kifuwarabe-uec13/presenter"
)

// SelfPlay - コンピューター同士の対局。
func SelfPlay(position *e.Position) {
	code.Console.Trace("# GoGo SelfPlay 自己対局開始☆（＾～＾）\n")

	var color = 1

	for {
		var z = GetComputerMoveDuringSelfPlay(position, color)

		var recItem = new(e.RecordItem)
		recItem.Z = z
		e.PutStoneOnRecord(position, z, color, recItem)
		p.PrintBoard(position, position.MovesNum)

		// パスで２手目以降で棋譜の１つ前（相手）もパスなら終了します。
		if z == 0 && 1 < position.MovesNum && position.Record[position.MovesNum-2].GetZ() == 0 {
			break
		}
		// 自己対局は400手で終了します。
		if 400 < position.MovesNum {
			break
		} // too long
		color = e.FlipColor(color)
	}

	p.PrintSgf(position, position.MovesNum, position.Record)
}

// GetComputerMoveDuringSelfPlay - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveDuringSelfPlay(position *e.Position, color int) int {

	var start = time.Now()
	pl.AllPlayouts = 0

	var z, winRate = pl.GetBestZByUct(
		position,
		color,
		pl.SearchingOfUct,
		createPrintingOfCalc(),
		createPrintingOfCalcFin())

	var sec = time.Since(start).Seconds()
	code.Console.Info("(GetComputerMoveDuringSelfPlay) %.1f sec, %.0f playout/sec, play_z=%04d,rate=%.4f,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(pl.AllPlayouts)/sec, position.GetZ4(z), winRate, position.MovesNum, color, pl.AllPlayouts)
	return z
}
