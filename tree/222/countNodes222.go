package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
判断树是否相同
*/
func main() {
	root := &TreeNode{Val: 0}
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6

	check := countNodes(root)
	println(check)
}

func countNodes(root *TreeNode) int {
	var l, r = root, root
	// 沿最左侧和最右侧分别计算高度

	var hl, hr = 0, 0
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hr++
	}
	// 如果左右侧计算的高度相同，则是一棵满二叉树
	if hl == hr {
		return int(math.Pow(2, float64(hl)) - 1)
	}
	// 如果左右侧的高度不同，则按照普通二叉树的逻辑计算
	return 1 + countNodes(root.Left) + countNodes(root.Right)

}
