package binary_tree_inorder_traversal

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	arrayLeft := inorderTraversal(root.Left)
	arrayRight := inorderTraversal(root.Right)
	l := append(arrayLeft, root.Val)
	r := append(l, arrayRight...)
	return r
}
