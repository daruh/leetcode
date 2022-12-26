package same_tree

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil || q == nil {
		return p == nil && q == nil
	}

	if p.Val != q.Val {
		return false
	}

	vLeft := isSameTree(p.Left, q.Left)
	if vLeft != true {
		return false
	}

	vRight := isSameTree(p.Right, q.Right)
	if vRight != true {
		return false
	}
	return true

}
