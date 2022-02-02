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
	getWinner *func(int) int,
	isDislike *func(int, int) bool) int {

	AllPlayouts++

	var color = turnColor
	var previousZ = 0
	var boardMax = e.SentinelBoardArea

	var playoutTrialCount = PlayoutTrialCount
	for trial := 0; trial < playoutTrialCount; trial++ {
		var empty = make([]int, boardMax)
		var emptyNum, z int

		// TODO 空点を差分更新できないか？ 毎回スキャンは重くないか？
		// 空点を記憶します
		var onPoint = func(z int) {
			if position.IsEmpty(z) { // 空点なら
				empty[emptyNum] = z
				emptyNum++
			}
		}
		position.IterateWithoutWall(onPoint)

		var r = 0
		var dislikeZ = e.Pass
		var randomPigeonX = GetRandomPigeonX(emptyNum) // 見切りを付ける試行回数を算出
		var i int
		for i = 0; i < randomPigeonX; i++ {
			if emptyNum == 0 { // 空点が無ければパスします
				z = e.Pass
			} else {
				r = rand.Intn(emptyNum) // 空点を適当に選びます
				z = empty[r]
			}

			var err = e.PutStone(position, z, color)
			if err == 0 { // 石が置けたなら

				// if !(*isDislike)(turnColor, z) { // 石を置きたくないわけでなければ
				break // 確定
				//}

				// dislikeZ = z // 候補が無かったときに使います
			}

			// 石を置かなかったら、その選択肢は削除します
			empty[r] = empty[emptyNum-1]
			emptyNum--
		}
		if i == randomPigeonX {
			z = dislikeZ
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
