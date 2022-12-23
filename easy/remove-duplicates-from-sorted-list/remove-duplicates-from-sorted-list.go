package remove_duplicates_from_sorted_list

//Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	var ptrNode, fwdNode *ListNode

	ptrNode = head
	fwdNode = ptrNode.Next

	for fwdNode != nil {
		valPtr := ptrNode.Val
		valFwd := fwdNode.Val

		if valPtr == valFwd {
			fwdNode = fwdNode.Next
		} else if valPtr != valFwd {
			ptrNode.Next = fwdNode
			ptrNode = fwdNode
			fwdNode = fwdNode.Next
		}
	}

	ptrNode.Next = nil

	return head
}
