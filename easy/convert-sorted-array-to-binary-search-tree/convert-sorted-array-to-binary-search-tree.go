package convert_sorted_array_to_binary_search_tree

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
        0
		0 1
		0 1 2
		0 1 2 3
		0 1 2 3 4
        0 1 2 3 4 5
*/

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	l := len(nums)
	i := l / 2
	var leftTree, rightTree *TreeNode
	if i > 0 {
		leftTree = sortedArrayToBST(nums[0:i])
	}
	if i+1 < l {
		rightTree = sortedArrayToBST(nums[i+1:])
	}

	return &TreeNode{
		Val:   nums[i],
		Left:  leftTree,
		Right: rightTree,
	}
}
