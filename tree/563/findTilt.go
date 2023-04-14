package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

	check := findTilt(root)
	println(check)
}

//https://leetcode.cn/problems/binary-tree-tilt/
//返回以node为根的树的结点之和。这个和也是供求ans使用的
// 也就是说一边求出这个和，一边求ans，ans在不断累加
// 后序遍历。 当前遍历的结点为node
func findTilt(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//第一次到达node
		//得到node的左子树之和
		sumLeft := dfs(node.Left)
		//第二次到达node
		//得到node的右子树之和
		sumRight := dfs(node.Right)
		//第三次到达node。算出node的坡度累加到ans，
		//并把以node作为根节点的树的结点之和往上反，供上层的结点使用
		ans += abs(sumLeft - sumRight)
		return sumLeft + sumRight + node.Val
	}
	dfs(root)
	return
}


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
