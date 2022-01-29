package main

import (
	"bufio"
	"os"
	"strings"
	"time"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-uec13/config_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
	pl "github.com/muzudho/kifuwarabe-uec13/play_algorithm"
	p "github.com/muzudho/kifuwarabe-uec13/presenter"
)

// RunGtpEngine - レッスン９a
// GTP2NNGS に対応しているのでは？
func RunGtpEngine() {
	code.Console.Trace("# GoGo RunGtpEngine プログラム開始☆（＾～＾）\n")

	// 設定は囲碁GUIから与えられて上書きされる想定です。設定ファイルはデフォルト設定です
	var config = cnf.LoadGameConf("input/game_conf.toml", OnFatal)
	var board = e.NewBoard(config.GetBoardArray(), config.BoardSize(), config.SentinelBoardArea(), config.Komi(), config.MaxMovesNum())
	board.InitBoard()

	e.AdjustParameters(board) // パラーメーター調整

	code.Console.Trace("# 何か標準入力しろだぜ☆（＾～＾）\n")

	// GUI から 囲碁エンジン へ入力があった、と考えてください
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var command = scanner.Text()
		code.Gtp.Log(command + "\n")
		code.ConsoleLog.Notice(command + "\n")

		var tokens = strings.Split(command, " ")
		switch tokens[0] {
		case "boardsize":
			// TODO 盤のサイズを変えたい
			// boardsize 19

			e.AdjustParameters(board) // パラーメーター再調整

			code.Gtp.Print("= \n\n")

		case "clear_board":
			board.InitBoard()
			code.Gtp.Print("= \n\n")

		case "quit":
			os.Exit(0)

		case "protocol_version":
			code.Gtp.Print("= 2\n\n")

		case "name":
			code.Gtp.Print("= KifuwarabeUEC13\n\n")

		case "version":
			code.Gtp.Print("= 0.0.1\n\n")

		case "list_commands":
			code.Gtp.Print("= boardsize\nclear_board\nquit\nprotocol_version\nundo\n" +
				"name\nversion\nlist_commands\nkomi\ngenmove\nplay\n\n")

		case "komi":
			code.Gtp.Print("= 6.5\n\n")

		case "undo":
			// 未実装
			code.Gtp.Print("= \n\n")

		case "genmove":
			// genmove black
			// genmove white
			var color int
			if 1 < len(tokens) && strings.ToLower(tokens[1][0:1]) == "w" {
				color = 2
			} else {
				color = 1
			}
			var z = PlayComputerMoveLesson09a(board, color)
			code.Gtp.Print("= %s\n\n", p.GetGtpZ(board, z))

		case "play":
			// play black A3
			// play white D4
			// play black D5
			// play white E5
			// play black E4
			// play white D6
			// play black F5
			// play white C5
			// play black PASS
			// play white PASS
			if 2 < len(tokens) {
				var color int
				if strings.ToLower(tokens[1][0:1]) == "w" {
					color = 2
				} else {
					color = 1
				}

				var z = p.GetZFromGtp(board, tokens[2])
				var recItem = new(e.RecordItemV02)
				recItem.Z = z
				recItem.Time = 0
				e.PutStoneOnRecord(board, z, color, recItem)
				p.PrintBoard(board, e.MovesNum)

				code.Gtp.Print("= \n\n")
			}

		default:
			code.Gtp.Print("? unknown_command\n\n")
		}
	}
}

// PlayComputerMoveLesson09a - コンピューター・プレイヤーの指し手。 SelfPlay, RunGtpEngine から呼び出されます。
func PlayComputerMoveLesson09a(
	board e.IBoard,
	color int) int {

	pl.GettingOfWinnerOnDuringUCTPlayout = pl.WrapGettingOfWinner(board)

	e.AdjustParameters(board)

	var z int
	var st = time.Now()
	pl.AllPlayouts = 0

	z = pl.GetBestZByUct(
		board,
		color,
		pl.WrapSearchUct(board))

	var sec = time.Since(st).Seconds()
	code.Console.Info("%.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(pl.AllPlayouts)/sec, board.GetZ4(z), e.MovesNum, color, pl.AllPlayouts)

	var recItem = new(e.RecordItemV02)
	recItem.Z = z
	recItem.Time = sec
	e.PutStoneOnRecord(board, z, color, recItem)
	p.PrintBoard(board, e.MovesNum)

	return z
}
