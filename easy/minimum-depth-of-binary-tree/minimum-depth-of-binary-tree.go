package binary_tree_inorder_traversal

import "math"

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	return dfsDepth(root) + 1
}

func dfsDepth(root *TreeNode) int {
	if root == nil {
		return -1
	}
	leftHeight := dfsDepth(root.Left)
	rightHeight := dfsDepth(root.Right)

	if leftHeight == -1 && rightHeight > -1 {
		return rightHeight + 1
	}
	if rightHeight == -1 && leftHeight > -1 {
		return leftHeight + 1
	}
	return int(math.Min(float64(leftHeight), float64(rightHeight)) + 1)
}
