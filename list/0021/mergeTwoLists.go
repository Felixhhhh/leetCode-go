package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	mergeTwoLists(Ints2List([]int{1, 3, 3, 4}), Ints2List([]int{1, 2, 3, 4}))

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := &ListNode{-1, nil}
	p := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next

}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//	if l1.Val < l2.Val {
//		l1.Next = mergeTwoLists(l1.Next, l2)
//		return l1
//	}
//	l2.Next = mergeTwoLists(l1, l2.Next)
//	return l2
//}

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
