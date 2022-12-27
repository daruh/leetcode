package maximum_depth_of_binary_tree

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	maxL := maxDepth(root.Left)
	maxR := maxDepth(root.Right)
	if maxL > maxR {
		return maxL + 1
	}
	return maxR + 1
}
