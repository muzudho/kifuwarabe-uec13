package entities

import (
	"os"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
)

// PutStoneOnRecord - SelfPlay, RunGtpEngine から呼び出されます
func PutStoneOnRecord(position *Position, z int, color int, recItem *RecordItem) {
	var err = PutStone(position, z, color)
	if err != 0 {
		code.Console.Error("(PutStoneOnRecord) Err!\n")
		os.Exit(0)
	}

	// 棋譜に記録
	position.Record[position.MovesNum] = recItem
	position.MovesNum++
}

// PutStone - 石を置きます。
// * `z` - 交点。壁有り盤の配列インデックス
//
// # Returns
// エラーコード
func PutStone(position *Position, z int, color int) int {
	var around = [4]*Ren{}          // 隣接する４つの交点
	var libertyArea int             // 呼吸点の数
	var renArea int                 // 連の石の数
	var oppColor = FlipColor(color) //相手(opponent)の石の色
	var space = 0
	var wall = 0
	var mycolSafe = 0 // TODO この変数は必要か？
	var captureSum = 0

	if z == 0 { // 投了なら、コウを消して関数を正常終了
		position.KoZ = 0
		return 0
	}

	for dir := 0; dir < 4; dir++ { // ４方向
		around[dir] = NewRen(0, 0, 0) // 呼吸点の数, 連の石の数, 石の色

		var z2 = z + Dir4[dir]              // 隣の交点
		var adjColor = position.ColorAt(z2) // 隣(adjacent)の交点の石の色
		if adjColor == 0 {                  // 空点
			space++
			continue
		} else if adjColor == 3 { // 壁
			wall++
			continue
		}
		position.CountLiberty(z2, &libertyArea, &renArea)
		around[dir].LibertyArea = libertyArea         // 呼吸点の数
		around[dir].StoneArea = renArea               // 連の意地の数
		around[dir].Color = adjColor                  // 石の色
		if adjColor == oppColor && libertyArea == 1 { // 相手の石で、呼吸点が１つで、その呼吸点に今石を置いたなら
			captureSum += renArea
		}
		if adjColor == color && 2 <= libertyArea { // 隣接するのが自分の石で、その石が呼吸点を２つ持ってるようなら、コウではない
			mycolSafe++
		}

	}

	// 石を置けないケースを判定します
	if z == position.KoZ { // コウには置けません
		return 2
	}
	if captureSum == 0 && space == 0 && mycolSafe == 0 { // コウを除く、自分の眼
		return 1
	}
	if wall+mycolSafe == 4 { // 壁
		return 3
	}
	if position.Exists(z) { // 石の上には置けません
		return 4
	}

	for dir := 0; dir < 4; dir++ {
		var lib = around[dir].LibertyArea // 隣接する連の呼吸点の数
		var adjColor = around[dir].Color  // 隣接する連の石の色
		if adjColor == oppColor &&        // 隣接する連が相手の石で
			lib == 1 && // その呼吸点は１つで、そこに今石を置いた
			position.Exists(z+Dir4[dir]) { // 石はまだあるなら（上と右の石がくっついている、といったことを除外）
			position.TakeStone(z+Dir4[dir], oppColor)

			// もし取った石の数が１個ならそこはコウ。また、図形上、コウは１個しか出現しません
			if around[dir].StoneArea == 1 {
				position.KoZ = z + Dir4[dir]
			}
		}
	}

	position.SetColor(z, color)
	position.CountLiberty(z, &libertyArea, &renArea)

	return 0
}
