package merge_two_sorted_lists

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var (
		head     *ListNode
		forward  *ListNode
		accessed bool
	)

	for {
		minimumNode, l1, l2 := findMinimumAndForward(list1, list2)
		list1 = l1
		list2 = l2
		if minimumNode == nil {
			break
		}
		if !accessed {
			head = minimumNode
			forward = head
			accessed = true
			continue
		}
		forward.Next = minimumNode
		forward = forward.Next
	}
	return head
}

func findMinimumAndForward(list1 *ListNode, list2 *ListNode) (minimunNode *ListNode, l1 *ListNode, l2 *ListNode) {

	var minimumNode *ListNode
	if list1 == nil && list2 == nil {
		return minimumNode, list1, list2
	}
	if list1 == nil && list2 != nil {
		minimumNode = list2
		list2 = list2.Next
		return minimumNode, list1, list2
	}
	if list1 != nil && list2 == nil {
		minimumNode = list1
		list1 = list1.Next
		return minimumNode, list1, list2
	}
	if list1.Val < list2.Val {
		minimumNode = list1
		list1 = list1.Next
		return minimumNode, list1, list2
	}
	minimumNode = list2
	list2 = list2.Next
	return minimumNode, list1, list2
}
