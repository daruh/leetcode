package add_two_numbers

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var current *ListNode
	var prevList *ListNode
	var move int

	for {
		if (l1 == nil && l2 == nil) && move == 0 {
			break
		}

		var a int
		if l1 == nil {
			a = 0
		} else {
			a = l1.Val
		}

		var b int
		if l2 == nil {
			b = 0
		} else {
			b = l2.Val
		}

		result := a + b + move
		if result >= 10 {
			result = result - 10
			move = 1
		} else {
			move = 0
		}

		if current != nil {
			current = &ListNode{Val: result}
			prevList.Next = current
		}
		if current == nil {
			current = &ListNode{Val: result}
		}
		if head == nil {
			head = current
		}

		prevList = current
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head
}
