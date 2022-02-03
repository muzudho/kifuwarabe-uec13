// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"flag"
	"math/rand"
	"time"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-uec13/config_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
	pl "github.com/muzudho/kifuwarabe-uec13/play_algorithm"
	p "github.com/muzudho/kifuwarabe-uec13/presenter"
)

func main() {
	flag.Parse()
	lessonVer := flag.Arg(0)

	// 乱数の種を設定
	rand.Seed(time.Now().UnixNano())

	// ログの書込み先設定
	code.GtpLog.SetPath("output/gtp_print.log")
	code.ConsoleLog.SetPath(
		"output/trace.log",
		"output/debug.log",
		"output/info.log",
		"output/notice.log",
		"output/warn.log",
		"output/error.log",
		"output/fatal.log",
		"output/print.log")

	code.Console.Trace("# Author: %s\n", e.Author)

	// 設定は囲碁GUIから与えられて上書きされる想定です。設定ファイルはデフォルト設定です
	var config = cnf.LoadGameConf("input/game_conf.toml", OnFatal)
	e.Komi = e.KomiType(config.Komi())
	e.MaxMovesNum = e.MovesNumType(config.MaxMovesNum())
	e.SetBoardSize(config.BoardSize())
	var position = e.NewPosition()
	pl.InitPosition(position)
	position.SetBoard(config.GetBoardArray())

	if lessonVer == "SelfPlay" {
		SelfPlay(position)
	} else {
		RunGtpEngine(position) // GTP
	}
}

func OnFatal(errorMessage string) {
	code.Console.Fatal(errorMessage)
}

func createPrintingOfCalc() *func(*e.Position, int, e.Point, float64, int) {
	// UCT計算中の表示
	var fn = func(position *e.Position, i int, z e.Point, rate float64, games int) {
		code.Console.Info("(UCT Calculating...) %2d:z=%s,rate=%.4f,games=%3d\n", i, p.GetGtpZ(position, z), rate, games)
	}

	return &fn
}

func createPrintingOfCalcFin() *func(*e.Position, e.Point, float64, int, int, int) {
	// UCT計算後の表示
	var fn = func(position *e.Position, bestZ e.Point, rate float64, max int, allPlayouts int, nodeNum int) {
		code.Console.Info("(UCT Calculated    ) bestZ=%s,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
			p.GetGtpZ(position, bestZ), rate, max, allPlayouts, nodeNum)

	}

	return &fn
}
