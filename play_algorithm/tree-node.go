package play_algorithm

import (
	"os"

	code "github.com/muzudho/kifuwarabe-uec13/coding_obj"
	e "github.com/muzudho/kifuwarabe-uec13/entities"
)

// Child - 子。
type Child struct {
	// table index. 盤の交点の配列のインデックス。
	Z     int
	Games int     // UCT検索をした回数？
	Rate  float64 // 勝率
	Next  int     // 配列のインデックス
}

// Node - ノード。
type Node struct {
	ChildNum     int
	Children     []Child
	ChildGameSum int
}

// Nodes -ノードの配列？
var Nodes = [NodeMax]Node{}

// NodeNum - ノード数？
var NodeNum = 0

// CreateNode から呼び出されます。
func addChild(pN *Node, z int) {
	var n = pN.ChildNum
	pN.Children[n].Z = z
	pN.Children[n].Games = 0
	pN.Children[n].Rate = 0.0
	pN.Children[n].Next = NodeEmpty
	pN.ChildNum++
}

// CreateNode - ノード作成。 searchUctV8, GetBestZByUct, searchUctLesson09 から呼び出されます。
func CreateNode(position *e.Position) int {

	if NodeNum == NodeMax {
		code.Console.Error("node over Err\n")
		os.Exit(0)
	}
	var pN = &Nodes[NodeNum]
	pN.ChildNum = 0
	pN.Children = make([]Child, position.UctChildrenSize())
	pN.ChildGameSum = 0

	var onPoint = func(z int) {
		if !position.Exists(z) {
			addChild(pN, z)
		}
	}

	position.IterateWithoutWall(onPoint)

	addChild(pN, 0)
	NodeNum++
	return NodeNum - 1
}
