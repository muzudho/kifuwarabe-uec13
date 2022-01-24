package main

import (
	"bufio"
	"os"
	"strings"
	"time"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-uec13/config_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
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
			code.Gtp.Print("= GoGo\n\n")

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
			// genmove b

			var color int
			if 1 < len(tokens) && strings.ToLower(tokens[1]) == "w" {
				color = 2
			} else {
				color = 1
			}
			var z = PlayComputerMoveLesson09a(board, color, p.PrintBoard)
			code.Gtp.Print("= %s\n\n", p.GetGtpZ(board, z))

		case "play":
			// play b a3
			// play w d4
			// play b d5
			// play w e5
			// play b e4
			// play w d6
			// play b f5
			// play w c5
			// play b pass
			// play w pass

			var color = 1
			if 1 < len(tokens) && strings.ToLower(tokens[1]) == "w" {
				color = 2
			} else {
				color = 1
			}

			if 2 < len(tokens) {
				var ax = strings.ToLower(tokens[2])
				code.Console.Trace("ax=%s\n", ax)
				var x = ax[0] - 'a' + 1
				if ax[0] >= 'i' {
					x--
				}
				var y = int(ax[1] - '0')
				var z = board.GetZFromXy(int(x)-1, board.BoardSize()-y)
				code.Console.Trace("x=%d y=%d z=%04d\n", x, y, board.GetZ4(z))
				if ax == "pass" {
					z = 0
				}

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
	color int,
	printBoardOutOfPlayout func(e.IBoard, int)) int {

	e.GettingOfWinnerOnDuringUCTPlayout = e.WrapGettingOfWinner(board)

	e.AdjustParameters(board)

	var z int
	var st = time.Now()
	e.AllPlayouts = 0

	z = e.GetBestZByUct(
		board,
		color,
		e.WrapSearchUct(board))

	var sec = time.Since(st).Seconds()
	code.Console.Info("%.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(e.AllPlayouts)/sec, board.GetZ4(z), e.MovesNum, color, e.AllPlayouts)

	var recItem = new(e.RecordItemV02)
	recItem.Z = z
	recItem.Time = sec
	e.PutStoneOnRecord(board, z, color, recItem)
	printBoardOutOfPlayout(board, e.MovesNum)

	return z
}