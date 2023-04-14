package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type question23 struct {
	para23
	ans23
}

// para 是参数
// one 代表第一个参数
type para23 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans23 struct {
	one []int
}

func main() {

	qs := []question23{

		{
			para23{[][]int{}},
			ans23{[]int{}},
		},

		{
			para23{[][]int{
				{1},
				{1},
			}},
			ans23{[]int{1, 1}},
		},

		{
			para23{[][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			}},
			ans23{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		{
			para23{[][]int{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			}},
			ans23{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		{
			para23{[][]int{
				{1},
				{9, 9, 9, 9, 9},
			}},
			ans23{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			para23{[][]int{
				{9, 9, 9, 9, 9},
				{1},
			}},
			ans23{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			para23{[][]int{
				{2, 3, 4},
				{4, 5, 6},
			}},
			ans23{[]int{2, 3, 4, 4, 5, 6}},
		},

		{
			para23{[][]int{
				{1, 3, 8},
				{1, 7},
			}},
			ans23{[]int{1, 1, 3, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 23------------------------\n")

	for _, q := range qs {
		var ls []*ListNode
		for _, qq := range q.para23.one {
			ls = append(ls, Ints2List(qq))
		}
		fmt.Printf("【input】:%v       【output】:%v\n", q.para23.one, List2Ints(mergeKLists(ls)))
	}
	fmt.Printf("\n\n\n")
}
func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

//使用优先队列
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) ==0 {
		return nil
	}
	h := new(IntHeap)
	dummy := &ListNode{-1, nil}
	result := dummy
	for _, v := range lists {
		if v != nil {
			heap.Push(h, v)
		}
	}
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		dummy.Next = node
		dummy = dummy.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return result.Next;
}

type IntHeap []*ListNode
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


//
//func mergeKLists(lists []*ListNode) *ListNode {
//	length := len(lists)
//	if length < 1 {
//		return nil
//	}
//	if length == 1 {
//		return lists[0]
//	}
//	num := length / 2
//	left := mergeKLists(lists[:num])
//	right := mergeKLists(lists[num:])
//	return mergeTwoLists1(left, right)
//}
//
//func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	if l1.Val < l2.Val {
//		l1.Next = mergeTwoLists1(l1.Next, l2)
//		return l1
//	}
//	l2.Next = mergeTwoLists1(l1, l2.Next)
//	return l2
//}


