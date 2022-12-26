package same_tree

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	if root == nil || (root.Right == nil && root.Left == nil) {
		return true
	}

	vLeft := isSymmetric(root.Left)
	vRight := isSymmetric(root.Right)
	if vLeft != vRight {
		return false
	}
	subTreeLeft := root.Left
	subTreeRight := root.Right

	if subTreeLeft == nil || subTreeRight == nil {
		return subTreeLeft == nil && subTreeRight == nil
	}
	if subTreeLeft.Val != subTreeRight.Val {
		return false
	}

	if subTreeLeft.Left == nil && subTreeRight.Right != nil {
		return false
	}
	if subTreeLeft.Left.Val != subTreeRight.Right.Val {
		return false
	}

	if subTreeLeft.Right == nil && subTreeRight.Left != nil {
		return false
	}
	if subTreeLeft.Right.Val != subTreeRight.Left.Val {
		return false
	}
	if subTreeLeft.Left.Val != subTreeRight.Right.Val && subTreeLeft.Right.Val != subTreeRight.Left.Val {
		return false
	}
	return true

}
