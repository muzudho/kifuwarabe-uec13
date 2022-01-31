package play_algorithm

import (
	"math/rand"

	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// Playout - 最後まで石を打ちます。得点を返します
// * `getWinner` - 地計算
//
// # Returns
//
// 手番が勝ったら 1、引分けなら 0、 相手が勝ったら -1
func Playout(
	position *e.Position,
	turnColor int,
	getWinner *func(int) int) int {

	AllPlayouts++

	var color = turnColor
	var previousZ = 0
	var boardMax = e.SentinelBoardArea

	var playoutTrialCount = PlayoutTrialCount
	for trial := 0; trial < playoutTrialCount; trial++ {
		var empty = make([]int, boardMax)
		var emptyNum, r, z int

		// TODO 空点を差分更新できないか？ 毎回スキャンは重くないか？
		// 空点を記憶します
		var onPoint = func(z int) {
			if !position.Exists(z) { // 空点なら
				empty[emptyNum] = z
				emptyNum++
			}
		}
		position.IterateWithoutWall(onPoint)

		r = 0
		var randomPigeonX = GetRandomPigeonX(emptyNum) // 見切りを付ける試行回数を算出
		var i int
		for i = 0; i < randomPigeonX; i++ {
			if emptyNum == 0 { // 空点が無ければ投了します
				z = 0
			} else {
				r = rand.Intn(emptyNum) // 空点を適当に選びます
				z = empty[r]
			}

			var err = e.PutStone(position, z, color)
			if err == 0 { // 石が置けたなら
				break
			}

			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if i == randomPigeonX {
			z = 0
		}

		// テストのときは棋譜を残します
		if FlagTestPlayout != 0 {
			position.Record[position.MovesNum].SetZ(z)
			position.MovesNum++
		}

		if z == 0 && previousZ == 0 {
			break
		}
		previousZ = z

		color = e.FlipColor(color)
	}

	return (*getWinner)(turnColor)
}
