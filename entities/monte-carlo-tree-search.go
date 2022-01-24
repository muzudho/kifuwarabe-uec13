package entities

import (
	"fmt"
	"os"
)

// Child - 子。
type Child struct {
	// table index. 盤の交点の配列のインデックス。
	Z     int
	Games int
	Rate  float64
	Next  int
}

// Node - ノード。
type Node struct {
	ChildNum     int
	Children     []Child
	ChildGameSum int
}

// Nodes -ノード？
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
func CreateNode(board IBoardV02) int {

	if NodeNum == NodeMax {
		fmt.Printf("node over Err\n")
		os.Exit(0)
	}
	var pN = &Nodes[NodeNum]
	pN.ChildNum = 0
	pN.Children = make([]Child, board.UctChildrenSize())
	pN.ChildGameSum = 0

	var onPoint = func(z int) {
		if !board.Exists(z) {
			addChild(pN, z)
		}
	}

	board.IterateWithoutWall(onPoint)

	addChild(pN, 0)
	NodeNum++
	return NodeNum - 1
}
