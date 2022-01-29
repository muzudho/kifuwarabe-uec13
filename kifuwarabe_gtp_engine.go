package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	e.Komi = config.Komi()
	e.MaxMovesNum = config.MaxMovesNum()
	e.SetBoardSize(config.BoardSize())
	var position = e.NewPosition(config.GetBoardArray())
	pl.InitPosition(position)

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
			// boardsize 19
			// 盤のサイズを変えます
			if 2 <= len(tokens) {
				var boardSize, err = strconv.Atoi(tokens[1])

				if err != nil {
					code.Console.Fatal(fmt.Sprintf("command=%s", command))
					panic(err)
				}

				e.SetBoardSize(boardSize)
				pl.InitPosition(position)

				code.Gtp.Print("= \n\n")
			} else {
				code.Gtp.Print("?What is a %s\n\n", command)
			}

		case "clear_board":
			pl.InitPosition(position)
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
			// komi 6.5
			if 2 <= len(tokens) {
				var komi, err = strconv.ParseFloat(tokens[1], 64)

				if err != nil {
					code.Console.Fatal(fmt.Sprintf("command=%s", command))
					panic(err)
				}

				e.Komi = komi
				code.Gtp.Print("= %d\n\n", e.Komi)
			} else {
				code.Gtp.Print("?What is a %s\n\n", command)
			}

			// TODO 消す code.Gtp.Print("= 6.5\n\n")

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
			var z = PlayComputerMoveLesson09a(position, color)
			code.Gtp.Print("= %s\n\n", p.GetGtpZ(position, z))

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

				var z = p.GetZFromGtp(position, tokens[2])
				var recItem = new(e.RecordItem)
				recItem.Z = z
				recItem.Time = 0
				e.PutStoneOnRecord(position, z, color, recItem)
				p.PrintBoard(position, position.MovesNum)

				code.Gtp.Print("= \n\n")
			}

		default:
			code.Gtp.Print("? unknown_command\n\n")
		}
	}
}

// PlayComputerMoveLesson09a - コンピューター・プレイヤーの指し手。 SelfPlay, RunGtpEngine から呼び出されます。
func PlayComputerMoveLesson09a(
	position *e.Position,
	color int) int {

	var z int
	var st = time.Now()
	pl.AllPlayouts = 0

	z = pl.GetBestZByUct(
		position,
		color,
		pl.SearchingOfUct)

	var sec = time.Since(st).Seconds()
	code.Console.Info("%.1f sec, %.0f playout/sec, play_z=%04d,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(pl.AllPlayouts)/sec, position.GetZ4(z), position.MovesNum, color, pl.AllPlayouts)

	var recItem = new(e.RecordItem)
	recItem.Z = z
	recItem.Time = sec
	e.PutStoneOnRecord(position, z, color, recItem)
	p.PrintBoard(position, position.MovesNum)

	return z
}
