// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"flag"
	"math/rand"
	"time"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

func main() {
	flag.Parse()
	lessonVer := flag.Arg(0)

	// 乱数の種を設定
	rand.Seed(time.Now().UnixNano())

	// ロガーの作成。
	code.Log.SetPath(
		"output/trace.log",
		"output/debug.log",
		"output/info.log",
		"output/notice.log",
		"output/warn.log",
		"output/error.log",
		"output/fatal.log",
		"output/print.log")

	// 標準出力への表示と、ログへの書き込みを同時に行います。
	code.Console.Trace("Author: %s\n", e.Author)

	if lessonVer == "Lesson01" {
		Lesson01()
	} else if lessonVer == "Lesson02" {
		Lesson02()
	} else if lessonVer == "Lesson03" {
		Lesson03()
	} else if lessonVer == "Lesson04" {
		Lesson04()
	} else if lessonVer == "Lesson05" {
		Lesson05()
	} else if lessonVer == "Lesson06" {
		Lesson06()
	} else if lessonVer == "Lesson07" {
		Lesson07()
	} else if lessonVer == "Lesson08" {
		Lesson08()
	} else if lessonVer == "Lesson09" {
		Lesson09()
	} else if lessonVer == "Lesson09a" {
		Lesson09a() // GTP
	} else {
		Lesson09a() // GTP
	}
	//KifuwarabeV1()
}

func OnFatal(errorMessage string) {
	code.Console.Fatal(errorMessage)
}
