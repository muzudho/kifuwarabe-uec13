// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"flag"
	"math/rand"
	"time"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
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

	if lessonVer == "SelfPlay" {
		SelfPlay()
	} else {
		RunGtpEngine() // GTP
	}
}

func OnFatal(errorMessage string) {
	code.Console.Fatal(errorMessage)
}

func createPrintingOfCalc() *func(*e.Position, int, int, float64, int) {
	// UCT計算中の表示
	var fn = func(position *e.Position, i int, z int, rate float64, games int) {
		code.Console.Info("(UCT Calculating...) %2d:z=%s,rate=%.4f,games=%3d\n", i, p.GetGtpZ(position, z), rate, games)
	}

	return &fn
}

func createPrintingOfCalcFin() *func(*e.Position, int, float64, int, int, int) {
	// UCT計算後の表示
	var fn = func(position *e.Position, bestZ int, rate float64, max int, allPlayouts int, nodeNum int) {
		code.Console.Info("(UCT Calculated    ) bestZ=%s,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
			p.GetGtpZ(position, bestZ), rate, max, allPlayouts, nodeNum)

	}

	return &fn
}
