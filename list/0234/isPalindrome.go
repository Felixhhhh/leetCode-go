package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	isPalindrome(Ints2List([]int{1, 3, 3, 1}))

}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	left := head
	right := reverse(slow)
	for right != nil {
		if right.Val != left.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}
	return true

	//len := 0
	//cur := head
	//for cur != nil {
	//	cur = cur.Next
	//	len++
	//
	//}
	//
	//if len == 0 || len == 1 {
	//	return true
	//}
	//
	//leftNode := head
	//for i := 0; i < len/2; i++ {
	//	leftNode = leftNode.Next
	//}
	//leftNode = reverse(leftNode)
	//for leftNode != nil {
	//	if leftNode.Val != head.Val {
	//		return false
	//	}
	//	leftNode = leftNode.Next
	//	head = head.Next
	//}
	//return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

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
