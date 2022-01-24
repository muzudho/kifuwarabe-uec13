package main

import (
	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-uec13/config_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
	p "github.com/muzudho/kifuwarabe-uec13/presenter"
)

// Lesson01 - レッスン１
func Lesson01() {
	code.Console.Trace("# GoGo Lesson01 プログラム開始☆（＾～＾）\n")
	var config = cnf.LoadGameConf("input/example-v1.gameConf.toml", OnFatal)

	code.Console.Trace("# Config読んだ☆（＾～＾）\n")
	code.Console.Trace("# Server=%s\n", config.Nngs.Server)
	code.Console.Trace("# Port=%d\n", config.Nngs.Port)
	code.Console.Trace("# User=%s\n", config.Nngs.User)
	code.Console.Trace("# Pass=%s\n", config.Nngs.Pass)
	code.Console.Trace("# Komi=%f\n", config.Game.Komi)
	code.Console.Trace("# BoardSize=%d\n", config.Game.BoardSize)
	code.Console.Trace("# MaxMoves=%d\n", config.Game.MaxMoves)
	code.Console.Trace("# BoardData=%s\n", config.Game.BoardData)
	code.Console.Trace("# SentinelBoardArea()=%d\n", config.SentinelBoardArea())

	var board = e.NewBoard( // 盤生成
		config.GetBoardArray(),     // 配列
		config.BoardSize(),         // 盤サイズ
		config.SentinelBoardArea(), // 番兵
		config.Komi(),              // コミ
		config.MaxMovesNum())       // 上限手数
	/*
		fmt.Println("board.BoardSize()=", board.BoardSize())
		fmt.Println("board.SentinelBoardArea()=", board.SentinelBoardArea())
		fmt.Println("board.GetData()=", board.GetData())
	*/

	p.PrintBoard(board, -1)
}
