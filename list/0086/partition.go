package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	partition(Ints2List([]int{1,4,3,2,5,2}),3)

}
func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{-1,nil}
	smallHead := small
	large := &ListNode{-1,nil}
	largeHead := large

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
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
