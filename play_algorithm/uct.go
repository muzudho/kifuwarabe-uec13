package play_algorithm

import (
	"math"
	"math/rand"
	"os"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// UCT
const (
	NodeMax   = 10000
	NodeEmpty = -1
	// Table index.
	IllegalZ = -1
)

// GetBestZByUct - Lesson08,09,09aで使用。 一番良いUCTである着手を選びます。 GetComputerMoveDuringSelfPlay などから呼び出されます。
//
// # Return
// (bestZ int, winRate float64)
func GetBestZByUct(
	position *e.Position,
	color int,
	searchUct *func(int, int) int) (int, float64) {

	// UCT計算フェーズ
	NodeNum = 0 // カウンターリセット
	var next = CreateNode(position)
	var uctLoopCount = UctLoopCount
	for i := 0; i < uctLoopCount; i++ {
		// 一時記憶
		var copiedBoard = position.CopyData()
		var copiedKoZ = position.KoZ

		(*searchUct)(color, next)

		// 復元
		position.KoZ = copiedKoZ
		position.ImportData(copiedBoard)
	}

	// ベスト値検索フェーズ
	var bestI = -1
	var max = -999
	var pN = &Nodes[next]
	for i := 0; i < pN.ChildNum; i++ {
		var c = &pN.Children[i]
		if c.Games > max {
			bestI = i
			max = c.Games
		}
		code.Console.Info("(UCT Calculating...) %2d:z=%04d,rate=%.4f,games=%3d\n", i, position.GetZ4(c.Z), c.Rate, c.Games)
	}

	// 結果
	var bestZ = pN.Children[bestI].Z
	code.Console.Info("(UCT Calculated    ) bestZ=%04d,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
		position.GetZ4(bestZ), pN.Children[bestI].Rate, max, AllPlayouts, NodeNum)
	return bestZ, pN.Children[bestI].Rate
}

// WrapSearchUct - 局面を束縛変数として与えます
func WrapSearchUct(position *e.Position) *func(int, int) int {
	var searchUct = func(color int, nodeN int) int {
		return SearchUct(position, color, nodeN)
	}

	return &searchUct
}

// SearchUct - 再帰関数。 GetBestZByUct() から呼び出されます
func SearchUct(
	position *e.Position,
	color int,
	nodeN int) int {

	var pN = &Nodes[nodeN]
	var c *Child

	for {
		var selectI = selectBestUcb(nodeN)
		c = &pN.Children[selectI]
		var z = c.Z

		var err = e.PutStone(position, z, color)
		if err == 0 { // 石が置けたなら
			break
		}

		c.Z = IllegalZ
		// code.Console.Debug("ILLEGAL:z=%04d\n", GetZ4(z))
	}

	var winner int // 手番が勝ちなら1、引分けなら0、手番の負けなら-1 としてください
	if c.Games <= 0 {
		winner = -Playout(position, e.FlipColor(color), GettingOfWinnerOnDuringUCTPlayout, IsDislike)
	} else {
		if c.Next == NodeEmpty {
			c.Next = CreateNode(position)
		}
		winner = -SearchUct(position, e.FlipColor(color), c.Next)
	}
	c.Rate = (c.Rate*float64(c.Games) + float64(winner)) / float64(c.Games+1)
	c.Games++
	pN.ChildGameSum++
	return winner
}

// 一番良い UCB を選びます。 SearchUct から呼び出されます。
func selectBestUcb(nodeN int) int {
	var pN = &Nodes[nodeN]
	var selectI = -1
	var maxUcb = -999.0
	var ucb = 0.0
	for i := 0; i < pN.ChildNum; i++ {
		var c = &pN.Children[i]
		if c.Z == IllegalZ {
			continue
		}
		if c.Games == 0 {
			ucb = 10000.0 + 32768.0*rand.Float64()
		} else {
			ucb = c.Rate + 1.0*math.Sqrt(math.Log(float64(pN.ChildGameSum))/float64(c.Games))
		}
		if maxUcb < ucb {
			maxUcb = ucb
			selectI = i
		}
	}

	// 異常終了
	if selectI == -1 {
		code.Console.Error("Err! select\n")
		os.Exit(0)
	}

	return selectI
}
