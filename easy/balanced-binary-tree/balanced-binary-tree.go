package binary_tree_inorder_traversal

import (
	"math"
)

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	if balanced(root) != -1 {
		return true
	}
	return false
}

func balanced(root *TreeNode) int {

	if root == nil {
		return 0
	}
	leftHeight := balanced(root.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := balanced(root.Right)
	if rightHeight == -1 {
		return -1
	}
	if math.Abs(float64(rightHeight-leftHeight)) > 1 {
		return -1
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight)) + 1)
}
