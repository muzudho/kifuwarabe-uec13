package game_domain

// Empty triangle (アキ三角)
//
// x.
// xx

import (
	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// WrapIsDislike - 盤を束縛変数として与えます
func WrapIsDislike(position *e.Position) *func(int, int) bool {
	// 「手番の勝ちなら1、引き分けなら0、手番の負けなら-1を返す関数（自分視点）」を作成します
	// * `turnColor` - 手番の石の色
	var isDislike = func(turnColor int, z int) bool {
		// 座標取得
		// 432
		// 5S1
		// 678
		var eastZ = z + e.Dir4[e.East]
		var northEastZ = z + e.Dir4[e.North] + 1
		var northZ = z + e.Dir4[e.North]
		var northWestZ = z + e.Dir4[e.North] - 1
		var westZ = z + e.Dir4[e.West]
		var southWestZ = z + e.Dir4[e.South] - 1
		var southZ = z + e.Dir4[e.South]
		var southEastZ = z + e.Dir4[e.South] + 1

		// 東北
		// **
		// S*
		if isEmptyTriangle(turnColor, [3]int{eastZ, northEastZ, northZ}) {
			return true
		}

		// 西北
		// **
		// *S
		if isEmptyTriangle(turnColor, [3]int{northZ, northWestZ, westZ}) {
			return true
		}

		// 西南
		// *S
		// **
		if isEmptyTriangle(turnColor, [3]int{westZ, southWestZ, southZ}) {
			return true
		}

		// 東南
		// S*
		// **
		if isEmptyTriangle(turnColor, [3]int{southZ, southEastZ, eastZ}) {
			return true
		}

		return false
	}

	return &isDislike
}

func isEmptyTriangle(myColor int, points [3]int) bool {
	var myColorNum = 0
	var emptyNum = 0

	for _, v := range points {
		if v == myColor {
			myColorNum++
		} else if v == e.Empty {
			emptyNum++
		}
	}

	return myColorNum == 2 && emptyNum == 1
}
