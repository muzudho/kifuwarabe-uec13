package entities

import (
	"os"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
)

// PutStoneOnRecord - SelfPlay, RunGtpEngine から呼び出されます
func PutStoneOnRecord(board *Board, z int, color int, recItem IRecordItemV01) {
	var err = PutStone(board, z, color)
	if err != 0 {
		code.Console.Error("(PutStoneOnRecord) Err!\n")
		os.Exit(0)
	}

	// 棋譜に記録
	Record[MovesNum] = recItem

	MovesNum++
}

// PutStone - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
func PutStone(board *Board, z int, color int) int {
	var around = [4][3]int{}
	var liberty, stone int
	var unCol = FlipColor(color)
	var space = 0
	var wall = 0
	var mycolSafe = 0
	var captureSum = 0
	var koMaybe = 0

	if z == 0 {
		KoZ = 0
		return 0
	}
	for dir := 0; dir < 4; dir++ {
		around[dir][0] = 0
		around[dir][1] = 0
		around[dir][2] = 0
		var z2 = z + Dir4[dir]
		var color2 = board.ColorAt(z2)
		if color2 == 0 {
			space++
		}
		if color2 == 3 {
			wall++
		}
		if color2 == 0 || color2 == 3 {
			continue
		}
		board.CountLiberty(z2, &liberty, &stone)
		around[dir][0] = liberty
		around[dir][1] = stone
		around[dir][2] = color2
		if color2 == unCol && liberty == 1 {
			captureSum += stone
			koMaybe = z2
		}
		if color2 == color && 2 <= liberty {
			mycolSafe++
		}

	}

	// 石を置けないケースを判定します
	if captureSum == 0 && space == 0 && mycolSafe == 0 {
		return 1
	}
	if z == KoZ {
		return 2
	}
	if wall+mycolSafe == 4 {
		return 3
	}
	if board.Exists(z) {
		return 4
	}

	for dir := 0; dir < 4; dir++ {
		var lib = around[dir][0]
		var color2 = around[dir][2]
		if color2 == unCol && lib == 1 && board.Exists(z+Dir4[dir]) {
			board.TakeStone(z+Dir4[dir], unCol)
		}
	}

	board.SetColor(z, color)
	board.CountLiberty(z, &liberty, &stone)

	if captureSum == 1 && stone == 1 && liberty == 1 {
		KoZ = koMaybe
	} else {
		KoZ = 0
	}

	return 0
}
